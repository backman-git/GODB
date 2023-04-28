package main

type RecordID struct {
	Pid     PageID
	TupleNo int
}

func NewRecordID(pid PageID, tupleNo int) *RecordID {
	return &RecordID{Pid: pid, TupleNo: tupleNo}
}

func (rid RecordID) getTupleNo() int {
	return rid.TupleNo
}

func (rid RecordID) getPageID() PageID {
	return rid.Pid
}
