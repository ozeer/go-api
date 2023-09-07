package utils

import (
	"fmt"
	"testing"
)

func TestFloatToString(t *testing.T) {
	var score float32 = 99.9
	fmt.Println(FloatToString(score))

	var size float64 = 89.999
	fmt.Println(FloatToString(size))
}
