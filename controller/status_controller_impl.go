package controller

import (
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
	"odissey-golang/odissey-reconciliation-restapi/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type StatusControllerImpl struct {
	StatusService service.StatusService
}

func NewStatusController(statusService service.StatusService) TypeController {
	return &StatusControllerImpl{
		StatusService: statusService,
	}
}

func (controller *StatusControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeCreateRequest := web.TypeCreateRequest{}
	helper.ReadFromRequestBody(request, &typeCreateRequest)

	response := controller.StatusService.Create(request.Context(), typeCreateRequest)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeUpdateRequest := web.TypeUpdateRequest{}
	helper.ReadFromRequestBody(request, &typeUpdateRequest)

	typeId := params.ByName("id")
	id, err := strconv.Atoi(typeId)
	helper.PanicIfError(err)

	typeUpdateRequest.Id = id

	response := controller.StatusService.Update(request.Context(), typeUpdateRequest)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeId := params.ByName("id")
	id, err := strconv.Atoi(typeId)
	helper.PanicIfError(err)

	response := controller.StatusService.Delete(request.Context(), id)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeId := params.ByName("id")
	id, err := strconv.Atoi(typeId)
	helper.PanicIfError(err)

	response := controller.StatusService.FindById(request.Context(), id)

	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StatusControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.StatusService.FindAll(request.Context())

	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
