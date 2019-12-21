// LIFO and FIFO in Golang
package main

import (
	"fmt"
)

type Node struct {
	Value string
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

//print stack
func (s *Stack) String() string {
	return fmt.Sprint(s.nodes)
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func main() {
	s := NewStack()
	s.Push(&Node{"A"})
	s.Push(&Node{"B"})
	s.Push(&Node{"C"})
	s.Push(&Node{"D"})
	fmt.Println(s)
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())

}
