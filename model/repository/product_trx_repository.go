package repository

import (
	"context"
	"database/sql"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
)

type ProductTrxRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaction entity.ProductTransaction) entity.ProductTransaction
	SaveProgress(ctx context.Context, tx *sql.Tx, progress entity.Progress) entity.Progress
	UpdateProgress(ctx context.Context, tx *sql.Tx, progress entity.Progress) entity.Progress
}
