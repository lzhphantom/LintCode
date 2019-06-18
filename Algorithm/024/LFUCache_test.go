package _24

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	l := LFUCache(3)
	l.Set(2, 2)
	l.Set(1, 1)
	fmt.Println(l.get(2))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	l.Set(3, 3)
	l.Set(4, 4)
	fmt.Println(l.get(3))
	fmt.Println(l.get(2))
	fmt.Println(l.get(1))
	fmt.Println(l.get(4))
}
