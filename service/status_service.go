package service

import (
	"context"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
)

type StatusService interface {
	Create(ctx context.Context, request web.TypeCreateRequest) web.StatusResponse
	Update(ctx context.Context, request web.TypeUpdateRequest) web.StatusResponse
	Delete(ctx context.Context, statusId int) web.StatusResponse
	FindById(ctx context.Context, statusId int) web.StatusResponse
	FindAll(ctx context.Context) []web.StatusResponse
}
