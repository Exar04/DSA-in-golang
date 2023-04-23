package main

import (
	"fmt"
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) int {
	num3 := []int {}
	num3 = append(num3, nums1...)
	num3 = append(num3, nums2...)

	sort.Ints(num3)
  	lenOfArr := len(num3)

	median := 0
	if(lenOfArr%2 == 0){
		mid := lenOfArr/2
		medianOx := math.Floor(float64(mid))
		// medianOx = int(medianOx)
		newoi := int16(medianOx)
		median = (num3[newoi] + num3[newoi + 1])/2
	}else {
		median = num3[lenOfArr/2]
	}
  return median
}

func main(){
	ok := []int {1,2,3}
	okis := []int {4,5,6,7, 8}
	fmt.Print(findMedianSortedArrays(ok, okis))
}