package _12

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

//@value is a number
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("out of index,len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (statck *Stack) min() (interface{}, error) {
	theStack := *statck
	if len(theStack) == 0 {
		return nil, errors.New("out of index,len is 0")
	}
	min, _ := theStack[0].(int)
	for index, value := range theStack {
		if index == 0 {
			continue
		}
		val, ok := value.(int)
		if ok {
			if val < min {
				min = val
			}
		}
	}
	return min, nil
}

func MinStack() {
	s := Stack{}
	s.Push(1)
	min, _ := s.min()
	fmt.Println(min)

	s.Push(2)
	min, _ = s.min()
	fmt.Println(min)

	s.Push(3)
	min, _ = s.min()
	fmt.Println(min)
}
