package test

import "time"

func Add(a int, b int) int {
	return a + b
}

func Skip(a, b int) int {
	time.Sleep(365 * 24 * time.Hour)
	return a + b
}

func Parallel(a, b int) int {
	time.Sleep(time.Duration(a+b) * time.Second)

	return a + b
}