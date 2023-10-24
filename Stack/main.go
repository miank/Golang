package main

import "fmt"

var stackSize int

type Stack struct {
	items []interface{}
	top   int
}

func (s *Stack) Push(value interface{}) {
	if s.top == stackSize-1 {
		fmt.Println("Stack Overflow, cannot insert more values")
		return
	}

	s.items = append(s.items, value)
	s.top++
	fmt.Printf("Inserted Value %v\n", value)
	return
}

func (s *Stack) Pop() interface{} {
	if s.top == -1 {
		fmt.Println("Stack underflow, No values to remove")
		return nil
	}

	poppedElement := s.items[s.top]
	s.items = s.items[:s.top]
	s.top--
	fmt.Printf("Popped element is %d \n", poppedElement)
	fmt.Printf("Current stack is %d \n", s)
	return poppedElement
}

func (s *Stack) IsEmpty() {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Stack is not empty")
}

func (s *Stack) IsFull() {
	if s.top == stackSize-1 {
		fmt.Println("Stack is full")
		return
	}

	fmt.Println("Stack is not full")
}

func (s *Stack) Peek() interface{} {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return -1
	}

	fmt.Printf("Removed value %v\n", s.items[s.top])
	element := s.items[s.top]
	return element
}

func (s *Stack) Display() {
	for _, v := range s.items {
		fmt.Println(v)
	}
}

func main() {
	s := Stack{
		items: []interface{}{1},
		top:   0,
	}

	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Pop()
	s.Display()
	s.Peek()
	s.Pop()
	s.Display()

}
