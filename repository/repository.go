// Package repository provides the database resources and their rules
package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	InsertOne(ctx context.Context, b bson.D) mongo.InsertOneResult
}
// Repository provides a main type to uses the database resources
type Repository struct {
	db DB
	tableName string
}

// New initiate a Repository pointer
func New(db DB, tableName string) *Repository {
	return &Repository{
		db: db,
		tableName: tableName,
	}
}

func (r Repository) Get(ctx context.Context, id string) (string, error) {


	panic("implement me")
}

func (r Repository) Shorten(ctx context.Context, u string) (string, error) {
	panic("implement me")
}

func (r Repository) List(ctx context.Context) ([]string, error) {
	panic("implement me")
}
