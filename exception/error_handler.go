package exception

import (
	"net/http"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"odissey-golang/odissey-reconciliation-restapi/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.BaseResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Status:  false,
		Data:    err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func notFoundErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Not Found",
			Status:  false,
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
