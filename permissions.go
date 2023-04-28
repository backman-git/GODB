package main

const (
	READONLY = iota
	READWRITE
)

type Permissions struct {
	permLevel int
}

var READ_ONLY = Permissions{permLevel: READONLY}
var READ_WRITE = Permissions{permLevel: READWRITE}
