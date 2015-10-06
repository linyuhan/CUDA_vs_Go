package main
 
import "fmt"
 
func main() {
	 
	for i := 0;; i++ {
		if  i > 3 && i <= 5 || i%2 != 0 { continue }
		if  i > 16 { break }
		fmt.Printf("%d\n",i)
	}
	 
}