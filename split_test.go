package split

import (
	"testing"
)

func TestSplitToChunks(t *testing.T) {
	var i = []int{1, 2, 3, 4, 5}
	si := SplitToChunks(i, 1)
	if len(si) != 5 {
		panic("fail")
	}
	var s = []map[string]string{
		{"k1": "v1"},
		{"k2": "v2"},
		{"k3": "v3"},
		{"k4": "v4"},
		{"k5": "v5"},
	}
	ss := SplitToChunks(s, 2)
	if len(ss) != 3 {
		panic("fail")
	}
	if len(ss[2]) != 1 {
		panic("fail")
	}
}
