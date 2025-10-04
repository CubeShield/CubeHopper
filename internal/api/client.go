package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/CubeShield/CubeHopper/internal/types"
)

type ApiClient struct {
	httpClient *http.Client
	apiUrl string
}

func NewApiClient(apiUrl string) *ApiClient {
	return &ApiClient{
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		apiUrl: apiUrl,
	}
}

func (c *ApiClient) GetInstance() (*types.Instance, error) {
	requestURL := fmt.Sprintf("%s/instances", c.apiUrl)
	
	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API вернуло невернуый статус: %s", resp.Status)
	}

	var instance types.Instance
	if err := json.NewDecoder(resp.Body).Decode(&instance); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON-ответа от API: %w", err)
	}

	return &instance, nil
}

func (c *ApiClient) DownloadFile(url string) (io.ReadCloser, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при скачивании файла с %s: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сервер вернул неверный статус для файла %s: %s", url, resp.Status)
	}

	return resp.Body, nil
}