package main

type TupleIterator struct {
	td     TupleDesc
	tuples []Tuple
	idx    int
}

func NewTupleIterator(td TupleDesc, tuples []Tuple) *TupleIterator {

	return &TupleIterator{td: td, tuples: tuples}
}

func (tIter *TupleIterator) open() {
	tIter.idx = 0
}

func (tIter TupleIterator) hasNext() bool {
	return tIter.idx < len(tIter.tuples)
}

func (tIter *TupleIterator) next() Tuple {
	if tIter.hasNext() {
		tuple := tIter.tuples[tIter.idx]
		tIter.idx++
		return tuple
	}
	return Tuple{}
}

func (tIter *TupleIterator) rewind() {
	tIter.idx = 0
}

func (tIter TupleIterator) getTupleDesc() TupleDesc {
	return tIter.td
}

func (tIter *TupleIterator) close() {
	tIter.idx = len(tIter.tuples)
}
