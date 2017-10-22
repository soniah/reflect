package main

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

// Portuguese Tarts - Mong Kok East, Kennedy Town, Cathay City.
func types2() {
	fmt.Printf("\n%s\n\n", "-- portugueseTart() -- more Types/Kinds --")
	w := tabwriter.NewWriter(os.Stdout, 40, 8, 0, ' ', 0)

	// inner struct
	type siT struct {
		i   int
		psi *siT
	}

	// outer struct
	type soT struct {
		f  float64
		s  string
		ai []int
		si *siT
	}

	si := siT{i: 9}
	so := soT{6.78, "xyzzy", []int{1, 2}, &si}

	fmt.Fprintf(w, "reflect.TypeOf(so)\t%s\n", reflect.TypeOf(so))
	fmt.Fprintf(w, "reflect.TypeOf(si)\t%s\n", reflect.TypeOf(si))

	fmt.Fprintf(w, "reflect.ValueOf(so)\t%v\n", reflect.ValueOf(so))
	fmt.Fprintf(w, "reflect.ValueOf(so).String()\t%s\n", reflect.ValueOf(so).String())

	fmt.Fprintf(w, "reflect.ValueOf(si)\t%v\n", reflect.ValueOf(si))
	fmt.Fprintf(w, "reflect.ValueOf(si).String()\t%s\n", reflect.ValueOf(si).String())

	fmt.Fprintf(w, "reflect.ValueOf(so).Type()\t%s\n", reflect.ValueOf(so).Type())
	fmt.Fprintf(w, "reflect.ValueOf(so).Kind()\t%s\n", reflect.ValueOf(so).Kind()) // struct

	fmt.Fprintf(w, "reflect.ValueOf(so.ai).Type()\t%s\n", reflect.ValueOf(so.ai).Type()) // []int
	fmt.Fprintf(w, "reflect.ValueOf(so.ai).Kind()\t%s\n", reflect.ValueOf(so.ai).Kind()) // slice

	fmt.Fprintf(w, "reflect.ValueOf(so.si).Type()\t%s\n", reflect.ValueOf(so.si).Type()) // *main.siT
	fmt.Fprintf(w, "reflect.ValueOf(so.si).Kind()\t%s\n", reflect.ValueOf(so.si).Kind()) // ptr

	fmt.Fprintf(w, "reflect.ValueOf(*so.si).Type()\t%s\n", reflect.ValueOf(*so.si).Type()) // main.siT
	fmt.Fprintf(w, "reflect.ValueOf(*so.si).Kind()\t%s\n", reflect.ValueOf(*so.si).Kind()) // struct

	w.Flush()
}
