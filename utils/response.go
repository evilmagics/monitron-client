package utils

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(data ...interface{}) ApiResponse {
	if len(data) == 0 {
		return ApiResponse{Message: "success"}
	}

	return ApiResponse{Message: "success", Data: data[0]}
}
func FailedResponse(err error, data ...interface{}) ApiResponse {
	if len(data) == 0 {
		return ApiResponse{Message: err.Error()}
	}
	return ApiResponse{Message: err.Error(), Data: data[0]}
}
