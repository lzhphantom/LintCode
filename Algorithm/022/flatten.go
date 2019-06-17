package _22

type NestedInteger interface {
	isInteger() bool
	getInteger() int
	getList() []NestedInteger
}

//递归
func flatten(ni []NestedInteger) []int {
	result := make([]int, 0)

	for _, integer := range ni {
		if integer.isInteger() {
			result = append(result, integer.getInteger())
		} else {
			result = append(result, flatten(integer.getList())...)
		}
	}
	return result
}

//非递归
func flattenNew(ni []NestedInteger) []int {

	isFlat := true
	ls := ni
	for isFlat {
		isFlat = false
		newLs := make([]NestedInteger, 0)
		for _, integer := range ls {
			if integer.isInteger() {
				newLs = append(newLs, integer)
			} else {
				newLs = append(newLs, integer.getList()...)
				isFlat = true
			}
		}
		ls = newLs
	}
	r := make([]int, 0)
	for _, ls := range ls {
		r = append(r, ls.getInteger())
	}
	return r
}
