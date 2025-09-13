package main

import (
	"fmt"
	"sync"
)

func Workers(workerid int, jobs <-chan int, results chan <- int) {
	for job := range jobs {
		fmt.Println("worker :", workerid, "started")
		results <- job * 2
	}

}

func main() {

	var numOfworkers = 5
	var numOfjobs = 15

	jobs := make(chan int, numOfjobs)
	results := make(chan int, numOfjobs)

	var wg sync.WaitGroup
	for i := 1; i <= numOfworkers; i++ {
		wg.Add(1)
		go func(workerid int) {
			defer wg.Done()
			Workers(workerid, jobs, results)
		}(i)
	}

	for i:=1;i<=numOfjobs;i++{
		jobs <-i
	}
    close(jobs)

	go func(){
       wg.Wait()
	   close(results)
	}()
   for result := range results{
	   fmt.Println("results--->",result)
   }

}
