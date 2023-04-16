package main

type RecordID struct {
	Pid      PageID
	TupleNum int
}

func NewRecordID(pid PageID, tupleNum int) *RecordID {
	return &RecordID{Pid: pid, TupleNum: tupleNum}
}

func (rid RecordID) getTupleNum() int {
	return rid.TupleNum
}

func (rid RecordID) getPageID() PageID {
	return rid.Pid
}
