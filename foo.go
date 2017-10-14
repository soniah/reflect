package main

import "fmt"

func foo() {
	fmt.Printf("\n%s\n\n", "-- foo() -- TODO --")

	// TODO

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
		si siT
	}

}
