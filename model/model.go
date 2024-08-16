package model

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
