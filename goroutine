package main

import (
	"fmt"
)
//func greeter(s string){
//	for i:= 0; i<6; i++{
//		time.Sleep(3 * time.Second)
//	fmt.Println(s)
//}
//}
func getstatuscode(endpoint string){
	res, err:=  http.Get(endpoint)
	if err != nil{
		fmt.Println("oops is the endpoint")
		}
	fmt.Printf("%d status code for %s", res.StatusCode, endpoint)
}


func main (){
	// go greeter ("Hello")
	// greeter ("ssup?")
	websitelist := []string{
		"https://at.dev",
        "https://go.dev",
		"https://gg.com",
		"https;//github.com",
	}
	for _, web := range websitelist{
		getstatuscode(web)
	}
}
