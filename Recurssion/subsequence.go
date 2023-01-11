package main

import "fmt"


func subRec(arr []int){
}

func sub(arr []int){
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++{
			if j != i {
				fmt.Print(arr[j])
			}
		}
		fmt.Println()
	}
	fmt.Println(arr)
}

func main(){
	var arr []int
	arr = append(arr, 1,2,3,4,5)
	sub(arr)

	var ao [][]int
	a1 := []int{1,2,3}
	a2 := []int{4,5,6}
	ao = append(ao,a1,a2)
	fmt.Println(ao)
}