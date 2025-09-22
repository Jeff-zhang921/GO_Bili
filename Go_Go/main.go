// 1.
package main //declare the dir that doc at
//every go need to inside a dir(package) which should be the name of （...) .go
import (
	"fmt"
	"strconv"
	"unsafe"
) //import the package that program need to run

//go build ... compile the program
//it will become an exe that can run independently
//go run .... will build and run together in one step while not create exe

//dos is the interface that have before windows
//change disk: D: change the disk
//dir: see the directory
//enviornment path: the commands that can execute in sdk bin :go version
//since bin has go.exe, but other place cannot, to make the commands that can
//execute in other place, set the enviornment path to the sdk bin, then any of path can execute

// to change the name of exe when compiling using go build -o hello.exe main.go
// no need to add column at the end
// everything include package and var that is not use will report error
func main() {
	fmt.Println("hello world")
	//tab all go back shift+tab all go front 4block
	//use gofmt -w main.go will help you make the formate more elegant
	//when there is too many in one line seperate by changing the line and add , at  back of every line
	var age int = 19
	fmt.Println(age)
	//declare or not the type of var
	var age1 = 19
	fmt.Println(age1)
	//different from =
	sex := "boy"
	fmt.Println(sex)

	//declare more var with one go
	var n1, n2, n3 = 10, "jeff", 78
	fmt.Println(n1, n2, n3) // , will use as a space in print

	n6, height := 6, 175
	fmt.Println(n6, height)

	var (
		n9  = 9
		n10 = 10
	)
	fmt.Println(n9, n10)

	//basic datatype: num, bool, char, string
	//complicated: pointer list,...
	//change from 10 to 2 by divide by 2 and revert the remaining
	// int8(-128-127) int16....: is signed number
	//unit8(0-255) unit16...: is unsigned number
	//golang by default using int
	//printf use as a format shown, it will fill the %T
	fmt.Printf("typeof num3 %T", n9)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(n9))
	//byte= unit8
	//float32=4 byte
	//float 64=8byte using float64 will more precise(golang default)
	var num1 float32 = 314e-2 //e means times 10 to the power of
	fmt.Println(num1)

	//char in golang is byte
	//ascii is the 0-128position of Unicode char set
	//unicode using utf-8 to encode
	var c1 byte = 'a' //byte is only 0-128 so it cannot translate chinese encode
	fmt.Println(c1)
	//if want to show the byte corresponding char must use format
	fmt.Printf("byte c1 change to char is %c", c1)
	// /n change line /b backspace one char /n backspace to the front
	fmt.Println("aaa/bbbbb")
	fmt.Println("aaaaa/rbbbbb")

	// \" :this is use for disable the function that " make it as a char or string

	var flag bool = true
	fmt.Println(flag)

	//after define a string whole string can change but cannot change only some of the string inside the whole string
	s1 := "gos"
	//	s1[0] = "s" wrong
	fmt.Println(s1)

	//when there are special char inside like " inside a string, use `` to represent string instead of ""
	s2 := `"you can see " inside a string and /n"`
	fmt.Println(s2)
	//---------------------------------------------------------------------
	//joint two string
	s3 := "abc" + "def"
	s3 += "jhi"
	fmt.Println(s3)
	//default value: int=0 float=0 bool=false string=""
	//---------------------------------------------------------------------
	//when changing the datatype of different data need to delcare every time no matter small to big
	var n8 int = 5
	// this will error var n6 float32 = n5
	//need to force change
	var n7 float32 = float32(n8)
	fmt.Println(n7)
	//when force change from a bigger format to smaller it might incorrect、
	//-----------------------------------------------------------------------------------------------------------------------------------------------
	//convert string
	var number1 int = 19
	var string1 string = fmt.Sprintf("%d", number1)
	//%d int Sprintf means translate using Decimal

	fmt.Printf("type of string1 is %T, s1=%v \n", string1, string1)
	//%T in printf use as a shown type
	//%V use as directly show the value of it

	var number2 float32 = 4.1
	var string2 = fmt.Sprintf("%f", number2)
	//if is bool, use %t

	fmt.Printf("type of string2 is %T, s2=%q \n", string2, string2)
	//%q if it is a string, it will use "" to show
	//---------------------------------------------------------------------------------------------------------------------------------
	// or use strconv to convert to string
	var number3 int = 20
	var string3 = strconv.FormatInt(int64(number3), 10)
	//the first number is the var that you want to convert
	//if is int then need to force convert to int64
	//second number is how to interprete. in decimal or in byte
	fmt.Printf("type of the string3 is %T, s3=%q \n", string3, string3)

	var number4 float64 = 2.456
	var string4 = strconv.FormatFloat(number4, 'f', 9, 64)
	//first is the var, second represent float shown in decimal third represent 9 decimal point, forth represent float64
	fmt.Printf("type of the string4 is %T, s4=%q\n", string4, string4)

	var number5 bool = true
	var string5 = strconv.FormatBool(number5)
	fmt.Printf("type of string5 is %T, s5=%q\n", string5, string5)
	//-------------------------------------------------------------------------------------- --------------------------------------
	//when parsing string to other datatype
	//they have two return value, value and error
	//since we don't care error, and have two return, just use ,  _ is to represent don't care
	var b, _ = strconv.ParseBool(string5)
	fmt.Printf("datatype of number5 is %T n5=%v\n", b, b)

	//when parse to an int need to have the in what base : decimal or byte, and what the size int8 or 16 or64...
	var c, _ = strconv.ParseInt(string1, 10, 64)
	//if is float then parsefloat(string2,32or64 ) second para means convert to float32 or float64
	fmt.Printf("datatype of number1 is %T n1=%v\n", c, c)
	//if the something goes wrong it will output the default value
	//------------------------------------------------------------------------------------------------------
	//pointer
	var count int = 20
	fmt.Println("the address of count is ", &count)
	//define a pointer
	var ptr *int = &count
	//ptr is a pointer point to the datatype of int
	//ptr itself has been store in another area
	//and the value inside ptr is the address value of count
	fmt.Println("the value that store in ptr is", ptr)
	fmt.Println("the address of ptr itself is ", &ptr)
	//to get the value(20)
	//*ptr is the value of the address inside the address where pointer point to
	fmt.Println("the value inside the address that ptr point to is ", *ptr)
	//the datatype inside the address must be the same when declare pointer what datatype to point to

	//------------------------------------------------------------------------------------------------------------------------------------
	//identifier
	//the _ use as ignore. anything after _ will be ignore
	//anything that is not important can use _ to represent ignore

	//package name can be different from dir name or document name
	//but package name must be main to run this code
	//it better to name the same of dir, doc, package

	//if the front letter in captical then it can be use in other package
	//if the front letterr is in lower case then it can be only access by this package

	//import after the package, the package that import need to use ""
	//when delcare a package that is local, you use lower case, outside inanyway cannot access
	//
}
