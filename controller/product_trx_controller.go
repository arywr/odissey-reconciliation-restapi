package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductTrxController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateFromCSV(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
