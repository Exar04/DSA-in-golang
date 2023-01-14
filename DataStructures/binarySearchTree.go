package main

import (
	"fmt"
)




type node struct{
	data int
	left *node
	right *node
}

type BinarySearchTree struct{
	root *node
	length int
}

func(b *BinarySearchTree) add(val int){
	var newNode node
	newNode.data = val
	newNode.left = nil
	newNode.right = nil

	temp := b.root

	if b.root == nil {
		b.root = &newNode
		return
	}
	for{
		if temp.data < val && temp.right != nil{
			temp = temp.right
		}
		if  temp.data > val && temp.left != nil {
			temp = temp.left	
		}
		if temp.data <= val && temp.right == nil {
			temp.right = &newNode	
			return
		}
		if  temp.data > val && temp.left == nil {
			temp.left = &newNode
			return
		}
	}

}

func PreOrderTraverseBST(b BinarySearchTree){
	temp := b.root
	fmt.Print(temp.data, " ")
	
	PreTraversalSub(*temp.left)
	PreTraversalSub(*temp.right)
}
func PreTraversalSub(temp node){
	if &temp == nil{
		return
	}
	print(temp.data, " ")
	if temp.left != nil {
		PreTraversalSub(*temp.left)
	}
	if temp.right != nil {
		PreTraversalSub(*temp.right)
	}
}

func InOrderTraverseBST(b BinarySearchTree){
	temp := b.root
	
	InTraversalSub(*temp.left)
	fmt.Print(temp.data, " ")
	InTraversalSub(*temp.right)
}
func InTraversalSub(temp node){
	if &temp == nil{
		return
	}
	
	if temp.left != nil {
		InTraversalSub(*temp.left)
	}
	print(temp.data, " ")
	if temp.right != nil {
		InTraversalSub(*temp.right)
	}
	
}


func PostOrderTraverseBST(b BinarySearchTree){
	temp := b.root
	
	PostTraversalSub(*temp.left)
	PostTraversalSub(*temp.right)
	fmt.Print(temp.data, " ")
}
func PostTraversalSub(temp node){
	if &temp == nil{
		return
	}
	
	if temp.left != nil {
		PostTraversalSub(*temp.left)
	}
	if temp.right != nil {
		PostTraversalSub(*temp.right)
	}
	print(temp.data, " ")
}

func main(){
	a := BinarySearchTree{}

	fmt.Println("Enter data in the BST :")
	n := 1	
	for n != -1 {
		fmt.Scan(&n)	
		if n != -1 {
			a.add(n)
		}
	}	

	PreOrderTraverseBST(a)
	fmt.Println()
	InOrderTraverseBST(a)
	fmt.Println()
	PostOrderTraverseBST(a)

}