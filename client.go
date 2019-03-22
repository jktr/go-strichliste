package strichliste

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jktr/go-strichliste/schema"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	LibName         = "go-strichliste"
	LibVersion      = "0.1.0"
	ApiVersion      = "v2:1.6.0"
	DefaultEndpoint = "https://demo.strichliste.org/api"
)

type (
	ClientOption func(*Client)
	Client       struct {
		endpoint   string
		httpClient *http.Client
		appName    string
		appVersion string
		userAgent  string

		User        UserClient
		Transaction TransactionClient
		Article     ArticleClient
	}

	ListOpts struct {
		Page    uint // 1-indexed (0 means all)
		PerPage uint // entries per page (0 means default)
	}
)

func (l *ListOpts) values() url.Values {
	v := url.Values{}

	if l == nil {
		return v
	}

	if l.Page > 0 {
		v.Add("page", string(l.Page))
	}
	if l.PerPage > 0 {
		v.Add("limit", string(l.PerPage))
	}
	return v
}

func WithEndpoint(endpoint string) ClientOption {
	return func(client *Client) {
		client.endpoint = strings.TrimRight(endpoint, "/")
	}
}

func WithApplication(name, version string) ClientOption {
	return func(client *Client) {
		client.appName = name
		client.appVersion = version
	}
}

func NewClient(options ...ClientOption) *Client {
	client := &Client{
		endpoint:   DefaultEndpoint,
		httpClient: &http.Client{},
		appName:    LibName,
		appVersion: fmt.Sprintf("%s (api:%s)", LibVersion, ApiVersion),
	}

	for _, option := range options {
		option(client)
	}

	client.userAgent = client.appName + "/" + client.appVersion

	client.Article = ArticleClient{client: client}
	client.User = UserClient{client: client}
	client.Transaction = TransactionClient{client: client}

	return client
}

func newJsonReader(obj interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if obj != nil {
		err := json.NewEncoder(buf).Encode(obj)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	reader, err := newJsonReader(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.endpoint+path, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

type Response struct {
	*http.Response
}

func (c *Client) Do(req *http.Request, obj interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return response, err
	}
	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewReader(body))

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		err = errorFromResponse(resp, body)
		if err == nil {
			err = fmt.Errorf("%s: server responded with status code %d",
				c.appName, resp.StatusCode)
		}
		return response, err
	}

	if obj != nil {
		if w, ok := obj.(io.Writer); ok {
			_, err = io.Copy(w, bytes.NewReader(body))
		} else {
			err = json.Unmarshal(body, obj)
		}
	}
	return response, err
}

func errorFromResponse(resp *http.Response, body []byte) error {
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil
	}
	var er schema.SingleErrorResponse
	if err := json.Unmarshal(body, &er); err != nil {
		return nil
	}
	if er.Error.Class == "" || er.Error.Code == 0 {
		return nil
	}
	if er.Error.Message == "" {
		er.Error.Message = string(er.Error.Class)
	}
	return &er.Error
}
