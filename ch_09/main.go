package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	err := makeTea(2)
	// unwrap
	wrapped := errors.Unwrap(err)
	fmt.Println(wrapped)

	// is : recursively call the unwrap method until either target error is found or nil is found
	baseErr := io.EOF
	wrapped = fmt.Errorf("reading file: %w", baseErr)
	fmt.Println(errors.Is(wrapped, io.EOF))

	// as : recursively call the unwrap to either find the target error or nil and assign it to the variable.
	base := &os.PathError{
		Op:   "open",
		Path: "/tmp/data.txt",
		Err:  errors.New("permission Denied"),
	}
	// wrap it
	err = fmt.Errorf("failed to read config: %w", base)
	// target need to be passed as a pointer
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("operation:", pathErr.Op)
		fmt.Println("file:", pathErr.Path)
	}

	// recover
	for _, v := range []int{1, 2, 0, 4, 5} {
		div(v)
	}
	// panic
	panic(nil)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero error")
	}
	return b / a, nil
}

var ErrOutOfTea = errors.New("no more tea available")

func makeTea(arg int) error {
	if arg == 2 {
		return fmt.Errorf("making tea: %w", ErrOutOfTea)
	}
	return nil
}

// wrap multiple errors: using errors.Join()
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ValidatePerson(p Person) error {
	var errs []error

	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("firstname cannot be empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("lastname cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("age cannot be nagavtive"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// panic and recover example

func div(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i) // illegal operation if i==0
} // go runtime will internally call panic which is caught by recover , reason why this works despite no explicit panic.
