package utils

import (
	"math/rand"
	"sync"
	"time"

	sq "github.com/sqids/sqids-go"
)

const (
	nickPrefix = "outpost_"
	// 设置前缀长度
	prefixLength = 4
	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitSet     = "0123456789"
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
	randNew := rand.New(rand.NewSource(time.Now().UnixNano()))

	prefix := make([]byte, prefixLength)
	for i := range prefix {
		prefix[i] = charset[randNew.Intn(len(charset))]
	}

	// 添加一个随机数字作为后缀，确保唯一性
	prefix = append(prefix, digitSet[randNew.Intn(len(digitSet))])

	return string(prefix)
}

func GenNick() string {
	generator := NewNicknameGenerator()
	nickname := nickPrefix + generator.GenerateNickname()

	return nickname
}

func GenNewNick(uid uint64) string {
	s, _ := sq.New(sq.Options{
		MinLength: 6,
	})
	id, _ := s.Encode([]uint64{uid})

	return id
}
