package main

import "fmt"

var (
	g     *Graph
	nodes []Node
	edges []Edge
)

const nodesCount int = 10

func setUp() {
	g = EmptyGraph()

	for i := 0; i < nodesCount; i++ {
		node := NewNode(i)
		nodes = append(nodes, *node)

		g.addNode(*node)
	}

	for i := 0; i < nodesCount-1; i++ {
		edge := NewEdge(nodes[i], nodes[i+1])
		edges = append(edges, *edge)

		g.addEdge(*edge)
	}
}

func printGraph() {
	nCount := len((*g).nodes)
	eCount := len((*g).edges)

	fmt.Println("*------------------*")
	for i := 0; i < nCount; i++ {
		fmt.Println("Node:", (*g).nodes[i].val)
	}
	for i := 0; i < eCount; i++ {
		fmt.Println("Edge:", (*g).edges[i].m.val, (*g).edges[i].n.val)
	}
	fmt.Println("*------------------*")
}

func main() {
	setUp()

	nodes := []Node{}
	for i := 0; i < nodesCount; i++ {
		nodes = append(nodes, (*g).nodes[i])
	}

	printGraph()
	g.blindCut()
	printGraph()
}
