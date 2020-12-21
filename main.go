package main

import "fmt"

type figure struct {
	width, height int
}

var (
	m = map[string]figure{
		"Square": {
			width: 20, height: 50,
		},
		"Triangle": {
			width: 15, height: 20,
		},
	}
	v  figure
	ok bool
)

func main() {
	fmt.Println(m["Square"])
	fmt.Println(m)

	m["Rectangle"] = figure{width: 20, height: 40}
	fmt.Println(m)

	m["Square"] = figure{width: 0, height: 0}
	fmt.Println(m)

	delete(m, "Triangle")
	fmt.Println(m)

	v, ok = m["Circle"]
	fmt.Println(v, ok)

	v, ok = m["Rectangle"]
	fmt.Println(v, ok)
}
