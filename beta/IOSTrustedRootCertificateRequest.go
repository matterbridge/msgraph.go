// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IOSTrustedRootCertificateRequestBuilder is request builder for IOSTrustedRootCertificate
type IOSTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSTrustedRootCertificateRequest
func (b *IOSTrustedRootCertificateRequestBuilder) Request() *IOSTrustedRootCertificateRequest {
	return &IOSTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSTrustedRootCertificateRequest is request for IOSTrustedRootCertificate
type IOSTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for IOSTrustedRootCertificate
func (r *IOSTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *IOSTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSTrustedRootCertificate
func (r *IOSTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *IOSTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSTrustedRootCertificate
func (r *IOSTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
