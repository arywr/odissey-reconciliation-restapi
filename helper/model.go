package helper

import (
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
)

func ToTypeResponse(types entity.Type) web.TypeResponse {
	return web.TypeResponse{
		Id:          types.Id,
		Name:        types.Name,
		Description: types.Description,
	}
}

func ToStatusResponse(status entity.Status) web.StatusResponse {
	return web.StatusResponse{
		Id:          status.Id,
		Name:        status.Name,
		Description: status.Description,
	}
}

func ToProductTrxResponse(transaction entity.ProductTransaction) web.ProductTrxResponse {
	return web.ProductTrxResponse{
		Id:                    transaction.Id,
		OwnerId:               transaction.OwnerId,
		TransactionId:         transaction.TransactionId,
		TransactionCode:       transaction.TransactionCode,
		TransactionStatus:     transaction.TransactionStatus,
		TransactionKey:        transaction.TransactionKey,
		TransactionDate:       transaction.TransactionDate,
		TransactionDateTime:   transaction.TransactionDateTime,
		TransactionTypeId:     transaction.TransactionTypeId,
		TelkomTransactionId:   transaction.TelkomTransactionId,
		MerchantTransactionId: transaction.MerchantTransactionId,
		ChannelTransactionId:  transaction.ChannelTransactionId,
		ProductCode:           transaction.ProductCode,
		ProductName:           transaction.ProductName,
		MerchantCode:          transaction.MerchantCode,
		MerchantName:          transaction.MerchantName,
		ChannelCode:           transaction.ChannelCode,
		ChannelName:           transaction.ChannelName,
		Nominal:               transaction.Nominal,
		CreatedAt:             transaction.CreatedAt,
		UpdatedAt:             transaction.UpdatedAt,
		DeletedAt:             transaction.DeletedAt,
	}
}
