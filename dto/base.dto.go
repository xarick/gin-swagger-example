package dto

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Msg string `json:"msg"`
}
