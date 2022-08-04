package web

type ProductTrxRequest struct {
	OwnerId               string  `json:"owner_id"`
	TransactionId         string  `json:"transaction_id"`
	TransactionCode       string  `json:"transaction_code"`
	TransactionStatus     int     `json:"transaction_status"`
	TransactionKey        string  `json:"transaction_key"`
	TransactionDate       string  `json:"transaction_date"`
	TransactionDateTime   string  `json:"transaction_datetime"`
	TransactionTypeId     int     `json:"transaction_type_id"`
	TelkomTransactionId   string  `json:"telkom_transaction_id"`
	MerchantTransactionId string  `json:"merchant_transaction_id"`
	ChannelTransactionId  string  `json:"channel_transaction_id"`
	ProductCode           string  `json:"product_code"`
	ProductName           string  `json:"product_name"`
	MerchantName          string  `json:"merchant_name"`
	MerchantCode          string  `json:"merchant_code"`
	ChannelCode           string  `json:"channel_code"`
	ChannelName           string  `json:"channel_name"`
	Nominal               float64 `json:"nominal"`
}
