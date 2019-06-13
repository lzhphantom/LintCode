package _01

//给出两个整数 a 和 b , 求他们的和
func aplusb1(a, b int) int {
	return a + b
}

//给出两个整数 a 和 b , 求他们的和
//用位运算
func aplusb2(a, b int) (result int) {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return aplusb2(sum, carry)
}
