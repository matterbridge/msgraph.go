// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AccessPackageAssignmentRequestBuilder is request builder for AccessPackageAssignment
type AccessPackageAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageAssignmentRequest
func (b *AccessPackageAssignmentRequestBuilder) Request() *AccessPackageAssignmentRequest {
	return &AccessPackageAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageAssignmentRequest is request for AccessPackageAssignment
type AccessPackageAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageAssignment
func (r *AccessPackageAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AccessPackageAssignment
func (r *AccessPackageAssignmentRequest) Get() (*AccessPackageAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AccessPackageAssignment
func (r *AccessPackageAssignmentRequest) Update(reqObj *AccessPackageAssignment) (*AccessPackageAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AccessPackageAssignment
func (r *AccessPackageAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AccessPackage is navigation property
func (b *AccessPackageAssignmentRequestBuilder) AccessPackage() *AccessPackageRequestBuilder {
	bb := &AccessPackageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackage"
	return bb
}

// AccessPackageAssignmentPolicy is navigation property
func (b *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentPolicy() *AccessPackageAssignmentPolicyRequestBuilder {
	bb := &AccessPackageAssignmentPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentPolicy"
	return bb
}

// AccessPackageAssignmentRequests returns request builder for AccessPackageAssignmentRequestObject collection
func (b *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentRequests() *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequestBuilder {
	bb := &AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentRequests"
	return bb
}

// AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequestBuilder is request builder for AccessPackageAssignmentRequestObject collection
type AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentRequestObject collection
func (b *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequestBuilder) Request() *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest {
	return &AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentRequestObject item
func (b *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentRequestObjectRequestBuilder {
	bb := &AccessPackageAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest is request for AccessPackageAssignmentRequestObject collection
type AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageAssignmentRequestObject collection
func (r *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageAssignmentRequestObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AccessPackageAssignmentRequestObject collection
func (r *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest) Paging(method, path string, obj interface{}) ([]AccessPackageAssignmentRequestObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AccessPackageAssignmentRequestObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AccessPackageAssignmentRequestObject
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

// Get performs GET request for AccessPackageAssignmentRequestObject collection
func (r *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest) Get() ([]AccessPackageAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AccessPackageAssignmentRequestObject collection
func (r *AccessPackageAssignmentAccessPackageAssignmentRequestsCollectionRequest) Add(reqObj *AccessPackageAssignmentRequestObject) (*AccessPackageAssignmentRequestObject, error) {
	return r.Do("POST", "", reqObj)
}

// AccessPackageAssignmentResourceRoles returns request builder for AccessPackageAssignmentResourceRole collection
func (b *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentResourceRoles() *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequestBuilder {
	bb := &AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentResourceRoles"
	return bb
}

// AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequestBuilder is request builder for AccessPackageAssignmentResourceRole collection
type AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentResourceRole collection
func (b *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequestBuilder) Request() *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest {
	return &AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentResourceRole item
func (b *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentResourceRoleRequestBuilder {
	bb := &AccessPackageAssignmentResourceRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest is request for AccessPackageAssignmentResourceRole collection
type AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageAssignmentResourceRole collection
func (r *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageAssignmentResourceRole, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AccessPackageAssignmentResourceRole collection
func (r *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest) Paging(method, path string, obj interface{}) ([]AccessPackageAssignmentResourceRole, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AccessPackageAssignmentResourceRole
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AccessPackageAssignmentResourceRole
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

// Get performs GET request for AccessPackageAssignmentResourceRole collection
func (r *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest) Get() ([]AccessPackageAssignmentResourceRole, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AccessPackageAssignmentResourceRole collection
func (r *AccessPackageAssignmentAccessPackageAssignmentResourceRolesCollectionRequest) Add(reqObj *AccessPackageAssignmentResourceRole) (*AccessPackageAssignmentResourceRole, error) {
	return r.Do("POST", "", reqObj)
}

// Target is navigation property
func (b *AccessPackageAssignmentRequestBuilder) Target() *AccessPackageSubjectRequestBuilder {
	bb := &AccessPackageSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/target"
	return bb
}
