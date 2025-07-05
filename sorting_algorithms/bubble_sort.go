package sortingalgorithms



// Here’s a basic outline of how the bubble sort algorithm works:

//1. Start at the beginning of the list.

//2. Compare the first two elements. If the first is greater than the second, swap them.

//3. Move to the next pair of elements and repeat the comparison and swap if necessary.

//4.Continue this process, moving one element to the right each time, until you reach the end of the list.

//5.Repeat steps 1–4 until no more swaps are needed, indicating that the list is sorted.

// Time complexity: O(n^2)

func BubbleSort(list []int) []int{
	for i:=0;i< len(list)-1;i++{
		for j:=0;j<len(list)-i-1;j++{
			if list[j]> list[j+1]{
				list[j],list[j+1] = list[j+1],list[j]
			}
			
		}
	}
   return list
}




