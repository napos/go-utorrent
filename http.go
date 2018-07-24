package utorrent

import (
	"bytes"
	"fmt"
	"net/http"
)

func (c *Client) url(path string) string {
	if path == "" || path[0:1] != "/" {
		path = fmt.Sprintf("/%s", path)
	}

	if c.token != "" {
		path = fmt.Sprintf("%s&token=%s", path, c.token)
	}
	return fmt.Sprintf("%s%s", c.API, path)
}

func (c *Client) request(method, path string, payload []byte, headers *http.Header) (*http.Response, error) {
	if c == nil {
		return nil, fmt.Errorf("Cannot make a request with a nil client")
	}
	in := bytes.NewBuffer(payload)
	req, err := http.NewRequest(method, c.url(path), in)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)

	if headers != nil {
		for header, values := range *headers {
			for _, value := range values {
				req.Header.Add(header, value)
			}
		}
	}

	res, err := c.user_agent.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) post(path string, payload []byte, headers *http.Header) (*http.Response, error) {
	return c.request("POST", path, payload, headers)
}

func (c *Client) put(path string, payload []byte, headers *http.Header) (*http.Response, error) {
	return c.request("PUT", path, payload, headers)
}

func (c *Client) get(path string, headers *http.Header) (*http.Response, error) {
	return c.request("GET", path, nil, headers)
}

func (c *Client) delete(path string, headers *http.Header) (*http.Response, error) {
	return c.request("DELETE", path, nil, headers)
}
