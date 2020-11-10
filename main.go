package main

import (
	"fmt"

	"github.com/chanpon2015/ci-test/usecase"
)

func main() {
	o := usecase.NewOutput()
	if err := o.Out(); err != nil {
		fmt.Println(err)
	}
}
