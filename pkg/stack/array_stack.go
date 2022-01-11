package array_stack

import (
	"errors"
	"sync"
)

type ArrayStack struct {
	array []int //底层切片
	size  int
	lock  sync.Mutex
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

// Push 入栈
func (stack *ArrayStack) Push(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)
	// 栈中元素数量+1
	stack.size = stack.size + 1
}

//Pop 出站
func (stack *ArrayStack) Pop() (int, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 栈中元素已空
	if stack.size == 0 {
		return 0, errors.New("stack empty")
	}
	// 栈顶元素
	v := stack.array[stack.size-1]

	// 切片收缩，但可能占用空间越来越大
	//stack.array = stack.array[0 : stack.size-1]

	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newArray := make([]int, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v, nil
}

func (stack *ArrayStack) IsEmpty() bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size == 0
}

// Peek 获取栈顶元素
func (stack *ArrayStack) Peek() int {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

//Size 栈大小
func (stack *ArrayStack) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	return stack.size
}

//Reverse 翻转
func Reverse(stack *ArrayStack) {

	if stack.IsEmpty() {
		return
	}

	i, err := getAndRemoveLastElement(stack)

	if err != nil {
		return
	}

	Reverse(stack)

	stack.Push(i)
}

func getAndRemoveLastElement(stack *ArrayStack) (int, error) {

	result, err := stack.Pop()

	if err != nil {
		return 0, err
	}

	if stack.IsEmpty() {
		return result, nil
	}

	last, err := getAndRemoveLastElement(stack)

	if err != nil {
		return 0, err
	}

	stack.Push(result)
	return last, nil
}
