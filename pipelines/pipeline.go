package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 4)
	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)
	fmt.Println(<-out)
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range nums {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}
