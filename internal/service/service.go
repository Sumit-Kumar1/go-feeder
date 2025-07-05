package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"go-feeder/internal/models"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Storer interface {
	GetFeeds(context.Context) error
}

type Service struct {
	Store Storer
}

func New(s Storer) *Service {
	return &Service{
		Store: s,
	}
}

func (s *Service) FetchFeeds(ctx context.Context, url string) (*models.RSS, error) {
	if err := validateURL(url); err != nil {
		return nil, err
	}

	return fetchFeed(url)
}

func fetchFeed(url string) (*models.RSS, error) {
	url = sanitizeURL(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("error while fetch feeds for url{%s} : %s", url, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status while fetching feeds code: %d", resp.StatusCode)
	}

	res, err := decodeFeedResp(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return res, nil
}

func decodeFeedResp(body io.ReadCloser) (*models.RSS, error) {
	var resp models.RSS

	if err := xml.NewDecoder(body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func sanitizeURL(uri string) string {
	if uri[:4] == "http" {
		return uri
	}

	return "https://" + uri
}

func validateURL(uri string) error {
	uri = strings.TrimSpace(uri)

	if uri == "" {
		return fmt.Errorf("invalid url {%s} provided", uri)
	}

	_, err := url.Parse(uri)
	if err != nil {
		return fmt.Errorf("Invalid url {%s} provided", uri)
	}

	return nil
}
