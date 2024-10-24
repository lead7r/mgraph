package mgraph

type EdgeID string

type Edge[T any] interface {
	Id() EdgeID
	From() VertexID
	To() VertexID
	StoreData(data T)
	Data() T
}

func newEdge[T any](id EdgeID, from, to VertexID) *edge[T] {
	return &edge[T]{
		id:   id,
		from: from,
		to:   to,
	}
}

type edge[T any] struct {
	id   EdgeID
	from VertexID
	to   VertexID
	data T
}

func (e *edge[T]) Id() EdgeID {
	return e.id
}

func (e *edge[T]) From() VertexID {
	return e.from
}

func (e *edge[T]) To() VertexID {
	return e.to
}

func (e *edge[T]) StoreData(data T) {
	e.data = data
}

func (e *edge[T]) Data() T {
	return e.data
}
