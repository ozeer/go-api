package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/ozeer/go-api/utils"
)

func TestGenNick(t *testing.T) {
	// 生成并打印唯一昵称
	for i := 0; i < 100000; i++ {
		fmt.Println(utils.GenNick())
	}
}

func TestGenNewNick(t *testing.T) {
	for i := 0; i < 100; i++ {
		log.Println(utils.GenNewNick(uint64(i)))
	}
}
