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

func main(){
	Lis := linkedList{}
	Lis.addAtHead(1)
	Lis.addAtHead(2)
	Lis.addAtHead(3)

	Lis.addAtEnd(4)
	Lis.addAtEnd(5)
	Lis.addAtEnd(6)

	Lis.deleteAtHead()
	Lis.deleteAtEnd()
	Lis.addAfterVal(4,3)
	Lis.addBeforeVal(4,9)

	Lis.printAllNodes()

}