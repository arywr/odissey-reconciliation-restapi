package web

type TypeCreateRequest struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
}
