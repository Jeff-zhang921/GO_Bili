// 3.
package main

import "fmt"

func main() {
	sum, div := cal(10, 20)
	fmt.Println(sum)
	fmt.Println(div)
	//if don't want to use one return,use _
	sum2, _ := cal(20, 10)
	fmt.Println(sum2)

	sum3, _ := cal1(1, 2)
	fmt.Println(sum3)

	var n1 = 5
	var n2 = 10
	swap(&n1, &n2)
	fmt.Println("n1:", n1, "n2:", n2)
	//function in go do not support overloading

	fmt.Println("------------------------------------------")
	test()
	test(2)
	test(1, 2, 3)

	//data and array in function won't affect data or array that in main
	//if you want to change data in main through function, use pointer to change the value in address
	//--------------------------------------------------------------------------------------------------------
	//function in golang is a datatype
	a := test1 //without para
	//if with para, then a is the return value of test
	fmt.Printf("datatype of a is %T, datatype of test1 is %T", a, test1)
	a(10)
	test2(10, 1.4, a)

	//self define a new var name, but it actually int
	type myInt int
	var num1 myInt = 30
	//a(num1)
	fmt.Println("num1:", num1)
	//although it myInt just a different name of int,
	//it still is different when go compile, myInt cannot mixed with int no matter when
	//deliver a para or reassign a 	myInt to int directly
	fmt.Println(int(num1)) //you need to convert to int to make it stays same with int

	//but if give a function different name, it is completely the same with that function even when you call the new name

}

// -------------------------------------------------------------------------------------------
// first paren is para,second is the return type
// if cal is Cal then every package can use
func cal(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

// or
func cal1(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2 //no need declare
	sub = num1 - num2
	return //rememeber to return
}

// if return value is null then don't need to declare return value
// -----------------------------------------------------------------------------------------
func swap(ad1 *int, ad2 *int) {
	//var ptr1 *int = &num1
	//var ptr2 *int = &num2

	//temp = num1
	//*ptr1 = *ptr2
	//*ptr2 = temp
	temp := *ad1
	*ad1 = *ad2
	*ad2 = temp
	//*ptr is the value of the address inside the address where pointer point to
	//but var ad1 *int is declare a pointer point to the location of ...
	//one is value inside location and one is location

	//when var is in new function, think it as it is now has different storage space(new area between main)
	//the function use the storage area is different from storage area that main function use

}

// -------------------------------------------------------------------------------------------------------------------
// function in go support different number of para that takes in everytime
// args...int allow to take in multiple different int type para
func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	} //see it as an array
}

// define a func
func test1(num int) {
	fmt.Println(num)

}

// declare a new name which exactly the same with test1
type myFunc func(int) //the second para is the datatype

// the third para, need to know the function datatype first
func test2(num int, num2 float32, testFunc func(int)) { //or myFunc instead of func(int)
	fmt.Println("test02-----")

}
