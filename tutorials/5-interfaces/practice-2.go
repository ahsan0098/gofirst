package main

import "fmt"

func employeeReport() {
	// Create a contractor
	contractor1 := contractor{
		name:         "Bob",
		hourlyPay:    50,
		hoursPerYear: 2000,
	}

	// Create a full-time employee
	fullTimeEmployee := fullTime{
		name:   "Alice",
		salary: 60000,
	}

	// Print reports
	printEmployeeReport(contractor1)
	printEmployeeReport(fullTimeEmployee)
}

func printEmployeeReport(e employee) {
	fmt.Printf("%s earns %d per year.\n", e.getName(), e.getSalary())
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}
