package main

type DB struct {
	catalog    *Catalog
	bufferPool *BufferPool
}

var db = NewDB()

func NewDB() *DB {
	catlog := NewCatalog()
	bufferPool := NewBufferPool(DEFAULTNUMPAGE)
	return &DB{catalog: catlog, bufferPool: bufferPool}
}

func (db DB) getCatalog() *Catalog {
	return db.catalog
}

func (db DB) getBufferPool() *BufferPool {
	return db.bufferPool
}
