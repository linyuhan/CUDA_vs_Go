package main
 
import (
    "fmt";
	"time";
	"math/rand";
	"math";
	"sort"
)
 
func main() {
	rand.Seed(time.Now().UnixNano())
    
	x := []float64{-10,-8,-7,-6,-5,-4,-3,-2,-1,0,1,2,3,4,5,6,7,8,9,10}
	X := [10]float64{}
	y := []float64{1,1,1,1,1,1,1,1,1,1} 
	
	
	
	
	for i:=0; i<10; i++{
		X[i] = x[rand.Intn(len(x))]
		y[i] = J(X[i])
	}
	
	
	fmt.Println("x: ", X)
	
	sort.Float64s(y)
	fmt.Println("J(x): ",y)
	
	for i:=0; i<3000; i++{
	
		fmt.Println("\n============= mutação ===============")
		
		for i:=0; i<10; i++{	
			X[i] = mutacao(X[i])
			aux := J(X[i])
			if aux < y[i]{
				y[i] = aux
			}
		}
		sort.Float64s(y)
		
		fmt.Println("J(x): ",i,y)
		
	}
	
	
}

func J(x float64) float64{
	return  math.Sin(x)*x*x - math.Sin(x)
}

func mutacao (x float64) float64{
	//x = x*math.Cos(x)
	x = x+1
	//x = x-1
	//x = x*math.Tan(x)
	//x = x*math.Sin(x)
	return x
}



