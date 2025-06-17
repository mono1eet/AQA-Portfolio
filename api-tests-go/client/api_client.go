package client

import (
    "github.com/go-resty/resty/v2"
)

// ApiClient предоставляет методы для отправки HTTP-запросов к API
type ApiClient struct {
    client  *resty.Client
    baseURL string
}

// NewApiClient создаёт новый экземпляр ApiClient
func NewApiClient(baseURL string) *ApiClient {
    return &ApiClient{
        client:  resty.New(),
        baseURL: baseURL,
    }
}

// Get отправляет GET-запрос к указанному эндпоинту
func (c *ApiClient) Get(endpoint string) (*resty.Response, error) {
    return c.client.R().Get(c.baseURL + endpoint)
}