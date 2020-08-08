package entry

//OpenApiResponseEntry OpenApi 的响应结构体
type OpenApiResponseEntry struct {
	Code int `json:"code"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewOpenApiSuccessResponse(msg string, data interface{}) OpenApiResponseEntry {
	return OpenApiResponseEntry{
		Code: OpenApiSuccessCode,
		Message: msg,
		Data: data,
	}
}

func NewOpenApiErrorResponse(msg string, data interface{}) OpenApiResponseEntry {
	return OpenApiResponseEntry{
		Code: OpenApiErrorCode,
		Message: msg,
		Data: data,
	}
}

func NewOpenApiFailResponse(msg string, data interface{}) OpenApiResponseEntry {
	return OpenApiResponseEntry{
		Code: OpenApiFailCode,
		Message: msg,
		Data: data,
	}
}

func NewOpenApiResponse(code int, msg string, data interface{}) OpenApiResponseEntry {
	return OpenApiResponseEntry{
		Code: code,
		Message: msg,
		Data: data,
	}
}