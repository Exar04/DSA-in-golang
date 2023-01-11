package main

import (
	"fmt"
)

func pSquare(n int){
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}	
		fmt.Println()
	} 
// ***
// ***
// ***
}

func pHalfPyrimidLeft(n int){
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}	
		fmt.Println()
	} 	
// *
// **
// ***
}

func pHalfPyrimidRight(n int){
	for i := 0; i < n; i++ {
		for j := 0; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}	
		fmt.Println()
	}
//  *
// **
//***
}

func pFullPyrimidUp(n int){
	for i := 0; i < n; i++ {
		for j := 0; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i*2 ; j++ {
			fmt.Print("*")
		}	
		fmt.Println()
	}
// 	*
// ***
//*****
}

func pFullPyrimidDown(n int){
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (n - i)*2 - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
// *****
//  ***
//   *
}

func main(){
	fmt.Print("Enter pattern length : ")
	var len int
	fmt.Scan(&len)
	pSquare(len)
	fmt.Println()
	pFullPyrimidDown(len)
	fmt.Println()
	pFullPyrimidUp(len)
	fmt.Println()
	pHalfPyrimidLeft(len)
	fmt.Println()
	pHalfPyrimidRight(len)
}