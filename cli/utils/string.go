package utils

import "strconv"

// string转int
func StringToInt(s string) (a int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	return i
}
