package main
 
import "fmt"
 
func duploRetorno() (int, int) {
	return 1, 2
}
 
func main() {
	 
	x, y := duploRetorno()
	fmt.Printf("%d\n",x)
	fmt.Printf("%d\n",y)
 
}