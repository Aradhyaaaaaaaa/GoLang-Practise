package main 

import "fmt" , "log", "net/http"
func handler(w http.ResponceWriter, R *HTTP.Request){
	fmt.Fprintf(w, "Hello, world from %sn", r.URLPath[1:])//handlin one request
}
func main(){
	http.HandleFunc("/", Handler)
	log.fatal(http.listenandserver(":8080", nill))
	}   