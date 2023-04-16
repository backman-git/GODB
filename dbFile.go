package main

type TransactionID interface {
}

type DbFile interface {
	readPage(pageID int) Page
	writePage(pageID int) Page
	//insertTuple(txid TransactionID)
}
