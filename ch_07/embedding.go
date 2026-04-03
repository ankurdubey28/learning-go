package main

import "fmt"

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	m.Description()          // method of employee accessible to manager instance directly
	m.Employee.Description() // same as above in this example , as manager does not have any method named Description .
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() {
	fmt.Println(e.Name + "_" + e.ID)
}

type Manager struct {
	Employee // embedded field
	Reports  []Employee
}
