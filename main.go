package main

import (
	"log"

	"strconv"

	"github.com/kr/pretty"
)

func catch(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func pp(description string, stuff ...interface{}) {
	for _, s := range stuff {
		log.Print("\"%s\"\n%+v\n", description, pretty.Formatter(s))
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	catch(err)
	return i
}

func main() {
	// dingding()
	// icedLemonTea()
	// portugueseTart()
	sinoCentre()
}
