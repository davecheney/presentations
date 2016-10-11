package main

import "fmt"

type Map func(key string) string

func (m Map) Add(k, v string) Map {
	return func(key string) string {
		if k == key {
			return v
		}
		return m(key)
	}
}

func main() {
	m := Map(func(string) string { return "" })
	fmt.Println(m("george")) // ""

	m = m.Add("george", "lion")
	fmt.Println(m("george")) // "lion"
}
