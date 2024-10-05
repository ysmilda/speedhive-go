package speedhive

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Client struct {
	c *http.Client

	base url.URL
}

func NewClient(c *http.Client, base string) (*Client, error) {
	if c == nil {
		c = http.DefaultClient
	}

	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	return &Client{
		c:    c,
		base: *u,
	}, nil
}

// request sends a request with the given method, path, data, and headers and returns the response body.
func (c Client) Request(method string, path string, opt any) (*http.Request, error) {
	u := c.base
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.RawPath = c.base.Path + path
	u.Path = c.base.Path + unescaped

	reqHeaders := make(http.Header)
	reqHeaders.Set("Origin", "https://sporthive.com")
	reqHeaders.Set("Accept", "application/json")

	var body []byte
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	request, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, values := range reqHeaders {
		request.Header[key] = values
	}

	return request, nil
}

func (c Client) Do(req *http.Request, v any) error {
	res, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("not found")
	} else if res.StatusCode < 200 || res.StatusCode >= 300 {
		message, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("status code %d", res.StatusCode)
		}
		return fmt.Errorf("status code %d: %s", res.StatusCode, message)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}

	return err
}

func (c Client) GetBody(path string, opt any) ([]byte, error) {
	req, err := c.Request("GET", path, opt)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = c.Do(req, bufio.NewWriter(&b))
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func Get[T any](c *Client, path string, opt any) (*T, error) {
	req, err := c.Request(http.MethodGet, path, opt)
	if err != nil {
		return nil, err
	}

	v := new(T)

	err = c.Do(req, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
