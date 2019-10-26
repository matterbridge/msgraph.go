// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// IOSEnterpriseWiFiConfigurationRequestBuilder is request builder for IOSEnterpriseWiFiConfiguration
type IOSEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSEnterpriseWiFiConfigurationRequest
func (b *IOSEnterpriseWiFiConfigurationRequestBuilder) Request() *IOSEnterpriseWiFiConfigurationRequest {
	return &IOSEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSEnterpriseWiFiConfigurationRequest is request for IOSEnterpriseWiFiConfiguration
type IOSEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSEnterpriseWiFiConfiguration
func (r *IOSEnterpriseWiFiConfigurationRequest) Get() (resObj *IOSEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSEnterpriseWiFiConfiguration
func (r *IOSEnterpriseWiFiConfigurationRequest) Update(reqObj *IOSEnterpriseWiFiConfiguration) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSEnterpriseWiFiConfiguration
func (r *IOSEnterpriseWiFiConfigurationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DerivedCredentialSettings is navigation property
func (b *IOSEnterpriseWiFiConfigurationRequestBuilder) DerivedCredentialSettings() *DeviceManagementDerivedCredentialSettingsRequestBuilder {
	bb := &DeviceManagementDerivedCredentialSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/derivedCredentialSettings"
	return bb
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *IOSEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *IOSCertificateProfileBaseRequestBuilder {
	bb := &IOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificatesForServerValidation returns request builder for IOSTrustedRootCertificate collection
func (b *IOSEnterpriseWiFiConfigurationRequestBuilder) RootCertificatesForServerValidation() *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequestBuilder {
	bb := &IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificatesForServerValidation"
	return bb
}

// IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequestBuilder is request builder for IOSTrustedRootCertificate collection
type IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IOSTrustedRootCertificate collection
func (b *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequestBuilder) Request() *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest {
	return &IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IOSTrustedRootCertificate item
func (b *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequestBuilder) ID(id string) *IOSTrustedRootCertificateRequestBuilder {
	bb := &IOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest is request for IOSTrustedRootCertificate collection
type IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IOSTrustedRootCertificate collection
func (r *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest) Paging(method, path string, obj interface{}) ([]IOSTrustedRootCertificate, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []IOSTrustedRootCertificate
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
			value  []IOSTrustedRootCertificate
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

// Get performs GET request for IOSTrustedRootCertificate collection
func (r *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest) Get() ([]IOSTrustedRootCertificate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for IOSTrustedRootCertificate collection
func (r *IOSEnterpriseWiFiConfigurationRootCertificatesForServerValidationCollectionRequest) Add(reqObj *IOSTrustedRootCertificate) (resObj *IOSTrustedRootCertificate, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
