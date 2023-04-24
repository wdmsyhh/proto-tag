package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	staff "proto-tag/test/example/tagReset"
)

func main() {
	req := staff.StringMessage{}
	validateStruct, err := govalidator.ValidateStruct(req)
	fmt.Println(validateStruct)
	fmt.Println(err)
}
