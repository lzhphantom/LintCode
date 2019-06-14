package _07

import (
	"reflect"
	"strconv"
	"strings"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

const strDouble = "#"

func serialize(tn *TreeNode) string {
	tmp := []*TreeNode{tn}

	strList := make([]string, 0)
	if tn == nil {
		return ""
	}
	double := 0

	for len(tmp) != 0 {

		now := tmp[0]
		strList = append(strList, strconv.Itoa(now.val))

		for double != 0 {
			strList = append(strList, strDouble)
			double--
		}

		double = 0
		if now.left != nil {
			tmp = append(tmp, now.left)
		} else {
			double++
		}

		if now.right != nil {
			tmp = append(tmp, now.right)
		} else {
			double++
		}
		tmp = append(tmp[1:])
	}

	for double != 0 {
		strList = append(strList, strDouble)
		double--
	}
	for reflect.DeepEqual(strList[len(strList)-1], strDouble) {
		strList = strList[:len(strList)-1]
	}

	return strings.Join(strList, ",")
}

func deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}
	tn := new(TreeNode)
	strArr := strings.Split(data, ",")
	for index, str := range strArr {
		if index == 0 {
			tn.val, _ = strconv.Atoi(str)
		}
		if index%2 != 0 {
			//left
			if reflect.DeepEqual(str, strDouble) {

			}
		} else {
			//right
		}

	}
	return tn
}
