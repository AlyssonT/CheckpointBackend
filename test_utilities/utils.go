package testutilities

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func ExtractAllMessagesFromResponse(w *httptest.ResponseRecorder) ([]string, error) {
	var responseJSON map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &responseJSON); err != nil {
		return nil, err
	}

	data, ok := responseJSON["data"].([]any)
	if !ok {
		return nil, errors.New("error parsing response")
	}

	var messages []string
	for _, item := range data {
		message, ok := item.(string)
		if !ok {
			return nil, errors.New("error parsing response")
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func addCookiesToRequest(req *http.Request, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
}

func MakeRequest(server *gin.Engine, method, path string, body any, cookies []*http.Cookie) *httptest.ResponseRecorder {
	var reqBody *bytes.Reader
	if body != nil {
		jsonData, _ := json.Marshal(body)
		reqBody = bytes.NewReader(jsonData)
	} else {
		reqBody = bytes.NewReader([]byte{})
	}

	req, _ := http.NewRequest(method, path, reqBody)
	if cookies != nil {
		addCookiesToRequest(req, cookies)
	}

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	return w
}

func ConvertDataFromResponse[T any](data any) (T, error) {
	var result T

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return result, err
	}

	return result, nil
}
