package sortingalgorithms

// Here’s a basic description of how the insertion sort algorithm works:

// 1.Start with the second element of the list (assuming the first element is already sorted).

// 2.Compare this element with the elements in the sorted sublist, moving from right to left.

// 3.If the current element is smaller than the element being compared, shift the compared element one position to the right.

// 4. Repeat step 3 until you find the correct position for the current element in the sorted sublist.

// 5.Insert the current element into its correct position in the sorted sublist.

// 6.Repeat steps 2–5 for each remaining unsorted element until the entire list is sorted.

// https://medium.com/@alexfoleydevops/dsa-in-golang-sorting-f896eb89ac9b

// Time Complexity  := 
// Insertion Sort is an efficient algorithm for sorting small lists or lists that are almost sorted. Its average and worst-case time complexity is O(n²), where n is the number of elements in the list. However, its best-case time complexity is O(n) when the list is already sorted. Additionally, insertion sort has a space complexity of O(1), making it suitable for situations with limited memory resources.

func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1

		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}

	return list
}
