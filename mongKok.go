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

	// Elem returns the value that the interface v contains or that the pointer
	// v points to. It panics if v's Kind is not Interface or Ptr. It returns
	// the zero Value if v is nil.
	//gv := reflect.ValueOf(result).Elem()
	gv := reflect.ValueOf(result)

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
			key: surname,	value: Hendrix,	kind: string
			key: year,	    value: 1942,	kind: int
			key: american,	value: true,	kind: bool
			key: rating,	value: 10.0,	kind: float32
			key: style,	    value: blues,	kind: invalid   // notice slice fields are invalid
			key: style,	    value: rock,	kind: invalid
			key: style,	    value: psychedelic,	kind: invalid

		*/

	}
	return result
}
