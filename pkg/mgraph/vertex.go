package mgraph

type VertexID string

type Vertex interface {
	Id() VertexID
	StoreData(data any)
	Data() any
	Incoming() []Edge
	Outgoing() []Edge
	Edges() []Edge
	BelongsTo(edge Edge) bool
	Degree() int
	Graph() Graph
}

func newVertex(id VertexID, graph *graph) *vertex {
	return &vertex{
		id:    id,
		graph: graph,
	}
}

type vertex struct {
	id    VertexID
	data  any
	graph *graph
}

func (v *vertex) Id() VertexID {
	return v.id
}

func (v *vertex) StoreData(data any) {
	v.data = data
}

func (v *vertex) Data() any {
	return v.data
}

func (v *vertex) Incoming() (edges []Edge) {
	for _, e := range v.graph.edgesTo[v.id] {
		edges = append(edges, e)
	}
	return
}

func (v *vertex) Outgoing() (edges []Edge) {
	for _, e := range v.graph.edgesFrom[v.id] {
		edges = append(edges, e)
	}
	return
}

func (v *vertex) Edges() []Edge {
	edges := make([]Edge, 0, v.Degree())
	for _, e := range v.graph.edgesTo[v.id] {
		edges = append(edges, e)
	}
	for _, e := range v.graph.edgesFrom[v.id] {
		if !e.IsLoop() {
			edges = append(edges, e)
		}
	}
	return edges
}

func (v *vertex) BelongsTo(edge Edge) bool {
	return v.id == edge.To() || v.id == edge.From()
}

func (v *vertex) Degree() int {
	return len(v.graph.edgesTo[v.id]) + len(v.graph.edgesFrom[v.id])
}

func (v *vertex) Graph() Graph {
	if v.graph != nil {
		return v.graph
	}
	return nil
}

func (v *vertex) onRemove() {
	v.data = nil
	v.graph = nil
}
