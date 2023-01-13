package main

import "fmt"

func mergeSort(arr []int) []int{
	mid := len(arr) / 2

	if mid != 0 {
		return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
	}
	return arr
}
func merge(first []int, second []int)[]int{
	var mixArr []int
	i := 0
	j := 0
	
	for i < len(first) && j < len(second){
		if first[i] < second[j] {
			mixArr = append(mixArr, first[i])	
			i++
		}
		if first[i] > second[j] {
			mixArr = append(mixArr, second[j])	
			j++
		}
	}
	if i >= len(first) {
		mixArr = append(mixArr, second...)
	}
	if j >= len(second) {
		mixArr = append(mixArr, first...)
	}
	return mixArr
}

func main() {
	fmt.Print("Length of array : ")
	var o int
	fmt.Scan(&o)
	var newArr []int
	for i := o; i > 0; i-- {
		newArr = append(newArr, int(i))
	}
	fmt.Println(newArr)

	newArr = mergeSort(newArr)
	fmt.Println(newArr)
}
