package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenNick(t *testing.T) {
	// 生成并打印唯一昵称
	for i := 0; i < 10; i++ {
		fmt.Println(GenNick())
	}
}

func TestSeed(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	prefix := make([]byte, prefixLength)
	for i := range prefix {
		prefix[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println("老的随机数", string(prefix))

	rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix1 := make([]byte, prefixLength)
	for j := range prefix1 {
		prefix1[j] = charset[rand.Intn(len(charset))]
	}
	fmt.Println("新的随机数", string(prefix1))

}
