package main

import "fmt"

type dynamicArray struct {
	array []int
	len   int
	cap   int
}

func New(cap int) *dynamicArray {
	return &dynamicArray{
		array: make([]int, cap),
		len: 0,
		cap: cap,
	}
}

func (a *dynamicArray) add(x int) {
	if a.len == a.cap {
		fmt.Println("Array resized")
		a.resize()
	}
	a.array[a.len] = x
	a.len++
}

func (a *dynamicArray) resize() {
	newCap := a.cap * 2
	newArray := make([]int, newCap)
	copy(newArray, a.array)

	a.array = newArray
	a.cap = newCap
}

func callDynamicArray() {
	a := New(2)
	for i := range 10 {
		a.add(i)
	}

	fmt.Println(a)
}