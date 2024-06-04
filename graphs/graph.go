package graphs

type Graph[T any] struct {
	// Vertexes should be auto populated by adding edges? Currently done for easier rendering.
	nodes []Node[T]
	edges []Edge[T]
}

func New[T any]() Graph[T] {
	return Graph[T]{nodes: make([]Node[T], 0)}
}

func (g *Graph[T]) AddNode(vertex Node[T]) {
	g.nodes = append(g.nodes, vertex)
}

func (g *Graph[T]) GetNode() []Node[T] {
	return g.nodes
}

func (g *Graph[T]) AddEdge(fromVertex Node[T], toVertex Node[T]) {
	// Check to not allow same edges
	g.edges = append(g.edges, NewEdge(fromVertex, toVertex))
}

func (g *Graph[T]) GetEdges() []Edge[T] {
	return g.edges
}
