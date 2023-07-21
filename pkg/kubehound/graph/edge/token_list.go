package edge

import (
	"context"
	"fmt"

	"github.com/DataDog/KubeHound/pkg/kubehound/graph/adapter"
	"github.com/DataDog/KubeHound/pkg/kubehound/graph/types"
	"github.com/DataDog/KubeHound/pkg/kubehound/models/converter"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/cache"
	"github.com/DataDog/KubeHound/pkg/kubehound/storage/storedb"
	"github.com/DataDog/KubeHound/pkg/kubehound/store/collections"
	gremlin "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	Register(&TokenList{}, RegisterGraphMutation)
}

type TokenList struct {
	BaseEdge
}

type tokenListGroup struct {
	Role primitive.ObjectID `bson:"_id" json:"role"`
}

func (e *TokenList) Label() string {
	return "TOKEN_LIST"
}

func (e *TokenList) Name() string {
	return "TokenListCluster"
}

func (e *TokenList) BatchSize() int {
	if e.cfg.LargeClusterOptimizations {
		// Under optimization this becomes a very cheap operation
		return e.cfg.BatchSize
	}

	return e.cfg.BatchSizeClusterImpact
}

func (e *TokenList) Processor(ctx context.Context, oic *converter.ObjectIDConverter, entry any) (any, error) {
	typed, ok := entry.(*tokenListGroup)
	if !ok {
		return nil, fmt.Errorf("invalid type passed to processor: %T", entry)
	}

	rid, err := oic.GraphID(ctx, typed.Role.Hex())
	if err != nil {
		return nil, fmt.Errorf("%s edge role id convert: %w", e.Label(), err)
	}

	return rid, nil
}

func (e *TokenList) Traversal() types.EdgeTraversal {
	return func(source *gremlin.GraphTraversalSource, inserts []any) *gremlin.GraphTraversal {
		g := source.GetGraphTraversal()
		if e.cfg.LargeClusterOptimizations {
			// For larger clusters simply target the system:masters group to reduce redundant attack paths
			g.V().
				HasLabel("Identity").
				Has("name", "system:masters").
				As("i").
				V(inserts...).
				Has("critical", false).
				AddE(e.Label()).
				To("i").
				Barrier().Limit(0)
		} else {
			// In smaller clusters we can still show the (large set of) attack paths generated by this attack
			g.V().
				HasLabel("Identity").
				Has("class", "Identity").
				As("i").
				V(inserts...).
				Has("critical", false).
				AddE(e.Label()).
				To("i").
				Barrier().Limit(0)
		}

		return g
	}
}

// Stream finds all roles that are NOT namespaced and have secrets/list or equivalent wildcard permissions.
func (e *TokenList) Stream(ctx context.Context, store storedb.Provider, _ cache.CacheReader,
	callback types.ProcessEntryCallback, complete types.CompleteQueryCallback) error {

	roles := adapter.MongoDB(store).Collection(collections.RoleName)
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"is_namespaced": false,
				"rules": bson.M{
					"$elemMatch": bson.M{
						"$and": bson.A{
							bson.M{"$or": bson.A{
								bson.M{"apigroups": ""},
								bson.M{"apigroups": "*"},
							}},
							bson.M{"$or": bson.A{
								bson.M{"resources": "secrets"},
								bson.M{"resources": "*"},
							}},
							bson.M{"$or": bson.A{
								bson.M{"verbs": "list"},
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

	return adapter.MongoCursorHandler[tokenListGroup](ctx, cur, callback, complete)
}
