// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"encoding/csv"
	"time"

	"fmt"
	"io"
	"reflect"
)

// A Writer writes records to a CSV encoded file.
//
// see encoding/csv package.
type Writer struct {
	*csv.Writer
	timeFormat string
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer:     csv.NewWriter(w),
		timeFormat: "2006-01-02 15:04:05",
	}
}

func (w *Writer) SetTimeFormat(format string) {
	w.timeFormat = format
}

// An UnsupportedTypeError is returned by Writer when attempting
// to encode an unsupported value type.
type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "csv: unsupported type: " + e.Type.String()
}

// WriteStruct writes a single CSV record to w along with any necessary quoting.
// A record is a struct of flat structure.
func (w *Writer) WriteStruct(v interface{}) (err error) {

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		panic("must be struct")
	}
	data := make([]string, rv.NumField())

	for s := 0; s < rv.NumField(); s++ {
		val := rv.Field(s)

		str, err := w.reflectValue(val)
		if err != nil {
			return err
		}
		data[s] = str
		//fmt.Sprintf("%v", val.Interface())
	}
	err = w.Write(data)
	return
}

// WriteStructAll write
func (w *Writer) WriteStructAll(v interface{}) (err error) {

	r := reflect.ValueOf(v)

	if r.Kind() != reflect.Slice {
		panic("must be slice")
	}

	for i := 0; i < r.Len(); i++ {
		rv := r.Index(i)
		err = w.WriteStruct(rv.Interface())
		if err != nil {
			return
		}
	}
	w.Flush()
	return
}

//
func (w *Writer) WriteStructHeader(v interface{}) (err error) {

	rt := reflect.TypeOf(v)
	if rt.Kind() != reflect.Struct {
		panic("must be struct")
	}
	data := make([]string, rt.NumField())

	for s := 0; s < rt.NumField(); s++ {
		t := rt.Field(s)

		data[s] = fmt.Sprintf("%s", t.Name)
	}
	err = w.Write(data)
	return
}

func (w *Writer) reflectValue(v reflect.Value) (str string, err error) {
	str = ""
	switch v.Kind() {
	case reflect.Bool:
		str = fmt.Sprintf("%v", v.Interface())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = fmt.Sprintf("%v", v.Interface())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		str = fmt.Sprintf("%v", v.Interface())
	case reflect.Float32, reflect.Float64:
		str = fmt.Sprintf("%v", v.Interface())
	case reflect.String:
		str = fmt.Sprintf("%v", v.Interface())
	case reflect.Struct:
		return w.structValue(v)
	case reflect.Map:
	case reflect.Slice:
	case reflect.Array:
	case reflect.Interface, reflect.Ptr:
	default:
		return "", &UnsupportedTypeError{v.Type()}
	}
	return
}

func (w *Writer) structValue(v reflect.Value) (str string, err error) {
	switch v.Type().String() {
	case "time.Time":
		return time.Time.Format(v.Interface().(time.Time), w.timeFormat), nil
	}
	return "", nil
}
