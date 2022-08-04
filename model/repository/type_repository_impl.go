package repository

import (
	"context"
	"database/sql"
	"errors"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
)

type TypeRepositoryImpl struct {
}

func NewTypeRepository() TypeRepository {
	return &TypeRepositoryImpl{}
}

func (repository *TypeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, types entity.Type) entity.Type {
	query := "INSERT INTO od_transaction_types(transaction_type_name, transaction_type_description) VALUES(?,?)"

	result, err := tx.ExecContext(ctx, query, types.Name, types.Description)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	types.Id = int(id)
	return types
}

func (repository *TypeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, types entity.Type) entity.Type {
	query := `
		UPDATE od_transaction_types 
		SET transaction_type_name = ?, transaction_type_description = ? 
		WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, types.Name, types.Description, types.Id)

	helper.PanicIfError(err)

	return types
}

func (repository *TypeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, types entity.Type) entity.Type {
	query := `
		DELETE FROM od_transaction_types 
		WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, types.Id)

	helper.PanicIfError(err)

	return types
}

func (repository *TypeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, typeId int) (entity.Type, error) {
	query := `
		SELECT id, transaction_type_name, transaction_type_description
		FROM od_transaction_types
		WHERE id = ?
		LIMIT 1
	`

	rows, err := tx.QueryContext(ctx, query, typeId)
	helper.PanicIfError(err)

	defer rows.Close()

	types := entity.Type{}
	if rows.Next() {
		err := rows.Scan(&types.Id, &types.Name, &types.Description)
		helper.PanicIfError(err)

		return types, nil
	} else {
		return types, errors.New("transaction type is not found")
	}
}

func (repository *TypeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Type {
	query := `
		SELECT id, transaction_type_name, transaction_type_description
		FROM od_transaction_types
	`

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	defer rows.Close()

	var types []entity.Type

	for rows.Next() {
		data := entity.Type{}
		err := rows.Scan(&data.Id, &data.Name, &data.Description)
		helper.PanicIfError(err)
		types = append(types, data)
	}

	return types
}
