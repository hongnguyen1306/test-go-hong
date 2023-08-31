package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

type point struct {
	x, y int
}

func main() {
	p("Containts: ", s.Contains("String", "Str"))
	p("Count:     ", s.Count("Stringiiii", "i"))
	p("HasPrefix: ", s.HasPrefix("String", "Str"))
	p("Index:     ", s.Index("String", "n"))
	p("Join:	  ", s.Join([]string{"String", "is"}, " "))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("Changegg", "g", "d", -1))
	p("Replace:   ", s.Replace("Change", "g", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	s := point{1, 2}

	fmt.Printf("struct1: %v\n", s)  // struct1: {1 2}
	fmt.Printf("struct2: %+v\n", s) // struct2: {x:1 y:2}
	fmt.Printf("struct3: %#v\n", s) // struct3: main.point{x:1, y:2}

	fmt.Printf("type: %T\n", s)    // type: main.point
	fmt.Printf("bool: %t\n", true) // bool: true
	fmt.Printf("int: %d\n", 123)   // int: 123

}
