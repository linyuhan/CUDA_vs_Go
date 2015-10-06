package main

import (
        "fmt"
        "runtime"
		"time"
)

func main() {
		canal1 := make(chan int)
        runtime.GOMAXPROCS(runtime.NumCPU())
		
		start := time.Now()
        for i := 1; i <= 100000; i++ {
                go func(num int) {
                    fmt.Println("Executando", i)
					canal1 <- 0
                }(i)
        }
		fmt.Println("==============================================")
		<-canal1
		fmt.Println(time.Since(start))
		
}