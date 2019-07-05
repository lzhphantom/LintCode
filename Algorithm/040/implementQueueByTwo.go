package _40

import "errors"

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	theStack := *s
	if len(theStack) == 0 {
		return nil, errors.New("out of index,len is 0")
	}
	value := theStack[len(theStack)-1]
	*s = theStack[:len(theStack)-1]
	return value, nil
}
