package main

import "fmt"

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

func main(){
	var o int
	fmt.Scan(&o)	
	var newArr []int
	for i := o; i > 0; i--{
		newArr = append(newArr, int(i))
	}	
	fmt.Println(newArr)

	newArr = bubbleSort(newArr)
	fmt.Println(newArr)
}