package main

import (
	"fmt"
)

type employee struct {
	id     int
	name   string
	salary int
}

func changeDetails(channel chan *int, done chan bool) {
	for emp := range channel {
		*emp += int(float64(*emp) * 3.4)
		done <- true
	}
}

func main() {
	slice := []*employee{}

	ch := make(chan *int)
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

	for emloyee := range slice {
		ch <- &slice[emloyee].salary
		<-done
	}

	close(ch)
	close(done)

	for _, value := range slice {
		fmt.Printf("Updated salary of employee with name - %s is %d\n", value.name, value.salary)
	}
}
