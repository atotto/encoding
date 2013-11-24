// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriter() {

	type Person struct {
		Name   string
		Age    int
		Height float32
		Weight float64
	}

	member := []Person{
		{"Gopher", 2, 12.3, 34.5},
		{"Ghost", 4, 12.3, 0.000034},
	}

	w := NewWriter(os.Stdout)
	w.WriteStructAll(member)

	// Output:
	// Gopher,2,12.3,34.5
	// Ghost,4,12.3,3.4e-05
}

func ExampleReader() {

	csvtext := "Gopher,2,12.3,34.5\nGhost,4,12.3,3.4e-05"

	type Person struct {
		Name   string
		Age    int
		Height float32
		Weight float64
	}

	var person = []Person{}

	w := NewReader(strings.NewReader(csvtext))
	err := w.ReadStructAll(&person)
	if err != nil {
		return
	}

	fmt.Printf("%v", person)

	// Output:
	// [{Gopher 2 12.3 34.5} {Ghost 4 12.3 3.4e-05}]
}
