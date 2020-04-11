package response

type SuccessResponse struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

type Json map[string]interface{}

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
	Status       int    `json:"status"`
}
