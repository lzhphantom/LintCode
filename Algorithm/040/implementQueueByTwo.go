package _40

import "errors"

type MyQuene []interface{}

func (s MyQuene) Len() int {
	return len(s)
}

func (s MyQuene) Cap() int {
	return cap(s)
}

func (s *MyQuene) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *MyQuene) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	return (*s)[0], nil
}

func (s *MyQuene) Pop() (interface{}, error) {
	theStack := *s
	if len(theStack) == 0 {
		return nil, errors.New("out of index,len is 0")
	}
	value := theStack[0]
	*s = theStack[1:]
	return value, nil
}
