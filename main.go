package main

import (
	"fmt"
	sortingalgorithms "golang_dsa_concepts/sorting_algorithms"
)


func main() {
	input := []int{6,8,9,1,0,5,4,2,3,7}

	fmt.Println("Bubble Sort  :", sortingalgorithms.BubbleSort(input))

	fmt.Println("Insertion Sort :" , sortingalgorithms.InsertionSort(input)) 


}