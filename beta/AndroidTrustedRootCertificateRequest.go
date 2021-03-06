// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidTrustedRootCertificateRequestBuilder is request builder for AndroidTrustedRootCertificate
type AndroidTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidTrustedRootCertificateRequest
func (b *AndroidTrustedRootCertificateRequestBuilder) Request() *AndroidTrustedRootCertificateRequest {
	return &AndroidTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidTrustedRootCertificateRequest is request for AndroidTrustedRootCertificate
type AndroidTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for AndroidTrustedRootCertificate
func (r *AndroidTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *AndroidTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidTrustedRootCertificate
func (r *AndroidTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *AndroidTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidTrustedRootCertificate
func (r *AndroidTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
