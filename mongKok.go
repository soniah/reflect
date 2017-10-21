package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

type guitaristT struct {
	// TODO tag manipulation
	Surname  string   `required=true`
	Year     int64    `required=false`
	American bool     // example of missing tag
	Rating   float32  `required=true`
	Styles   []string `required=true,minsize=1`
}

// Mong Kok, Prince Edward, Sham Shui Po, Cheung Sha Wan.
func mongKok() {

	// example parameters taken out of an environment variable

	jimiEnvvar := "surname=Hendrix|year=1942|american=true|rating=9.99|styles=blues|styles=rock|styles=psychedelic"
	jimiStruct := fillStruct(jimiEnvvar)
	fmt.Printf("\n\n%# v", pretty.Formatter(jimiStruct))

}

func fillStruct(allParameters string) guitaristT {

	result := guitaristT{}
	result.Styles = make([]string, 10)

	// If we want to modify x by reflection, we must give the reflection library
	// a pointer to the value we want to modify. Think of passing x to a
	// function:
	//
	// f(x)
	//
	// We would not expect f to be able to modify x because we passed a copy of
	// x's value, not x itself. If we want f to modify x directly we must pass
	// our function the address of x (that is, a pointer to x):
	//
	// f(&x)

	p := reflect.ValueOf(&result) // p is of type reflect.Value (reflect.TypeOf)

	// The reflection object p isn't settable, but it's not p we want to set,
	// it's (in effect) *p. To get to what p points to, we call the Elem method
	// of Value, which indirects through the pointer, and save the result in a
	// reflection Value called v:

	// Elem returns the value that the interface v contains or that the pointer
	// v points to. It panics if v's Kind is not Interface or Ptr. It returns
	// the zero Value if v is nil.

	v := p.Elem() // v is also of type reflect.Value (reflect.TypeOf)

	// simpler: v := reflect.ValueOf(&result).Elem()

	fmt.Println()
	parameters := strings.Split(allParameters, "|")
	if len(parameters) == 0 {
		log.Fatalln("input parameters are malformed", allParameters)
	}

	for _, parameter := range parameters {
		kv := strings.Split(parameter, "=")
		if len(kv) != 2 {
			log.Fatalln("malformed parameter", parameter)
		}
		key := strings.Title(kv[0])
		value := kv[1]
		field := v.FieldByName(key)

		switch field.Kind() {

		case reflect.String:
			fmt.Println("key:", key, "is String, has value:", value)
			field.SetString(value)

		case reflect.Int64:
			fmt.Println("key:", key, "is Int64, has value:", value)
			i, err := strconv.ParseInt(value, 10, 64)
			catch(err)
			field.SetInt(i)

		case reflect.Bool:
			fmt.Println("key:", key, "is Bool, has value:", value)
			b, err := strconv.ParseBool(value)
			catch(err)
			field.SetBool(b)

		case reflect.Float32:
			fmt.Println("key:", key, "is Float32, has value:", value)
			f, err := strconv.ParseFloat(value, 32)
			catch(err)
			field.SetFloat(f)

		case reflect.Slice:
			fmt.Println("key:", key, "is Slice, has value:", value)
			stringValue := reflect.ValueOf(value)
			field = reflect.Append(field, stringValue) // FAILS

		default:
			fmt.Println("XXX: key:", key, "is", field.Kind(), "has value:", value)
		}
		fmt.Println()
	}
	return result
}
