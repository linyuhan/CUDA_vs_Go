package main

import(
	."fmt"
	."io/ioutil"
	
)


func main(){
	
	transicoes, a := ReadFile("C:/Users/Lin/AppData/Local/liteide/goplay/transicoes.txt")
	if a != nil{panic(a)}
	
	palavra, a := ReadFile("C:/Users/Lin/AppData/Local/liteide/goplay/palavra.txt")
	if a != nil{panic(a)}
	
	tr := string(transicoes)
	p := string(palavra)
	
	tam := len(p)

	aux := 3

	for tr[aux] != p[1] && string(tr[aux-1]) != "-" && string(tr[aux+1])!="-"{ //procura o primeiro alfabeto da palavra no automato
		aux++
	}
	//Println(string(tr[aux-1]))
	//Println(string(tr[aux]))
	//Println(string(tr[aux+1]))
	if string(tr[aux-2]) != "N"{
		Println("entro")
		if string(tr[aux-1]) == "-" && string(tr[aux+1])=="-"{
			
			Println(string(tr[aux]))
			
			prox := string(tr[aux+2:aux+4])
			ant := string(tr[aux-3:aux-1])
			
			for i:=1; i<tam; i++{
				for tr[aux] != p[i]{ //procura o primeiro alfabeto da palavra no automato
					aux++
				}
				ant = string(tr[aux-3:aux-1])
				Println("Anterior: "+ant)
				Println("Próximo: "+prox)
				if ant == prox{
					prox = string(tr[aux+2:aux+4])
					
					if i==tam-1{
						Println("Palavra Aceita! Estado final: "+prox)
						break
					}
				}else{
					Println(p)
					Println("Palavra não aceita pelo autômato!")
					break
				}
			
			}
			
		}
		
		
	}else{
		Println("Palavra não aceita pelo autômato!")
		
	}
		
}
