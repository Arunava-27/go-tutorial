package main

import "fmt"

func main(){
	// array
	var intArr [3]int32 // intArr := [3]int32{132, 232, 332}
	intArr[0] = 132 
	intArr[1] = 232
	intArr[2] = 332
	fmt.Println(intArr[0]) 
	fmt.Println(intArr[1:3]) // [232 332]

	fmt.Println(&intArr[0]) // address of first element

	// slice
	var intSlice []int32 = intArr[:] // intSlice := []int32{132, 232, 332}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 432)
	fmt.Println(intSlice)

	var slice2 []int32 = []int32{4, 5, 6}
	intSlice = append(intSlice, slice2...) // intSlice = append(intSlice, 4, 5, 6) same as spread operator in JS
	fmt.Println(intSlice)

	// slice with make
	var slice3 []int32 = make([]int32, 5, 10) // make([]T, length, capacity)
	fmt.Println(slice3)
	fmt.Println("Length",len(slice3))
	fmt.Println("Capacity",cap(slice3))


}