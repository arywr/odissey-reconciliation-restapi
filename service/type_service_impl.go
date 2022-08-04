package service

import (
	"context"
	"database/sql"
	"odissey-golang/odissey-reconciliation-restapi/exception"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
	"odissey-golang/odissey-reconciliation-restapi/model/repository"
	"odissey-golang/odissey-reconciliation-restapi/model/web"

	"github.com/go-playground/validator/v10"
)

type TypeServiceImpl struct {
	TypeRepository repository.TypeRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTypeService(typeRepository repository.TypeRepository, DB *sql.DB, validate *validator.Validate) *TypeServiceImpl {
	return &TypeServiceImpl{
		TypeRepository: typeRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TypeServiceImpl) Create(ctx context.Context, request web.TypeCreateRequest) web.TypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transactionType := entity.Type{
		Name:        request.Name,
		Description: request.Description,
	}

	transactionType = service.TypeRepository.Save(ctx, tx, transactionType)

	return helper.ToTypeResponse(transactionType)
}

func (service *TypeServiceImpl) Update(ctx context.Context, request web.TypeUpdateRequest) web.TypeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	transactionType, err := service.TypeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.Name != "" {
		transactionType.Name = request.Name
	}

	if request.Description != "" {
		transactionType.Description = request.Description
	}

	defer helper.CommitOrRollback(tx)

	transactionType = service.TypeRepository.Update(ctx, tx, transactionType)

	return helper.ToTypeResponse(transactionType)
}

func (service *TypeServiceImpl) Delete(ctx context.Context, typeId int) web.TypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	transactionType, err := service.TypeRepository.FindById(ctx, tx, typeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	defer helper.CommitOrRollback(tx)

	transactionType = service.TypeRepository.Delete(ctx, tx, transactionType)

	return helper.ToTypeResponse(transactionType)
}

func (service *TypeServiceImpl) FindById(ctx context.Context, typeId int) web.TypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	types, err := service.TypeRepository.FindById(ctx, tx, typeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTypeResponse(types)
}

func (service *TypeServiceImpl) FindAll(ctx context.Context) []web.TypeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	types := service.TypeRepository.FindAll(ctx, tx)

	var typeResponse []web.TypeResponse
	for _, row := range types {
		typeResponse = append(typeResponse, helper.ToTypeResponse(row))
	}

	return typeResponse
}
