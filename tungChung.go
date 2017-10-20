package main

import "fmt"

/* Nothing new or interesting here, just like Tung Chung railway station.
   - https://tour.golang.org/methods/16
   - https://golang.org/doc/effective_go.html#type_switch
*/

// The Marriott Hotel - several months trapped on Prison Island.  Airport
// barbed wire on 3 sides and building sites on the 4th, the sea so close yet
// so far.
func prisonIsland(t interface{}) {
	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean:\t\t%t\n", t) // t has type bool
	case int:
		fmt.Printf("integer:\t\t%d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean:\t%t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer:\t%d\n", *t) // t has type *int
	default:
		fmt.Printf("unexpected type:\t%T\n", t) // %T prints whatever type t has
	}
}

func tungChung() {
	fmt.Printf("\n%s\n\n", "-- tungChung() -- type switches --")
	fmt.Println("Note that t returns the concrete value of `t interface{}`")
	fmt.Println()

	s := "abc"
	i := 26
	b := true

	prisonIsland(i)
	prisonIsland(&i)
	prisonIsland(b)
	prisonIsland(&b)
	prisonIsland(s)
}
