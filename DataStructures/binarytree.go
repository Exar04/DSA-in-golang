package main

import "fmt"

type node struct{
	data int
	left *node
	right *node
}

type BinaryTree struct{
	root *node
}

func(b *BinaryTree) add(val int){
}


func main(){
	a := BinaryTree{}
	fmt.Println(a)
}