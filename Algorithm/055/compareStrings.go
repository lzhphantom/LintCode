package _55

//输入

//"ABCD"
//"ABCC"
//输出
//true
//期望答案
//false
//比较两个字符串A和B，确定A中是否包含B中所有的字符。字符串A和B中的字符都是 大写字母
func compareStrings(a, b string) bool {

	if a == b {
		return true
	}

	for _, vb := range b {
		isContain := false
		for _, va := range a {
			if va == vb {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}
	return true
}
