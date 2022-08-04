package entity

type ProductTransaction struct {
	Id                    int
	OwnerId               string
	TransactionId         string
	TransactionCode       string
	TransactionStatus     int
	TransactionKey        string
	TransactionDate       string
	TransactionDateTime   string
	TransactionTypeId     int
	TelkomTransactionId   string
	MerchantTransactionId string
	ChannelTransactionId  string
	ProductCode           string
	ProductName           string
	MerchantName          string
	MerchantCode          string
	ChannelCode           string
	ChannelName           string
	Nominal               float64
	CreatedAt             string
	UpdatedAt             string
	DeletedAt             *string
}
