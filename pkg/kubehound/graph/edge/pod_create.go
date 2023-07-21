package edge

import (
	"context"
	"fmt"

	"github.com/DataDog/KubeHound/pkg/kubehound/graph/adapter"
	"github.com/DataDog/KubeHound/pkg/kubehound/graph/types"
	"github.com/DataDog/KubeHound/pkg/kubehound/graph/vertex"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/converter"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/cache"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/storedb"
	"github.com/DataDog/KubeHound/pkg/kubehound/store/collections"
	gremlin "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	Register(&PodCreate{}, RegisterGraphMutation)
}

type PodCreate struct {
	BaseEdge
}

type podCreateGroup struct {
	Role primitive.ObjectID `bson:"_id" json:"role"`
}

func (e *PodCreate) Label() string {
	return "POD_CREATE"
}

func (e *PodCreate) Name() string {
	return "PodCreate"
}

func (e *PodCreate) BatchSize() int {
	if e.cfg.LargeClusterOptimizations {
		// Under optimization this becomes a very cheap operation
		return e.cfg.BatchSize
	}

	return e.cfg.BatchSizeClusterImpact
}

func (e *PodCreate) Processor(ctx context.Context, oic *converter.ObjectIDConverter, entry any) (any, error) {
	typed, ok := entry.(*podCreateGroup)
	if !ok {
		return nil, fmt.Errorf("invalid type passed to processor: %T", entry)
	}

	rid, err := oic.GraphID(ctx, typed.Role.Hex())
	if err != nil {
		return nil, fmt.Errorf("%s edge role id convert: %w", e.Label(), err)
	}

	if e.cfg.LargeClusterOptimizations {
		return map[any]any{
			gremlin.T.Label: vertex.RoleLabel,
			gremlin.T.Id:    rid,
		}, nil
	}

	return rid, nil
}

func (e *PodCreate) Traversal() types.EdgeTraversal {
	return func(source *gremlin.GraphTraversalSource, inserts []any) *gremlin.GraphTraversal {
		g := source.GetGraphTraversal()
		if e.cfg.LargeClusterOptimizations {
			// In large clusters this can explode the number of edges and we can safely assume this is a critical issue
			g.
				Inject(inserts).
				Unfold().
				As("rpc").
				MergeV(__.Select("rpc")).
				Option(gremlin.Merge.OnCreate, __.Fail("missing role vertex on POD_CREATE insert")).
				Option(gremlin.Merge.OnMatch, map[any]any{
					"critical": true,
				}).
				AddE(e.Label()).
				Barrier().Limit(0)
		} else {
			// In smaller clusters we can still show the (large set of) attack paths generated by this attack
			g.V().
				HasLabel("Node").
				Has("class", "Node").
				As("n").
				V(inserts...).
				Has("critical", false).
				AddE(e.Label()).
				To("n").
				Barrier().Limit(0)
		}

		return g
	}
}

// Stream finds all roles that have pod/create or equivalent wildcard permissions.
func (e *PodCreate) Stream(ctx context.Context, store storedb.Provider, _ cache.CacheReader,
	callback types.ProcessEntryCallback, complete types.CompleteQueryCallback) error {

	roles := adapter.MongoDB(store).Collection(collections.RoleName)
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"rules": bson.M{
					"$elemMatch": bson.M{
						"$and": bson.A{
							bson.M{"$or": bson.A{
								bson.M{"apigroups": ""},
								bson.M{"apigroups": "*"},
							}},
							bson.M{"$or": bson.A{
								bson.M{"resources": "pods"},
								bson.M{"resources": "*"},
							}},
							bson.M{"$or": bson.A{
								bson.M{"verbs": "create"},
								bson.M{"verbs": "*"},
							}},
						},
					},
				},
			},
		},
		{
			"$project": bson.M{
				"_id": 1,
			},
		},
	}

	cur, err := roles.Aggregate(context.Background(), pipeline)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	return adapter.MongoCursorHandler[podCreateGroup](ctx, cur, callback, complete)
}
