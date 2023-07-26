package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
)

// JSONEq returns true if s1 and s2 are equal
func JSONEq(expected, actual []byte) (bool, error) {

	var o1, o2 interface{}

	err := json.Unmarshal(expected, &o1)
	if err != nil {
		return false, errors.New("expected parameter value is not valid json")
	}
	err = json.Unmarshal(actual, &o2)
	if err != nil {
		return false, errors.New("actual parameter value is not valid json")
	}

	return reflect.DeepEqual(o1, o2), nil

}

func main() {

	s1 := []byte(`{"x":10, "y":42}`)
	s2 := []byte(`{"y":42, "x":10}`)
	fmt.Println("s1: ", string(s1))
	fmt.Println("s2: ", string(s2))

	eq, err := JSONEq(s1, s2)
	if err != nil {
		fmt.Println("error occurred: ", err)
	}
	fmt.Println("jsoneq s1, s2:", eq)

	eq2 := cmp.Diff(s1, s2)
	fmt.Printf("cmp equal s1, s2: %s", eq2)
}
