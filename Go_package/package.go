// 4.
package main

import "fmt"

//package name under same directory better be the same as directory name, can be different file name
// under working dir is main.go and other subdir, under other subdir all file has same package name with that subdir

//you can rename the package, but after rename, you must need to use that name
//import(//rename the test to gocode/...package
//      test   gocode/.../.../.../
//         )

//reason for multi package,
//1. not put all function to one file
//2. function want to have same name
//package and file and dir better be the same

//main function under main package is the executable package

//when using another function in different local package
//you need to locate to the dir of that file
//you can set the enviornment variable to make it easy
//and when use you need to goPackage.function

//function name cannot be the same under same package!
//

// initfunction is apply before main been apply
// global var init first than init function
// other package (if have import to this package) will execute first than global var
var num = testing()

func testing() int {
	fmt.Println("test func has executed")
	return 10

}
func init() {
	fmt.Println("init func has executed")
}
func main() {
	fmt.Println("main has executed")
	//anonymous function
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(result)
	fmt.Println(Sub(20, 30))
	f := getSum()
	//f:=getSum() get the return value fuction
	//f:=getSum f is the reference of getSum like rename

	fmt.Println(f(1))
	fmt.Println(f(2)) //this is three not two! closures!
	fmt.Println(adding(30, 50))
}

// outside the main declare an anonymous
// anonymous function to a var: sub=anonymous
var Sub = func(num1 int, num2 int) int {
	return num1 - num2
}

// --------------------------------------------------------------------------------------------------------
// function return a function
// a closure function: it an anonymous function that must use a var that outside that function
// and the value won't vanish it will add on to the next calculation
func getSum() func(int) int {
	var n int = 10
	return func(num int) int {
		n = n + num
		return n
	}
}

// -------------------------------------------------------------------------------------------------------
// when go meets defer, it will put the var that has defer keyword to stack ,and execute it later
// so what ever the var change in later the function, the var that in stack will execute later won't change
// it execute defer after the function has been executed, then they execute in stack way LIFO
func adding(num1 int, num2 int) int {
	defer fmt.Println("num1:", num1)
	defer fmt.Println("num2", num2)
	sums := num1 + num2
	fmt.Println("sums:", sums)
	return num1 + num2

}

//when you want to close a resource, you add a defer when using it, it will auto close after the function executed
