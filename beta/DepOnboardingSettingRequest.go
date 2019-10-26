// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DepOnboardingSettingRequestBuilder is request builder for DepOnboardingSetting
type DepOnboardingSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepOnboardingSettingRequest
func (b *DepOnboardingSettingRequestBuilder) Request() *DepOnboardingSettingRequest {
	return &DepOnboardingSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepOnboardingSettingRequest is request for DepOnboardingSetting
type DepOnboardingSettingRequest struct{ BaseRequest }

// Get performs GET request for DepOnboardingSetting
func (r *DepOnboardingSettingRequest) Get() (resObj *DepOnboardingSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DepOnboardingSetting
func (r *DepOnboardingSettingRequest) Update(reqObj *DepOnboardingSetting) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DepOnboardingSetting
func (r *DepOnboardingSettingRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DefaultIOSEnrollmentProfile is navigation property
func (b *DepOnboardingSettingRequestBuilder) DefaultIOSEnrollmentProfile() *DepIOSEnrollmentProfileRequestBuilder {
	bb := &DepIOSEnrollmentProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultIosEnrollmentProfile"
	return bb
}

// DefaultMacOsEnrollmentProfile is navigation property
func (b *DepOnboardingSettingRequestBuilder) DefaultMacOsEnrollmentProfile() *DepMacOSEnrollmentProfileRequestBuilder {
	bb := &DepMacOSEnrollmentProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultMacOsEnrollmentProfile"
	return bb
}

// EnrollmentProfiles returns request builder for EnrollmentProfile collection
func (b *DepOnboardingSettingRequestBuilder) EnrollmentProfiles() *DepOnboardingSettingEnrollmentProfilesCollectionRequestBuilder {
	bb := &DepOnboardingSettingEnrollmentProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/enrollmentProfiles"
	return bb
}

// DepOnboardingSettingEnrollmentProfilesCollectionRequestBuilder is request builder for EnrollmentProfile collection
type DepOnboardingSettingEnrollmentProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EnrollmentProfile collection
func (b *DepOnboardingSettingEnrollmentProfilesCollectionRequestBuilder) Request() *DepOnboardingSettingEnrollmentProfilesCollectionRequest {
	return &DepOnboardingSettingEnrollmentProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EnrollmentProfile item
func (b *DepOnboardingSettingEnrollmentProfilesCollectionRequestBuilder) ID(id string) *EnrollmentProfileRequestBuilder {
	bb := &EnrollmentProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DepOnboardingSettingEnrollmentProfilesCollectionRequest is request for EnrollmentProfile collection
type DepOnboardingSettingEnrollmentProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EnrollmentProfile collection
func (r *DepOnboardingSettingEnrollmentProfilesCollectionRequest) Paging(method, path string, obj interface{}) ([]EnrollmentProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EnrollmentProfile
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
			value  []EnrollmentProfile
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

// Get performs GET request for EnrollmentProfile collection
func (r *DepOnboardingSettingEnrollmentProfilesCollectionRequest) Get() ([]EnrollmentProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EnrollmentProfile collection
func (r *DepOnboardingSettingEnrollmentProfilesCollectionRequest) Add(reqObj *EnrollmentProfile) (resObj *EnrollmentProfile, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ImportedAppleDeviceIdentities returns request builder for ImportedAppleDeviceIdentity collection
func (b *DepOnboardingSettingRequestBuilder) ImportedAppleDeviceIdentities() *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequestBuilder {
	bb := &DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/importedAppleDeviceIdentities"
	return bb
}

// DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequestBuilder is request builder for ImportedAppleDeviceIdentity collection
type DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ImportedAppleDeviceIdentity collection
func (b *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequestBuilder) Request() *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest {
	return &DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ImportedAppleDeviceIdentity item
func (b *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequestBuilder) ID(id string) *ImportedAppleDeviceIdentityRequestBuilder {
	bb := &ImportedAppleDeviceIdentityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest is request for ImportedAppleDeviceIdentity collection
type DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ImportedAppleDeviceIdentity collection
func (r *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest) Paging(method, path string, obj interface{}) ([]ImportedAppleDeviceIdentity, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ImportedAppleDeviceIdentity
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
			value  []ImportedAppleDeviceIdentity
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

// Get performs GET request for ImportedAppleDeviceIdentity collection
func (r *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest) Get() ([]ImportedAppleDeviceIdentity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ImportedAppleDeviceIdentity collection
func (r *DepOnboardingSettingImportedAppleDeviceIdentitiesCollectionRequest) Add(reqObj *ImportedAppleDeviceIdentity) (resObj *ImportedAppleDeviceIdentity, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
