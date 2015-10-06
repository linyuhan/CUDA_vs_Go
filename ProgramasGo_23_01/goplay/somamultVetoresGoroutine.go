package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	
	cpu := runtime.NumCPU()
	
	runtime.GOMAXPROCS(cpu)
	
	arraySize := 501

	a := make([]int, arraySize)
	b := make([]int, arraySize)
	//c := make([]int, arraySize)
	//d := make([]int, arraySize)
	
	sum := make(chan int)
	mult := make(chan int)

	for i := 0; i < arraySize; i++ {
		a[i] = i
		b[i] = i
	}

	start := time.Now()

	for i := 0; i < arraySize; i++ {
		go func(i int){
			//c[i] = a[i] + b[i]
			sum <- a[i] + b[i]
		}(i)
		go func(i int){
			//d[i] = a[i] * b[i]
			mult <- a[i] * b[i]
		}(i)
		fmt.Println(i,<-sum, <-mult)
		//fmt.Println(<-mult)
	}
	
	//fmt.Println(c)
	
	fmt.Println("\n================================================================================================")
	
	//fmt.Println(d)
	//fmt.Println(<-mult)

	fmt.Println(time.Since(start))
}
