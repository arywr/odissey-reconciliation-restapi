package repository

import (
	"context"
	"database/sql"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
)

type TypeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, types entity.Type) entity.Type
	Update(ctx context.Context, tx *sql.Tx, types entity.Type) entity.Type
	Delete(ctx context.Context, tx *sql.Tx, types entity.Type) entity.Type
	FindById(ctx context.Context, tx *sql.Tx, typeId int) (entity.Type, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Type
}
