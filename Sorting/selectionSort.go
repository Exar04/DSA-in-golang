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


func main(){
	fmt.Println("Length of array :")
	var o int
	fmt.Scan(&o)	
	var newArr []int
	for i := o; i > 0; i--{
		newArr = append(newArr, int(i))
	}	
	fmt.Println(newArr)

	newArr = SelectionSort(newArr)
	fmt.Println(newArr)

}