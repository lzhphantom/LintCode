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
	deserializeMo(strArr, tn, true)
	return tn
}

func deserializeMo(strArr []string, treeNode *TreeNode, first bool) {
	if len(strArr) == 0 {
		return
	}
	if !reflect.DeepEqual(strArr[0], strDouble) && first {
		treeNode.val, _ = strconv.Atoi(strArr[0])
	}

	if len(strArr) < 2 {
		return
	}
	if !reflect.DeepEqual(strArr[1], strDouble) {
		val, _ := strconv.Atoi(strArr[1])
		treeNode.left = &TreeNode{
			val: val,
		}
	}

	if len(strArr) < 3 {
		return
	}
	if !reflect.DeepEqual(strArr[2], strDouble) {
		val, _ := strconv.Atoi(strArr[2])
		treeNode.right = &TreeNode{
			val: val,
		}
	}

	strArr = strArr[3:]

	if treeNode.left != nil {
		deserializeMo(strArr, treeNode.left, false)
	}

	if treeNode.right != nil {
		deserializeMo(strArr, treeNode.right, false)
	}

}
