package vertex

import gremlingo "github.com/apache/tinkerpop/gremlin-go/driver"

const (
	identityLabel = "Identity"
)

var _ Builder = (*Identity)(nil)

type Identity struct {
}

func (v Identity) Label() string {
	return identityLabel
}

func (v Identity) BatchSize() int {
	return DefaultBatchSize
}

func (v Identity) Traversal() VertexTraversal {
	return func(g *gremlingo.GraphTraversalSource, inserts []TraversalInput) *gremlingo.GraphTraversal {
		return g.Inject(inserts).Unfold().As("c").
			AddV("Identity").
			Property("storeId", gremlingo.T__.Select("c").Select("store_id")).
			Property("name", gremlingo.T__.Select("c").Select("name")).
			Property("isNamespaced", gremlingo.T__.Select("c").Select("is_namespaced")).
			Property("namespace", gremlingo.T__.Select("c").Select("namespace")).
			Property("type", gremlingo.T__.Select("c").Select("type"))
	}
}
