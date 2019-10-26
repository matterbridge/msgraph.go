// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GroupPolicyConfigurationAssignRequestParameter undocumented
type GroupPolicyConfigurationAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []GroupPolicyConfigurationAssignment `json:"assignments,omitempty"`
}

//
type GroupPolicyConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *GroupPolicyConfigurationRequestBuilder) Assign(reqObj *GroupPolicyConfigurationAssignRequestParameter) *GroupPolicyConfigurationAssignRequestBuilder {
	bb := &GroupPolicyConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupPolicyConfigurationAssignRequest struct{ BaseRequest }

//
func (b *GroupPolicyConfigurationAssignRequestBuilder) Request() *GroupPolicyConfigurationAssignRequest {
	return &GroupPolicyConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupPolicyConfigurationAssignRequest) Paging(method, path string, obj interface{}) ([][]GroupPolicyConfigurationAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]GroupPolicyConfigurationAssignment
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
			value  [][]GroupPolicyConfigurationAssignment
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
func (r *GroupPolicyConfigurationAssignRequest) Get() ([][]GroupPolicyConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
