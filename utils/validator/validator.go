package validator

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.SetTagName("v")
}

func outRes(tag string, err *error) {
	log.Println("--------tag start---------")
	log.Println(tag, *err)
	log.Println("--------tag end-----------")
	fmt.Println()

	err = nil
}
