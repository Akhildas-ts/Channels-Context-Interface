package main

import "fmt"


func removeElement(slice []int, i int) []int {
    return append(slice[:i], slice[i+1:]...)
}

func main() {
    mySlice := []int{10, 20, 30, 40, 50}
    indexToRemove := 2 // Remove the element at index 2 (value 30)

    updatedSlice := removeElement(mySlice, indexToRemove)
    fmt.Println(updatedSlice) // [10 20 40 50]
}


//1 . find the length of the slice


// func lengthOfSlice( ){

  
// 	length := [...]int{1,3,5,6,7,8}
// 	fmt.Println(length)


// }



// // 2. so when we change the value into the slice it will change at the original 
// //   [1 2 7 4 5]  // array[2] = 7
// // [1 2 7 4 5]
// //  but in  the case of append its not changing the original array 
// // 

// 	func main() {

// 		array := []int{1, 2, 3, 4, 5}

// 		hi(array)

// 		fmt.Println(array)

// 	}

// 	func hi(array []int) {

 
// 		fmt.Println(array)

// 	}




