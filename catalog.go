package main

import (
	"fmt"
	"os"
)

type Catalog struct {
	ID2File map[int]DbFile
	ID2Name map[int]string
	ID2PKey map[int]string
	Name2ID map[string]int
}

func NewCatalog() *Catalog {

	ID2File := make(map[int]DbFile, 0)
	ID2Name := make(map[int]string, 0)
	ID2PKey := make(map[int]string, 0)
	Name2ID := make(map[string]int, 0)
	catalog := &Catalog{ID2File: ID2File, ID2Name: ID2Name, ID2PKey: ID2PKey, Name2ID: Name2ID}

	catalog.loadSchema()
	return catalog
}

func (c Catalog) addTable(file DbFile, tName string, pKeyField string) {
	fid := file.getID()
	c.ID2Name[fid] = tName
	c.ID2PKey[fid] = pKeyField
	c.Name2ID[tName] = fid
	c.ID2File[fid] = file
}

func (c Catalog) getTableID(name string) int {

	if v, ok := c.Name2ID[name]; ok {
		return v
	}
	return -1
}

func (c Catalog) getTupleDesc(tableID int) (*TupleDesc, error) {
	if file, ok := c.ID2File[tableID]; ok {
		return file.getTupleDesc(), nil
	}
	return &TupleDesc{}, fmt.Errorf("Can't found the TupleDesc")
}

func getPrimaryKey(tableID int) {

}

func (c Catalog) getDBFile(tableID int) DbFile {
	if f, ok := c.ID2File[tableID]; ok {
		return f
	}
	return nil
}

// TODO read from file
func (c *Catalog) loadSchema() {
	tableName := "hello_DB"
	types := []Type{INT_TYPE{}, INT_TYPE{}, INT_TYPE{}}
	names := []string{"f1", "f2", "f3"}

	td := NewTupleDesc(types, names)

	absFileName := fmt.Sprintf("/home/backman/GODB/DATA/%s.data", tableName)
	fd, err := os.Create(absFileName)
	check(err)

	// Create a HeapFile
	hf := NewHeapFile(fd, td)
	c.addTable(hf, tableName, names[0])
}
