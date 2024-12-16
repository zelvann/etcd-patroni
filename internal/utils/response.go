package utils

type (
	SuccessResponse struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload,omitempty"`
		Meta    interface{} `json:"meta,omitempty"`
	}

	FailedResponse struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
)

func NewSucessResponse(message string) SuccessResponse {
	return SuccessResponse{
		Status:  true,
		Message: message,
	}
}

func (r SuccessResponse) WithPayload(payload interface{}) SuccessResponse {
	r.Payload = &payload
	return r
}

func (r SuccessResponse) WithMeta(meta interface{}) SuccessResponse {
	r.Meta = &meta
	return r
}

func NewFailedResponse(message string, err interface{}) FailedResponse {
	return FailedResponse{
		Status:  false,
		Message: message,
		Error:   err,
	}
}
