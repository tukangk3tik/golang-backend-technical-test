package helper

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type FailResponse struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	DataFail interface{} `json:"data_fail"`
}

type EmptyObject []string

func BuildSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Status: "success",
		Data:   data,
	}
}

func BuildFailResponse(message string, dataFail interface{}) FailResponse {
	return FailResponse{
		Status:   "fail",
		Message:  message,
		DataFail: dataFail,
	}
}
