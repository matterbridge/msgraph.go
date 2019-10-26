// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ServicePrincipalRequestBuilder is request builder for ServicePrincipal
type ServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ServicePrincipalRequest
func (b *ServicePrincipalRequestBuilder) Request() *ServicePrincipalRequest {
	return &ServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ServicePrincipalRequest is request for ServicePrincipal
type ServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for ServicePrincipal
func (r *ServicePrincipalRequest) Get() (resObj *ServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ServicePrincipal
func (r *ServicePrincipalRequest) Update(reqObj *ServicePrincipal) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ServicePrincipal
func (r *ServicePrincipalRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AppRoleAssignedTo returns request builder for AppRoleAssignment collection
func (b *ServicePrincipalRequestBuilder) AppRoleAssignedTo() *ServicePrincipalAppRoleAssignedToCollectionRequestBuilder {
	bb := &ServicePrincipalAppRoleAssignedToCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appRoleAssignedTo"
	return bb
}

// ServicePrincipalAppRoleAssignedToCollectionRequestBuilder is request builder for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignedToCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppRoleAssignment collection
func (b *ServicePrincipalAppRoleAssignedToCollectionRequestBuilder) Request() *ServicePrincipalAppRoleAssignedToCollectionRequest {
	return &ServicePrincipalAppRoleAssignedToCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppRoleAssignment item
func (b *ServicePrincipalAppRoleAssignedToCollectionRequestBuilder) ID(id string) *AppRoleAssignmentRequestBuilder {
	bb := &AppRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalAppRoleAssignedToCollectionRequest is request for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignedToCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) Paging(method, path string, obj interface{}) ([]AppRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AppRoleAssignment
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
			value  []AppRoleAssignment
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

// Get performs GET request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) Get() ([]AppRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) Add(reqObj *AppRoleAssignment) (resObj *AppRoleAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// AppRoleAssignments returns request builder for AppRoleAssignment collection
func (b *ServicePrincipalRequestBuilder) AppRoleAssignments() *ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder {
	bb := &ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appRoleAssignments"
	return bb
}

// ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder is request builder for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppRoleAssignment collection
func (b *ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder) Request() *ServicePrincipalAppRoleAssignmentsCollectionRequest {
	return &ServicePrincipalAppRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppRoleAssignment item
func (b *ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder) ID(id string) *AppRoleAssignmentRequestBuilder {
	bb := &AppRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalAppRoleAssignmentsCollectionRequest is request for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]AppRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AppRoleAssignment
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
			value  []AppRoleAssignment
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

// Get performs GET request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) Get() ([]AppRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) Add(reqObj *AppRoleAssignment) (resObj *AppRoleAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// CreatedObjects returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) CreatedObjects() *ServicePrincipalCreatedObjectsCollectionRequestBuilder {
	bb := &ServicePrincipalCreatedObjectsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdObjects"
	return bb
}

// ServicePrincipalCreatedObjectsCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalCreatedObjectsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalCreatedObjectsCollectionRequestBuilder) Request() *ServicePrincipalCreatedObjectsCollectionRequest {
	return &ServicePrincipalCreatedObjectsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalCreatedObjectsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalCreatedObjectsCollectionRequest is request for DirectoryObject collection
type ServicePrincipalCreatedObjectsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalCreatedObjectsCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalCreatedObjectsCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalCreatedObjectsCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// LicenseDetails returns request builder for LicenseDetails collection
func (b *ServicePrincipalRequestBuilder) LicenseDetails() *ServicePrincipalLicenseDetailsCollectionRequestBuilder {
	bb := &ServicePrincipalLicenseDetailsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/licenseDetails"
	return bb
}

// ServicePrincipalLicenseDetailsCollectionRequestBuilder is request builder for LicenseDetails collection
type ServicePrincipalLicenseDetailsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for LicenseDetails collection
func (b *ServicePrincipalLicenseDetailsCollectionRequestBuilder) Request() *ServicePrincipalLicenseDetailsCollectionRequest {
	return &ServicePrincipalLicenseDetailsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for LicenseDetails item
func (b *ServicePrincipalLicenseDetailsCollectionRequestBuilder) ID(id string) *LicenseDetailsRequestBuilder {
	bb := &LicenseDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalLicenseDetailsCollectionRequest is request for LicenseDetails collection
type ServicePrincipalLicenseDetailsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for LicenseDetails collection
func (r *ServicePrincipalLicenseDetailsCollectionRequest) Paging(method, path string, obj interface{}) ([]LicenseDetails, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []LicenseDetails
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
			value  []LicenseDetails
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

// Get performs GET request for LicenseDetails collection
func (r *ServicePrincipalLicenseDetailsCollectionRequest) Get() ([]LicenseDetails, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for LicenseDetails collection
func (r *ServicePrincipalLicenseDetailsCollectionRequest) Add(reqObj *LicenseDetails) (resObj *LicenseDetails, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// MemberOf returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) MemberOf() *ServicePrincipalMemberOfCollectionRequestBuilder {
	bb := &ServicePrincipalMemberOfCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/memberOf"
	return bb
}

// ServicePrincipalMemberOfCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalMemberOfCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalMemberOfCollectionRequestBuilder) Request() *ServicePrincipalMemberOfCollectionRequest {
	return &ServicePrincipalMemberOfCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalMemberOfCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalMemberOfCollectionRequest is request for DirectoryObject collection
type ServicePrincipalMemberOfCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalMemberOfCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalMemberOfCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalMemberOfCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Oauth2PermissionGrants returns request builder for OAuth2PermissionGrant collection
func (b *ServicePrincipalRequestBuilder) Oauth2PermissionGrants() *ServicePrincipalOauth2PermissionGrantsCollectionRequestBuilder {
	bb := &ServicePrincipalOauth2PermissionGrantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/oauth2PermissionGrants"
	return bb
}

// ServicePrincipalOauth2PermissionGrantsCollectionRequestBuilder is request builder for OAuth2PermissionGrant collection
type ServicePrincipalOauth2PermissionGrantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OAuth2PermissionGrant collection
func (b *ServicePrincipalOauth2PermissionGrantsCollectionRequestBuilder) Request() *ServicePrincipalOauth2PermissionGrantsCollectionRequest {
	return &ServicePrincipalOauth2PermissionGrantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OAuth2PermissionGrant item
func (b *ServicePrincipalOauth2PermissionGrantsCollectionRequestBuilder) ID(id string) *OAuth2PermissionGrantRequestBuilder {
	bb := &OAuth2PermissionGrantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalOauth2PermissionGrantsCollectionRequest is request for OAuth2PermissionGrant collection
type ServicePrincipalOauth2PermissionGrantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OAuth2PermissionGrant collection
func (r *ServicePrincipalOauth2PermissionGrantsCollectionRequest) Paging(method, path string, obj interface{}) ([]OAuth2PermissionGrant, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OAuth2PermissionGrant
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
			value  []OAuth2PermissionGrant
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

// Get performs GET request for OAuth2PermissionGrant collection
func (r *ServicePrincipalOauth2PermissionGrantsCollectionRequest) Get() ([]OAuth2PermissionGrant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OAuth2PermissionGrant collection
func (r *ServicePrincipalOauth2PermissionGrantsCollectionRequest) Add(reqObj *OAuth2PermissionGrant) (resObj *OAuth2PermissionGrant, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// OwnedObjects returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) OwnedObjects() *ServicePrincipalOwnedObjectsCollectionRequestBuilder {
	bb := &ServicePrincipalOwnedObjectsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/ownedObjects"
	return bb
}

// ServicePrincipalOwnedObjectsCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalOwnedObjectsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalOwnedObjectsCollectionRequestBuilder) Request() *ServicePrincipalOwnedObjectsCollectionRequest {
	return &ServicePrincipalOwnedObjectsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalOwnedObjectsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalOwnedObjectsCollectionRequest is request for DirectoryObject collection
type ServicePrincipalOwnedObjectsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalOwnedObjectsCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalOwnedObjectsCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalOwnedObjectsCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Owners returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) Owners() *ServicePrincipalOwnersCollectionRequestBuilder {
	bb := &ServicePrincipalOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/owners"
	return bb
}

// ServicePrincipalOwnersCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalOwnersCollectionRequestBuilder) Request() *ServicePrincipalOwnersCollectionRequest {
	return &ServicePrincipalOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalOwnersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalOwnersCollectionRequest is request for DirectoryObject collection
type ServicePrincipalOwnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalOwnersCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalOwnersCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalOwnersCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Policies returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) Policies() *ServicePrincipalPoliciesCollectionRequestBuilder {
	bb := &ServicePrincipalPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policies"
	return bb
}

// ServicePrincipalPoliciesCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalPoliciesCollectionRequestBuilder) Request() *ServicePrincipalPoliciesCollectionRequest {
	return &ServicePrincipalPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalPoliciesCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalPoliciesCollectionRequest is request for DirectoryObject collection
type ServicePrincipalPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalPoliciesCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalPoliciesCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalPoliciesCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Synchronization is navigation property
func (b *ServicePrincipalRequestBuilder) Synchronization() *SynchronizationRequestBuilder {
	bb := &SynchronizationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/synchronization"
	return bb
}

// TransitiveMemberOf returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) TransitiveMemberOf() *ServicePrincipalTransitiveMemberOfCollectionRequestBuilder {
	bb := &ServicePrincipalTransitiveMemberOfCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/transitiveMemberOf"
	return bb
}

// ServicePrincipalTransitiveMemberOfCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalTransitiveMemberOfCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalTransitiveMemberOfCollectionRequestBuilder) Request() *ServicePrincipalTransitiveMemberOfCollectionRequest {
	return &ServicePrincipalTransitiveMemberOfCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalTransitiveMemberOfCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalTransitiveMemberOfCollectionRequest is request for DirectoryObject collection
type ServicePrincipalTransitiveMemberOfCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
