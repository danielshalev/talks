package main

import (
    "fmt"
)

func main() {
    // START OMIT
    // slice formal initialization
    myListFormal := make([]int, 5, 10) //make(type, length, capacity)
    fmt.Printf("%d, %d\n", len(myListFormal), cap(myListFormal))

    // slice literal initialization
    myListLiteral := []int{1, 2, 3, 4, 5}
    fmt.Printf("%d, %d\n", len(myListLiteral), cap(myListLiteral))

    // slice a slice
    sliceOfSlice := myListLiteral[2:3]
    fmt.Printf("%d, %d\n", len(sliceOfSlice), cap(sliceOfSlice))
    // END OMIT
}

