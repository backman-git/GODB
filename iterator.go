package main

type Iterator interface {
	open()
	hasNext() bool
	next() Tuple
	//rewind()
	getTupleDesc()
	close()
}
