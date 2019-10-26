// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PrivilegedAccessRequestBuilder is request builder for PrivilegedAccess
type PrivilegedAccessRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedAccessRequest
func (b *PrivilegedAccessRequestBuilder) Request() *PrivilegedAccessRequest {
	return &PrivilegedAccessRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedAccessRequest is request for PrivilegedAccess
type PrivilegedAccessRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedAccess
func (r *PrivilegedAccessRequest) Get() (resObj *PrivilegedAccess, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedAccess
func (r *PrivilegedAccessRequest) Update(reqObj *PrivilegedAccess) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedAccess
func (r *PrivilegedAccessRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Resources returns request builder for GovernanceResource collection
func (b *PrivilegedAccessRequestBuilder) Resources() *PrivilegedAccessResourcesCollectionRequestBuilder {
	bb := &PrivilegedAccessResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resources"
	return bb
}

// PrivilegedAccessResourcesCollectionRequestBuilder is request builder for GovernanceResource collection
type PrivilegedAccessResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceResource collection
func (b *PrivilegedAccessResourcesCollectionRequestBuilder) Request() *PrivilegedAccessResourcesCollectionRequest {
	return &PrivilegedAccessResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceResource item
func (b *PrivilegedAccessResourcesCollectionRequestBuilder) ID(id string) *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessResourcesCollectionRequest is request for GovernanceResource collection
type PrivilegedAccessResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceResource collection
func (r *PrivilegedAccessResourcesCollectionRequest) Paging(method, path string, obj interface{}) ([]GovernanceResource, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceResource
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
			value  []GovernanceResource
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

// Get performs GET request for GovernanceResource collection
func (r *PrivilegedAccessResourcesCollectionRequest) Get() ([]GovernanceResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GovernanceResource collection
func (r *PrivilegedAccessResourcesCollectionRequest) Add(reqObj *GovernanceResource) (resObj *GovernanceResource, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// RoleAssignmentRequests returns request builder for GovernanceRoleAssignmentRequestObject collection
func (b *PrivilegedAccessRequestBuilder) RoleAssignmentRequests() *PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignmentRequests"
	return bb
}

// PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder is request builder for GovernanceRoleAssignmentRequestObject collection
type PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignmentRequestObject collection
func (b *PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder) Request() *PrivilegedAccessRoleAssignmentRequestsCollectionRequest {
	return &PrivilegedAccessRoleAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignmentRequestObject item
func (b *PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestObjectRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleAssignmentRequestsCollectionRequest is request for GovernanceRoleAssignmentRequestObject collection
type PrivilegedAccessRoleAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignmentRequestObject collection
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) Paging(method, path string, obj interface{}) ([]GovernanceRoleAssignmentRequestObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleAssignmentRequestObject
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
			value  []GovernanceRoleAssignmentRequestObject
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

// Get performs GET request for GovernanceRoleAssignmentRequestObject collection
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) Get() ([]GovernanceRoleAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GovernanceRoleAssignmentRequestObject collection
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) Add(reqObj *GovernanceRoleAssignmentRequestObject) (resObj *GovernanceRoleAssignmentRequestObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// RoleAssignments returns request builder for GovernanceRoleAssignment collection
func (b *PrivilegedAccessRequestBuilder) RoleAssignments() *PrivilegedAccessRoleAssignmentsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// PrivilegedAccessRoleAssignmentsCollectionRequestBuilder is request builder for GovernanceRoleAssignment collection
type PrivilegedAccessRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignment collection
func (b *PrivilegedAccessRoleAssignmentsCollectionRequestBuilder) Request() *PrivilegedAccessRoleAssignmentsCollectionRequest {
	return &PrivilegedAccessRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignment item
func (b *PrivilegedAccessRoleAssignmentsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleAssignmentsCollectionRequest is request for GovernanceRoleAssignment collection
type PrivilegedAccessRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignment collection
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]GovernanceRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleAssignment
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
			value  []GovernanceRoleAssignment
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

// Get performs GET request for GovernanceRoleAssignment collection
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) Get() ([]GovernanceRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GovernanceRoleAssignment collection
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) Add(reqObj *GovernanceRoleAssignment) (resObj *GovernanceRoleAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// RoleDefinitions returns request builder for GovernanceRoleDefinition collection
func (b *PrivilegedAccessRequestBuilder) RoleDefinitions() *PrivilegedAccessRoleDefinitionsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinitions"
	return bb
}

// PrivilegedAccessRoleDefinitionsCollectionRequestBuilder is request builder for GovernanceRoleDefinition collection
type PrivilegedAccessRoleDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleDefinition collection
func (b *PrivilegedAccessRoleDefinitionsCollectionRequestBuilder) Request() *PrivilegedAccessRoleDefinitionsCollectionRequest {
	return &PrivilegedAccessRoleDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleDefinition item
func (b *PrivilegedAccessRoleDefinitionsCollectionRequestBuilder) ID(id string) *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleDefinitionsCollectionRequest is request for GovernanceRoleDefinition collection
type PrivilegedAccessRoleDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleDefinition collection
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) Paging(method, path string, obj interface{}) ([]GovernanceRoleDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleDefinition
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
			value  []GovernanceRoleDefinition
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

// Get performs GET request for GovernanceRoleDefinition collection
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) Get() ([]GovernanceRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GovernanceRoleDefinition collection
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) Add(reqObj *GovernanceRoleDefinition) (resObj *GovernanceRoleDefinition, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// RoleSettings returns request builder for GovernanceRoleSetting collection
func (b *PrivilegedAccessRequestBuilder) RoleSettings() *PrivilegedAccessRoleSettingsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleSettingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleSettings"
	return bb
}

// PrivilegedAccessRoleSettingsCollectionRequestBuilder is request builder for GovernanceRoleSetting collection
type PrivilegedAccessRoleSettingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleSetting collection
func (b *PrivilegedAccessRoleSettingsCollectionRequestBuilder) Request() *PrivilegedAccessRoleSettingsCollectionRequest {
	return &PrivilegedAccessRoleSettingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleSetting item
func (b *PrivilegedAccessRoleSettingsCollectionRequestBuilder) ID(id string) *GovernanceRoleSettingRequestBuilder {
	bb := &GovernanceRoleSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleSettingsCollectionRequest is request for GovernanceRoleSetting collection
type PrivilegedAccessRoleSettingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleSetting collection
func (r *PrivilegedAccessRoleSettingsCollectionRequest) Paging(method, path string, obj interface{}) ([]GovernanceRoleSetting, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleSetting
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
			value  []GovernanceRoleSetting
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

// Get performs GET request for GovernanceRoleSetting collection
func (r *PrivilegedAccessRoleSettingsCollectionRequest) Get() ([]GovernanceRoleSetting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GovernanceRoleSetting collection
func (r *PrivilegedAccessRoleSettingsCollectionRequest) Add(reqObj *GovernanceRoleSetting) (resObj *GovernanceRoleSetting, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
