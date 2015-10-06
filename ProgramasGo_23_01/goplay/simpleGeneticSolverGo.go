package main
 
import (
    "fmt";
    "math/rand";
    "time";
	"math"
)
 
func main() {
    rand.Seed(time.Now().UnixNano())
	
    const geneSet = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!."
    target := "Studing Go and CUDA! Final Project coming soon"
	
	
    calc := func (candidate string) int {
        return getFitness(target, candidate)
		//calc vai receber a quantidade de letras que está igual
    }
     
    start := time.Now()
     
    disp := func (candidate string) {
        fmt.Print(candidate) //imprime a palavra candidato
        fmt.Print("\t")
        fmt.Print(calc(candidate)) //imprime a quantidade de letras que está igual
        fmt.Print("\t")
        fmt.Println(time.Since(start))
    }
	
    var best = getBest(calc, disp, geneSet, len(target))
	//best recebe
    println(best)
     
    fmt.Print("Total time: ")
    fmt.Println(time.Since(start))
}

 
func getBest(getFitness func(string) int, display func(string), geneSet string, length int) string { // rotina executada apenas 1x

	i := 0
    var bestParent = generateParent(geneSet, length)
	//bestParent vai receber uma variação qualquer de string
    value := getFitness(bestParent)
	//value vai receber a quantidade de letras que está igual do bestParent com o target
    var bestFitness = value
     
		for bestFitness < length {//enquanto a quantidade de letras iguais não for o mesmo que o target
			
        	child := mutateParent(bestParent, geneSet)
        	fitness := getFitness(child) //calcula a quantidade de palvras iguais 
			
        	if fitness > bestFitness { //se a nova quantidade for maior que a quantidade antiga
            	display(child)//imprime o atual
            	bestFitness = fitness //quantidade antiga recebe a nova
            	bestParent = child //novo filho
        	}
			i++
    	}
		fmt.Println("Total loops:", i)
    return bestParent //retorna o filho = target
}
 
func mutateParent(parent, geneSet string) string {
    geneIndex := rand.Intn(len(geneSet)) //recebe uma posição qualquer do gene [ex: posição 4]
    parentIndex := rand.Intn(len(parent)) //recebe uma posição qualquer do pai [ex: posição 7]
    candidate := ""
    if parentIndex > 0 { //se a posição não for 0
        candidate += parent[:parentIndex] //candidato recebe as letras das 4 primeiras posições [do 0~3]
    }
    candidate += geneSet[geneIndex:1+geneIndex] //candidato recebe no adicional 
    if parentIndex+1 < len(parent) {
        candidate += parent[parentIndex+1:]
    }
    return candidate
}
 
func generateParent(geneSet string, length int) string {
	//recebe por parâmetro o gene e o tamanho do target
    s := ""
    for i := 0; i < length; i++ {
        index := rand.Intn(len(geneSet)) //randomiza uma letra do gene
        s += geneSet[index:1+index] //joga na string
    }
    return s
}
 
func getFitness(target, candidate string) int {
    differenceCount := 0
	//calcula a diferença entre o objetivo e o candidato
    for i := 0; i < len(target); i++ {
        if target[i] != candidate[i] {
            differenceCount++
        }
    }
	//retorna um int com a quantidade de letras que estão iguais
    return len(target) - differenceCount
}