package main

import "math/rand"

const minCutSize int = 2

func contains(slice []Node, element Node) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i].val == element.val {
			return true
		}
	}
	return false
}

func (g *Graph) collapse(a, b Node, id int) *Graph {
	nodes := []Node{}

	// Collapse nodes
	for i := 0; i < len((*g).nodes); i++ {
		node := &((*g).nodes[i])

		if !(node.val == a.val || node.val == b.val) {
			if !contains(nodes, *node) {
				nodes = append(nodes, *node)
			}
		}
	}
	// New collapsed node
	collapsedNode := NewNode(id)

	if !contains(nodes, *collapsedNode) {
		nodes = append(nodes, *collapsedNode)
	}

	// Collapse edges
	edges := []Edge{}
	for i := 0; i < len(g.edges); i++ {
		edge := g.edges[i]
		mNode := edge.m
		nNode := edge.n

		switch {
		case mNode.val == nNode.val:

		// NC - NC | Edge between not collapsing nodes
		case !(mNode.val == a.val || mNode.val == b.val) && !(nNode.val == a.val || nNode.val == b.val):
			edges = append(edges, edge)

		// C - NC | Edge on collapsed nodes: delete it
		case (mNode.val == a.val || mNode.val == b.val) && !(nNode.val == a.val || nNode.val == b.val) &&
			mNode.val != collapsedNode.val:
			newEdge := NewEdge(*collapsedNode, nNode)
			edges = append(edges, *newEdge)

		// NC - C | Edge from collapsed node to not collapsed node
		case !(mNode.val == a.val || mNode.val == b.val) && (nNode.val == a.val || nNode.val == b.val) &&
			mNode.val != collapsedNode.val:
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
	survivorNodesIndex := rand.Intn(len(g.nodes) - 1)

	for i, id := 0, len(g.nodes)*len(g.nodes); i < len(nodes)-minCutSize; i++ {
		if i == survivorNodesIndex {
			continue
		}

		aNode := nodes[i]
		bNode := nodes[i+1]

		g.collapse(aNode, bNode, id)
	}
}
