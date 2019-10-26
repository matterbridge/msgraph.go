// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ExactMatchDataStoreLookupRequestParameter undocumented
type ExactMatchDataStoreLookupRequestParameter struct {
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// Values undocumented
	Values []string `json:"values,omitempty"`
	// ResultColumnNames undocumented
	ResultColumnNames []string `json:"resultColumnNames,omitempty"`
}

//
type ExactMatchDataStoreLookupRequestBuilder struct{ BaseRequestBuilder }

// Lookup action undocumented
func (b *ExactMatchDataStoreRequestBuilder) Lookup(reqObj *ExactMatchDataStoreLookupRequestParameter) *ExactMatchDataStoreLookupRequestBuilder {
	bb := &ExactMatchDataStoreLookupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/lookup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchDataStoreLookupRequest struct{ BaseRequest }

//
func (b *ExactMatchDataStoreLookupRequestBuilder) Request() *ExactMatchDataStoreLookupRequest {
	return &ExactMatchDataStoreLookupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchDataStoreLookupRequest) Paging(method, path string, obj interface{}) ([][]string, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]string
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  [][]string
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *ExactMatchDataStoreLookupRequest) Get() ([][]string, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
