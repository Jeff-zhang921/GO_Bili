// 8 struct
// struct is a datatype too
package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	School string
}
type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var stu Student
	fmt.Println(stu)
	stu.Name = "Jack"
	stu.Age = 18
	fmt.Println(stu)

	//create instance of go
	var stu1 Student = Student{}
	fmt.Println(stu1)

	//new is create a pointer
	var stu2 *Student = new(Student)
	//stu2 is a pointer
	(*stu2).Name = "Jeff"

	var stu3 *Student = &Student{}
	(*stu3).Name = "Jeff"

	stu4 := new(Student)
	(*stu4).Name = "Jeff"

	//---------------------------------------------
	var s Student
	var t Teacher
	//s=t// this is wrong
	//change teacher to student
	s = Student(t)
	fmt.Println(s)
	//the datatype and the value(2 thing) store in teacher instance will go to student instance

	//accross package
}
