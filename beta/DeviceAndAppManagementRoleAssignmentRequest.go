// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceAndAppManagementRoleAssignmentRequestBuilder is request builder for DeviceAndAppManagementRoleAssignment
type DeviceAndAppManagementRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceAndAppManagementRoleAssignmentRequest
func (b *DeviceAndAppManagementRoleAssignmentRequestBuilder) Request() *DeviceAndAppManagementRoleAssignmentRequest {
	return &DeviceAndAppManagementRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceAndAppManagementRoleAssignmentRequest is request for DeviceAndAppManagementRoleAssignment
type DeviceAndAppManagementRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for DeviceAndAppManagementRoleAssignment
func (r *DeviceAndAppManagementRoleAssignmentRequest) Get() (resObj *DeviceAndAppManagementRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceAndAppManagementRoleAssignment
func (r *DeviceAndAppManagementRoleAssignmentRequest) Update(reqObj *DeviceAndAppManagementRoleAssignment) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceAndAppManagementRoleAssignment
func (r *DeviceAndAppManagementRoleAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// RoleScopeTags returns request builder for RoleScopeTag collection
func (b *DeviceAndAppManagementRoleAssignmentRequestBuilder) RoleScopeTags() *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder {
	bb := &DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleScopeTags"
	return bb
}

// DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder is request builder for RoleScopeTag collection
type DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RoleScopeTag collection
func (b *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder) Request() *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest {
	return &DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RoleScopeTag item
func (b *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder) ID(id string) *RoleScopeTagRequestBuilder {
	bb := &RoleScopeTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest is request for RoleScopeTag collection
type DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RoleScopeTag collection
func (r *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest) Paging(method, path string, obj interface{}) ([]RoleScopeTag, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []RoleScopeTag
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
			value  []RoleScopeTag
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

// Get performs GET request for RoleScopeTag collection
func (r *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest) Get() ([]RoleScopeTag, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for RoleScopeTag collection
func (r *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequest) Add(reqObj *RoleScopeTag) (resObj *RoleScopeTag, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
