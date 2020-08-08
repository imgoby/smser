package entry

const (
	OpenApiSuccessCode = 20000

	OpenApiErrorCode = 50000

	OpenApiFailCode = 40000
)

var openApiResponseMessages = map[int]string{
	OpenApiSuccessCode: "success",
	OpenApiErrorCode: "error",
	OpenApiFailCode: "fail",
}

func GetOpenApiResponseMessageByCode(code int) string {
	str, ok := openApiResponseMessages[code]
	if !ok {
		return ""
	}

	return str
}