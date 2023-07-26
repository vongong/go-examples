package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func splitX(s string, x int) ([]string, error) {
	if x < 1 {
		return nil, fmt.Errorf("x needs to larger then 1")
	}
	sLen := len(s)
	if sLen <= x {
		return nil, fmt.Errorf("string length is less then x")
	}
	xDiv := sLen / x
	xMod := sLen % x
	size := xDiv
	if xMod > 0 {
		size++
	}

	arr := make([]string, size)
	for i := 0; i < xDiv; i++ {
		arr[i] = s[i*x : i*x+x]
	}
	if xMod > 0 {
		arr[xDiv] = s[sLen-xMod:]
	}
	return arr, nil
}

func strToInts(s string) []int {
	ints := make([]int, 0)
	if len(s) > 0 {
		codes := strings.Split(s, ",")
		for _, cd := range codes {
			cd = strings.Replace(cd, " ", "", -1) // just in case
			v, err := strconv.Atoi(cd)
			if err != nil {
				log.Printf("Error converting to integers: %s", err.Error())
				//return nil, err
			} else {
				ints = append(ints, v)
			}
		}
	}
	return ints
}

func IntsToStr(arr []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")

}

func main() {
	//fmt.Println(splitX("abcdefg", 4))
	fmt.Println(strToInts("1,2,3"))
	fmt.Println(strToInts("10,a,12"))
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(IntsToStr(A, ","))

}
