package main

import "fmt"

const minCutSize int = 2

func (g *Graph) collapse(a, b Node) *Graph {
	fmt.Println("Collapsing", a.val, b.val)
	nodes := []Node{}

	// Collapse nodes
	for i := 0; i < len((*g).nodes); i++ {
		node := &((*g).nodes[i])

		if !(node.val == a.val || node.val == b.val) {
			nodes = append(nodes, *node)
		}
	}
	// New collapsed node
	collapsedNode := NewNode(len((*g).nodes))
	nodes = append(nodes, *collapsedNode)

	// Collapse edges
	edges := []Edge{}
	for i := 0; i < len((*g).edges); i++ {
		edge := ((*g).edges[i])
		mNode := edge.m
		nNode := edge.n

		switch {
		// NC - NC | Edge between not collapsing nodes
		case !(mNode.val == a.val || mNode.val == b.val) && !(nNode.val == a.val || nNode.val == b.val):
			edges = append(edges, edge)

		// C - NC | Edge on collapsed nodes: delete it
		case (mNode.val == a.val || mNode.val == b.val) && !(nNode.val == a.val || nNode.val == b.val):
			newEdge := NewEdge(*collapsedNode, nNode)
			edges = append(edges, *newEdge)

		// NC - C | Edge from collapsed node to not collapsed node
		case !(mNode.val == a.val || mNode.val == b.val) && (nNode.val == a.val || nNode.val == b.val):
			newEdge := NewEdge(mNode, *collapsedNode)
			edges = append(edges, *newEdge)

		// C - C | Edge from not collapsed node to collapsed node
		case (mNode.val == a.val || mNode.val == b.val) && (nNode.val == a.val || nNode.val == b.val):

		}
	}

	// Refresh the graph
	(*g).nodes = nodes
	(*g).edges = edges
	return g
}

func (g *Graph) blindCut() {
	if len(g.nodes) < minCutSize {
		return
	}

	nodes := g.nodes

	// Cut all but the survivor nodes
	for i := 0; i < len(nodes)-minCutSize; i++ {
		aNode := nodes[i]
		bNode := nodes[i+1]

		g.collapse(aNode, bNode)
	}
}
