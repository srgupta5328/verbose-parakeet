package model

import (
	"database/sql"
	"errors"
)

type Employee struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	RoleName  string `json:"role_name,omitempty"`
}

func (e *Employee) createEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) updateEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) deleteEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) getEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}
