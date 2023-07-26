package main

import (
	"fmt"
	"time"
)

func RetTimePtr(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func main() {
	inStr := "2016-03-11 12:19:48.490"

	fmt.Println("time convert from:", inStr)
	fmt.Println("to:Mar 3, 2016 12:19:48 PM")

	inLayout := "2006-01-02 15:04:05.000"
	outLayout := "Jan 01, 2006 15:04:05 PM"
	t, err := time.Parse(inLayout, inStr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Format(outLayout))
}
