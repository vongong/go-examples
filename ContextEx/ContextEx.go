package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()

	//ctx, cancel := context.WithCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//v1
	// go func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()

	//v2 - same as v1
	//time.AfterFunc(time.Second, cancel)

	sleepAndTalk(ctx, 3*time.Second, "Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {

	//v1
	//time.Sleep(d)
	//fmt.Println(s)

	//v2
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
	}
}
