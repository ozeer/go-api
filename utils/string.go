package utils

import (
	"fmt"
	"strconv"
)

// int转string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// int64转string
func Int64ToString(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// string转int
func StringToInt(s string) (a int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	return i
}

// string转int64
func StringToInt64(s string) (a int64) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return
	}
	return i
}

func FloatToString(f any) string {
	return fmt.Sprintf("%f", f)
}
