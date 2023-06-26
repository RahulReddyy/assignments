package main

import (
	"fmt"
)

type employee struct {
	id     int
	name   string
	salary int
}

func main() {
	slice := []*employee{}

	ch := make(chan *employee)
	done := make(chan bool)
	NumOfWorkers := 3

	emp1 := &employee{id: 1, name: "employee1", salary: 1000}
	emp2 := &employee{id: 2, name: "employee2", salary: 2000}
	emp3 := &employee{id: 3, name: "employee3", salary: 3000}
	emp4 := &employee{id: 4, name: "employee4", salary: 4000}
	emp5 := &employee{id: 5, name: "employee5", salary: 5000}

	slice = append(slice, emp1, emp2, emp3, emp4, emp5)

	for i := 0; i < NumOfWorkers; i++ {
		go changeDetails(ch, done)
	}

	for index, value := range slice {
		ch <- slice[index]
		<-done
		fmt.Printf("salary of employee with name - %s is %d", value.name, value.salary)
		fmt.Println()
	}

	close(ch)
	close(done)

}

func changeDetails(channel chan *employee, done chan bool) {
	for emp := range channel {
		change := float64(emp.salary) * 3.4
		emp.salary += int(change)
		done <- true
	}
}
