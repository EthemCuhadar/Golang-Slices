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

----------------------------------------

## Slice of a slice

* If you want to get part of a slice, it can be defined with range in brackets **[start: end]**. Let's look at an example.

```go
mySlice := []int{1, 2, 3, 4, 5}
mySlice = mySlice[0: 3]

// Print slice of a slice
fmt.Println(mySlice)
```

```console
[1 2 3]
```

----------------------------------------

## Creating a slice with make

* **make()** is a built-in function to create slices. The function takes **length** and **capacity** as parameters. Let's look at an example below.

```go
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
```

```console
[0 0 0 0 0]
[11 22 33 44 55]
```

* Now let's print the capacity and length of the slice.

```go
// Print the length and capacity of the slice.
fmt.Println("Length of a is: ", len(a))
fmt.Println("Capacity of a is: ", cap(a))
```

```console
Length of a is:  5
Capacity of a is:  10
```

--------------------------------------

## Appending to a slice

* Append operation in Go programming language is provided by **append()** built-in function. It takes the slice that we want into append and the values which is to be appended as parameters. Let's look at an example below.

```go
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
```

```console
[1]
len:  1
cap:  1
[1 2 3 4]
len:  4
cap:  4
[1 2 3 4 5]
len:  5
cap:  8
[1 2 3 4 5 6]
len:  6
cap:  8
```
