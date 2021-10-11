package main

import "fmt"

/**
验证 Golang append 多个值
 */

func main() {
	//for i := 0; i < 1000; i++ {
	//	sliceIn := make([]int, i, i)
	//	var slice []int
	//	slice = append(slice, sliceIn...)
	//	fmt.Printf("len=%d, cap=%d, %d\n", len(slice), cap(slice), i)
	//
	//} // 5 len=5, cap=8
	array := [...]int{1,2,3}
	slice := []int{1,2,3}
	fmt.Println(array[3:], slice[3:]) // [] []


}
