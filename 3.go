package main

import "fmt"

type emp interface{
	ft()
	cont()
	fl()
}

type salary struct{
	ftBasic int
	ctBasic int
	flBasic int
}

func (s salary) ft(){
	var salaryFt int = s.ftBasic * 28
	fmt.Printf("Monthly salary of Full time Employee: %v\n", salaryFt)
}

func (s salary) cont(){
	var salaryCont int = s.ctBasic * 28
	fmt.Printf("Monthly salary of Contract Employee: %v\n", salaryCont)
}

func (s salary) fl(){
	var salaryFl int = s.flBasic * 20
	fmt.Printf("Salary of freelancer if he works for 20 hours: %v\n", salaryFl)
}

func main() {
	var e emp = salary{500,100,10}
	e.ft()
	e.fl()
	e.cont()
}
