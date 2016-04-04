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

func (g *Graph) copy() *Graph {
	h := EmptyGraph()

	for i := 0; i < len(g.nodes); i++ {
		val := g.nodes[i].val
		node := NewNode(val)

		h.addNode(*node)
	}

	for i := 0; i < len(g.edges); i++ {
		edge := g.edges[i]
		mNode := edge.m
		nNode := edge.n

		var hMNode Node
		var hNNode Node

		for j := 0; j < len(h.nodes); j++ {
			if h.nodes[j].val == mNode.val {
				hMNode = mNode
			}
			if h.nodes[j].val == nNode.val {
				hNNode = nNode
			}
		}

		hEdge := NewEdge(hMNode, hNNode)
		h.addEdge(*hEdge)
	}

	return h
}
