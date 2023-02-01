package models

import (
	"context"
)

// GeneralRepository -
type GeneralRepository interface {
	CreateTables() error
	DeleteByContract(indices []string, address string) error
	GetByID(output Model) error
	GetAll(index string) ([]Model, error)
	UpdateDoc(model Model) (err error)
	IsRecordNotFound(err error) bool

	// Save - performs insert or update items.
	Save(ctx context.Context, items []Model) error
	BulkDelete(context.Context, []Model) error

	TablesExist() bool
	// Drop - drops full database
	Drop(ctx context.Context) error
}
