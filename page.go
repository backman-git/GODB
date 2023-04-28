package main

import "io"

type Page interface {
	getNumTuples() int
	getHeaderSize() int
	getID() PageID
	readTuple(reader io.Reader, slotID int) *Tuple
	getPageData() []byte
	//createEmptyPageData()
	//deleteTuple()
	insertTuple(t *Tuple)
	markDirty(bool, TransactionID)
	isDirty() TransactionID
}
