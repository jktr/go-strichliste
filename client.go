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
	LibVersion      = "0.2.0"
	ApiVersion      = "2:1.6.1" // compatible with this Strichliste API version
	DefaultEndpoint = "https://demo.strichliste.org/api"
)

type (
	ClientOption func(*Client)

	// Contains the necessary context to make API requests.
	Client struct {
		endpoint   string
		httpClient *http.Client
		appName    string
		appVersion string
		userAgent  string // derived via appName/appVersion

		User        UserClient
		Transaction TransactionClient
		Article     ArticleClient
		Settings    SettingsClient
		Metrics     MetricsClient
	}

	// Allows for result pagination and result limitation.
	//
	// Pagination works by passing a page number and the number of
	// items on each page.
	//
	// Limiting results works by keeping Page at 0 and setting
	// PerPage to the desired number of total entries.
	//
	// Note that endpoint semantics differ; some use "most recent
	// items" and some "oldest items". TODO document this behaviour.
	ListOpts struct {
		Page    uint // page number (1-indexed, 0 means all on one page)
		PerPage uint // items per page (0 means default size)
	}
)

// Pagination and limit is handled via query parameters,
// so this converts ListOpts into corresponding url.Values.
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

// Configure the URL of the API endpoint.
// Not setting this option will default to DefaultEndpoint.
func WithEndpoint(endpoint string) ClientOption {
	return func(client *Client) {
		client.endpoint = strings.TrimRight(endpoint, "/")
	}
}

// Configure the application's name and version to use when contacting
// the API endpoint. The User Agent for requests derives from this value.
// Not setting this option will default to LibName and LibVersion.
func WithApplication(name, version string) ClientOption {
	return func(client *Client) {
		client.appName = name
		client.appVersion = version
	}
}

// Create a new API client. This is the library's entrypoint.
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
	client.Settings = SettingsClient{client: client}
	client.Metrics = MetricsClient{client: client}

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

// Builds a http.Request suitable for contacting the API endpoint.
// Method should be GET/POST/DELETE/etc.
// Body may be any one of the …Request strichliste.schema structs.
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

// Wraps a http.Response with additional Metadata.
type Response struct {
	*http.Response
}

// Executes an API call. Parses the response and any errors.
// A suitable reqest may be prepared via NewRequest.
// Obj may be any one of the …Single-/MultiResponse strichliste.schema structs.
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
	if er.Error.Class == "" { //|| er.Error.Code == 0 {
		return nil
	}
	if er.Error.Message == "" {
		er.Error.Message = string(er.Error.Class)
	}
	return &er.Error
}
