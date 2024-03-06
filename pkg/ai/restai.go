package ai

import (
	"context"
	"io"
	"net/http"
	"strings"
)

const restAIClientName = "restai"

type RestAIClient struct {
	token   string
	url     string
	request string
}

func (c *RestAIClient) Configure(config IAIConfig) error {
	c.token = config.GetPassword()
	c.url = config.GetBaseURL()
	c.request = config.GetRequestBody()
	return nil
}

func (c *RestAIClient) GetCompletion(ctx context.Context, prompt string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, c.url, strings.NewReader(c.request))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

func (c *RestAIClient) GetName() string {
	return restAIClientName
}

func (c *RestAIClient) Close() {
}
