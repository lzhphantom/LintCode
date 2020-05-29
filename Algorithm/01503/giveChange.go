/*
@Time : 2020/5/29 12:47
@Author : lzhphantom
@File : giveChange
@Software: GoLand
*/
package _1503

type change struct {
	One     int //1元
	Four    int //4元
	OneSix  int //16元
	SixFour int //64元
	All     int //1024元
}

func giveChange(amount int) (result change) {
	total := 1024 - amount
	if total == 1024 {
		result.All = 1
	} else {
		if total >= 64 {
			result.SixFour = total / 64
		}
		if total%64 >= 16 {
			result.OneSix = total % 64 / 16
		}
		if total%64%16 >= 4 {
			result.Four = total % 64 % 16 / 4
		}

		if total%64%16%4 > 0 {
			result.One = total % 64 % 16 % 4
		}
	}
	return result
}
