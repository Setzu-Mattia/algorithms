package it.msetzu.scala.algII


case class Vertex[T](val held:T) {
  override def == (o:Object) = o match {
    case ver:Vertex[T] if ver.held == held => true
    case _ => false
  }
}


object Vertex {
  def apply (held:T) = new Vertex(held)
}


case class Edge[T](val aVertex:Vertex[T], val bVertex:Vertex[T]) {
  override def == (o:Object) = o match {
    case e(v1 @ Vertex(val1), v2 @ Vertex(val2))
          if (val1 == aVertex.held && val2 == bVertex.held) => true
    case _ => false
  }
}


object Edge {
  def apply (aVertex:Vertex, bVertex:Vertex) = new Edge(aVertex, bVertex)
}


case class Graph[T](val vertexes:List[Vertex[T]], val edges:List[Edge[T]]) extends AbstractSeq[T] {
  def + (vertexes:List[Vertex[T]]) = Graph(this.vertexes :: vertexes, edges)

  def - (vertex:Vertex[T]) = {
    var found = false

    Graph(vertexes filter {
        if (!found && _ == vertex)
          found = true
      }, edges)
  }

  def + (edges:List[Edge[T]]) = Graph(vertexes, this.edges :: edges)

  def - (edge:Edge[T]) = {
    var found = false

    Graph(vertexes, edges filter {
        if (!found && _ == edge)
          found = true
      })
  }
  def + (vertexes:List[Vertex[T]], edges:List[Edge[T]]) = Graph(this.vertexes :: vertexes, this.edges :: edges)
}


object Graph {
  def apply (vertexes:List[Vertex[T]]) = new Graph[T](vertexes, {})
  def apply (vertexes:List[Vertex[T]], edges:List[Edge[T]]) = new Graph[T](vertexes, edges)
}


object MinCut extends scala.App {

}
