package main

import (
	"fmt"
	"math"
)

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func factorial(n int)int{
	if n <= 1 {
		return 1	
	}
	return n* factorial(n - 1)
}

func sumAll(n int) int{
	if n <= 1 {
		return 1	
	}
	return n + sumAll(n -1)
}

func printAllNum(n int){
	if n == 0{
		return
	}
	printAllNum(n - 1)	
	fmt.Print(n, " ")
}

func reverseNumber(n float64)int{

	var digit int
	digit = int(math.Floor(math.Log10(n)))
	return revNumHelp(int(n), digit)

	//1234
	//4321
}
func revNumHelp(n int, digit int) int{
	if n%10 == n {
		return n	
	}
	rem := n%10
	return rem * int(math.Pow10(digit)) + revNumHelp(n/10, digit - 1)
}

func NoOfZeros(n float64) int{
	var digit int
	digit = int(math.Floor(math.Log10(n)))
	return NumZero(int(n), digit)	
}

func NumZero(n int, digit int)int{
	rem := n%10
	if n%10 == n && n != 0{
		return 0
	}
	if rem == 0 {
		return 1 + NumZero(n/10, digit - 1)	
	}else {
		return NumZero(n/10, digit - 1)
	}
	
}

func main() {
	// fmt.Println(factorial(4))
	// fmt.Println(sumAll(4))
	// printAllNum(4)
	fmt.Println(reverseNumber(1234567))
	fmt.Println(NoOfZeros(1002))

}
