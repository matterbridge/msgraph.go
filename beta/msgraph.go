// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/yaegashi/msgraph.go/jsonx"
)

const (
	defaultBaseURL = "https://graph.microsoft.com/beta"
)

// Binary is type alias for Edm.Binary
type Binary []byte

// Stream is type alias for Edm.Stream
type Stream []byte

// UUID is type alias for Edm.Guid
type UUID string

// Object is the common ancestor of all models
type Object struct {
	// AdditionalData contains all other fields not defined above
	AdditionalData map[string]interface{} `json:"-" jsonx:"true"`
}

// SetAdditionalData sets object's additional data
func (o *Object) SetAdditionalData(key string, val interface{}) {
	if o.AdditionalData == nil {
		o.AdditionalData = map[string]interface{}{key: val}
	} else {
		o.AdditionalData[key] = val
	}
}

// GetAdditionalData gets object's additional data
func (o *Object) GetAdditionalData(key string) (interface{}, bool) {
	if o.AdditionalData == nil {
		return nil, false
	} else {
		val, ok := o.AdditionalData[key]
		return val, ok
	}
}

// Paging is sturct returned to paging requests
type Paging struct {
	NextLink string          `json:"@odata.nextLink"`
	Value    json.RawMessage `json:"value"`
}

// BaseRequestBuilder is base reuqest builder
type BaseRequestBuilder struct {
	baseURL       string
	client        *http.Client
	requestObject interface{}
}

// URL returns URL
func (r *BaseRequestBuilder) URL() string {
	return r.baseURL
}

// BaseRequest is base request
type BaseRequest struct {
	baseURL       string
	client        *http.Client
	requestObject interface{}
	query         url.Values
}

// URL returns URL with queries
func (r *BaseRequest) URL() string {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.baseURL + query
}

// Client returns HTTP client
func (r *BaseRequest) Client() *http.Client {
	return r.client
}

// Query returns queries for URL
func (r *BaseRequest) Query() url.Values {
	return r.query
}

// Expand adds $expand query
func (r *BaseRequest) Expand(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$expand", value)
}

// Select adds $select query
func (r *BaseRequest) Select(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$select", value)
}

// Top adds $top query
func (r *BaseRequest) Top(value int) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$top", strconv.Itoa(value))
}

// Filter adds $filter query
func (r *BaseRequest) Filter(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$filter", value)
}

// Skip adds $skip query
func (r *BaseRequest) Skip(value int) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$skip", strconv.Itoa(value))
}

// OrderBy adds $orderby query
func (r *BaseRequest) OrderBy(value string) {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add("$orderby", value)
}

// NewRequest returns new HTTP request
func (r *BaseRequest) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, r.baseURL+path, body)
}

// NewJSONRequest returns new HTTP request with JSON payload
func (r *BaseRequest) NewJSONRequest(method, path string, obj interface{}) (*http.Request, error) {
	buf := &bytes.Buffer{}
	if obj != nil {
		err := jsonx.NewEncoder(buf).Encode(obj)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, r.baseURL+path, buf)
	if err != nil {
		return nil, err
	}
	if obj != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}

// DecodeJSONResponse decodes HTTP response with JSON payload
func (r *BaseRequest) DecodeJSONResponse(res *http.Response, obj interface{}) error {
	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if obj != nil {
			err := jsonx.NewDecoder(res.Body).Decode(obj)
			if err != nil {
				return err
			}
		}
		return nil
	case http.StatusNoContent:
		return nil
	default:
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("%s: %s", res.Status, string(b))
	}
}

// JSONRequest issues HTTP request with JSON payload
func (r *BaseRequest) JSONRequest(method, path string, reqObj, resObj interface{}) error {
	req, err := r.NewJSONRequest(method, path, reqObj)
	if err != nil {
		return err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return r.DecodeJSONResponse(res, resObj)
}

// GraphServiceRequestBuilder is GraphService reuqest builder
type GraphServiceRequestBuilder struct {
	BaseRequestBuilder
}

// NewClient returns GraphService request builder with default base URL
func NewClient(cli *http.Client) *GraphServiceRequestBuilder {
	return &GraphServiceRequestBuilder{
		BaseRequestBuilder: BaseRequestBuilder{baseURL: defaultBaseURL, client: cli},
	}
}
