package	main
 
import(. "fmt"; "os";"bufio")
 
func main() {
	 
	Printf("Escrever alguma coisa: ")
	line, _:= (bufio.NewReader(os.Stdin)).ReadString('\n')
	Printf("Imprimindo: \"%s\"\n",line[0 : len(line)-1]) // - "\n" 
	 
}