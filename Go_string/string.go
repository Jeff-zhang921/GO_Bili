// 5
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//len(str)
	str := "golang"
	fmt.Println(len(str))
	//chinese use 3 byte

	//loop through string using for range
	//On each iteration you get two values:
	//i → the starting byte index of this rune inside the string
	//value → the rune (Unicode code point) at that position

	for i, value := range str {
		fmt.Printf("index:%v, var:%c", i, value)
		fmt.Println()

	}
	fmt.Println("-------------------------------------------------------------")
	//loop through string using切片
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	//convert from a string to int
	nums, _ := strconv.Atoi("666")
	fmt.Println()
	fmt.Println(nums)
	//convert from a in to string
	str1 := strconv.Itoa(666)
	fmt.Println(str1)

	//how many substring are there in string
	counts := strings.Count("golanggo", "go")
	fmt.Println(counts)

	//don't care the capital or not to compare two string
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)
	//if care the capital use ==
	//search the first appearance of an substring
	fmt.Println(strings.Index("golanggo", "go"))

	//change some substring
	//the last number represent how many want to subsitute
	//-1 represent subsitute all
	//1 or 2 or... represent subsitute 1 or 2 substring
	fmt.Println(strings.Replace("golanggo", "go", "java", -1))
	//to split by specific point
	arra := strings.Split("go-python-java", "-")
	fmt.Println(arra)
	//strings.ToLower()
	//strings.ToUpper()\
	//trim two side space
	strings.TrimSpace("  go   java   ")

	//---------------------------------------------------------------------------------------------------------------
	//time
	fmt.Println(time.Now())
	//Now return as a struct
	fmt.Printf(" type:%T", time.Now())
	fmt.Println()
	fmt.Printf("year:%v\n", time.Now().Year())
	//different between printf and sprintf: sprintf can get the string wile printf directly print to the terminal

	//change the format of how to print year,month...
	datestr := time.Now().Format("2006/01/02 15/04/05 ")
	datestr2 := time.Now().Format("2006, 15:04")
	//every number must be the specific, i.e.what i wrote upside(no matter where is year it must be 2006)
	fmt.Println(datestr)
	fmt.Println(datestr2)

	//new(int) is a pointer point to the datatype of int and it return value is a pointer
	ptr := new(int)
	fmt.Printf("datatype of ptr is:%T,the address of ptr is%v, ptr=: %v, ptr point to the value%v", ptr, &ptr, ptr, *ptr)

	//----------------------------------------------------------------------------------------------------------------------------------------
	test10()
	fmt.Println("logic  continue ")
	err := test11()
	if err != nil {
		fmt.Println("self define error:", err)
		//if there is a point if something go wrong,you don't have to run following code
		panic(err)
	}

}
func test10() {
	//to use defer with an anonymous function
	defer func() {
		err := recover() //if there is no error it will return nil
		if err != nil {
			fmt.Println("error capture")
			fmt.Println("error is:", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
	return
	//if something goes wrong it will stop unless use defer recovery
	//after the function is executed, the error will capture by recovery, since defer will execute after the
	//function executed, then anything goes run in function can always been capture, make sure code won't stop
}

// declare a self define error
func test11() error {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("num2 cannot =0")

	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
		//return nil if nothing run

	}
}
