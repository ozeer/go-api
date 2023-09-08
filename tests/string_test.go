package tests

import (
	"fmt"
	"testing"

	"github.com/ozeer/go-api/utils"
)

func TestFloatToString(t *testing.T) {
	var score float32 = 99.9
	fmt.Println(utils.FloatToString(score))

	var size float64 = 89.999
	fmt.Println(utils.FloatToString(size))
}
