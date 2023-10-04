package external_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) RecallDataSource(_ context.Context, timeout time.Duration) error {
	url := os.Getenv("DATA_SOURCE_RECALL_URL")
	if url == "" {
		return fmt.Errorf("no url specified in enviroment variables")
	}

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		return fmt.Errorf("cannot perform request: %v", err)
	}

	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("cannot parse response body: %v", err)
	}

	expectedResponse, _ := json.Marshal("OK")

	if string(result) != string(expectedResponse) {
		return fmt.Errorf("wrong response: %v", string(result))
	}

	return nil
}
