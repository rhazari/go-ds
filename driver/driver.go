package main

import (
	"fmt"
	"raihazar/utils"
)

func main() {
	// Get a greeting message and print it.
	stk := utils.NewStack[int]()

	arr := []int{1, 2, 3, 4, 5, 6}

	for _, value := range arr {
		stk.Push(value)
	}

	// Print the content of the stack
	stk.Print()
	// Iterate over the stack
	ok := stk.IsEmpty()
	for !ok {
		fmt.Println(stk.Pop())
		ok = stk.IsEmpty()
	}
}
