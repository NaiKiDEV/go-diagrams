package graphs

type Vertex[T any] struct {
	id   int16
	data T // This will be useful for different types of Vertexes (e.g. UML Class, start node, end node, decision, etc.)
}

func NewVertex[T any](id int16, data T) Vertex[T] {
	return Vertex[T]{id: id, data: data}
}

func (v Vertex[T]) GetId() int16 {
	return v.id
}

func (v Vertex[T]) GetData() T {
	return v.data
}