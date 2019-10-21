// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PolicySetUpdateRequestParameter undocumented
type PolicySetUpdateRequestParameter struct {
	// AddedPolicySetItems undocumented
	AddedPolicySetItems []PolicySetItem `json:"addedPolicySetItems,omitempty"`
	// UpdatedPolicySetItems undocumented
	UpdatedPolicySetItems []PolicySetItem `json:"updatedPolicySetItems,omitempty"`
	// DeletedPolicySetItems undocumented
	DeletedPolicySetItems []string `json:"deletedPolicySetItems,omitempty"`
	// Assignments undocumented
	Assignments []PolicySetAssignment `json:"assignments,omitempty"`
}

// PolicySetCollectionGetPolicySetsRequestParameter undocumented
type PolicySetCollectionGetPolicySetsRequestParameter struct {
	// PolicySetIDs undocumented
	PolicySetIDs []string `json:"policySetIds,omitempty"`
}

//
type PolicySetUpdateRequestBuilder struct{ BaseRequestBuilder }

// Update action undocumented
func (b *PolicySetRequestBuilder) Update(reqObj *PolicySetUpdateRequestParameter) *PolicySetUpdateRequestBuilder {
	bb := &PolicySetUpdateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/update"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PolicySetUpdateRequest struct{ BaseRequest }

//
func (b *PolicySetUpdateRequestBuilder) Request() *PolicySetUpdateRequest {
	return &PolicySetUpdateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PolicySetUpdateRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *PolicySetUpdateRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type PolicySetCollectionGetPolicySetsRequestBuilder struct{ BaseRequestBuilder }

// GetPolicySets action undocumented
func (b *DeviceAppManagementPolicySetsCollectionRequestBuilder) GetPolicySets(reqObj *PolicySetCollectionGetPolicySetsRequestParameter) *PolicySetCollectionGetPolicySetsRequestBuilder {
	bb := &PolicySetCollectionGetPolicySetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPolicySets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PolicySetCollectionGetPolicySetsRequest struct{ BaseRequest }

//
func (b *PolicySetCollectionGetPolicySetsRequestBuilder) Request() *PolicySetCollectionGetPolicySetsRequest {
	return &PolicySetCollectionGetPolicySetsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PolicySetCollectionGetPolicySetsRequest) Do(method, path string, reqObj interface{}) (resObj *[]PolicySet, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *PolicySetCollectionGetPolicySetsRequest) Paging(method, path string, obj interface{}) ([][]PolicySet, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]PolicySet
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]PolicySet
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
func (r *PolicySetCollectionGetPolicySetsRequest) Get() ([][]PolicySet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
