package main

import(
	"fmt"
	"runtime"
	
)

var h float64
var pi float64
var n = 9

func f(a float64) float64{
	return 4.0/(1.0 + a*a)
}

func chunk (start, end int, c chan float64){
	var sum float64 = 0.0
	for i:=start; i<end; i++{
		x:= h * (float64(i) + 0.5)
		sum += f(x)
	}
	c <- sum *h
}

func main(){
	
	np := runtime.NumCPU()
	
	runtime.GOMAXPROCS(np)

	h = 1.0/float64(n)

	c := make(chan float64, np)

	for i:=0; i<np; i++{
		go chunk(i*n/np, (i+1)*n/np, c)
	}
	for i:=0; i<np; i++{
		pi+= <-c
	}
	
	fmt.Println(pi)
}