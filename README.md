# Golang

![Golang Image](golang.png)

---------------------------------------------------------------------

## Slices

Arrays in Go programming language have fixed size. This situation has some pros and cons. For instance, when an element has to be appended in an array which has full length, another array should be created. An so on, all the elements have to be transferred to the new array like C programming language. On the other hand, slices are another composite type in Golang which is dinamically-sized and flexible. Since it has no fixed-size, slices are preferred commanly in practice. 

* Like arrays, there are 2 decleration of slices.
1. Defining with **var** keyword.

2. Short-hand definition.

```go
package main


import(
    "fmt"
)


func main(){
    
    // Define a slice with "var".
    var newSlice1 = []string{"cat", "car", "dog"}
    fmt.Println(newSlice1)

    // Define another slice with short-hand definition.
    newSlice2 := []string{"table", "house"}
    fmt.Println(newSlice2)
}
```

```console
[cat car dog]
[table house]
```

* Let's look at the types of the slices that are defined above.

```go
// Print the types of slices
fmt.Printf("The type of newSlice1 is: %T\n", newSlice1)
fmt.Printf("The type of newSlice2 is: %T\n", newSlice2)
```

```console
The type of newSlice1 is: []string
The type of newSlice2 is: []string
```

* Now let's look at the types of the slices are equal or not?

```go
// Check the types are equal or not.
fmt.Println(reflect.TypeOf(newSlice1) == reflect.TypeOf(newSlice2))
```

```console
true
```

---------------------------------------

## Length and capacity

* A slice has **length** and **capacity**. The expressions for length and capacity for any slice **s** are **len(s)** and **cap(s)**.

```go
// Define a slice
s := []float64{10.1, 20.2, 30.3}
	
// Print the capacity and length of the slice
fmt.Printf("Length of slice s is: %d\n", cap(s))
fmt.Printf("Capacity of slice s is: %d\n", len(s))
```

```console
Length of slice s is: 3
Capacity of slice s is: 3
```
