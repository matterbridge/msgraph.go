// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidEnterpriseWiFiConfigurationRequestBuilder is request builder for AndroidEnterpriseWiFiConfiguration
type AndroidEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidEnterpriseWiFiConfigurationRequest
func (b *AndroidEnterpriseWiFiConfigurationRequestBuilder) Request() *AndroidEnterpriseWiFiConfigurationRequest {
	return &AndroidEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidEnterpriseWiFiConfigurationRequest is request for AndroidEnterpriseWiFiConfiguration
type AndroidEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Get(ctx context.Context) (resObj *AndroidEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Update(ctx context.Context, reqObj *AndroidEnterpriseWiFiConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *AndroidEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *AndroidCertificateProfileBaseRequestBuilder {
	bb := &AndroidCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *AndroidEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *AndroidTrustedRootCertificateRequestBuilder {
	bb := &AndroidTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}
