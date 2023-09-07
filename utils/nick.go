package utils

import (
	"math/rand"
	"sync"
	"time"
)

const (
	nickPrefix   = "outpost_"
	prefixLength = 4 // 设置前缀长度
	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	mutex         sync.Mutex
	usedNicknames = make(map[string]bool)
)

type NicknameGenerator struct {
	prefix string
}

func NewNicknameGenerator() *NicknameGenerator {
	return &NicknameGenerator{}
}

func (ng *NicknameGenerator) GenerateNickname() string {
	prefix := ng.generateRandomPrefix()
	nickname := prefix

	// 确保生成的昵称唯一
	mutex.Lock()
	defer mutex.Unlock()
	for usedNicknames[nickname] {
		prefix = ng.generateRandomPrefix()
		nickname = prefix
	}

	usedNicknames[nickname] = true
	return nickname
}

func (ng *NicknameGenerator) generateRandomPrefix() string {
	rand.Seed(time.Now().UnixNano())

	prefix := make([]byte, prefixLength)
	for i := range prefix {
		prefix[i] = charset[rand.Intn(len(charset))]
	}

	return string(prefix)
}

func GenNick() string {
	generator := NewNicknameGenerator()
	nickname := nickPrefix + generator.GenerateNickname()

	return nickname
}
