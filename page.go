package main

type Page interface {
	//getNumTuples()
	//getHeaderSize()
	getID() PageID
	//readNextTuple()
	getPageData() []byte
	//createEmptyPageData()
	//deleteTuple()
	insertTuple(t *Tuple)
	//markDirty()
	//isDirty()
}
