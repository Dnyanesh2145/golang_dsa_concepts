package main

import "fmt"




func main(){

	data := []int{1,2,3,4,5}

	firstChan := make(chan int,len(data))
	for _, v := range data {
		firstChan <- v
	}
	close(firstChan)

	doubleinput := make(chan int,len(data))

	go func(){
		defer close(doubleinput)
		for num := range firstChan{
			doubleinput <- num * 2
		}

	}()

	tripleinput :=  make(chan int, len(data))

	go func(){
		defer close(tripleinput)
		for double :=  range doubleinput{
             tripleinput <- double * 3
		}
	}()

    for result :=  range tripleinput{
		fmt.Println("result",result)
	}



}