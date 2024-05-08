package graphs

type Edge[T any] struct {
	fromVertex Vertex[T]
	toVertex   Vertex[T]
	// fromCardinality - 1, 0..1, 0..*, 1..*, etc.
	// toCardinality - 1, 0..1, 0..*, 1..*, etc.
	// description - e.g. "From has many to"
}

func NewEdge[T any](fromVertex Vertex[T], toVertex Vertex[T]) Edge[T] {
	return Edge[T]{fromVertex: fromVertex, toVertex: toVertex}
}

func (e Edge[T]) GetFromVertex() Vertex[T] {
	return e.fromVertex
}

func (e Edge[T]) GetToVertex() Vertex[T] {
	return e.toVertex
}