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

	// testing the declaration of Address type
	var address magazine.Address
	address.City = "Ile-Ife"
	address.PostalCode = "220005"
	address.State = "Osun"
	address.Street = "OAU SSQ"
	fmt.Println(address)

	// initializing Subscriber by adding existing Address
	subscriber2 := magazine.Subscriber{Name: "akinlua olorunfemi", Rate: 3.99, Active: true}
	subscriber2.Address = address
	fmt.Println(subscriber2)

	// initializing Subscriber by default
	subscriber3 := magazine.Subscriber{Name: "akinlua olorunfemi", Rate: 4.99, Active: true, Address: magazine.Address{Street: "lukman", City: "lagos"}}
	fmt.Println(subscriber3)

	// testing
	fmt.Println(subscriber3.City)

	// assigning new values to subscriber variable
	subscriber.City = "san francisco"
	fmt.Println(subscriber)
	subscriber.Address = magazine.Address{PostalCode: "68111"} // overwrites the previous line of assignment
	fmt.Println(subscriber.Address)

	// checking if you can add more than one anonymous struct fields
	employee.Address = magazine.Address{Street: "herbert macaulay street", City: "ikorodu", State: "lagos"}
	employee.School = "obafemi awolowo university"
	employee.Work = "microsoft nigeria"
	fmt.Println(employee)

	// second employee to test knowledge
	employee2 := magazine.Employee{Name: "jacob arigbede", Salary: 123000, Address: magazine.Address{Street: "102 MFM Main One Zone", City: "Majodo", State: "lagos", PostalCode: "22001"}, History: magazine.History{School: "OAU", Work: "Turing"}}
	fmt.Println(employee2)

	// printing embedded structs
	jacobState := employee2.State
	jacobSchool := employee2.School
	fmt.Println("State and School: ", jacobState, jacobSchool)

}
