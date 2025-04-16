package services

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/AlyssonT/CheckpointBackend/configs"
)

type SteamApiHelper struct {
	route  string
	url    string
	client *http.Client
}

func NewSteamApiHelper() *SteamApiHelper {
	helper := &SteamApiHelper{
		url:    configs.GetConfigs().SteamApiURL,
		client: &http.Client{},
	}
	return helper
}

func NewSteamStoreApiHelper() *SteamApiHelper {
	helper := &SteamApiHelper{
		url:    configs.GetConfigs().SteamStoreApiURL,
		client: &http.Client{},
	}
	return helper
}

func (helper *SteamApiHelper) createReq() (*http.Request, error) {
	route := strings.Trim(helper.route, " /")
	apiUrl := strings.Trim(helper.url, " /")

	fmt.Println(apiUrl + "/" + route)
	req, err := http.NewRequest("GET", apiUrl+"/"+route, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (helper *SteamApiHelper) Route(route string) *SteamApiHelper {
	helper.route = route
	return helper
}

func (helper *SteamApiHelper) Run() ([]byte, error) {
	req, err := helper.createReq()
	if err != nil {
		return nil, err
	}
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

	return body, nil
}
