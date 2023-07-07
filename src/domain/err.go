package domain

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}