// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"encoding/csv"

	"fmt"
	"io"
	"reflect"
)

// A Writer writes records to a CSV encoded file.
//
// see encoding/csv package.
type Writer struct {
	*csv.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		csv.NewWriter(w),
	}
}

// WriteStruct writes a single CSV record to w along with any necessary quoting.
// A record is a struct of flat structure.
func (w *Writer) WriteStruct(v interface{}) {

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		panic("must be struct")
	}
	data := make([]string, rv.NumField())

	for s := 0; s < rv.NumField(); s++ {
		val := rv.Field(s)

		data[s] = fmt.Sprintf("%v", val.Interface())
	}
	w.Write(data)
}

// WriteStructAll write
func (w *Writer) WriteStructAll(v interface{}) {

	r := reflect.ValueOf(v)

	if r.Kind() != reflect.Slice {
		panic("must be slice")
	}

	for i := 0; i < r.Len(); i++ {
		rv := r.Index(i)
		w.WriteStruct(rv.Interface())
	}
	w.Flush()
}

//
func (w *Writer) WriteStructHeader(v interface{}) {

	rt := reflect.TypeOf(v)
	if rt.Kind() != reflect.Struct {
		panic("must be struct")
	}
	data := make([]string, rt.NumField())

	for s := 0; s < rt.NumField(); s++ {
		t := rt.Field(s)

		data[s] = fmt.Sprintf("%v", t.Name)
	}
	w.Write(data)
}

func reflectValue2string(v *reflect.Value) string {
	switch v.Type {
	}
	return ""
}
