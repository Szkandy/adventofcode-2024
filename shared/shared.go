package shared

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile(path string) (lines []string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	lines = strings.Split(string(file), "\n")
	return
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
