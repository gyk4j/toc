package utils

import (
	"net/http"

	"github.com/gyk4j/toc/toc-backend/models"
)

type APIResponseFactory struct {
	responses map[int]*models.APIResponse
}

func NewAPIResponseFactory() *APIResponseFactory {
	f := APIResponseFactory{}

	f.responses = make(map[int]*models.APIResponse)

	f.responses[http.StatusNotFound] = &models.APIResponse{
		Code:    http.StatusNotFound,
		Message: "Not found",
		Type:    models.APIResponseTypeError,
	}

	f.responses[http.StatusInternalServerError] = &models.APIResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Type:    models.APIResponseTypeError,
	}

	f.responses[http.StatusServiceUnavailable] = &models.APIResponse{
		Code:    http.StatusServiceUnavailable,
		Message: "Service unavailable",
		Type:    models.APIResponseTypeError,
	}

	f.responses[http.StatusNotImplemented] = &models.APIResponse{
		Code:    http.StatusNotImplemented,
		Message: "Not implemented",
		Type:    models.APIResponseTypeError,
	}

	return &f
}

func (f *APIResponseFactory) GetAPIResponse(status int) *models.APIResponse {
	return f.responses[status]
}
