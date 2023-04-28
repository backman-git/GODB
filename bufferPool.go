package main

import "fmt"

// TODO
const (
	PAGESIZE       = 4096
	DEFAULTNUMPAGE = 50
)

type BufferPool struct {
	// page table
	pages        map[PageID]Page
	NumPageLimit int
}

func NewBufferPool(numPage int) *BufferPool {
	pages := make(map[PageID]Page, numPage)
	return &BufferPool{pages: pages, NumPageLimit: numPage}
}

func (bp BufferPool) getPageSize() int {
	return PAGESIZE
}

func (bp BufferPool) getPage(tid TransactionID, pid PageID, perm Permissions) Page {

	var page Page
	if p, ok := bp.pages[pid]; ok {
		page = p
	} else {
		// new page
		file := db.getCatalog().getDBFile(pid.getTableID())
		if pid.pageno() < file.getNumPage() {
			page = file.readPage(pid)
		} else {
			page = NewHeapPage(pid, createEmptyPageData())

		}

		// When the page pool is full, evict page back to disk
		if len(bp.pages) == bp.NumPageLimit {
			bp.evictPage()
		}
		bp.pages[pid] = page
	}

	return page
}

// One Tuple find one page
func (bp BufferPool) insertTuple(tid TransactionID, tableId int, t *Tuple) {
	page := db.getCatalog().getDBFile(tableId).insertTuple(tid, t)
	page.markDirty(true, tid)
}

func deleteTuple(tid TransactionID, t *Tuple) {

	pid := t.getRecordID().getPageID()

	if pid == nil {
		check(fmt.Errorf("No Tuple to delete"))
	}

	db.getCatalog().getDBFile(pid.getTableID()).deleteTuple(tid, t).markDirty(true, tid)
}

func (bp BufferPool) flushPage(pid PageID) {
	var page Page
	if p, ok := bp.pages[pid]; ok {
		page = p
	} else {
		return
	}

	file := db.getCatalog().getDBFile(pid.getTableID())
	file.writePage(page)
	page.markDirty(false, 0)
}

func (bp BufferPool) flushPagesByTID(tid TransactionID) {

	//
}

func (bp *BufferPool) evictPage() {

	for k, v := range bp.pages {
		if v.isDirty() == nil {
			bp.flushPage(k)
			delete(bp.pages, k)
			return
		}
	}

	check(fmt.Errorf("No page can be evicted..."))
}

func (bp *BufferPool) discardPage(pid PageID) {
	delete(bp.pages, pid)
}
