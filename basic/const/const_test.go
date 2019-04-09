package const_test

import "testing"

const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestGetYear(t *testing.T) {
	t.Log(Jan)
	t.Log(Feb)
	t.Log(Mar)
	t.Log(Apr)
	t.Log(May)
	t.Log(Jun)
	t.Log(Jul)
	t.Log(Aug)
	t.Log(Sep)
	t.Log(Oct)
	t.Log(Nov)
	t.Log(Dec)
}

func TestFilePerm(t *testing.T) {
	perm := 7 //0111
	t.Logf("Readable: %d\nWriteable: %d\nExecutable: %d\n", Readable, Writeable, Executable)
	t.Log(perm&Readable == Readable, perm&Writeable == Writeable, perm&Executable == Executable)
}
