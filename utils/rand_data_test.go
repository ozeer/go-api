package utils

import (
	"log"
	"testing"
)

func TestGenerateRandomDate(t *testing.T) {
	for i := 0; i < 200; i++ {
		date := GenerateRandomDate(1949, 2023)
		log.Println("随机日期:", date)
	}
}

func TestGenRandomNumber(t *testing.T) {
	for i := 0; i < 200; i++ {
		number := GenRandomNumber(20)
		log.Println("随机整数:", number)
	}
}
