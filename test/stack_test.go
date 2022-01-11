package test

import (
	"fmt"
	"gitbuh.com/better-maksim/leetcode/array_stack"
	"sync"
	"testing"
)

func TestSize(T *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	stack := array_stack.ArrayStack{}
	go func() {
		fmt.Println(stack.Size())
		wg.Done()
	}()

	go func() {
		stack.Push("1")
		wg.Done()
	}()


	go func() {
		stack.Pop()
		wg.Done()
	}()


	wg.Wait()
}
