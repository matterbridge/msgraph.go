// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidWorkProfileCertificateProfileBaseRequestBuilder is request builder for AndroidWorkProfileCertificateProfileBase
type AndroidWorkProfileCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidWorkProfileCertificateProfileBaseRequest
func (b *AndroidWorkProfileCertificateProfileBaseRequestBuilder) Request() *AndroidWorkProfileCertificateProfileBaseRequest {
	return &AndroidWorkProfileCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidWorkProfileCertificateProfileBaseRequest is request for AndroidWorkProfileCertificateProfileBase
type AndroidWorkProfileCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *AndroidWorkProfileCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *AndroidWorkProfileCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidWorkProfileCertificateProfileBase
func (r *AndroidWorkProfileCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RootCertificate is navigation property
func (b *AndroidWorkProfileCertificateProfileBaseRequestBuilder) RootCertificate() *AndroidWorkProfileTrustedRootCertificateRequestBuilder {
	bb := &AndroidWorkProfileTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
