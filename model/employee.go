package model

import (
	"errors"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Employee struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	RoleName  string `json:"role_name,omitempty"`
}

func (e *Employee) CreateEmployee(db *scribble.Driver) error {
	employee := Employee{FirstName: "Rohan", LastName: "Gupta", ID: "1", Email: "TestEmail@gmail.dom", RoleName: "Software Engineer 1"}
	err := db.Write("empolyees", employee.FirstName, Employee{FirstName: employee.FirstName})
	if err != nil {
		return errors.New("Not implemented yet")
	}
	return nil

}

func (e *Employee) updateEmployee(db *scribble.Driver) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) deleteEmployee(db *scribble.Driver) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) getEmployee(db *scribble.Driver) error {
	return errors.New("Not implemented yet")
}
