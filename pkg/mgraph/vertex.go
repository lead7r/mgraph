package mgraph

type VertexID string

type Vertex[T any] interface {
	Id() VertexID
	StoreData(data T)
	Data() T
}

func newVertex[T any](id VertexID) *vertex[T] {
	return &vertex[T]{
		id: id,
	}
}

type vertex[T any] struct {
	id   VertexID
	data T
}

func (v *vertex[T]) Id() VertexID {
	return v.id
}

func (v *vertex[T]) StoreData(data T) {
	v.data = data
}

func (v *vertex[T]) Data() T {
	return v.data
}
