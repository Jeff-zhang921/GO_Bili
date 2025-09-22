// 6
package main

import (
	"fmt"
)

func main() {

	//define a array
	var score [5]int
	for i := 0; i < 5; i++ {
		score[i] = i + 20
	}
	var sum int
	for i := 0; i < len(score); i++ {
		sum = sum + score[i]
	}
	fmt.Println(sum)

	//array address is 0th position
	//array initial is 0 if is int
	//get arrays position by using &array
	//arrays not like other language cannot directly show address

	for i, value := range score { //don't need value or i use _ to subsitute
		fmt.Printf("the index is %v, score[%v]=%v\n", i, i, value)
	}
	//directly init
	var arr1 [3]int = [3]int{3, 6, 9}
	//array name is arr1, array datatype is [3]int
	fmt.Println(arr1)
	var arr2 = [3]int{1, 4, 7}
	fmt.Println(arr2)
	var arr3 = [...]int{4, 5, 6, 7}
	fmt.Println(arr3)
	var arr4 = [...]int{2: 66, 0: 33, 1: 99, 3: 88}
	fmt.Println(arr4)

	//length fo array is one part of array datatype
	fmt.Printf("the datatype of the arr[1] is %T\n", arr1)
	//the datatype of the arr1 is [3]int

	//send the array to the function is just sending copy of it
	//therefore when call array won't be the address
	test20(arr1)
	fmt.Println("test20:", arr1)
	test21(&arr1)
	fmt.Println("test21:", arr1)

	//2D array---------------------------------------------------------------------------------
	var arr2D1 [2][3]int16

	//the element inside the array is an array
	//you can see it as [2] [3]int16
	//the first bracket represent the  column
	//the second braket represent the rows
	arr2D1[0][1] = 27
	arr2D1[1][1] = 56
	fmt.Println(arr2D1)
	//init a 2D array
	var arr2D2 [3][2]int = [3][2]int{{1, 2}, {3, 4}, {4, 5}}
	fmt.Println(arr2D2)

	//loop through a 2d array
	for i := 0; i < len(arr2D2); i++ {
		for j := 0; j < len(arr2D2[i]); j++ {
			fmt.Print(arr2D2[i][j], "\t")
		}
		fmt.Println()
	}
	//second way
	fmt.Println("---------------------")
	for i, value := range arr2D2 {
		for k, v := range value {
			fmt.Printf("arra[%v][%v] is %v\t", i, k, v)

		}
		fmt.Println()
	}

}
func test20(arr1 [3]int) {
	arr1[0] = 7
}
func test21(ptr *[3]int) {
	(*ptr)[0] = 7 //need the braket
}
