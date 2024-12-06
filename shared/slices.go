package shared

func WithoutIndex(s []int, index int) []int {
	var dst []int
	dst = append(dst, s[:index]...)
	return append(dst, s[index+1:]...)
}
