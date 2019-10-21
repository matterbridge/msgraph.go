// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TargetedManagedAppConfigurationRequestBuilder is request builder for TargetedManagedAppConfiguration
type TargetedManagedAppConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppConfigurationRequest
func (b *TargetedManagedAppConfigurationRequestBuilder) Request() *TargetedManagedAppConfigurationRequest {
	return &TargetedManagedAppConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppConfigurationRequest is request for TargetedManagedAppConfiguration
type TargetedManagedAppConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *TargetedManagedAppConfiguration, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Get() (*TargetedManagedAppConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Update(reqObj *TargetedManagedAppConfiguration) (*TargetedManagedAppConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TargetedManagedAppConfiguration
func (r *TargetedManagedAppConfigurationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Apps returns request builder for ManagedMobileApp collection
func (b *TargetedManagedAppConfigurationRequestBuilder) Apps() *TargetedManagedAppConfigurationAppsCollectionRequestBuilder {
	bb := &TargetedManagedAppConfigurationAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/apps"
	return bb
}

// TargetedManagedAppConfigurationAppsCollectionRequestBuilder is request builder for ManagedMobileApp collection
type TargetedManagedAppConfigurationAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedMobileApp collection
func (b *TargetedManagedAppConfigurationAppsCollectionRequestBuilder) Request() *TargetedManagedAppConfigurationAppsCollectionRequest {
	return &TargetedManagedAppConfigurationAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedMobileApp item
func (b *TargetedManagedAppConfigurationAppsCollectionRequestBuilder) ID(id string) *ManagedMobileAppRequestBuilder {
	bb := &ManagedMobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppConfigurationAppsCollectionRequest is request for ManagedMobileApp collection
type TargetedManagedAppConfigurationAppsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedMobileApp, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedMobileApp, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedMobileApp
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ManagedMobileApp
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

// Get performs GET request for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Get() ([]ManagedMobileApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedMobileApp collection
func (r *TargetedManagedAppConfigurationAppsCollectionRequest) Add(reqObj *ManagedMobileApp) (*ManagedMobileApp, error) {
	return r.Do("POST", "", reqObj)
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppConfigurationRequestBuilder) Assignments() *TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder {
	bb := &TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder) Request() *TargetedManagedAppConfigurationAssignmentsCollectionRequest {
	return &TargetedManagedAppConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *TargetedManagedAppConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppConfigurationAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]TargetedManagedAppPolicyAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TargetedManagedAppPolicyAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TargetedManagedAppPolicyAssignment
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

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Get() ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppConfigurationAssignmentsCollectionRequest) Add(reqObj *TargetedManagedAppPolicyAssignment) (*TargetedManagedAppPolicyAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// DeploymentSummary is navigation property
func (b *TargetedManagedAppConfigurationRequestBuilder) DeploymentSummary() *ManagedAppPolicyDeploymentSummaryRequestBuilder {
	bb := &ManagedAppPolicyDeploymentSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deploymentSummary"
	return bb
}
