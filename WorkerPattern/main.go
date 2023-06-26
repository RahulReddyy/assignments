package main

import (
	"fmt"
)

type employee struct {
	id     int
	name   string
	salary int
}

func (e employee) changeDetails(channel chan *employee) {
	for emp := range channel {
		emp.salary = emp.salary + 103
	}
}

func main() {
	slice := []*employee{}

	NumOfWorkers := 3
	ch := make(chan *employee)

	empcheck := employee{}
	emp1 := &employee{id: 1, name: "employee1", salary: 1000}
	emp2 := &employee{id: 2, name: "employee2", salary: 2000}
	emp3 := &employee{id: 3, name: "employee3", salary: 3000}
	emp4 := &employee{id: 4, name: "employee4", salary: 4000}
	emp5 := &employee{id: 5, name: "employee5", salary: 5000}

	slice = append(slice, emp1, emp2, emp3, emp4, emp5)

	for i := 0; i < NumOfWorkers; i++ {
		go empcheck.changeDetails(ch)
	}

	for emloyee := range slice {
		ch <- slice[emloyee]

	}

	close(ch)

	for _, value := range slice {
		fmt.Printf("Updated salary of employee with name - %s is %d\n", value.name, value.salary)
	}
}
