package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type APIError struct {
	StatusCode int `json:"status_code"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

func InvalidQueryParams() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid query params"))
}

func InvalidPermissions() APIError {
	return NewAPIError(http.StatusForbidden, fmt.Errorf("invalid permissions"))
}

func PermissionDenied(w http.ResponseWriter) {
	writeJSON(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	response := Response{
		Status: status,
		Data:   data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}
