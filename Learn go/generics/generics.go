package main

import "fmt"

func printSlice[T comparable,V string](items []T, name V) {  // t can be any , int , string , bool ,float32 or interface{}
	for _, v := range items {
		fmt.Println(v,name)
	}
}
type stack[T any] struct {
	element []T
}
func main() {
	arr := []int{1, 2, 3}
	arr2 := []string{"hello", "world"}
	printSlice(arr,"int")
	printSlice(arr2,"String")
	stack := stack[int]{element: arr}
	fmt.Println(stack)
}
