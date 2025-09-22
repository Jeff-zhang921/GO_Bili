// 7
package main

import "fmt"

// slice based on array,
func main() {
	var arra [5]int = [5]int{3, 6, 9, 12, 15}

	//slice is  dynamic length
	//start from arra[1]position to arra[3]position
	//exclude arra[3]
	var slice []int = arra[1:3]
	//or slice:=arra[1:3]
	fmt.Println("slice:", slice)
	fmt.Println("length of slice:", len(slice))
	fmt.Println("slice capacity:", cap(slice)) //approximate 2 times the real used

	//slice is actually a struct:
	//the pointer to the slice
	//the length of slice
	//the capacity of slice
	fmt.Printf("position of arra[1]:%p\n", &arra[1])
	fmt.Printf("position of slice[0]:%p", &slice[0])

	//change the slice[x] will change the correspodence value under that array address
	slice[1] = 16
	fmt.Println("intarr:", arra)
	fmt.Println("slice:", slice)

	//use make directly create slice without array
	//make function 3 para:slice type, slice length, slice capacity

	slice1 := make([]int, 4, 20)
	fmt.Println(slice1)
	fmt.Println("slice capacity:", cap(slice1))
	slice1[0] = 66
	slice1[1] = 88
	//make actually create an array but is private

	//
	slice3 := []int{1, 4, 6}
	fmt.Println("slice3:", slice3)
	fmt.Printf("slice3 capacity:%v\n", cap(slice3))

	//loop through slice-----------------------------------------------------
	slice4 := []int{6, 7, 8, 9, 0}
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice[%v] value is%v\n", i, slice4[i])

	}
	fmt.Println("---------------")
	for i, value := range slice4 {
		fmt.Printf("slice[%v] value is%v\n", i, value)
	}
	//slice can slice a slice
	slice5 := slice4[1:2]
	fmt.Println(slice5)
	//slice can append but append function is a new function
	append6 := append(slice5, 88, 20) //add 88 and 20
	fmt.Println("append6:", append6)
	//therefore can but is = not :=
	slice5 = append(slice5, 88, 20)
	fmt.Println("slice5:", slice5)
	slice6 := []int{99, 44}
	slice5 = append(slice6, slice5...)
	fmt.Println(slice5)

	//slice copy
	var a []int = []int{1, 2, 3, 4}
	var b []int = make([]int, 5, 10)
	copy(b, a) //first is dest second is source
	//it copy the element one by one, the array under slice datatype won't chnage
	//print:12340 not only 1234
	fmt.Println(b)

	//------------------------------------------------------------------------------------------------------------------------------------------

	//map
	//use key to get value
	//name:value

	//key value can be bool int pointer channel
	//key can be int string
	//value can be int string map struct
	//key cannot ne slice map function

	//define a map var
	var m map[int]string //key is int value is string
	//only declare won't allocate space
	//must use make to init, to allocae space
	m = make(map[int]string, 10) //map can allocate space for 10
	//the second var of make can omit
	m[2606056] = "jeff"
	m[2606057] = "taby"
	m[2606058] = "david"
	//it is unsequence when store to memory
	//key cannot duplicate, latter same key will subsitute before same key
	//value can be the same
	fmt.Println(m)


	//three type of init a map var
	n := make(map[int]string)
	n[2606057] = "jeff"
	n[2606058] = "taby"

	//operating on map
	//to change the value under certain key just rewrite key value pair with the old key
	//delete map
	delete(n, 2606058)
	fmt.Println(n)
	//search a key value pair through key to search value
	//value bool=map[key]
	val, _ := m[2606058]
	fmt.Println(val)

	//getting the length of map
	fmt.Println("length of map m:", len(b))

	//loop through map only allow for range
	for i, value := range m {
		//i here is the index which is the key of map
		//value is the value of that key
		fmt.Printf("map[%v] is %v\n", i, value)
	}
	fmt.Println("---------------------------------------------------------------")
	r := make(map[string]map[int]string)
	r["class1"] = make(map[int]string)
	r["class1"][2606056] = "jeff"
	r["class1"][2606057] = "taby"
	r["class1"][2606058] = "david"
	r["class2"] = make(map[int]string)
	r["class2"][2606059] = "alex"
	//treat it as a 2D array
	
	for k, _ := range r {
		for x, y := range r[k] {
			fmt.Printf("class: %v  id:%v name%v\n", k, x, y)
		}
	}

}
