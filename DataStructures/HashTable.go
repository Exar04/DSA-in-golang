package main

import "fmt"

const ArraySize = 5

type HashTabel struct{
	array [ArraySize]*Bucket
}

type Bucket struct{
	head *bucketNode
}

type bucketNode struct{
	key int
	next *bucketNode
}

func Init() *HashTabel{
	result := &HashTabel{}
	for i := range result.array{
			result.array[i] = &Bucket{}
	}
	return result
}

func main(){
	hasho := Init()
	fmt.Println(hasho)

}