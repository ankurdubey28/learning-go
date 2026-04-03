package main

func main() {
	firstName := "ankur"
	lastName := "dwivedi"
	age := 22
	makePerson(firstName, lastName, age)
	makePersonPointer(firstName, lastName, age)

}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson(firstName, secondName string, age int) Person {
	return Person{
		firstName,
		secondName,
		age,
	}
}

func makePersonPointer(firstName, secondName string, age int) *Person {
	return &Person{
		firstName,
		secondName,
		age,
	}
}
