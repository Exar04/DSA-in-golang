package main

import "fmt"

type Stack struct{
	items []int
}

func (s *Stack) push(val int){
	s.items = append(s.items,val)
}

func (s *Stack) pop() []int{
	toRemove := s.items[:len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}

func (s *Stack) peek() {
	fmt.Println(s.items[len(s.items)-1:])
}

func main(){
	shelf := Stack{}
	shelf.push(1)
	shelf.push(2)
	shelf.push(3)
	shelf.peek()
	shelf.pop()
	fmt.Println(shelf)
}