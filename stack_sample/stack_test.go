package main

import "testing"

func TestStack(t *testing.T) {
	value := make([]string, 4)
	var stack = Stack{}
	stack.Value = value
	stack.Push("dataA")
	stack.Push("dataB")
	stack.Push("dataC")
	stack.Pop()
	stack.Pop()

	// stackの期待値
	expect1 := "dataA"
	expect2 := "dataD"

	if result1 := stack.Pop(); result1 != expect1 {
		t.Errorf("result1 wants %v, but actuary %v", expect1, result1)
	}

	stack.Push("dataD")

	if result2 := stack.Pop(); result2 != expect2 {
		t.Errorf("result1 wants %v, but actuary %v", expect2, result2)
	}
}

func TestStackLimit(t *testing.T) {
	value := make([]string, 4)
	var stack = Stack{limit: 2}
	stack.Value = value
	stack.Push("dataA")
	stack.Push("dataB")
	stack.Push("dataC")

	// stackの期待値
	expect1 := "dataC"
	expect2 := "dataB"

	if result1 := stack.Value[len(stack.Value)-1]; result1 != expect1 {
		t.Errorf("result1 wants %v, but actuary %v", expect1, result1)
	}

	if result2 := stack.Value[len(stack.Value)-2]; result2 != expect2 {
		t.Errorf("result1 wants %v, but actuary %v", expect2, result2)
	}
}

func TestChangeLimit(t *testing.T) {
	value := make([]string, 4)
	var stack = Stack{limit: 2}
	stack.Value = value
	stack.ChangeLimit(3)

	// 期待値
	expect := 3
	if stack.limit != expect {
		t.Errorf("result1 wants %v, but actuary %v", expect, stack.limit)
	}
}
