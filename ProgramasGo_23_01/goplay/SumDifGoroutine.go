package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate two random numbers.
	x := rand.Int()
	y := rand.Int()
	// Create channels for sum and difference.
	sum := make(chan int)
	dif := make(chan int)
	// Start a goroutine to calculate the sum.
	go func() {
		sum <- x + y // Send the result to channel.
	}()
	// Start a goroutine to calculate the difference.
	go func() {
		dif <- x - y // Send the result to channel.
	}()
	// Receive the results from the channels and print.
	fmt.Printf("sum=%v dif=%v\n", <-sum, <-dif)
}