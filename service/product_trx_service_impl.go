package service

import (
	"context"
	"database/sql"
	"fmt"
	"io"
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
	current := time.Now().Format("2006-01-02 15:04:05")

	createRequest := web.ProductTrxCreateExcel{}
	helper.ReadFromFormData(request, &createRequest)

	file, err := helper.UploadFile(request)
	helper.PanicIfError(err)

	csvReader, csvFile, err := helper.ReadCsvFile(createRequest, file)
	helper.PanicIfError(err)

	defer csvFile.Close()

	jobs := make(chan entity.ProductTransaction, runtime.NumCPU())
	wg := new(sync.WaitGroup)

	go dispatchWorkers(service, jobs, wg, ctx)

	isHeader := true

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

		if s, err := strconv.ParseFloat(row[11], 64); err == nil {
			amount = s
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

		wg.Add(1)
		jobs <- transaction
	}
	close(jobs)
	wg.Wait()

	helper.DestroyFile(file)
}

func dispatchWorkers(service *ProductTrxServiceImpl, jobs <-chan entity.ProductTransaction, wg *sync.WaitGroup, ctx context.Context) {
	for workerIndex := 0; workerIndex <= runtime.NumCPU(); workerIndex++ {
		go func(workerIndex int, jobs <-chan entity.ProductTransaction, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJob(workerIndex, counter, service, job, ctx)
				wg.Done()
				counter++
			}
		}(workerIndex, jobs, wg)
	}
}

func doTheJob(workerIndex, counter int, service *ProductTrxServiceImpl, transaction entity.ProductTransaction, ctx context.Context) {
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

			service.Repository.Save(ctx, tx, transaction)
		}(&outerError)
		if outerError == nil {
			break
		}
	}
}
