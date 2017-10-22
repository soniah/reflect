package main

import (
	"fmt"
	"reflect"
)

// Iced Lemon Tea, Char Siu, Mong Kok.
func types() {
	fmt.Printf("\n%s\n", "-- types --")

	var inter interface{}
	var i int = 3
	type iT int
	var it iT = 4

	fmt.Println(`var inter interface{}
var i int = 3
type iT int
var it iT = 4`)

	fmt.Println()
	fmt.Println("reflect.TypeOf(inter):", reflect.TypeOf(inter)) // <nil>
	fmt.Println("reflect.TypeOf(i)    :", reflect.TypeOf(i))     // int
	fmt.Println("reflect.TypeOf(it)   :", reflect.TypeOf(it))    // main.iT

	fmt.Println()
	// ValueOf without String() shows the concrete value inside
	fmt.Println("reflect.ValueOf(inter):", reflect.ValueOf(inter)) // <invalid reflect.Value>
	fmt.Println("reflect.ValueOf(i)    :", reflect.ValueOf(i))     // 3
	fmt.Println("reflect.ValueOf(it)   :", reflect.ValueOf(it))    // 4

	fmt.Println()
	fmt.Println("reflect.ValueOf(inter).String()):", reflect.ValueOf(inter).String()) // <invalid Value>
	fmt.Println("reflect.ValueOf(i).String())    :", reflect.ValueOf(i).String())     // <int Value>
	fmt.Println("reflect.ValueOf(it).String())   :", reflect.ValueOf(it).String())    // <main.iT Value>

	fmt.Println()
	// fmt.Println("reflect.ValueOf(inter).Type():", reflect.ValueOf(inter).Type()) // panic: reflect: call of reflect.Value.Type on zero Value
	fmt.Println("reflect.ValueOf(inter).Kind():", reflect.ValueOf(inter).Kind()) // reflect.ValueOf(inter).Kind(): invalid
	// fmt.Println("reflect.ValueOf(inter).Int():", reflect.ValueOf(inter).Int())   // panic: reflect: call of reflect.Value.Int on zero Value

	fmt.Println()
	fmt.Println("reflect.ValueOf(i).Type():", reflect.ValueOf(i).Type()) // int
	fmt.Println("reflect.ValueOf(i).Kind():", reflect.ValueOf(i).Kind()) // int
	fmt.Println("reflect.ValueOf(i).Int() :", reflect.ValueOf(i).Int())  // 3

	fmt.Println()
	fmt.Println("reflect.ValueOf(it).Type():", reflect.ValueOf(it).Type()) // main.iT
	fmt.Println("reflect.ValueOf(it).Kind():", reflect.ValueOf(it).Kind()) // int
	fmt.Println("reflect.ValueOf(it).Int() :", reflect.ValueOf(it).Int())  // 4
}
