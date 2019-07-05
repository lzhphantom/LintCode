package _36

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		ln   *ListNode
		m, n int
		want *ListNode
	}{
		{&ListNode{
			4,
			&ListNode{
				3,
				&ListNode{
					2,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
			2, 4,
			&ListNode{
				4,
				&ListNode{
					1,
					&ListNode{
						2,
						&ListNode{
							3,
							nil,
						},
					},
				},
			},
		},
		{
			&ListNode{
				4,
				&ListNode{
					3,
					&ListNode{
						2,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
			3, 7,
			nil,
		},
	}
	for _, test := range tests {
		result := reverse(test.ln, test.m, test.n)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("reverse(%v) => %v,want %v", test.ln, result, test.want)
		}
	}

}
