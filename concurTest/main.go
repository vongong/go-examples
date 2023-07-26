package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu    sync.RWMutex
	ilist []int
}

func NewContainer() Container {
	return Container{
		ilist: make([]int, 0),
	}
}

func (c *Container) append(name ...int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ilist = append(c.ilist, name...)
	// t := rand.Intn(9)
	// time.Sleep(time.Duration(t) * time.Second)
}

func (c *Container) Print() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println(c.ilist)
}
func main() {
	c := NewContainer()
	var wg sync.WaitGroup
	maxGoroutines := 3
	guard := make(chan struct{}, maxGoroutines)

	for i := 1; i < 10; i++ {
		wg.Add(1)
		guard <- struct{}{} // would block if guard channel is already filled
		go func(in int) {
			defer wg.Done()
			c.append(in)
			<-guard
		}(i)
	}
	wg.Wait()
	c.Print()
}
