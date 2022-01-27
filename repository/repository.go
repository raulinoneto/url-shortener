// Package repository provides the database resources and their rules
package repository

// Repository provides a main type to uses the database resources
type Repository struct {
	DB interface{}
	TableName string
}

// New initiate a Repository pointer
func New(db interface{}, tableName string) *Repository {
	return &Repository{
		DB: db,
		TableName: tableName,
	}
}
