package main

import (
	"fmt"
)

func bubbleSort(arr []int) ([]int){
	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr) - 1; i++{
			if arr[i] > arr[i + 1] {
				temp := arr[i]
				arr[i] = arr[i + 1]
				arr[i + 1] = temp
			}
		}
	}
	return arr
}

func bubbleSortRecursive(arr []int)[]int{
	return bubbleSortRecursiveSubFunction(arr,len(arr), 0)
}
func bubbleSortRecursiveSubFunction(arr []int,n int, i int)[]int{
	if i >= len(arr)-1 && n == 0{
		return arr
	}
	
	if arr[i]>arr[i+1] {
		temp := arr[i]
		arr[i] = arr[i+1]
		arr[i+1]= temp
	}
	if i >= len(arr)-2 && n != 0 {
		return bubbleSortRecursiveSubFunction(arr, n - 1, 0)
	}
	return bubbleSortRecursiveSubFunction(arr,n , i + 1)
}

func main(){
	var o int
	fmt.Print("Insert the length of array : ")
	fmt.Scan(&o)	
	var newArr []int
	for i := o; i > 0; i--{
		newArr = append(newArr, int(i))
	}	
	fmt.Println(newArr)
	// newArr = bubbleSortRecursive(newArr)
	newArr = bubbleSort(newArr)
	fmt.Println(newArr)
}