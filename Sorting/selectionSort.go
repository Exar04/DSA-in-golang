package main

import (
	"fmt"
)

func SelectionSort(arr []int) ([]int){
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++{
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func SelectionSortRecursive(arr []int)[]int{
	return SelectionSortRecursiveSubFunction(arr, len(arr), 0, 0)
}

func SelectionSortRecursiveSubFunction(arr []int, n int, a int, b int)[]int{
	if n == 0 {
		return arr	
	}
	if arr[a]<arr[b] {
		temp := arr[a]
		arr[a] = arr[b]	
		arr[b] = temp
	}
	if b < len(arr) -1 {
		return SelectionSortRecursiveSubFunction(arr,n, a, b + 1)
	}
	return SelectionSortRecursiveSubFunction(arr, n -1 , a + 1, 0)
}

func main(){
	fmt.Print("Length of array : ")
	var o int
	fmt.Scan(&o)	
	var newArr []int
	for i := o; i > 0; i--{
		newArr = append(newArr, int(i))
	}	
	fmt.Println(newArr)

	// newArr = SelectionSort(newArr)
	newArr = SelectionSortRecursive(newArr)
	fmt.Println(newArr)

}