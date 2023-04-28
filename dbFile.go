package main

type TransactionID interface {
}

type DbFile interface {
	readPage(pageID PageID) Page
	writePage(page Page)
	insertTuple(tid TransactionID, t *Tuple) Page
	deleteTuple(tid TransactionID, t *Tuple) Page
	getID() int
	getTupleDesc() *TupleDesc
	getNumPage() int
	getIterator() DBIterator
}
