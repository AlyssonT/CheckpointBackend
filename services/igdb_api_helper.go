package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/AlyssonT/CheckpointBackend/configs"
)

type IGDBApiHelper struct {
	body    string
	route   string
	configs *configs.Configs
	client  *http.Client
}

func NewIGDBApiHelper() *IGDBApiHelper {
	helper := &IGDBApiHelper{
		configs: configs.GetConfigs(),
		client:  &http.Client{},
	}
	return helper
}

func (helper *IGDBApiHelper) createReq() (*http.Request, error) {
	route := strings.Trim(helper.route, " /")
	apiUrl := strings.Trim(helper.configs.IGDBApiURL, " /")

	req, err := http.NewRequest("POST", apiUrl+"/"+route, bytes.NewBuffer([]byte(helper.body)))
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Client-ID", helper.configs.IGDBClientId)
	req.Header.Set("Authorization", "Bearer "+helper.configs.IGDBAuthorization)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (helper *IGDBApiHelper) Route(route string) *IGDBApiHelper {
	helper.route = route
	return helper
}

func (helper *IGDBApiHelper) Req(body string) *IGDBApiHelper {
	helper.body = body
	return helper
}

func (helper *IGDBApiHelper) Run() (any, error) {
	req, err := helper.createReq()
	if err != nil {
		return nil, err
	}
	helper.body = ""
	helper.route = ""

	resp, err := helper.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseBody any
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
