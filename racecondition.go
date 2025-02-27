package main
import "fmt"
func main(){
	fmt.println("Raceconditions")
	wg := &sync.WaitGroup{}
	var score =[]int{0}


	wg.Add(1)
	go func(wg *sync.WaitGroup){
		fmt.Println("One R")
		score = append(score, 1)
	}(wg)
	go func(wg *sync.WaitGroup){
		fmt.Println("Two R")
		score = append(score, 2)
	}(wg)
 	go func(){}()
}