package graphs

type Node[T any] struct {
	id   int
	data T // This will be useful for different types of Vertexes (e.g. UML Class, start node, end node, decision, etc.)
}

func NewNode[T any](id int, data T) Node[T] {
	return Node[T]{id: id, data: data}
}

func (v Node[T]) GetId() int {
	return v.id
}

func (v Node[T]) GetData() T {
	return v.data
}
