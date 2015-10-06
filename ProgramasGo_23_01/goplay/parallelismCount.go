package main

import (
        "fmt"
        "runtime"
		//"math/rand"
        //"time"
)

func main() {
		c := make(chan string)
        runtime.GOMAXPROCS(runtime.NumCPU())
        for i := 0; i <= 450000; i++ {
                go func(num int) {
			//		    time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
						c <- fmt.Sprintln("Hi from goroutine", num)
                }(i)
        }

        // Wait for finish
        //time.Sleep(1 * time.Second)
		for i := 0; i <= 450000; i++ {
			fmt.Print(<-c)
		}
}