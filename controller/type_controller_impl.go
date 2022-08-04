package controller

import (
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
	"odissey-golang/odissey-reconciliation-restapi/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TypeControllerImpl struct {
	TypeService service.TypeService
}

func NewTypeController(typeService service.TypeService) TypeController {
	return &TypeControllerImpl{
		TypeService: typeService,
	}
}

func (controller *TypeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeCreateRequest := web.TypeCreateRequest{}
	helper.ReadFromRequestBody(request, &typeCreateRequest)

	response := controller.TypeService.Create(request.Context(), typeCreateRequest)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TypeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeUpdateRequest := web.TypeUpdateRequest{}
	helper.ReadFromRequestBody(request, &typeUpdateRequest)

	typeId := params.ByName("id")
	id, err := strconv.Atoi(typeId)
	helper.PanicIfError(err)

	typeUpdateRequest.Id = id

	response := controller.TypeService.Update(request.Context(), typeUpdateRequest)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TypeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeId := params.ByName("id")
	id, err := strconv.Atoi(typeId)
	helper.PanicIfError(err)

	response := controller.TypeService.Delete(request.Context(), id)
	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TypeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	typeId := params.ByName("id")
	id, err := strconv.Atoi(typeId)
	helper.PanicIfError(err)

	response := controller.TypeService.FindById(request.Context(), id)

	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TypeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.TypeService.FindAll(request.Context())

	webResponse := web.BaseResponse{
		Code:    200,
		Status:  true,
		Message: "Successfully",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
