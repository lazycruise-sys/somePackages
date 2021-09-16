package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {

	// declaration of s as a Subscriber type
	var s magazine.Subscriber
	s.Rate = 3.99
	fmt.Println(s.Rate)

	// initialization of subscriber
	subscriber := magazine.Subscriber{Name: "aman singh", Rate: 4.99, Active: true}
	fmt.Println("name:", subscriber.Name)
	fmt.Println("rate:", subscriber.Rate)

	// declaration of employee as a Employee type
	var employee magazine.Employee
	employee.Name = "joy carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
