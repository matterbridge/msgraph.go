// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OnPremisesAgentGroupRequestBuilder is request builder for OnPremisesAgentGroup
type OnPremisesAgentGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesAgentGroupRequest
func (b *OnPremisesAgentGroupRequestBuilder) Request() *OnPremisesAgentGroupRequest {
	return &OnPremisesAgentGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesAgentGroupRequest is request for OnPremisesAgentGroup
type OnPremisesAgentGroupRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesAgentGroup
func (r *OnPremisesAgentGroupRequest) Get() (resObj *OnPremisesAgentGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesAgentGroup
func (r *OnPremisesAgentGroupRequest) Update(reqObj *OnPremisesAgentGroup) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesAgentGroup
func (r *OnPremisesAgentGroupRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Agents returns request builder for OnPremisesAgent collection
func (b *OnPremisesAgentGroupRequestBuilder) Agents() *OnPremisesAgentGroupAgentsCollectionRequestBuilder {
	bb := &OnPremisesAgentGroupAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agents"
	return bb
}

// OnPremisesAgentGroupAgentsCollectionRequestBuilder is request builder for OnPremisesAgent collection
type OnPremisesAgentGroupAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgent collection
func (b *OnPremisesAgentGroupAgentsCollectionRequestBuilder) Request() *OnPremisesAgentGroupAgentsCollectionRequest {
	return &OnPremisesAgentGroupAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgent item
func (b *OnPremisesAgentGroupAgentsCollectionRequestBuilder) ID(id string) *OnPremisesAgentRequestBuilder {
	bb := &OnPremisesAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentGroupAgentsCollectionRequest is request for OnPremisesAgent collection
type OnPremisesAgentGroupAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Paging(method, path string, obj interface{}) ([]OnPremisesAgent, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnPremisesAgent
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
			value  []OnPremisesAgent
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

// Get performs GET request for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Get() ([]OnPremisesAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Add(reqObj *OnPremisesAgent) (resObj *OnPremisesAgent, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// PublishedResources returns request builder for PublishedResource collection
func (b *OnPremisesAgentGroupRequestBuilder) PublishedResources() *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder {
	bb := &OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/publishedResources"
	return bb
}

// OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder is request builder for PublishedResource collection
type OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PublishedResource collection
func (b *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder) Request() *OnPremisesAgentGroupPublishedResourcesCollectionRequest {
	return &OnPremisesAgentGroupPublishedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PublishedResource item
func (b *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder) ID(id string) *PublishedResourceRequestBuilder {
	bb := &PublishedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentGroupPublishedResourcesCollectionRequest is request for PublishedResource collection
type OnPremisesAgentGroupPublishedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Paging(method, path string, obj interface{}) ([]PublishedResource, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PublishedResource
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
			value  []PublishedResource
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

// Get performs GET request for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Get() ([]PublishedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Add(reqObj *PublishedResource) (resObj *PublishedResource, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
