package tests

import (
	"log"
	"testing"

	"github.com/ozeer/go-api/utils"
)

func TestGenerateRandomDate(t *testing.T) {
	for i := 0; i < 200; i++ {
		date := utils.GenerateRandomDate(1949, 2023)
		log.Println("随机日期:", date)
	}
}

func TestGenRandomNumber(t *testing.T) {
	for i := 0; i < 200; i++ {
		number := utils.GenRandomNumber(20)
		log.Println("随机整数:", number)
	}
}
