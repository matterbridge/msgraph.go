// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AccessPackageResourceRequestBuilder is request builder for AccessPackageResource
type AccessPackageResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageResourceRequest
func (b *AccessPackageResourceRequestBuilder) Request() *AccessPackageResourceRequest {
	return &AccessPackageResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageResourceRequest is request for AccessPackageResource
type AccessPackageResourceRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageResource
func (r *AccessPackageResourceRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageResource, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AccessPackageResource
func (r *AccessPackageResourceRequest) Get() (*AccessPackageResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AccessPackageResource
func (r *AccessPackageResourceRequest) Update(reqObj *AccessPackageResource) (*AccessPackageResource, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AccessPackageResource
func (r *AccessPackageResourceRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AccessPackageResourceRoles returns request builder for AccessPackageResourceRole collection
func (b *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles() *AccessPackageResourceAccessPackageResourceRolesCollectionRequestBuilder {
	bb := &AccessPackageResourceAccessPackageResourceRolesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceRoles"
	return bb
}

// AccessPackageResourceAccessPackageResourceRolesCollectionRequestBuilder is request builder for AccessPackageResourceRole collection
type AccessPackageResourceAccessPackageResourceRolesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceRole collection
func (b *AccessPackageResourceAccessPackageResourceRolesCollectionRequestBuilder) Request() *AccessPackageResourceAccessPackageResourceRolesCollectionRequest {
	return &AccessPackageResourceAccessPackageResourceRolesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceRole item
func (b *AccessPackageResourceAccessPackageResourceRolesCollectionRequestBuilder) ID(id string) *AccessPackageResourceRoleRequestBuilder {
	bb := &AccessPackageResourceRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageResourceAccessPackageResourceRolesCollectionRequest is request for AccessPackageResourceRole collection
type AccessPackageResourceAccessPackageResourceRolesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageResourceRole collection
func (r *AccessPackageResourceAccessPackageResourceRolesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageResourceRole, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AccessPackageResourceRole collection
func (r *AccessPackageResourceAccessPackageResourceRolesCollectionRequest) Paging(method, path string, obj interface{}) ([]AccessPackageResourceRole, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AccessPackageResourceRole
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AccessPackageResourceRole
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

// Get performs GET request for AccessPackageResourceRole collection
func (r *AccessPackageResourceAccessPackageResourceRolesCollectionRequest) Get() ([]AccessPackageResourceRole, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AccessPackageResourceRole collection
func (r *AccessPackageResourceAccessPackageResourceRolesCollectionRequest) Add(reqObj *AccessPackageResourceRole) (*AccessPackageResourceRole, error) {
	return r.Do("POST", "", reqObj)
}

// AccessPackageResourceScopes returns request builder for AccessPackageResourceScope collection
func (b *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes() *AccessPackageResourceAccessPackageResourceScopesCollectionRequestBuilder {
	bb := &AccessPackageResourceAccessPackageResourceScopesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceScopes"
	return bb
}

// AccessPackageResourceAccessPackageResourceScopesCollectionRequestBuilder is request builder for AccessPackageResourceScope collection
type AccessPackageResourceAccessPackageResourceScopesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceScope collection
func (b *AccessPackageResourceAccessPackageResourceScopesCollectionRequestBuilder) Request() *AccessPackageResourceAccessPackageResourceScopesCollectionRequest {
	return &AccessPackageResourceAccessPackageResourceScopesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceScope item
func (b *AccessPackageResourceAccessPackageResourceScopesCollectionRequestBuilder) ID(id string) *AccessPackageResourceScopeRequestBuilder {
	bb := &AccessPackageResourceScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageResourceAccessPackageResourceScopesCollectionRequest is request for AccessPackageResourceScope collection
type AccessPackageResourceAccessPackageResourceScopesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageResourceScope collection
func (r *AccessPackageResourceAccessPackageResourceScopesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageResourceScope, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AccessPackageResourceScope collection
func (r *AccessPackageResourceAccessPackageResourceScopesCollectionRequest) Paging(method, path string, obj interface{}) ([]AccessPackageResourceScope, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AccessPackageResourceScope
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AccessPackageResourceScope
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

// Get performs GET request for AccessPackageResourceScope collection
func (r *AccessPackageResourceAccessPackageResourceScopesCollectionRequest) Get() ([]AccessPackageResourceScope, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AccessPackageResourceScope collection
func (r *AccessPackageResourceAccessPackageResourceScopesCollectionRequest) Add(reqObj *AccessPackageResourceScope) (*AccessPackageResourceScope, error) {
	return r.Do("POST", "", reqObj)
}
