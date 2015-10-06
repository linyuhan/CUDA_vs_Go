package main

import (
    "fmt"
    "math/rand"
    "sort"
    "time"
	"runtime"
    //"bufio"
)

type gene struct {
    x float32   //posicao x
    y float32   //posicao y
    v float32   //valor da funcao
}

const QThreads = 2
const Quantos = 10
const Quantos2 = 2 * Quantos

type Tpop [Quantos]gene

var Gpop1 Tpop //ilha 1
var Gpop2 Tpop //ilha 2

var Gtemperatura, Gtemperatura2 float32

var c chan int

const NCPU = 4



//------------------------------
// Methods required by sort.Interface.
func (s Tpop) Len() int {
    return Quantos
}
func (s Tpop) Less(i, j int) bool {
    return s[i].v < s[j].v
}
func (s Tpop) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

//------------------------------------------
// esta funcao recebe um gene, aplica mutacao
//  e devolve um novo gene
func radiacao_gama1(entra gene) gene {

    entra.x *= (1.0 - Gtemperatura) + (2 * Gtemperatura * rand.Float32())
    entra.y *= (1.0 - Gtemperatura) + (2 * Gtemperatura * rand.Float32())

    return entra

}

func radiacao_gama2(entra gene) gene {

    entra.x *= (1.0 - Gtemperatura2) + (2 * Gtemperatura2 * rand.Float32())
    entra.y *= (1.0 - Gtemperatura2) + (2 * Gtemperatura2 * rand.Float32())

    return entra

}

//------------------------------------------
// esta funcao aplica mutacao a varios genes,
//  ou seja uma populacao inteira.
// depois cada gene eh avaliado e a populacao
//  e organizada, do pior ate o melhor.
func (apop *Tpop) chernobyl1(sizepop int) {

    for i := 0; i < sizepop/2; i++ {
        apop[i*2] = apop[i]
    }

    for i := 0; i < sizepop; i++ {
        apop[i] = radiacao_gama2(apop[i])
        apop[i].v = apop[i].x*apop[i].x + apop[i].y*apop[i].y
    }
	
    sort.Sort(apop)
	c<-1
}

func (apop *Tpop) chernobyl2(sizepop int, i int, n int, c chan int) {

    for ; i < sizepop/2; i++ {
        apop[i*2] = apop[i]
    }

    for ; i < n; i++ {
        apop[i] = radiacao_gama1(apop[i])
        apop[i].v = apop[i].x*apop[i].x + apop[i].y*apop[i].y
    }
	
    sort.Sort(apop)
	c<-1
}

//------------------------------------------
func f(a *gene, i *float32) { a.x = *i }

//------------------------------------------
func g(a gene, i float32) { a.x = i }

//------------------------------------------
func migracao(){
	
	var aux gene
	migracao := 1
	if (Gpop1[0].v < 0.001 || Gpop2[0].v < 0.001) && migracao==1{ //migração
			fmt.Println("migracao")
			aux = Gpop1[0]
			Gpop1[0] = Gpop2[0]
			Gpop2[0] = aux
			migracao = 0
		}
}


//------------------------------------------
func main() {

    var a [32]int32
    var i int32
	
	

    i = 4

    a[i] = 5

    rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(NCPU)
	
	
	
    Gtemperatura = 0.1
	Gtemperatura2 = 0.2

    var x1 gene

    x1.x = 10
    x1.y = 10
    x1.v = 100

    for i := 0; i < Quantos; i++ {
        Gpop1[i] = x1
		Gpop2[i] = x1
    }

	start := time.Now()

	//genetic()
	counter := 0
	c := make(chan int, NCPU)
	//
	for Gpop1[0].v > 0.0001{
	for i:=0; i<NCPU; i++{
		
		go Gpop1.chernobyl2(Quantos, counter*Quantos/NCPU, (counter+1)*Quantos/NCPU, c)
		
		counter++
		}
	}
	
	<-c

	fmt.Printf("1-%d) %f %f %f \n", counter, Gpop1[0].x, Gpop1[0].y, Gpop1[0].v)
	//fmt.Printf("2-%d) %f %f %f \n", counter, Gpop2[0].x, Gpop2[0].y, Gpop2[0].v)
	
	fmt.Println(time.Since(start))

}