package main

import (
	"fmt"
	"reflect"
)

func main(){
	
	// Define a slice with "var".
    var newSlice1 = []string{"cat", "car", "dog"}
    fmt.Println(newSlice1)

    // Define another slice with short-hand definition.
    newSlice2 := []string{"table", "house"}
    fmt.Println(newSlice2)
	
	// Print the types of slices
	fmt.Printf("The type of newSlice1 is: %T\n", newSlice1)
	fmt.Printf("The type of newSlice2 is: %T\n", newSlice2)
	
	// Check the types are equal or not.
	fmt.Println(reflect.TypeOf(newSlice1) == reflect.TypeOf(newSlice2))
	
	// CAPACITY AND LENGTH
	// Define a slice
	s := []float64{10.1, 20.2, 30.3}
	
	// Print the capacity and length of the slice
	fmt.Printf("Length of slice s is: %d\n", cap(s))
	fmt.Printf("Capacity of slice s is: %d\n", len(s))
	
	// SLICE of a SLICE
	mySlice := []int{1, 2, 3, 4, 5}
	mySlice = mySlice[0: 3]
	
	// Print slice of a slice
	fmt.Println(mySlice)
	
	// CREATE a SLICE with make()
	
	// Define a slice.
	a := make([]int, 5, 10)
	
	// Print the slice
	fmt.Println(a)
	
	// Reassign the elements of the slice.
	a[0] = 11
	a[1] = 22
	a[2] = 33
	a[3] = 44
	a[4] = 55
	
	// Print the assigned slice.
	fmt.Println(a)
	
	// Print the length and capacity of the slice.
	fmt.Println("Length of a is: ", len(a))
	fmt.Println("Capacity of a is: ", cap(a))
	
	// APPEND OPERATION
	
	// Define a slice.
	var b []int
	
	// Apply append operation.
	b = append(b, 1)
	
	// Print the slice with capacity and length
	fmt.Println(b)
	fmt.Println("len: ", len(b))
	fmt.Println("cap: ", cap(b))
	
	// Apply another append operation.
	b = append(b, 2, 3, 4)
	
	// Print the slice with capacity and length
	fmt.Println(b)
	fmt.Println("len: ", len(b))
	fmt.Println("cap: ", cap(b))
	
	// Apply another append operation.
	b = append(b, 5)
	
	// Print the slice with capacity and length
	fmt.Println(b)
	fmt.Println("len: ", len(b))
	fmt.Println("cap: ", cap(b))
	
	// Apply another append operation.
	b = append(b, 6)
	
	// Print the slice with capacity and length
	fmt.Println(b)
	fmt.Println("len: ", len(b))
	fmt.Println("cap: ", cap(b))
	
}
