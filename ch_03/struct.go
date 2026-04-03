package main

import "fmt"

func main() {
	var ankur person // zero value struct , each field set to its zero value.
	fmt.Println(ankur)
	fmt.Println(ankur.name)
	fmt.Println(ankur.age)

	pet := struct {
		name string
		kind string
	}{
		name: "sheru",
		kind: "golden retriever",
	}
	fmt.Println(pet.name)

	emp1 := Employee{
		firstname: "a",
		lastname:  "b",
		id:        0,
	}
	emp3 := Employee{
		"x",
		"y",
		1,
	}
	var emp2 Employee
	emp2.firstname = "one"
	emp2.lastname = "plus"
	emp2.id = 2
	fmt.Println(emp1, emp2, emp3)
}

type person struct {
	name string
	age  int
}
type Employee struct {
	firstname string
	lastname  string
	id        int
}
