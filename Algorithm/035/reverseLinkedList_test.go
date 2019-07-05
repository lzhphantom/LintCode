package _35

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		ln   *ListNode
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
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							nil,
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		result := reverse(test.ln)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("reverse(%v) => %v,want %v", test.ln, result, test.want)
		}
	}

}
