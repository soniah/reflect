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
		log.Fatalln("parameters input is malformed", allParameters)
	}

	result := guitaristT{}

	// If we want to modify x by reflection, we must give the reflection library
	// a pointer to the value we want to modify. Think of passing x to a
	// function:
	//
	// f(x)
	//
	// 	We would not expect f to be able to modify x because we passed a copy of
	// 	x's value, not x itself. If we want f to modify x directly we must pass
	// 	our function the address of x (that is, a pointer to x):
	//
	// f(&x)

	p := reflect.ValueOf(&result)

	// The reflection object p isn't settable, but it's not p we want to set,
	// it's (in effect) *p. To get to what p points to, we call the Elem method
	// of Value, which indirects through the pointer, and save the result in a
	// reflection Value called v:

	// Elem returns the value that the interface v contains or that the pointer
	// v points to. It panics if v's Kind is not Interface or Ptr. It returns
	// the zero Value if v is nil.

	v := p.Elem()

	// usually:
	// v := reflect.ValueOf(&result).Elem()

	for _, parameter := range parameters {
		kv := strings.Split(parameter, "=")
		if len(kv) != 2 {
			log.Fatalln("malformed parameter", parameter)
		}
		key := kv[0]
		value := kv[1]
		field := v.FieldByName(key)

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
			// slice fields
			fmt.Println("key type **UNKNOWN**:", key)
		}
	}

	return result
}
