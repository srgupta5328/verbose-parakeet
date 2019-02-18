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

func (e *Employee) createEmployee(db *scribble.Driver) error {
	return errors.New("Not implemented yet")
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
