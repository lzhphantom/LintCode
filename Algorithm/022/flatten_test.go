package _22

import (
	"reflect"
	"testing"
)

type Myint int
type MyList []NestedInteger

func (mi Myint) isInteger() bool {
	return true
}

func (mi Myint) getInteger() int {
	return int(reflect.ValueOf(mi).Int())
}

func (mi Myint) getList() []NestedInteger {
	return nil
}

func (ml MyList) isInteger() bool {
	return false
}

func (ml MyList) getInteger() int {
	return -1
}

func (ml MyList) getList() []NestedInteger {
	return ml
}

func TestFlatten(t *testing.T) {

	tests := []struct {
		ni   []NestedInteger
		want []int
	}{
		{[]NestedInteger{
			Myint(1),
			MyList{
				Myint(2),
				MyList{
					Myint(3), Myint(4),
				},
			},
			Myint(5),
		},
			[]int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		result := flatten(test.ni)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("flatten(%v) => %v,want %v", test.ni, result, test.want)
		}
	}
}
