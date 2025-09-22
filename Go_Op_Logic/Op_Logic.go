// 2.
package main

import (
	"fmt"
)

func main() {
	//+: add, joint string
	// /: divide, with the float the outcome will be float
	//% is the remaining
	fmt.Println(5 % 3)
	// ++: add itself
	n := 10
	n++
	//++ or--  you can only put them on their own line — you can’t use them inside another expression. like under another operation or printf
	//++or-- cannot use before var
	//num+=20 ==  num=num+20

	//logic operator && and logic operator
	// || is or logic operator
	//! not operator
	//& get the address of the var
	//* get the value in that address

	//get the value from terminal
	//scanln everytime the user put in, they need to line feed

	var age int
	fmt.Printf("input age:")
	//the point of input the address is to really change the age,
	//anything that want to access that var(age) need to go through it address to get it
	fmt.Scanln(&age)
	fmt.Println("age of the student is ", age)
	//scanf
	var name string
	var score int
	fmt.Printf("enter your name and score using space between:")
	fmt.Scanf("%s %d", &name, &score)
	fmt.Println("name:", name, "score:", score)
	//-----------------------------------------------------------------------------------------------------------------------------------------------------------
	if score < 40 { //{} cannot omit
		fmt.Println("you are not pass")
	} else if score > 70 {
		fmt.Println("you are excellent!!")
	} else {
		fmt.Println("you are pass!")
	} //else or else if must follow with }

	//you can define a var in if statement
	if counts := 20; counts < 30 {
		fmt.Println("not enough")
	}
	//----------------------------------------------------------------
	//switch case case default
	var scores int = 80
	switch scores / 10 {
	case 10:

		fmt.Println("Alevel")
	case 9:

		fmt.Println("Blevel")
	case 8:
		fmt.Println("Clevel")
	case 7, 6, 5, 4, 3, 2, 1:
		fmt.Println("not pass")
	default:
		fmt.Println("not pass")
	}

	//the func after switch must return a known value
	//var after case datatype must be the same
	//you can use more than one var after case
	//it allow no var after switch(use like if statement)
	//by using fallthrough, if satisify a,with a fallthrough it will print b too

	//for------------------------------------------------------------------------------
	var sum int = 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println(sum)
	} //for cannot use var i int=1 use i:=1 instead
	j := 0
	for j < 6 {
		fmt.Printf("%v", j)
		j++

	}
	fmt.Println()
	//dead loop
	//for{
	//	fmt.Println("hi")
	//}
	//----------------------------------------------------------------------------------------
	//for range
	var string1 string = "Hello Go"
	for i := 0; i < len(string1); i++ { //this cannot print chinese yet
		fmt.Printf("%c", string1[i]) //not %v since it will output byte
	}
	fmt.Println()
	var string2 string = "你好，go"
	for i, value := range string2 {
		fmt.Printf("index:%d, letter:%c", i, value)
	}
	//loop through string2, index using i to get, value using value to get
	var sum1 int = 0
	for i := 0; i < 100; i++ {
		sum1 += i
		fmt.Println(sum1)
		if sum1 >= 300 {
			break
		}
	}

	//break: break the loop, but only break one loop
	//you can add label to the loop and break that label to break specific loop
	//label2:
	//	for i := 0; i < 5; i++ {
	//	label1:
	//		for j := 0; j < 4; j++ {
	//
	//        if j==3{
	//			break label2
	//		}
	//
	//		}
	//	}

	for i := 0; i < 100; i++ {
		//if i%6==0{
		//	fmt.Println(i)
		//}
		if i%6 != 0 {
			continue //continue over current th loop
		}
		fmt.Println(i)

	}
	//continue only apply to the nearest loop
	//continue can also use label

	//goto label will directly jump to that label line

	//return when meet return in anycas, the code after that won't function
}
