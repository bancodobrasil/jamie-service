package featws

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/bancodobrasil/jamie-service/dtos"
	log "github.com/sirupsen/logrus"
)

type RullerClient struct {
	Url    string
	ApiKey string
}

func NewRullerClient(url, apiKey string) *RullerClient {
	return &RullerClient{
		Url:    url,
		ApiKey: apiKey,
	}
}

func (c *RullerClient) GetFeatures(knowledgeBase string, version string, parameters map[string]string) (*dtos.Eval, error) {
	parametersJson, err := json.Marshal(parameters)
	if err != nil {
		log.Errorf("Error marshaling parameters: %s", err)
		return nil, err
	}
	body := []byte(`{"knowledgeBase": "` + knowledgeBase + `", "version": "` + version + `", "parameters": ` + string(parametersJson) + `}`)
	log.Debugf("Request body: %s", body)
	bodyReader := bytes.NewReader(body)
	request, err := http.NewRequest("POST", c.Url, bodyReader)
	if err != nil {
		log.Errorf("Error creating request: %s", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-key", c.ApiKey)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Errorf("Error making request: %s", err)
		return nil, err
	}
	defer response.Body.Close()
	log.Debugf("Response status: %s", response.Status)
	log.Debugf("Response headers: %s", response.Header)
	responseBody := &dtos.Eval{}
	err = json.NewDecoder(response.Body).Decode(responseBody)
	if err != nil {
		log.Errorf("Error decoding response: %s", err)
		return nil, err
	}
	log.Debugf("Response body: %s", responseBody)
	return responseBody, nil
}
