package utils

import (
	"testing"
)

func TestEmptyStack(t *testing.T) {
	stk := NewStack[int]()
	if stk.IsEmpty() == false {
		t.Fatal("stack.IsEmpty() not returning correct status")
	}
}

func TestPushAndPop(t *testing.T) {
	stk := NewStack[int]()
	arr := []int{1, 2, 3, 4, 5, 6}

	for _, value := range arr {
		stk.Push(value)
	}

	for index := len(stk.items) - 1; index >= 0; index-- {
		if v, ok := stk.Top(); ok {
			if v != arr[index] {
				t.Fatal("stack.Top() not returning the right value")
			}
		}
		if v, ok := stk.Pop(); ok {
			if v != arr[index] {
				t.Fatal("stack.{Push/Pop}() not operating properly")
			}
		} else {
			t.Fatal("stack.Pop() on non-empty stack returns not ok!")
		}
	}

	_, ok := stk.Pop()
	if ok {
		t.Fatal("stack.Pop() on empty stack return ok!")
	}
}
