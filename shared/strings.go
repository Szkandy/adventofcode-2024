package shared

import (
	"strconv"
)

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func GetChar(str string, index int) rune {
	return []rune(str)[index]
}
