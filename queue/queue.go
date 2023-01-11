package main

import "fmt"

type Queue struct{
	items []int
}

func (q *Queue) push(val int){
	q.items = append(q.items, val)
}

func (q *Queue) pop() {
	q.items = q.items[1:]
}

func main(){
	fmt.Println("Start!")
	uk := Queue{}
	
	uk.push(1)	
	uk.push(2)	
	uk.push(3)	
	uk.push(4)	
	uk.push(5)	
	uk.pop()
	uk.pop()
	fmt.Println(uk)
}