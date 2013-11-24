// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"encoding/csv"

	"io"
	"reflect"
	"strconv"
)

// A Reader reads records from a CSV-encoded file.
type Reader struct {
	*csv.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		csv.NewReader(r),
	}
}

func (r *Reader) ReadStruct(v interface{}) (err error) {
	record, err := r.Read()
	if err != nil {
		return
	}

	rv := reflect.ValueOf(v).Elem()
	// TODO: error

	for s := 0; s < rv.NumField(); s++ {
		val := rv.Field(s)
		x := record[s]

		setValue(&val, x)
	}

	return
}

func (r *Reader) ReadStructAll(v interface{}) (err error) {

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem().Elem()
	// TODO: error

	records := reflect.MakeSlice(reflect.TypeOf(v).Elem(), 0, 0)

	for {
		record := reflect.New(rt)
		err = r.ReadStruct(record.Interface())
		if err == io.EOF {
			rv.Set(records)
			return nil
		}
		if err != nil {
			return err
		}
		records = reflect.Append(records, record.Elem())
	}
}

func setValue(v *reflect.Value, x string) (err error) {
	switch v.Kind() {
	case reflect.Bool:
		val, err := strconv.ParseBool(x)
		if err != nil {
			return err
		}
		v.SetBool(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(x, 10, v.Type().Bits())
		if err != nil {
			return err
		}
		v.SetInt(val)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		val, err := strconv.ParseUint(x, 10, v.Type().Bits())
		if err != nil {
			return err
		}
		v.SetUint(val)

	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(x, v.Type().Bits())
		if err != nil {
			return err
		}
		v.SetFloat(val)

	case reflect.String:
		v.SetString(x)
	case reflect.Struct:
	case reflect.Map:
	case reflect.Slice:
	case reflect.Array:
	case reflect.Interface, reflect.Ptr:
	default:
		return &UnsupportedTypeError{v.Type()}
	}
	return
}
