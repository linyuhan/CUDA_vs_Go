//http://codereview.stackexchange.com/questions/28386/fibonacci-generator-with-golang
package main

import(
	"fmt"
	"time"
)

func fib_generator() chan int {
	
	//Canal de comunicação/sincronização: conecta duas computações concorrentes.
	//Passam informações como troca de mensagens ordenada entre as gorotinas ou
	//realizam a sincronização entre elas.
	//"Reduzem a complexidade de trabalhar com memória compartilhada"
  	c := make(chan int)


  	go func() { 
    	for i, j := 0, 1; ; i, j = i+j,i {
        	c <- i
    	}
  	}()

  	return c
}

func main() {
	
	start := time.Now()
    c1 := fib_generator()


	
	
    // read first 10 numbers from 1st channel
    for n := 0; n < 50 ; n++ { 
		fmt.Println(" ", <- c1) 
		}
    fmt.Println()


	
	
	fmt.Println(time.Since(start))
}