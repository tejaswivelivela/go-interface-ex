package main

import (
	"fmt"
	"go-interface-ex/internal/myutility"
)

func main() {
	create := make(myutility.Mapping, 0)
	create["FirstNumber"] = 10
	create["SecondNumber"] = 20
	fmt.Println(create.Add())
	form := myutility.MyOwn{}
	form.FirstNumber = 5
	form.SecondNumber = 10
	sub, err := form.Subtract()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Subtract: %v\n", sub)
	star := myutility.MyOwn{}
	star.FirstNumber = 12
	star.SecondNumber = 12
	fmt.Printf("Multiplication: %v\n", star.Multiply())
}
