package graphs

type Graph[T any] struct {
	// Vertexes should be auto populated by adding edges? Currently done for easier rendering.
	vertexes []Vertex[T]
	edges    []Edge[T]
}

func New[T any]() Graph[T] {
	return Graph[T]{vertexes: make([]Vertex[T], 0)}
}

func (g *Graph[T]) AddVertex(vertex Vertex[T]) {
	g.vertexes = append(g.vertexes, vertex)
}

func (g *Graph[T]) GetVertexes() []Vertex[T] {
	return g.vertexes
}

func (g *Graph[T]) AddEdge(fromVertex Vertex[T], toVertex Vertex[T]) {
	g.edges = append(g.edges, NewEdge(fromVertex, toVertex))
}

func (g *Graph[T]) GetEdges() []Edge[T] {
	return g.edges
}
