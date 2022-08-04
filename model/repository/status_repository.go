package repository

import (
	"context"
	"database/sql"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
)

type StatusRepository interface {
	Save(ctx context.Context, tx *sql.Tx, status entity.Status) entity.Status
	Update(ctx context.Context, tx *sql.Tx, status entity.Status) entity.Status
	Delete(ctx context.Context, tx *sql.Tx, status entity.Status) entity.Status
	FindById(ctx context.Context, tx *sql.Tx, statusId int) (entity.Status, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Status
}
