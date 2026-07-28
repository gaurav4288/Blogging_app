package utils

func SuccessResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(message string, err error) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": message,
		"error":   err,
	}
}
