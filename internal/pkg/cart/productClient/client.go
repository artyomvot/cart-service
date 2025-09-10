package productClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

type productRequest struct {
	Token string `json:"token"`
	SKU   int64  `json:"sku"`
}

type productResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func New(baseURL, token string) *Client {
	return &Client{
		baseURL:    baseURL,
		token:      token,
		httpClient: &http.Client{},
	}
}

type Item struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *Client) GetProductInfo(sku int64) (*Item, error) {
	url := fmt.Sprintf("%s/get_product", c.baseURL)

	requestBody := productRequest{
		Token: c.token,
		SKU:   sku,
	}

	jsonBody, err := json.Marshal(requestBody)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status: %d", resp.StatusCode)
	}

	var response productResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	var item Item

	item = Item{
		Name:  response.Name,
		Price: response.Price,
	}

	return &item, nil
}
