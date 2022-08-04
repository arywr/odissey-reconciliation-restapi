package main

import (
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/app"
	"odissey-golang/odissey-reconciliation-restapi/controller"
	"odissey-golang/odissey-reconciliation-restapi/exception"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/repository"
	"odissey-golang/odissey-reconciliation-restapi/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	typeRepository := repository.NewTypeRepository()
	typeService := service.NewTypeService(typeRepository, db, validate)
	typeController := controller.NewTypeController(typeService)

	statusRepository := repository.NewStatusRepository()
	statusService := service.NewStatusService(statusRepository, db, validate)
	statusController := controller.NewStatusController(statusService)

	productTrxRepository := repository.NewProductTrxRepository()
	productTrxService := service.NewProductTrxService(productTrxRepository, db, validate)
	productTrxControler := controller.NewProductTrxController(productTrxService)

	router := httprouter.New()

	// API for Transaction Types
	// Path: transaction-types
	router.GET("/api/transaction-types", typeController.FindAll)
	router.GET("/api/transaction-types/:id", typeController.FindById)
	router.POST("/api/transaction-types", typeController.Create)
	router.PUT("/api/transaction-types/:id", typeController.Update)
	router.DELETE("/api/transaction-types/:id", typeController.Delete)

	// API for Transaction Statuses
	// Path: transaction-statuses
	router.GET("/api/transaction-statuses", statusController.FindAll)
	router.GET("/api/transaction-statuses/:id", statusController.FindById)
	router.POST("/api/transaction-statuses", statusController.Create)
	router.PUT("/api/transaction-statuses/:id", statusController.Update)
	router.DELETE("/api/transaction-statuses/:id", statusController.Delete)

	// API for Product Transactions
	// Path: transaction/product
	router.POST("/api/transaction/products", productTrxControler.Create)
	router.POST("/api/transaction/products-from-csv", productTrxControler.CreateFromCSV)

	router.PanicHandler = exception.ErrorHandler

	handler := cors.Default().Handler(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
