package inter

type Caluclator interface {
	Add() int
	Subtract() (int, error)
	Multiply() int
}
