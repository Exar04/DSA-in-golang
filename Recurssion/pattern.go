package main

import "fmt"

func triangle1(n int)int{
	i := n
	if n == 0  {
		return 0
	}
	if n != 0 && i != 0 {
		return triangle1SubFunc(n,i)	
	}
	return 0
}
func triangle1SubFunc(n int, i int)int {
	if i != 0 {
		fmt.Print("*")	
		return triangle1SubFunc(n,i - 1)
	}
	fmt.Println()
	return triangle1(n-1)
}


func triangle2(n int){
	i := n
	if n == 0  {
		return
	}
	if n != 0 && i != 0 {
		triangle2SubFunc(n,i)
		return	
	}
}
func triangle2SubFunc(n int, i int){
	if i != 0 {
		triangle2SubFunc(n,i - 1)
		fmt.Print("*")	
		return
	}
	triangle2(n-1)
	fmt.Println()
	return
}
func main(){
	// triangle1(4)
	triangle2(4)
}