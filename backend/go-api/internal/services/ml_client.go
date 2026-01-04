package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type MLClient interface {
	Analyze(ctx context.Context, text string) (float64, bool, error)
}

type mlClient struct {
	baseURL string
}

func NewMLCLient(url string) MLClient {
	return &mlClient{baseURL: url}
}

func (c *mlClient) Analyze(ctx context.Context, text string) (float64, bool, error) {
	body, _ := json.Marshal(map[string]string{
		"text": text,
	})

	req, _ := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/predict", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, false, err
	}
	defer resp.Body.Close()

	var result struct {
		Score  float64 `json:"score"`
		IsFake bool    `json:"is_fake"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	return result.Score, result.IsFake, err
}
