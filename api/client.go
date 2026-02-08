package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MUlt1mate/foursquare-exporter/model"
)

const (
	checkinsURL = "https://api.foursquare.com/v2/users/self/checkins?limit=%d&oauth_token=%s&v=20260208&offset=%d"
	limit       = 250
)

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetCheckins(offset int) ([]model.Checkin, error) {
	url := fmt.Sprintf(checkinsURL, limit, c.token, offset)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var result model.CheckinsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	return result.Response.Checkins.Items, nil
}

func (c *Client) GetAllCheckins() ([]model.Checkin, error) {
	var all []model.Checkin
	offset := 0

	for {
		items, err := c.GetCheckins(offset)
		if err != nil {
			return nil, err
		}
		if len(items) == 0 {
			break
		}
		all = append(all, items...)
		offset += limit
	}

	return all, nil
}
