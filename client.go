package infisicalclient

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	cnf Config
}

type Config struct {
	HostURL      string
	serviceToken string
	apiKey       string
	httpClient   *resty.Client // By default a client will be created
}

func NewClient(cnf Config) (*Client, error) {
	if cnf.apiKey == "" && cnf.serviceToken == "" {
		return nil, fmt.Errorf("You must enter either a API Key or Service token for authentication with Infisical API")
	}

	if cnf.httpClient == nil {
		cnf.httpClient = resty.New()
		cnf.httpClient.SetBaseURL(cnf.HostURL)
	}

	if cnf.serviceToken != "" {
		cnf.httpClient.SetAuthToken(cnf.serviceToken)
	}

	if cnf.apiKey != "" {
		cnf.httpClient.SetHeader("X-API-KEY", cnf.apiKey)
	}

	cnf.httpClient.SetHeader("Accept", "application/json")

	return &Client{cnf}, nil
}
