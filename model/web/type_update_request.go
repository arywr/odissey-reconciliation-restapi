package web

type TypeUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
