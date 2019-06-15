package _07

import (
	"reflect"
	"testing"
)

func TestSs(t *testing.T) {
	tests := []struct {
		tn   *TreeNode
		want string
	}{
		{
			&TreeNode{
				1,
				nil, nil,
			},
			"1",
		},
		{
			&TreeNode{
				2,
				&TreeNode{
					3,
					nil,
					nil,
				},
				&TreeNode{
					7,
					nil,
					nil,
				},
			},
			"2,3,7",
		},
		{
			&TreeNode{
				3,
				&TreeNode{
					9,
					nil, nil,
				},
				&TreeNode{
					20,
					&TreeNode{
						15,
						nil, nil,
					},
					&TreeNode{
						7,
						nil, nil,
					},
				},
			}, "3,9,20,#,#,15,7",
		},
	}

	for _, test := range tests {
		result := serialize(test.tn)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("serialize(%v) => %s,want %s", *test.tn, result, test.want)
		}
	}

}

func TestDs(t *testing.T) {
	tests := []struct {
		str  string
		want *TreeNode
	}{
		{"1",
			&TreeNode{
				1,
				nil, nil,
			},
		},
		{
			"1,3,7",
			&TreeNode{
				1,
				&TreeNode{
					3,
					nil, nil,
				},
				&TreeNode{
					7,
					nil, nil,
				},
			},
		},
	}
	for _, test := range tests {
		result := deserialize(test.str)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("deserialize(%s) => %v,want %v", test.str, result, test.want)
		}
	}
}
