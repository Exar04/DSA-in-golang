package main

import (
	"fmt"
)

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func (l *linkedList) addAtHead(val int){
	newNode := &node{data: val}
	second := l.head
	l.head = newNode
	l.head.next = second
	l.length++
}

func (l *linkedList) addAtEnd(val int){

	newNode := &node{data: val}
	temp := l.head
	if temp == nil {
		l.head = newNode
		return	
	}
	for temp.next != nil {
		temp = temp.next	
	}
	temp.next = newNode
}

func (l *linkedList) deleteAtEnd(){
	temp := l.head
	for temp.next.next != nil{
		temp = temp.next
	}
	temp.next = nil
}

func (l *linkedList) deleteAtHead(){
	temp := l.head.next
	l.head = temp
}

func (l *linkedList) addAfterVal(AtVal int, val int){
	newNode := &node{data: val}
	temp := l.head
	for temp.data != AtVal {
		temp = temp.next	
	}
	
	newNode.next, temp.next = temp.next, newNode
}

func (l *linkedList) addBeforeVal(AtVal int, val int){
	newNode := &node{data: val}
	temp := l.head
	for temp.next.data != AtVal {
		temp = temp.next	
	}
	
	newNode.next, temp.next = temp.next, newNode
}

func (l *linkedList) printAllNodes(){
	temp := l.head
	for temp != nil{
		fmt.Print(temp.data)
		temp = temp.next	
	}
}

func add2LinkedLists(lis1 *linkedList, lis2 *linkedList){
	var list = []int{}

	temp1 := lis1.head
	temp2 := lis2.head

	for temp1 != nil{
		list = append(list, temp1.data)
		temp1 = temp1.next	
	}

	for i, value := range list{
		if temp2 == nil {
			goto end	
		}
		newNum := value + temp2.data
		list[i] = newNum
		temp2 = temp2.next
	}
	for temp2 != nil{
		list = append(list, temp2.data)
		temp2 = temp2.next	
	}

	end:
	fmt.Println(list)
}

func main(){
	Lis1:= linkedList{}
	// Lis1.addAtHead(1)
	Lis1.addAtHead(2)
	Lis1.addAtHead(3)
	Lis1.addAtHead(3)
	Lis1.addAtHead(1)

	Lis2 := linkedList{}
	Lis2.addAtEnd(5)
	Lis2.addAtEnd(6)
	Lis2.addAtEnd(2)

	Lis1.printAllNodes()
	fmt.Println()
	Lis2.printAllNodes()
	fmt.Println()
	add2LinkedLists(&Lis1, &Lis2)

}