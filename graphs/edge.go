package graphs

type Edge[T any] struct {
	fromVertex Node[T]
	toVertex   Node[T]
	// fromCardinality - 1, 0..1, 0..*, 1..*, etc.
	// toCardinality - 1, 0..1, 0..*, 1..*, etc.
	// description - e.g. "From has many to"
}

func NewEdge[T any](fromVertex Node[T], toVertex Node[T]) Edge[T] {
	return Edge[T]{fromVertex: fromVertex, toVertex: toVertex}
}

func (e Edge[T]) GetFromVertex() Node[T] {
	return e.fromVertex
}

func (e Edge[T]) GetToVertex() Node[T] {
	return e.toVertex
}
