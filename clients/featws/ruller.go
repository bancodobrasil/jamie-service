package featws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bancodobrasil/jamie-service/dtos"
	log "github.com/sirupsen/logrus"
)

type EvalRequest map[string]interface{}

func NewEvalRequest(dto dtos.Process) EvalRequest {
	return EvalRequest(dto)
}

type EvalPayload map[string]interface{}

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

func (c *RullerClient) Eval(knowledgeBase string, version string, parameters EvalRequest) (*EvalPayload, error) {
	url := fmt.Sprintf("%s/api/v1/eval/%s/%s", c.Url, knowledgeBase, version)
	log.Debugf("Request url: %s", url)
	parametersJson, err := json.Marshal(parameters)
	if err != nil {
		log.Errorf("Error marshaling parameters: %s", err)
		return nil, err
	}
	body := []byte(string(parametersJson))
	log.Debugf("Request body: %s", body)
	bodyReader := bytes.NewReader(body)
	request, err := http.NewRequest("POST", url, bodyReader)
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
	responseBody := &EvalPayload{}
	err = json.NewDecoder(response.Body).Decode(responseBody)
	if err != nil {
		log.Errorf("Error decoding response: %s", err)
		return nil, err
	}
	log.Debugf("Response body: %s", responseBody)
	return responseBody, nil
}
