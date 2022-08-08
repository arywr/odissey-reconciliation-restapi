package service

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/entity"
	"odissey-golang/odissey-reconciliation-restapi/model/repository"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductTrxServiceImpl struct {
	Repository repository.ProductTrxRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewProductTrxService(repository repository.ProductTrxRepository, DB *sql.DB, validate *validator.Validate) *ProductTrxServiceImpl {
	return &ProductTrxServiceImpl{
		Repository: repository,
		DB:         DB,
		Validate:   validate,
	}
}

func (service *ProductTrxServiceImpl) Create(ctx context.Context, request web.ProductTrxRequest) web.ProductTrxResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transaction := entity.ProductTransaction{
		OwnerId:               request.OwnerId,
		TransactionId:         request.TransactionId,
		TransactionCode:       request.TransactionCode,
		TransactionStatus:     request.TransactionStatus,
		TransactionKey:        request.TransactionKey,
		TransactionDate:       request.TransactionDate,
		TransactionDateTime:   request.TransactionDateTime,
		TransactionTypeId:     request.TransactionTypeId,
		TelkomTransactionId:   request.TelkomTransactionId,
		MerchantTransactionId: request.MerchantTransactionId,
		ChannelTransactionId:  request.ChannelTransactionId,
		ProductCode:           request.ProductCode,
		ProductName:           request.ProductName,
		MerchantCode:          request.MerchantCode,
		MerchantName:          request.MerchantName,
		ChannelCode:           request.ChannelCode,
		ChannelName:           request.ChannelName,
		Nominal:               request.Nominal,
	}

	transaction = service.Repository.Save(ctx, tx, transaction)

	return helper.ToProductTrxResponse(transaction)
}

func (service *ProductTrxServiceImpl) CreateFromCSV(ctx context.Context, request *http.Request) {
	var counterTotal int
	current := time.Now().Format("2006-01-02 15:04:05")

	createRequest := web.ProductTrxCreateExcel{}
	helper.ReadFromFormData(request, &createRequest)

	file, err := helper.UploadFile(request)
	helper.PanicIfError(err)

	csvReader, csvFile, err := helper.ReadCsvFile(createRequest, file)
	helper.PanicIfError(err)

	defer csvFile.Close()

	isHeader := true
	var store []entity.ProductTransaction

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if isHeader {
			isHeader = false
			continue
		}

		var amount float64

		if floatAmount, err := strconv.ParseFloat(row[11], 64); err == nil {
			amount = floatAmount
		}

		transaction := entity.ProductTransaction{
			OwnerId:               createRequest.OwnerId,
			TransactionId:         row[4],
			TransactionCode:       "acq",
			TransactionStatus:     1,
			TransactionKey:        row[4],
			TransactionDate:       row[2],
			TransactionDateTime:   row[2] + " " + row[3],
			TransactionTypeId:     1,
			TelkomTransactionId:   row[4],
			MerchantTransactionId: "",
			ChannelTransactionId:  "",
			ProductCode:           "",
			ProductName:           "",
			MerchantCode:          row[10],
			MerchantName:          row[15],
			ChannelCode:           "",
			ChannelName:           "",
			Nominal:               amount,
			CreatedAt:             current,
			UpdatedAt:             current,
		}

		counterTotal++
		store = append(store, transaction)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	log.Println(createRequest.OwnerId)

	progress := entity.Progress{
		ProgressName:    fmt.Sprintf("Upload: %s", createRequest.OwnerId),
		ProgressEventId: 1,
		File:            file,
		Status:          "on process",
		Percentage:      0,
		CreatedAt:       current,
		UpdatedAt:       current,
		DeletedAt:       nil,
	}

	progressResult := service.Repository.SaveProgress(ctx, tx, progress)

	jobs := generateIndex(store)

	worker := runtime.NumCPU()
	result := TestingInsert(service, ctx, jobs, worker)

	counterSuccess := 0
	for res := range result {
		if res.Id == 0 {
			log.Println("Has Error")
		} else {
			counterSuccess++
		}

		if counterSuccess%100 == 0 {
			// Updating Current Progress
			progressResult.Percentage = float64(counterSuccess) / float64(counterTotal) * 100
			service.Repository.UpdateProgress(ctx, tx, progressResult)
		}
	}

	helper.DestroyFile(file)

	progressResult.Percentage = 100
	progressResult.Status = "completed"
	service.Repository.UpdateProgress(ctx, tx, progressResult)
}

func generateIndex(store []entity.ProductTransaction) <-chan entity.ProductTransaction {
	result := make(chan entity.ProductTransaction)

	go func() {
		for _, transaction := range store {
			result <- transaction
		}

		close(result)
	}()

	return result
}

func TestingInsert(
	service *ProductTrxServiceImpl,
	ctx context.Context,
	jobs <-chan entity.ProductTransaction,
	worker int,
) <-chan entity.ProductTransaction {
	result := make(chan entity.ProductTransaction)

	wg := new(sync.WaitGroup)
	wg.Add(worker)

	go func() {
		for i := 0; i < worker; i++ {
			go func() {
				for job := range jobs {
					response := InsertFromExcel(service, job, ctx)
					result <- response
				}
				wg.Done()
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func InsertFromExcel(
	service *ProductTrxServiceImpl,
	transaction entity.ProductTransaction,
	ctx context.Context,
) entity.ProductTransaction {
	var response entity.ProductTransaction

	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
					helper.PanicIfError(*outerError)
				}
			}()

			tx, err := service.DB.Begin()
			helper.PanicIfError(err)

			defer helper.CommitOrRollback(tx)

			response = service.Repository.Save(ctx, tx, transaction)
		}(&outerError)
		if outerError == nil {
			break
		}
	}

	return response

}
