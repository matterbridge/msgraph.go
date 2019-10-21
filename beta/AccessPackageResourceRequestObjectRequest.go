// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AccessPackageResourceRequestObjectRequestBuilder is request builder for AccessPackageResourceRequestObject
type AccessPackageResourceRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageResourceRequestObjectRequest
func (b *AccessPackageResourceRequestObjectRequestBuilder) Request() *AccessPackageResourceRequestObjectRequest {
	return &AccessPackageResourceRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageResourceRequestObjectRequest is request for AccessPackageResourceRequestObject
type AccessPackageResourceRequestObjectRequest struct{ BaseRequest }

// Do performs HTTP request for AccessPackageResourceRequestObject
func (r *AccessPackageResourceRequestObjectRequest) Do(method, path string, reqObj interface{}) (resObj *AccessPackageResourceRequestObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AccessPackageResourceRequestObject
func (r *AccessPackageResourceRequestObjectRequest) Get() (*AccessPackageResourceRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AccessPackageResourceRequestObject
func (r *AccessPackageResourceRequestObjectRequest) Update(reqObj *AccessPackageResourceRequestObject) (*AccessPackageResourceRequestObject, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AccessPackageResourceRequestObject
func (r *AccessPackageResourceRequestObjectRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AccessPackageResource is navigation property
func (b *AccessPackageResourceRequestObjectRequestBuilder) AccessPackageResource() *AccessPackageResourceRequestBuilder {
	bb := &AccessPackageResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResource"
	return bb
}

// Requestor is navigation property
func (b *AccessPackageResourceRequestObjectRequestBuilder) Requestor() *AccessPackageSubjectRequestBuilder {
	bb := &AccessPackageSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/requestor"
	return bb
}