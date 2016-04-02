package main

type Node struct {
	val int
}

type Edge struct {
	m Node
	n Node
}

type Graph struct {
	nodes []Node
	edges []Edge
}

func NewNode(val int) *Node {
	return &(Node{val})
}

func NewEdge(m, n Node) *Edge {
	return &(Edge{m, n})
}

func EmptyGraph() *Graph {
	return &(Graph{[]Node{}, []Edge{}})
}

func (g *Graph) addNode(n Node) {
	(*g).nodes = append((*g).nodes, n)
	return
}

func (g *Graph) addNodes(n []Node) {
	for i := 0; i < len(n); i++ {
		(*g).nodes = append((*g).nodes, n[i])
	}
}

func (g *Graph) addEdge(e Edge) {
	(*g).edges = append((*g).edges, e)
}

func (g *Graph) addEdges(e []Edge) {
	for i := 0; i < len(e); i++ {
		(*g).edges = append((*g).edges, e[i])
	}
}
