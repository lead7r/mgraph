package mgraph

type EdgeID string

type Edge interface {
	Id() EdgeID
	From() VertexID
	To() VertexID
	StoreData(data any)
	Data() any
	Tail() Vertex
	Head() Vertex
	Endpoints() [2]Vertex
	IsIncident(vertex Vertex) bool
	IsInverted(edge Edge) bool
	IsParallel(edge Edge) bool
	IsLoop() bool
	Graph() Graph
}

func newEdge(id EdgeID, from, to VertexID, graph *graph) *edge {
	return &edge{
		id:    id,
		from:  from,
		to:    to,
		graph: graph,
	}
}

type edge struct {
	id    EdgeID
	from  VertexID
	to    VertexID
	data  any
	graph *graph
}

func (e *edge) Id() EdgeID {
	return e.id
}

func (e *edge) From() VertexID {
	return e.from
}

func (e *edge) To() VertexID {
	return e.to
}

func (e *edge) StoreData(data any) {
	e.data = data
}

func (e *edge) Data() any {
	return e.data
}

func (e *edge) Tail() Vertex {
	return e.graph.Vertex(e.from)
}

func (e *edge) Head() Vertex {
	return e.graph.Vertex(e.to)
}

func (e *edge) Endpoints() [2]Vertex {
	return [2]Vertex{e.graph.Vertex(e.from), e.graph.Vertex(e.to)}
}

func (e *edge) IsIncident(vertex Vertex) bool {
	return e.from == vertex.Id() || e.to == vertex.Id()
}

func (e *edge) IsInverted(edge Edge) bool {
	return e.from == edge.To() && e.to == edge.From()
}

func (e *edge) IsParallel(edge Edge) bool {
	return e.from == edge.From() && e.to == edge.To()
}

func (e *edge) IsLoop() bool {
	return e.from == e.to
}

func (e *edge) Graph() Graph {
	if e.graph != nil {
		return e.graph
	}
	return nil
}

func (e *edge) onRemove() {
	e.data = nil
	e.graph = nil
}
