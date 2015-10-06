package main
 
import ("io/ioutil")
 
func main() {
	 
	ioutil.WriteFile("C:/Users/Lin/AppData/Local/liteide/goplay/arquivo.txt",
	[]byte("Um texto qualquer"),0644)
	 
}