package api

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *errorResponse {
	return &errorResponse{
		Code:    code,
		Message: message,
	}
}

func Success() map[string]bool {
	return map[string]bool{
		"success": true,
	}
}

func InvalidJson() *errorResponse {
	return NewError(CodeInvalidJson, "Invalid json")
}

func InvalidCredentials() *errorResponse {
	return NewError(CodeInvalidCredentials, "Invalid credentials")
}

func Unknown(msg ...string) *errorResponse {
	message := "Unknown error"
	if len(msg) > 0 {
		message = msg[0]
	}
	return NewError(CodeUnknown, message)
}
