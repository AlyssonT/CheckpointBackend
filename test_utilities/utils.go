package testutilities

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
)

func ExtractAllMessagesFromResponse(w *httptest.ResponseRecorder) ([]string, error) {
	var responseJSON map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &responseJSON); err != nil {
		return nil, err
	}

	data, ok := responseJSON["Data"].([]any)
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
