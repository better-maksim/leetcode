package main

import (
	"fmt"
	"gitbuh.com/better-maksim/leetcode/array_stack"
)

func main() {

	stack := array_stack.ArrayStack{}
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	array_stack.Reverse(&stack)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
