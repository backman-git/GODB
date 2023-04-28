package main

type SeqScan struct {
	tableID int
}

func NewSeqScan(tid TransactionID, tableID int) *SeqScan {
	return &SeqScan{tableID: tableID}
}

func open() {

}

// Return a TupleDesc with table Name
func (ss SeqScan) getTupleDesc() *TupleDesc {

	//td, err := db.getCatalog().getTupleDesc(ss.tableID)
	//check(err)

}
