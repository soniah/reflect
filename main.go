package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kr/pretty"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	// interfaces()
	// types()
	// types2()
	// xml()
	// typeSwitches()
	structs()
}

func catch(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func pp(description string, stuff ...interface{}) {
	for _, s := range stuff {
		fmt.Print("\"%s\"\n%+v\n", description, pretty.Formatter(s))
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	catch(err)
	return i
}
