package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestParse(t *testing.T) {

	intType := INT_TYPE{}
	byteReader := bytes.NewReader([]byte{123, 0, 0, 0})
	intField := intType.parse(byteReader)

	assert.Equal(t, intField.(IntField).getValue(), 123)
}
