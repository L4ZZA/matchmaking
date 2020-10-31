package data

// // shiftLeft moves all element one element to the left
// // the first element will move to the back of the slice
// // mantains the size of the slice
// func shiftLeft(i int, slice *[]int){
// 	fmt.Sprintf("Slice Before - %v\n", slice)
// 	first := slice[i]
// 	copy(slice[i:], slice[i+1:]) // Shift slice[i+1:] left one index.
// 	last := len(slice)-1
// 	slice[last] = -1     // Erase last element (write zero value).
// 	// slice = slice[:last]     // Truncate slice.
// 	slice[last] = first
// 	fmt.Sprintf("Slice After - %v\n", slice)
// }

func remove(i int, list Sessions) int{
	// Remove the element at index i from list.
	removed := i
	copy(list[i:], list[i+1:]) // Shift list[i+1:] left one index.
	last := len(list)-1
	list[last] = nil     // Erase last element (write zero value).
	list = list[:last]     // Truncate list.
	return removed
}