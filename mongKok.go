package main

import (
	"fmt"
	"log"
	"strings"

	"reflect"

	"github.com/kr/pretty"
)

type guitaristT struct {
	surname  string   `required=true`
	year     int      `required=false`
	american bool     // example of missing field
	rating   float32  `required=true`
	styles   []string `required=true,minsize=1`
}

// Mong Kok, Prince Edward, Sham Shui Po, Cheung Sha Wan
func mongKok() {
	// example parameters taken out of an environment variable

	jimiEnvvar := "surname=Hendrix|year=1942|american=true|rating=10.0|style=blues|style=rock|style=psychedelic"
	jimiStruct := fillStruct(jimiEnvvar)
	fmt.Printf("\n\n%# v", pretty.Formatter(jimiStruct))

	//sonia := "surname=Hamilton|rating=0.9|style=blues|style=reggae"

}

func fillStruct(allParameters string) guitaristT {
	fmt.Println()
	parameters := strings.Split(allParameters, "|")
	if len(parameters) == 0 {
		log.Fatalln("parameters input is  malformed", allParameters)
	}

	result := guitaristT{}

	// working this way means I'll always be working with a copy. I think.
	// gv := reflect.ValueOf(result) // since haven't taken address of result, don't need Elem()

	// Working with a pointer to the struct means I can modify it. I think.

	// If we want to modify x by reflection, we must give the reflection library
	// a pointer to the value we want to modify.  Think of passing x to a
	// function:
	//
	// f(x)
	//
	// 	We would not expect f to be able to modify x because we passed a copy of
	// 	x's value, not x itself. If we want f to modify x directly we must pass
	// 	our function the address of x (that is, a pointer to x):
	//
	// f(&x)
	//
	// Elem returns the value that the interface v contains or that the pointer
	// v points to. It panics if v's Kind is not Interface or Ptr. It returns
	// the zero Value if v is nil.

	gv := reflect.ValueOf(&result).Elem()

	for _, parameter := range parameters {
		kv := strings.Split(parameter, "=")
		if len(kv) != 2 {
			log.Fatalln("malformed parameter", parameter)
		}
		key := kv[0]
		value := kv[1]
		field := gv.FieldByName(key)

		fmt.Printf("key: %s,\tvalue: %s,\tkind: %s\n", key, value, field.Kind())
		/*
			key: surname,   value: Hendrix, kind: string
			key: year,      value: 1942,    kind: int
			key: american,  value: true,    kind: bool
			key: rating,    value: 10.0,    kind: float32
			key: style,     value: blues,   kind: invalid     // notice array fields are invalid
			key: style,     value: rock,    kind: invalid
			key: style,     value: psychedelic,     kind: invalid
		*/

		switch field.Kind() {
		case reflect.String:
			fmt.Println("key is String:", key)
		case reflect.Int:
			fmt.Println("key is Int:", key)
		case reflect.Bool:
			fmt.Println("key is Bool:", key)
		case reflect.Float32:
			fmt.Println("key is Float32:", key)
		default:
			fmt.Println("key type **UNKNOWN**:", key)
		}

		/* cannot type switch on non-interface value field (type reflect.Value)
		switch field := field.(type) {
		case bool:
			fmt.Printf("boolean:\t\t%t\n", field)
		case int:
			fmt.Printf("integer:\t\t%d\n", field)
		default:
			fmt.Printf("unexpected type:\t%T\n", field) // %T prints whatever type t has
		}
		*/
	}
	return result
}
