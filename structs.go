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
	Cities   map[string]int
}

// Mong Kok, Prince Edward, Sham Shui Po, Cheung Sha Wan.
func structs() {
	fmt.Printf("\n%s\n", "-- structs --")

	// example parameters taken out of an environment variable

	jimiEnvvar := "surname=Hendrix|year=1942|american=true|rating=9.99|styles=blues|styles=rock|styles=psychedelic"
	jimiEnvvar += "|cities=New York^17|cities=Los Angeles^14|cities=London^11|cities=Bay Area^9"

	jimiStruct := fillStruct(jimiEnvvar)
	fmt.Printf("\n%# v", pretty.Formatter(jimiStruct))

}

func fillStruct(allParameters string) guitaristT {

	result := guitaristT{}
	result.Styles = make([]string, 0) // other values will put empty elements at start
	result.Cities = make(map[string]int)

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
		keyAsString := strings.Title(kv[0])
		valueAsString := kv[1]
		fieldAsValue := v.FieldByName(keyAsString)

		switch fieldAsValue.Kind() {

		case reflect.String:
			fmt.Println("keyAsString:", keyAsString, "is String, has value:", valueAsString)
			fieldAsValue.SetString(valueAsString)

		case reflect.Int64:
			fmt.Println("keyAsString:", keyAsString, "is Int64, has value:", valueAsString)
			i, err := strconv.ParseInt(valueAsString, 10, 64)
			catch(err)
			fieldAsValue.SetInt(i)

		case reflect.Bool:
			fmt.Println("keyAsString:", keyAsString, "is Bool, has value:", valueAsString)
			b, err := strconv.ParseBool(valueAsString)
			catch(err)
			fieldAsValue.SetBool(b)

		case reflect.Float32:
			fmt.Println("keyAsString:", keyAsString, "is Float32, has value:", valueAsString)
			f, err := strconv.ParseFloat(valueAsString, 32)
			catch(err)
			fieldAsValue.SetFloat(f)

		case reflect.Slice:
			fmt.Println("keyAsString:", keyAsString, "is Slice, has value:", valueAsString)
			valueAsValue := reflect.ValueOf(valueAsString)
			fieldAsValue.Set(reflect.Append(fieldAsValue, valueAsValue))

		case reflect.Map:
			fmt.Println("keyAsString", keyAsString, "is Map, has value:", valueAsString)
			mapKV := strings.Split(valueAsString, "^")
			if len(mapKV) != 2 {
				log.Fatalln("malformed map key/value:", mapKV)
			}
			mapK := mapKV[0]
			mapV := mapKV[1]
			thisMap := fieldAsValue.Interface().(map[string]int)
			thisMap[mapK] = atoi(mapV)
			thisMapAsValue := reflect.ValueOf(thisMap)
			fieldAsValue.Set(thisMapAsValue)

		default:
			fmt.Println("XXX: keyAsString:", keyAsString, "is", fieldAsValue.Kind(), "has value:", valueAsString)
		}
		fmt.Println()
	}
	return result
}
