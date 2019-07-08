package _40

import (
	"fmt"
	"testing"
)

func TestImplementQueueByTwo(t *testing.T) {
	mq := MyQuene{}
	mq.Push(1)
	num, _ := mq.Pop()
	fmt.Println(num)
	mq.Push(2)
	mq.Push(3)
	num, _ = mq.Top()
	fmt.Println(num)
	num, _ = mq.Pop()
	fmt.Println(num)
}
