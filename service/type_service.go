package service

import (
	"context"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
)

type TypeService interface {
	Create(ctx context.Context, request web.TypeCreateRequest) web.TypeResponse
	Update(ctx context.Context, request web.TypeUpdateRequest) web.TypeResponse
	Delete(ctx context.Context, typeId int) web.TypeResponse
	FindById(ctx context.Context, typeId int) web.TypeResponse
	FindAll(ctx context.Context) []web.TypeResponse
}
