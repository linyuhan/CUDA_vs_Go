package main
 
import "fmt"
 
func main() {
	 
	fmt.Printf("\nContando:\n\n")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n",i)
	}
	 
	fmt.Printf("\nContando decrescente:\n\n")
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d\n",i)
	}
	 
	fmt.Printf("\nDe 2 em 2:\n\n")
	for i := 2; i <= 10; i+=2 {
		fmt.Printf("%d\n",i)
	}
	 
	fmt.Printf("\nDe 3 em 3 descrecente a partir do 12:\n\n")
	for i := 12; i >= 1; i-=3 {
		fmt.Printf("%d\n",i)
	}
}