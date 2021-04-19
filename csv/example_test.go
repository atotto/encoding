// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv_test

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/encoding/csv"
)

func ExampleWriter() {

	type Person struct {
		Name       string
		Age        int
		Height     float32
		Weight     float64
		CreateTime time.Time
	}

	member := []Person{
		{"Gopher", 2, 12.3, 34.5, time.Date(2020, 1, 1, 12, 20, 2, 0, time.UTC)},
		{"Ghost", 4, 12.3, 0.000034, time.Date(2020, 2, 3, 15, 20, 2, 0, time.UTC)},
	}

	w := csv.NewWriter(os.Stdout)
	w.SetTimeLocation(time.UTC)
	w.WriteStructAll(member)

	// Output:
	// Gopher,2,12.3,34.5,2020-01-01 12:20:02
	// Ghost,4,12.3,3.4e-05,2020-02-03 15:20:02
}

func ExampleReader() {

	csvtext := "Gopher,2,12.3,34.5,2020/01/01 12:20:02\nGhost,4,12.3,3.4e-05,2020/02/03 15:20:02"

	type Person struct {
		Name       string
		Age        int
		Height     float32
		Weight     float64
		CreateTime time.Time
	}

	var person = []Person{}

	r := csv.NewReader(strings.NewReader(csvtext))
	r.SetTimeFormat("2006/01/02 15:04:05")
	r.SetTimeLocation(time.UTC)
	err := r.ReadStructAll(&person)
	if err != nil {
		return
	}

	fmt.Printf("%v", person)

	// Output:
	// [{Gopher 2 12.3 34.5 2020-01-01 12:20:02 +0000 UTC} {Ghost 4 12.3 3.4e-05 2020-02-03 15:20:02 +0000 UTC}]
}
