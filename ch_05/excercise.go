package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	n, err := fileLen("C:\\Users\\ASUS\\GolangProjects\\learning-go\\ch_05\\excercise.go")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(n)

	helloPrefix := prefixer("hello")
	fmt.Println(helloPrefix("bob"))
}

// function to read number of bytes in file
func fileLen(fname string) (int, error) {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		return 0, errors.New("error Opening File")
	}
	b := make([]byte, 1024)
	size, err := file.Read(b)
	if err != nil {
		return 0, errors.New("error Reading File")
	}
	return size, nil
}

// prefixer function
func prefixer(str string) func(string) string {
	return func(s string) string {
		return str + " " + s + "!"
	}
}
