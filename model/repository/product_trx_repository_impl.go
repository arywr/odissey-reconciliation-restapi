package repository

import (
	"context"
	"database/sql"
	"fmt"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
	"strings"
	"time"
)

type ProductTrxRepositoryImpl struct {
}

var dataHeaders = []string{
	"owner_id",
	"transaction_id",
	"transaction_code",
	"transaction_status",
	"transaction_key",
	"transaction_date",
	"transaction_datetime",
	"transaction_type_id",
	"telkom_transaction_id",
	"merchant_transaction_id",
	"channel_transaction_id",
	"product_code",
	"product_name",
	"merchant_code",
	"merchant_name",
	"channel_code",
	"channel_name",
	"nominal",
	"created_at",
	"updated_at",
	"deleted_at",
}

func NewProductTrxRepository() ProductTrxRepository {
	return &ProductTrxRepositoryImpl{}
}

func (repository *ProductTrxRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction entity.ProductTransaction) entity.ProductTransaction {
	query := fmt.Sprintf(`
		INSERT INTO od_product_transactions (%s)
		VALUES (%s)
	`, strings.Join(dataHeaders, ","), strings.Join(helper.GenerateQuestionsMark(len(dataHeaders)), ","))

	var current string
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	if transaction.CreatedAt != "" {
		current = transaction.CreatedAt
	} else {
		current = currentTime
	}

	result, err := tx.ExecContext(
		ctx,
		query,
		transaction.OwnerId, transaction.TransactionId, transaction.TransactionCode,
		transaction.TransactionStatus, transaction.TransactionKey, transaction.TransactionDate,
		transaction.TransactionDateTime, transaction.TransactionTypeId, transaction.TelkomTransactionId,
		transaction.MerchantTransactionId, transaction.ChannelTransactionId, transaction.ProductCode,
		transaction.ProductName, transaction.MerchantCode, transaction.MerchantName,
		transaction.ChannelCode, transaction.ChannelName, transaction.Nominal,
		current, current, nil,
	)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	transaction.Id = int(id)
	transaction.CreatedAt = current
	transaction.UpdatedAt = current
	return transaction
}

func (ProductTrxRepository *ProductTrxRepositoryImpl) SaveEmptyReturn(ctx context.Context, tx *sql.Tx, transaction entity.ProductTransaction) {
	query := fmt.Sprintf(`
		INSERT INTO od_product_transactions (%s)
		VALUES (%s)
	`, strings.Join(dataHeaders, ","), strings.Join(helper.GenerateQuestionsMark(len(dataHeaders)), ","))

	var current string
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	if transaction.CreatedAt != "" {
		current = transaction.CreatedAt
	} else {
		current = currentTime
	}

	_, err := tx.ExecContext(
		ctx,
		query,
		transaction.OwnerId, transaction.TransactionId, transaction.TransactionCode,
		transaction.TransactionStatus, transaction.TransactionKey, transaction.TransactionDate,
		transaction.TransactionDateTime, transaction.TransactionTypeId, transaction.TelkomTransactionId,
		transaction.MerchantTransactionId, transaction.ChannelTransactionId, transaction.ProductCode,
		transaction.ProductName, transaction.MerchantCode, transaction.MerchantName,
		transaction.ChannelCode, transaction.ChannelName, transaction.Nominal,
		current, current, nil,
	)

	helper.PanicIfError(err)
}
