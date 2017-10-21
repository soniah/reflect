package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
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

	jimi := "surname=Hendrix|year=1942|american=true|rating=10.0|style=blues|style=rock|style=psychedelic"
	fmt.Println(jimi)

	//sonia := "surname=Hamilton|rating=0.9|style=blues|style=reggae"

}

func (guitarist *guitaristT) fillStruct(allParameters string) {
	parameters := strings.Split(allParameters, "|")
	if len(splits) == 0 {
		log.Fatalln("parameters empty")
	}

	// Elem returns the value that the interface v contains or that the pointer
	// v points to. It panics if v's Kind is not Interface or Ptr. It returns
	// the zero Value if v is nil.
	gv := reflect.ValueOf(guitarist).Elem()

	for _, parameter := range parameters {
		kv := strings.Split("=")
		if len(kv) != 2 {
			log.Fatalln("malformed param", parameter)
		}
		key := kv[0]
		value := kv[1]
		field := gv.FieldByName(key)
	}
}
