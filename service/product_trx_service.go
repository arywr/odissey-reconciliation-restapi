package service

import (
	"context"
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
)

type ProductTrxService interface {
	Create(ctx context.Context, request web.ProductTrxRequest) web.ProductTrxResponse
	CreateFromCSV(ctx context.Context, request *http.Request)
}
