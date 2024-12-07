package shared

import "strconv"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ToString(i int) string {
	return strconv.Itoa(i)
}
