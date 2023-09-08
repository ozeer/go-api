package utils

import (
	"fmt"
	"testing"
)

func TestGenNick(t *testing.T) {
	// 生成并打印唯一昵称
	for i := 0; i < 100000; i++ {
		fmt.Println(GenNick())
	}
}
