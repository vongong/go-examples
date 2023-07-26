package main

import (
	"fmt"
)

func main() {
	value := func(p, q string) string {
		return p + q + "Geeks"
	}("All", "Hail")

	fmt.Println(value)
}
