package main

import "fmt"

func main() {
    for i, j := 0, 1; j < 100000000000; i, j = i+j,i {
        fmt.Println(i)
    }
}