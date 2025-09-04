package productClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type IItem struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *Client) GetProductInfo(sku int64) (*IItem, error) {
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

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var responsee productResponse

	err = json.Unmarshal(body, &responsee)

	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	var item IItem

	item = IItem{
		Name:  responsee.Name,
		Price: responsee.Price,
	}

	return &item, nil
}
