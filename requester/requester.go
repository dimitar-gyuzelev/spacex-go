package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// HTTPRequester implements Requester interface.
// Provides functions for performing HTTP requests.
type HTTPRequester struct {
	timeout time.Duration
}

func NewHTTPRequester(timeout int) *HTTPRequester {
	return &HTTPRequester{
		timeout: time.Duration(timeout) * time.Second,
	}
}

func (r *HTTPRequester) Get(config Config, result Result) error {
	if len(config.QueryStrings) > 0 {
		config.URL = fmt.Sprintf("%s?%s", config.URL, config.QueryStrings.String())
	}

	title := fmt.Sprintf("GET %s", config.URL)

	req, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return wrapError(err, title, "cannot create request")
	}

	for k, v := range config.Headers {
		req.Header.Add(k, v)
	}

	if err := r.do(req, result); err != nil {
		return wrapError(err, title)
	}

	return nil
}

func (r *HTTPRequester) Post(config Config, payload Payload, result Result) error {
	if len(config.QueryStrings) > 0 {
		config.URL = fmt.Sprintf("%s?%s", config.URL, config.QueryStrings.String())
	}

	title := fmt.Sprintf("POST %s", config.URL)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return wrapError(err, title, "cannot marshal payload")
	}

	req, err := http.NewRequest("POST", config.URL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return wrapError(err, title, "cannot create request")
	}

	for k, v := range config.Headers {
		req.Header.Add(k, v)
	}

	if err := r.do(req, result); err != nil {
		return wrapError(err, title)
	}

	return nil
}

func (r *HTTPRequester) do(req *http.Request, result interface{}) error {
	client := &http.Client{
		Timeout: r.timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	if !success(resp.StatusCode) {
		return fmt.Errorf("unsuccessful request: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("cannot decode response: %s", err)
	}

	return nil
}

func wrapError(err error, title string, msg ...string) error {
	return fmt.Errorf("[FAILED] %s: %s: %v", title, strings.Join(msg, ","), err)
}

func success(statusCode int) bool {
	return 200 <= statusCode && statusCode < 300
}
