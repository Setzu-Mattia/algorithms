package main

import (
	"math"
	"math/rand"
)

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

func (g *Graph) blindCut(survivors int) int {
	if len(g.nodes) < survivors {
		return 0
	}

	nodes := g.nodes
	survivorNodesIndex := rand.Intn(len(g.nodes) - 1)

	for i, id := 0, len(g.nodes)*len(g.nodes); i < len(nodes)-survivors; i++ {
		if i == survivorNodesIndex {
			continue
		}

		aNode := nodes[i]
		bNode := nodes[i+1]

		g.collapse(aNode, bNode, id)
	}

	return len(g.edges)
}

func (g *Graph) kargerMinCut(cuts int) int {
	var h *Graph
	paths := math.MaxInt8

	for i := 0; i < cuts; i++ {
		h = g.copy()

		h.blindCut(minCutSize)
		if len(h.edges) < paths {
			paths = len(g.edges)
		}
	}

	return paths
}

func (g *Graph) notSoBlindCut(contractLimit int, survivors int) int {
	// Blind cut up to contractLimit
	h := g.copy()

	// Contract: cut up to contract
	h.blindCut(contractLimit)
	j, k := g.copy(), g.copy()

	// Cut to minimum size
	jLen := j.blindCut(minCutSize)
	kLen := k.blindCut(minCutSize)

	if jLen <= kLen {
		return jLen
	}

	return kLen
}
