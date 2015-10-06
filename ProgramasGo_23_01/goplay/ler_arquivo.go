package main
 
import (. "fmt"; "io/ioutil")
 
func main() {

	doc, _ := ioutil.ReadFile("C:/Users/Lin/AppData/Local/liteide/goplay/arquivo.txt")
	Printf(string(doc) + "\n")

}