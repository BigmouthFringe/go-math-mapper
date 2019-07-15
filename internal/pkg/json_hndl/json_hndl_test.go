package json_hndl

import (
	"testing"
)

var (
	invParamsTable [][]string
	valParamsTable [][]string
)

func init() {
	invParamsTable = [][]string{
		{"", ""},
		{""},
		{"sample"},
		{"sample", ""},
	}
	valParamsTable = [][]string{
		{"sample", "sample", ""},
		{"sample", "sample"},
	}
}

func TestAssertValidParams(t *testing.T) {
	for _, invParams := range invParamsTable {
		if err := assertValidParams(invParams); err == nil {
			t.Fail()
		}
	}
	for _, valParams := range valParamsTable {
		if err := assertValidParams(valParams); err != nil {
			t.Fail()
		}
	}
}