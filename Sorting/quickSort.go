package main
// This code is broken
import "fmt"

func quickSort(arr []int, low int, high int){

	if low >= high {
		return
	}
	s := low
	e := high
	mid := low + (high - low)/2
	povit := arr[mid]

	for s <= e {
		fmt.Print("o")
		for arr[s]<= povit && s <= e{
		fmt.Print("e")
			s++		
		}
		for arr[e] >= povit && e > s{
			e--
		}	
		if s <= e {
			temp := arr[s]
			arr[s] = arr[e]
			arr[e] = temp
			e--
			s++
		}
	}
	quickSort(arr, low, e)
	quickSort(arr, s, high)
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

	quickSort(newArr, 0, len(newArr) - 1)
	fmt.Println(newArr)	
}