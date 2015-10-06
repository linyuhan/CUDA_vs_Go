package main

import (
	"fmt"
	"time"
)

func main() {
	arraySize := 500

	a := make([]int, arraySize)
	b := make([]int, arraySize)
	c := make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		a[i] = i
		b[i] = i
	}

	start := time.Now()

	for i := 0; i < arraySize; i++ {
		c[i] = a[i] + b[i]
	}
	fmt.Println(time.Since(start))
	fmt.Println(c)

	
}
