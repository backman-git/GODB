package main

type HeapPageID struct {
	tableID int
	pageNum int
}

func (id HeapPageID) getTableID() int {
	return id.tableID
}

func (id HeapPageID) pageno() int {
	return id.pageNum
}

func (id HeapPageID) equals(v interface{}) bool {
	if v == nil {
		return false
	}
	vOther := v.(HeapPage)
	if id.pageNum == vOther.pid.pageno() && id.tableID == vOther.pid.getTableID() {
		return true
	}
	return false
}
