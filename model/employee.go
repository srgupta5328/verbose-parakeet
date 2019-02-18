package model

import (
	"database/sql"
	"errors"
)

type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        string `json:"id"`
	Email     string `json:"email"`
	RoleName  string `json:"role_name"`
}

func (e *Employee) CreateEmployee(db *sql.DB) error {
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
