package it.unipi.di.msetzu


import scala.collection.mutable


class BinaryRangeTree(v:Int) extends mutable.ArraySeq(v) {
    def this(v:Int) = {this();
        new mutable.ArraySeq(v)}
}