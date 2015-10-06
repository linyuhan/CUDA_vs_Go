package main
 
import "fmt"
 
func main() {
	 
	x := "abc"
	 
	for { // É um "loop" puro
		 
		fmt.Printf("%s\n",x)
		 
		switch x {
		case "abc":			 x += " def"; continue
		case "abc def":      x += " ghi"; continue
		case "abc def ghi":  x += " jkl"; continue// os 'continue' é do for, e não do switch
		}
		 
		break
		 
	}
	 
}