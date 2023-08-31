package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // Co key: value
	m["one"] = 1
	m["two"] = 2
	fmt.Println("m ", m)
	fmt.Println("len m ", len(m))

	v1 := m["one"]
	// v2 := m["two"]
	println(v1)
	delete(m, "one")
	fmt.Println("m ", m)

	clear(m)
	fmt.Println("m ", m)

	m["one"] = 1
	m["two"] = 2
	_, prs := m["two"] // isExist ? => prs:  true
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n: ", n)

	value, ok := n["bar"]
	fmt.Println("value, ok: ", value, ok) // check key exist => value, ok:  2 true

	// Map la kieu tham chieu
	newMap := n
	newMap["bar"] = 100
	fmt.Println("newMap: ", newMap, "originalMap n: ", n) // newMap:  map[bar:100 foo:1] originalMap n:  map[bar:100 foo:1]

	// So sanh 2 map
	fmt.Println(maps.Equal(n, newMap)) // true
}
