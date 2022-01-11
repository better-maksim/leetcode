package main

import (
	"errors"
	"fmt"
	"gitbuh.com/better-maksim/leetcode/array_stack"
)

type MinStack struct {
	StackData *array_stack.ArrayStack
	MinStack  *array_stack.ArrayStack
}

func NewMinStack() *MinStack {
	return &MinStack{
		StackData: &array_stack.ArrayStack{},
		MinStack:  &array_stack.ArrayStack{},
	}
}

func (stack *MinStack) Push(num int) error {
	if stack.MinStack.IsEmpty() {
		stack.MinStack.Push(num)
	} else {
		currentMin, err := stack.GetMin()
		if err != nil {
			return err
		}
		if num <= currentMin {
			stack.MinStack.Push(num)
		}

	}
	stack.StackData.Push(num)
	return nil
}

func (stack *MinStack) Pop() (int, error) {

	if stack.StackData.IsEmpty() {
		return 0, errors.New("栈数据为空")
	}
	value, err := stack.StackData.Pop()
	if err != nil {
		return 0, err
	}

	min, err := stack.GetMin()

	if err != nil {
		return 0, err
	}

	if value == min {
		_, _ = stack.MinStack.Pop()
	}
	return value, err
}

func (stack *MinStack) GetMin() (int, error) {
	if stack.MinStack.IsEmpty() {
		return 0, errors.New("站数据为空")
	}
	return stack.MinStack.Peek(), nil
}
func main() {
	stack := NewMinStack()

	stack.Push(3)
	fmt.Println(stack.GetMin())
	stack.Push(4)
	fmt.Println(stack.GetMin())
	stack.Push(5)
	fmt.Println(stack.GetMin())
	stack.Push(2)
	fmt.Println(stack.GetMin())
	stack.Push(1)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
}
