package main

import (
	"fmt"
)

type MemberInfo struct {
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

func main() {
	// * Test 1
	//var m map[string]string
	m := make(map[string]string)
	v, ok := m["1234"]
	if !ok {
		v = "apple"
		m["1234"] = v
	}

	v2, ok := m["1234"]

	if ok {
		fmt.Println("1234:", v2)
	} else {
		fmt.Println("ERROR:")
	}

	// * Test 2
	m2 := make(map[string]MemberInfo)
	data := MemberInfo{
		FirstName:  "John",
		MiddleName: "T",
		LastName:   "Wick",
	}
	m2["abc"] = data

	var data2 MemberInfo
	m2["bbb"] = data2

	v3, ok := m2["abc"]

	if ok {
		fmt.Println("m2[abc]:", v3)
	} else {
		fmt.Println("ERROR:")
	}

	v4, ok := m2["bbb"]
	if v4 == (MemberInfo{}) {
		fmt.Println("m2[bbb] is empty")
	}

}
