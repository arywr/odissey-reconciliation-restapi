package controller

import (
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
	"odissey-golang/odissey-reconciliation-restapi/service"

	"github.com/julienschmidt/httprouter"
)

type ProductTrxControllerImpl struct {
	ProductTrxService service.ProductTrxService
}

func NewProductTrxController(service service.ProductTrxService) ProductTrxController {
	return &ProductTrxControllerImpl{
		ProductTrxService: service,
	}
}

func (controller *ProductTrxControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createRequest := web.ProductTrxRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	response := controller.ProductTrxService.Create(request.Context(), createRequest)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductTrxControllerImpl) CreateFromCSV(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.ProductTrxService.CreateFromCSV(request.Context(), request)

	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
