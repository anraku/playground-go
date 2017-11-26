package main

import "fmt"

func main() {
	var stack = &Stack{}
	value := make([]string, 4)
	stack.Value = value
	stack.Push("dataA")
	stack.Push("dataB")
	stack.Push("dataC")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push("dataD")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
