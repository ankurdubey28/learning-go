package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAge = errors.New("invalid age")

func main() {
	e := Employee{
		Name: "ankur",
		Age:  -1,
	}
	_, err := ValidateEmp(&e)
	// unwrap
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}

	// errors.As
	var me MyError
	if ok := errors.As(err, &me); ok == true {
		fmt.Println(me.Field)
		fmt.Println(me.Message)
	}

	// errors.Is + sentinel errors
	if errors.Is(err, ErrInvalidAge) {
		fmt.Println(err)
	}

	// panic error boundary
	err = SafeValidate(nil)
	if err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	Field   string
	Message string
}

func (m MyError) Error() string {
	return fmt.Sprintf("Validation failed for %s : %s", m.Field, m.Message)
}

type Employee struct {
	Name string
	Age  int
}

func ValidateEmp(e *Employee) (bool, error) {
	var errs []error
	if e == nil {
		panic("employee is nil") // manually panic
	}
	if e.Name == "" {
		errs = append(errs, fmt.Errorf("employee validation failed :%w",
			MyError{
				Field:   "Name",
				Message: "empty name now allowed",
			}))
	}
	if e.Age < 0 {
		errs = append(errs, fmt.Errorf("employee validatoin failed: %w",
			ErrInvalidAge))
	}

	if len(errs) > 0 {
		return false, errors.Join(errs...)
	}
	return true, nil
}

func SafeValidate(e *Employee) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered : %v", r)
		}
	}()
	_, err = ValidateEmp(e)
	return
}
