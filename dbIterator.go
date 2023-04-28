package main

// Iterator Model
type DBIterator interface {
	open()
	hasNext() bool
	next() *Tuple
	rewind()
	getTupleDesc()
	close()
}
