package myutility

import "fmt"

func (m Mapping) Add() int {
	num1 := m["FirstNumber"]
	num2 := m["SecondNumber"]
	return num1 + num2

}
func (o MyOwn) Subtract() (int, error) {
	if o.FirstNumber > o.SecondNumber {
		num1 := o.FirstNumber
		num2 := o.SecondNumber
		return num1 - num2, nil
	}

	Err1 := fmt.Errorf("first number is less than second")
	return 0, Err1
}

func (l MyOwn) Multiply() int {
	num1 := l.FirstNumber
	num2 := l.SecondNumber
	return num1 * num2
}
