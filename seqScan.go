package main

type SeqScan struct {
	tableID   int
	itr       DBFileIterator
	aliasName string
}

func NewSeqScan(tid TransactionID, tableID int) *SeqScan {
	dbFileItr := db.getCatalog().getDBFile(tableID).getIterator(tid)
	return &SeqScan{tableID: tableID, itr: dbFileItr}
}

func (ss SeqScan) open() {
	ss.itr.open()
}

// Return a TupleDesc of table with alias name
func (ss SeqScan) getTupleDesc() *TupleDesc {

	td, err := db.getCatalog().getTupleDesc(ss.tableID)
	check(err)

	types := make([]Type, td.numFields())
	filedNames := make([]string, td.numFields())

	for idx, f := range td.FieldDescs {
		types[idx] = f.TypeValue
		filedNames[idx] = ss.aliasName + "." + f.FieldName
	}
	return NewTupleDesc(types, filedNames)
}

func (ss SeqScan) hasNext() bool {
	return ss.itr.hasNext()
}

func (ss SeqScan) next() Tuple {
	return ss.itr.next()
}

func (ss SeqScan) close() {
	ss.itr.close()
}
