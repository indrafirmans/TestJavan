package dto

type Response struct {
	Status  int
	Message string
	Detail  string
	Data    any
}
