package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	r := strings.NewReader("my request")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("b: ", b)
	fmt.Println("b-str: ", string(b))

}
