package _29

//给出三个字符串:s1、s2、s3，判断s3是否由s1和s2交叉构成。
func isInterleave(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	interleaved := make([][]bool, 0)
	for i := 0; i <= len(s1); i++ {
		interleaved = append(interleaved, make([]bool, len(s2)+1))
	}
	interleaved[0][0] = true

	for i := 1; i <= len(s1); i++ {
		if s3[i-1] == s1[i-1] && interleaved[i-1][0] {
			interleaved[i][0] = true
		}
	}

	for i := 1; i <= len(s2); i++ {
		if s3[i-1] == s2[i-1] && interleaved[0][i-1] {
			interleaved[0][i] = true
		}
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if (s3[i+j-1] == s1[i-1] && interleaved[i-1][j]) || (s3[i+j-1] == s2[j-1] && interleaved[i][j-1]) {
				interleaved[i][j] = true
			}
		}
	}

	return interleaved[len(s1)][len(s2)]

}
