package main

type HeapPageID struct {
	tableID int
	pageNo  int
}

func (id HeapPageID) getTableID() int {
	return id.tableID
}

func (id HeapPageID) pageno() int {
	return id.pageNo
}

func (id HeapPageID) equals(v interface{}) bool {
	if v == nil {
		return false
	}
	vOther := v.(HeapPage)
	if id.pageNo == vOther.pid.pageno() && id.tableID == vOther.pid.getTableID() {
		return true
	}
	return false
}
