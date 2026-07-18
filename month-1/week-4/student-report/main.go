package main

import (
	"fmt"
)

type Student struct {
	Name       string `json:"name"`
	Class      string `json:"class"`
	RollNumber int    `json:"rollnumber"`
	Adderess
}

type Adderess struct {
	AddressLineOne string `json:"addLineOne"`
	Pincode        int    `json:"piccode"`
}

func (stud *Student) UpdateStudentDetail() {
	stud.Name = "John Doe"
	stud.Class = "Class 4"
	stud.RollNumber = 23
	stud.AddressLineOne = "12 Main 14 Cross"
	stud.Pincode = 100102
}

func main() {
	stu := Student{}
	stu.UpdateStudentDetail()
	fmt.Println("Student Data: ", stu)
}
