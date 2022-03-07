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
	
}
