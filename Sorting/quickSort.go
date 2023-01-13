package main

import "fmt"

// This code is broken
// func quickSort(arr []int, low int, high int){

// 	if low >= high {
// 		return
// 	}
// 	s := low
// 	e := high
// 	mid := low + (high - low)/2
// 	povit := arr[mid]

// 	for s <= e {
// 		fmt.Print("o")
// 		for arr[s]<= povit && s <= e{
// 		fmt.Print("e")
// 			s++		
// 		}
// 		for arr[e] >= povit && e > s{
// 			e--
// 		}	
// 		if s <= e {
// 			temp := arr[s]
// 			arr[s] = arr[e]
// 			arr[e] = temp
// 			e--
// 			s++
// 		}
// 	}
// 	quickSort(arr, low, e)
// 	quickSort(arr, s, high)
// }


func quickSort(arr []int) []int {
	return quickSortSubFunc(arr, 0, len(arr)-1)
}
func quickSortSubFunc(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortSubFunc(arr, low, p-1)
		arr = quickSortSubFunc(arr, p+1, high)
	}
	return arr
}
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func main(){
	fmt.Print("Length of array : ")
	var o int
	fmt.Scan(&o)
	var newArr []int
	for i := o; i > 0; i-- {
		newArr = append(newArr, int(i))
	}
	fmt.Println(newArr)

	quickSort(newArr)
	fmt.Println(newArr)	
}