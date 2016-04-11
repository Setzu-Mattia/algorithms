//package it.msetzu.scala.algII

import scala.language.postfixOps
import scala.util.Random

import scala.collection.immutable.Set

case class Vertex(val held:Any)

case class Edge(val aVertex:Vertex, val bVertex:Vertex)

case class Graph(val vertexes:Set[Vertex], val edges:List[Edge]) {
	def + [E](plus:List[E]) = plus headOption match {
		case Some(Vertex(v)) => Graph(vertexes ++ plus.asInstanceOf[List[Vertex]], edges)
		case Some(Edge(a, b)) => Graph(vertexes, plus.asInstanceOf[List[Edge]] ::: edges)
		case _ => this
	}

	def + (vertexes:List[Vertex], edges:List[Edge]) = Graph(this.vertexes ++ vertexes, this.edges ::: edges)

	def - (vertex:Vertex) =	Graph(this.vertexes filter { _ != vertex }, edges)

	def - (edge:Edge) =	Graph(this.vertexes, this.edges filter { _ != edge })

	private def isCollapsed (vertex:Vertex, edge:Edge):Boolean = edge.aVertex == vertex || edge.bVertex == vertex

	def collapse (edge:Edge) = {
		val collapsedVertex = Vertex(edge.aVertex.held + " - " + edge.bVertex.held)
		var h = Graph(this.vertexes, this.edges)

		h = h + List(collapsedVertex)
		h = h - edge.aVertex
		h = h - edge.bVertex
		h = h - edge

		val aVertex = edge.aVertex
		val bVertex = edge.bVertex

		var edges = List[Edge]()
		edges = edges ++ (edges filter { edge => (isCollapsed(aVertex, edge) && !isCollapsed(bVertex, edge)) } map {edge => Edge(collapsedVertex, bVertex) })
		edges = edges ++ (edges filter { edge => !isCollapsed(aVertex, edge) && isCollapsed(bVertex, edge) } map { edge => Edge(aVertex, collapsedVertex) })
		edges = edges ++ (edges filter { edge => !isCollapsed(aVertex, edge) && !isCollapsed(bVertex, edge) }) //map { edge => Edge(edge.aVertex, collapsedVertex) })

		h
	}
}

object MinCut extends scala.App {
	var g = Graph(Set[Vertex](), List[Edge]())
	val vertexes = Random.nextInt(5) + 1

	for (i <- 0 to vertexes) g = g + List(Vertex(i))

	for (i <- 0 to Random.nextInt(15)) {
		val edges = Random.nextInt(3) + 1
		val aVertex = Vertex(Random.nextInt(vertexes))
		val bVertex = Vertex(Random.nextInt(vertexes))

		for (edge <- 0 to edges if aVertex != bVertex) {
			g = g + List(Edge(aVertex, bVertex))
		}
	}

	g.vertexes foreach {println(_)}
	g.edges foreach {edge => println(edge.aVertex + " - " + edge.bVertex)}
}
