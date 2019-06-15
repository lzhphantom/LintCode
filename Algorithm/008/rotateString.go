package _08

func rotateString(str *string, offset int) {
	if len(*str) == 0 {
		return
	}
	if offset > len(*str) {
		offset %= len(*str)
	}
	s := (*str)[len(*str)-offset:] + (*str)[:len(*str)-offset]
	*str = s
}
