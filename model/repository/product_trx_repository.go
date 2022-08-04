package repository

import (
	"context"
	"database/sql"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
)

type ProductTrxRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaction entity.ProductTransaction) entity.ProductTransaction
	SaveEmptyReturn(ctx context.Context, db *sql.Tx, transaction entity.ProductTransaction)
}
