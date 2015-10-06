package main
 
import ("io/ioutil"; "regexp")
 
func main() {
	 
	oldfile := "C:/Users/Lin/AppData/Local/liteide/goplay/arquivo.txt" //<= "abcdefghijkl" 
	newfile := "C:/Users/Lin/AppData/Local/liteide/goplay/arquivo2.txt"
	 
	docu, _ := ioutil.ReadFile(oldfile) // docu = []uint8 
	doc := string(docu)                 // doc  = string 
	 
	re := regexp.MustCompile("[b-c]|[i-j]")
	docu = []byte(re.ReplaceAllString(doc, "X"))
	 
	ioutil.WriteFile(newfile,docu,0644) //=> "aXXdefghXXkl" 
	 
}