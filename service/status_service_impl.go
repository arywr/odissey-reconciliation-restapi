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

type StatusServiceImpl struct {
	StatusRepository repository.StatusRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewStatusService(statusRepository repository.StatusRepository, DB *sql.DB, validate *validator.Validate) *StatusServiceImpl {
	return &StatusServiceImpl{
		StatusRepository: statusRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *StatusServiceImpl) Create(ctx context.Context, request web.TypeCreateRequest) web.StatusResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transactionStatus := entity.Status{
		Name:        request.Name,
		Description: request.Description,
	}

	transactionStatus = service.StatusRepository.Save(ctx, tx, transactionStatus)

	return helper.ToStatusResponse(transactionStatus)
}

func (service *StatusServiceImpl) Update(ctx context.Context, request web.TypeUpdateRequest) web.StatusResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	transactionStatus, err := service.StatusRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.Name != "" {
		transactionStatus.Name = request.Name
	}

	if request.Description != "" {
		transactionStatus.Description = request.Description
	}

	defer helper.CommitOrRollback(tx)

	transactionStatus = service.StatusRepository.Update(ctx, tx, transactionStatus)

	return helper.ToStatusResponse(transactionStatus)
}

func (service *StatusServiceImpl) Delete(ctx context.Context, serviceId int) web.StatusResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	transactionStatus, err := service.StatusRepository.FindById(ctx, tx, serviceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	defer helper.CommitOrRollback(tx)

	transactionStatus = service.StatusRepository.Delete(ctx, tx, transactionStatus)

	return helper.ToStatusResponse(transactionStatus)
}

func (service *StatusServiceImpl) FindById(ctx context.Context, serviceId int) web.StatusResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	status, err := service.StatusRepository.FindById(ctx, tx, serviceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStatusResponse(status)
}

func (service *StatusServiceImpl) FindAll(ctx context.Context) []web.StatusResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	status := service.StatusRepository.FindAll(ctx, tx)

	var statusResponse []web.StatusResponse

	for _, row := range status {
		statusResponse = append(statusResponse, helper.ToStatusResponse(row))
	}

	return statusResponse
}
