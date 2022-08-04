package web

type ProductTrxCreateExcel struct {
	Day     int         `formam:"day"`
	OwnerId string      `formam:"platformId"`
	File    interface{} `formam:"file"`
}
