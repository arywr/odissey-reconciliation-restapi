package repository

import (
	"context"
	"database/sql"
	"errors"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
)

type StatusRepositoryImpl struct {
}

func NewStatusRepository() StatusRepository {
	return &StatusRepositoryImpl{}
}

func (repository *StatusRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, types entity.Status) entity.Status {
	query := "INSERT INTO od_transaction_statuses(status_description, status_description_name) VALUES(?,?)"

	result, err := tx.ExecContext(ctx, query, types.Description, types.Name)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	types.Id = int(id)
	return types
}

func (repository *StatusRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, types entity.Status) entity.Status {
	query := `
		UPDATE od_transaction_statuses 
		SET status_description = ?, status_description_name = ? 
		WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, types.Name, types.Description, types.Id)

	helper.PanicIfError(err)

	return types
}

func (repository *StatusRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, types entity.Status) entity.Status {
	query := `
		DELETE FROM od_transaction_statuses 
		WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, types.Id)

	helper.PanicIfError(err)

	return types
}

func (repository *StatusRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, statusId int) (entity.Status, error) {
	query := `
		SELECT id, status_description, status_description_name
		FROM od_transaction_statuses
		WHERE id = ?
		LIMIT 1
	`

	rows, err := tx.QueryContext(ctx, query, statusId)
	helper.PanicIfError(err)

	defer rows.Close()

	statuses := entity.Status{}
	if rows.Next() {
		err := rows.Scan(&statuses.Id, &statuses.Description, &statuses.Name)
		helper.PanicIfError(err)

		return statuses, nil
	} else {
		return statuses, errors.New("transaction status is not found")
	}
}

func (repository *StatusRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Status {
	query := `
		SELECT id, status_description, status_description_name
		FROM od_transaction_statuses
	`

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	defer rows.Close()

	var statuses []entity.Status

	for rows.Next() {
		status := entity.Status{}
		err := rows.Scan(&status.Id, &status.Description, &status.Name)
		helper.PanicIfError(err)
		statuses = append(statuses, status)
	}

	return statuses
}
