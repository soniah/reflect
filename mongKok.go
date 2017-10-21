package main

import "fmt"

type guitaristT struct {
	surname  string
	year     int
	american bool
	rating   float32
	styles   []string
}

func mongKok() {
	// example parameters taken out of an environment variable

	jimi := "surname=Hendrix|year=1942|american=true|rating=10.0|style=blues|style=rock|style=psychedelic"
	fmt.Println(jimi)

	//sonia := "surname=Hamilton|rating=0.9|style=blues|style=reggae"

}
