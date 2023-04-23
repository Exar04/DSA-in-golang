package main

import "fmt"

func main(){
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		if i%3 == 0  {
			fmt.Println("fizz")	
		}
		if i%5 == 0  {
			fmt.Println("Buzz")	
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}
}