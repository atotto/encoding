package csv

import (
	"bytes"
	"testing"
)

var writeTests = []struct {
	Input  interface{}
	Output string
}{
	{
		Input: []struct {
			Int   int
			Float float32
		}{
			{2, 22.00003},
			{2, 0.00003}},
		Output: "2,22.0003",
	},
}

type testStruct struct {
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Float32 float32
	Float64 float64
	String  string
	// Struct  s
	// Array []int
	// Map   map[string]int
	// Slice
}

type s struct {
	Int    int
	String string
}

var tests = []struct {
	Struct []testStruct
	CSV    string
}{
	{
		[]testStruct{{
			Bool:    false,
			Int:     234,
			Int8:    127,
			Int16:   345,
			Int32:   -234,
			Int64:   123,
			Uint:    123,
			Uint8:   255,
			Uint16:  23456,
			Uint32:  34567,
			Uint64:  102345,
			Float32: 0.23456,
			Float64: 0.00000023,
			String:  "hello",
			//Struct:  s{Int: 1, String: "日本語"},
		}},
		"false,234,127,345,-234,123,123,255,23456,34567,102345,0.23456,2.3e-07,hello\n",
	},
	{
		[]testStruct{{
			Bool:    true,
			Int:     234,
			Int8:    127,
			Int16:   345,
			Int32:   -234,
			Int64:   123,
			Uint:    123,
			Uint8:   255,
			Uint16:  23456,
			Uint32:  34567,
			Uint64:  102345,
			Float32: 0.23456,
			Float64: 0.00000023,
			String:  "hello",
			//Struct:  s{Int: 1, String: "日本語"},
		}},
		"true,234,127,345,-234,123,123,255,23456,34567,102345,0.23456,2.3e-07,hello\n",
	},
}

func TestWriteStruct(t *testing.T) {
	for n, tt := range tests {
		b := &bytes.Buffer{}
		f := NewWriter(b)
		err := f.WriteStructAll(tt.Struct)
		if err != nil {
			t.Error(err)
		}

		out := b.String()
		if out != tt.CSV {
			t.Errorf("#%d: out:%q want %q", n, out, tt.CSV)
		}
	}
}

func TestWriteStructHeader(t *testing.T) {
	b := &bytes.Buffer{}
	f := NewWriter(b)
	s := testStruct{}
	err := f.WriteStructHeader(s)
	if err != nil {
		t.Error(err)
	}
	f.Flush()

	out := b.String()
	expect := "Bool,Int,Int8,Int16,Int32,Int64,Uint,Uint8,Uint16,Uint32,Uint64,Float32,Float64,String\n"
	if out != expect {
		t.Errorf("header: out:%q want %q", out, expect)
	}
}
