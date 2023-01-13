package main

import "fmt"


func isArraySorted(arr []int)bool{
	i := 0
	return checkSorted(arr, i)
}
func checkSorted(arr[]int, i int) bool{

	if i > len(arr) - 2 {
		return true	
	}

	if arr[i] > arr[i+1] {
		return false	
	}
	return checkSorted(arr, i + 1)
}


func checkElement(arr []int, n int) bool{
	i := 0
	return checkElementSubfunction(arr, n, i)
}
func checkElementSubfunction(arr []int, n int, i int)bool{
	if i > len(arr) - 1 {
		return false	
	}	
	if arr[i] == n {
		return true	
	}
	return checkElementSubfunction(arr, n, i+1)
}

func checkElementRepeatedNumber(arr []int, n int)int{
	i := 0
	return checkElementRepeatedNumberSubFunction(arr, n, i)
}
func checkElementRepeatedNumberSubFunction(arr []int, n int, i int)int{
	if i > len(arr) - 1 {
		return 0	
	}	
	if arr[i] == n {
		return 1 + checkElementRepeatedNumberSubFunction(arr, n, i+1)	
	}else {
		return checkElementRepeatedNumberSubFunction(arr, n , i + 1)
	}
}

// func binarySearch(arr []int, n int)bool{

// 	return binarySearchSubFunction(arr, n, 0, len(arr))
// }

// func binarySearchSubFunction(arr []int, n int, start int, end int)bool{
// 	var mid int
// 	mid = (end - start) / 2
// 	if arr[mid] == n {
// 		return true	
// 	}
// 	if arr[mid] < n {
// 		return binarySearchSubFunction(arr, n, mid + 1, end)	
// 	}
// 	if arr[mid] > n {
// 		return binarySearchSubFunction(arr, n, start, mid - 1)	
// 	}
// 	return false
// }


func main(){

	var arr []int
	arr = append(arr, 1,2,3,4,21,33)
	fmt.Println(isArraySorted(arr))
	fmt.Println(checkElement(arr,2))
	fmt.Println(checkElementRepeatedNumber(arr, 1))
}