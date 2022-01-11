package main

import (
	"errors"
	"fmt"
	"gitbuh.com/better-maksim/leetcode/array_stack"
)

type StackQueue struct {
	PopStack *array_stack.ArrayStack
	popStack *array_stack.ArrayStack
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		PopStack: &array_stack.ArrayStack{},
		popStack: &array_stack.ArrayStack{},
	}
}

func (queue *StackQueue) Add(num int) {
	queue.PopStack.Push(num)
}

func (queue *StackQueue) Poll() (int, error) {
	if queue.PopStack.IsEmpty() && queue.popStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	if queue.popStack.IsEmpty() {
		for {
			value, err := queue.PopStack.Pop()
			if err != nil {
				break
			}
			queue.popStack.Push(value)
		}
	}
	return queue.popStack.Pop()
}

func (queue *StackQueue) Peek() (int, error) {
	if queue.PopStack.IsEmpty() && queue.popStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	if queue.popStack.IsEmpty() {
		for {
			value, err := queue.PopStack.Pop()
			if err != nil {
				break
			}
			queue.popStack.Push(value)
		}
	}
	return queue.popStack.Peek(), nil
}

func main() {
	queue := NewStackQueue()

	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Add(4)
	queue.Add(5)

	fmt.Println(queue.Peek())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Peek())
}
