package main

import "fmt"

func revArray(arr []int, l int, r int) []int{
	if l >= r {
		return arr
	}
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		revArray(arr, l+1, r-1)
	return arr
}

func main() {
	var arr []int
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)	
	}
	fmt.Println("OK")
	arr = revArray(arr, 0, len(arr) -1)
	fmt.Println(arr)
}
