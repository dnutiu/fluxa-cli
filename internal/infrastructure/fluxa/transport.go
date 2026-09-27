package fluxa

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const maxResponseBytes = 16 << 20

type Repository struct {
	baseURL *url.URL
	key     string
	http    *http.Client
}

func NewRepository(baseURL, key string, httpClient *http.Client) (*Repository, error) {
	u, err := parseBaseURL(baseURL)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("FLUXA_API_KEY is required; create a Pro API key in Tools → API access")
	}
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second, CheckRedirect: func(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse }}
	}
	return &Repository{baseURL: u, key: strings.TrimSpace(key), http: httpClient}, nil
}

func (r *Repository) request(ctx context.Context, method, path string, query url.Values, body []byte, key string, result any) error {
	if !strings.HasPrefix(path, "/api/v1/") || strings.Contains(path, "..") {
		return errors.New("invalid API path")
	}
	u := *r.baseURL
	u.Path = path
	u.RawQuery = query.Encode()
	var reader io.Reader
	if body != nil {
		if !json.Valid(body) {
			return errors.New("request body must be valid JSON")
		}
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "fluxa-cli/0.1")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if key != "" {
		req.Header.Set("Idempotency-Key", key)
	}
	resp, err := r.http.Do(req)
	if err != nil {
		return fmt.Errorf("Fluxa request failed: %w", err)
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBytes+1))
	if err != nil {
		return fmt.Errorf("read Fluxa response: %w", err)
	}
	if len(content) > maxResponseBytes {
		return errors.New("Fluxa response exceeded the size limit")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return responseError(resp, content)
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if !json.Valid(content) {
		return errors.New("Fluxa returned invalid JSON")
	}
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(content, &envelope); err != nil {
		return err
	}
	if len(envelope["data"]) == 0 {
		return errors.New("Fluxa response has no data envelope")
	}
	return json.Unmarshal(content, result)
}

func responseError(resp *http.Response, content []byte) error {
	var envelope struct {
		Error struct {
			Code    string          `json:"code"`
			Message string          `json:"message"`
			Details json.RawMessage `json:"details"`
		} `json:"error"`
	}
	_ = json.Unmarshal(content, &envelope)
	apiErr := &Error{Status: resp.StatusCode, Code: envelope.Error.Code, Message: envelope.Error.Message, Details: envelope.Error.Details}
	if resp.StatusCode == http.StatusTooManyRequests && resp.Header.Get("Retry-After") != "" {
		apiErr.Message += " (retry after " + resp.Header.Get("Retry-After") + " seconds)"
	}
	return apiErr
}
