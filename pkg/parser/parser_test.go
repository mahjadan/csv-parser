package parser

import "testing"

func TestColumnIndices_GetIndex(t *testing.T) {
	indices := ColumnIndices{}
	index := indices.GetIndex("sss")
	t.Log(index)
}
