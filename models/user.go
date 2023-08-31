package models

type User struct {
	Id           int `pg:"type:serial, pk"`
	FullName     string
	EmployeeCode int
	Email        string
	Role         string
}
