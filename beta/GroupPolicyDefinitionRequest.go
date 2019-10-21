// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GroupPolicyDefinitionRequestBuilder is request builder for GroupPolicyDefinition
type GroupPolicyDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyDefinitionRequest
func (b *GroupPolicyDefinitionRequestBuilder) Request() *GroupPolicyDefinitionRequest {
	return &GroupPolicyDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyDefinitionRequest is request for GroupPolicyDefinition
type GroupPolicyDefinitionRequest struct{ BaseRequest }

// Do performs HTTP request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Do(method, path string, reqObj interface{}) (resObj *GroupPolicyDefinition, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Get() (*GroupPolicyDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Update(reqObj *GroupPolicyDefinition) (*GroupPolicyDefinition, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for GroupPolicyDefinition
func (r *GroupPolicyDefinitionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DefinitionFile is navigation property
func (b *GroupPolicyDefinitionRequestBuilder) DefinitionFile() *GroupPolicyDefinitionFileRequestBuilder {
	bb := &GroupPolicyDefinitionFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definitionFile"
	return bb
}

// Presentations returns request builder for GroupPolicyPresentation collection
func (b *GroupPolicyDefinitionRequestBuilder) Presentations() *GroupPolicyDefinitionPresentationsCollectionRequestBuilder {
	bb := &GroupPolicyDefinitionPresentationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/presentations"
	return bb
}

// GroupPolicyDefinitionPresentationsCollectionRequestBuilder is request builder for GroupPolicyPresentation collection
type GroupPolicyDefinitionPresentationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GroupPolicyPresentation collection
func (b *GroupPolicyDefinitionPresentationsCollectionRequestBuilder) Request() *GroupPolicyDefinitionPresentationsCollectionRequest {
	return &GroupPolicyDefinitionPresentationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GroupPolicyPresentation item
func (b *GroupPolicyDefinitionPresentationsCollectionRequestBuilder) ID(id string) *GroupPolicyPresentationRequestBuilder {
	bb := &GroupPolicyPresentationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GroupPolicyDefinitionPresentationsCollectionRequest is request for GroupPolicyPresentation collection
type GroupPolicyDefinitionPresentationsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *GroupPolicyPresentation, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Paging(method, path string, obj interface{}) ([]GroupPolicyPresentation, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GroupPolicyPresentation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []GroupPolicyPresentation
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

// Get performs GET request for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Get() ([]GroupPolicyPresentation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GroupPolicyPresentation collection
func (r *GroupPolicyDefinitionPresentationsCollectionRequest) Add(reqObj *GroupPolicyPresentation) (*GroupPolicyPresentation, error) {
	return r.Do("POST", "", reqObj)
}
