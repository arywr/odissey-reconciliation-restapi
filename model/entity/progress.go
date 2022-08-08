package entity

type Progress struct {
	Id              int
	ProgressName    interface{}
	ProgressEventId int
	Status          string
	MaximumCounter  int
	Percentage      float64
	File            string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       *string
}
