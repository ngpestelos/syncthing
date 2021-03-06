package xdr

import (
	"bytes"
	"testing"
	"testing/quick"
)

func TestPad(t *testing.T) {
	tests := [][]int{
		{0, 0},
		{1, 3},
		{2, 2},
		{3, 1},
		{4, 0},
		{32, 0},
		{33, 3},
	}
	for _, tc := range tests {
		if p := pad(tc[0]); p != tc[1] {
			t.Errorf("Incorrect padding for %d bytes, %d != %d", tc[0], p, tc[1])
		}
	}
}

func TestBytesNil(t *testing.T) {
	fn := func(bs []byte) bool {
		var b = new(bytes.Buffer)
		var w = NewWriter(b)
		var r = NewReader(b)
		w.WriteBytes(bs)
		w.WriteBytes(bs)
		r.ReadBytes()
		res := r.ReadBytes()
		return bytes.Compare(bs, res) == 0
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestBytesGiven(t *testing.T) {
	fn := func(bs []byte) bool {
		var b = new(bytes.Buffer)
		var w = NewWriter(b)
		var r = NewReader(b)
		w.WriteBytes(bs)
		w.WriteBytes(bs)
		res := make([]byte, 12)
		res = r.ReadBytesInto(res)
		res = r.ReadBytesInto(res)
		return bytes.Compare(bs, res) == 0
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}
