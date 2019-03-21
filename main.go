package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err == nil {
		for i := 0; i < n; i++ {
			fmt.Println(fib(i))
		}
	} else {
		fmt.Println("Error!", err)
	}
}

func fib(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + b
	}

	return a
}
