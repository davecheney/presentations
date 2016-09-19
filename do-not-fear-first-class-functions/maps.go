package main

import "fmt"

type M func(key string) string

func (m M) Add(k, v string) M {
	return func(key string) string {
		if k == key {
			return v
		}
		return m(key)
	}
}

func main() {
	m := M(func(string) string { return "" })
	fmt.Println(m("george"))

	m = m.Add("george", "lion")
	fmt.Println(m("george"))
}
