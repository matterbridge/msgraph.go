// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppInstallStatusRequestBuilder is request builder for MobileAppInstallStatus
type MobileAppInstallStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppInstallStatusRequest
func (b *MobileAppInstallStatusRequestBuilder) Request() *MobileAppInstallStatusRequest {
	return &MobileAppInstallStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppInstallStatusRequest is request for MobileAppInstallStatus
type MobileAppInstallStatusRequest struct{ BaseRequest }

// Do performs HTTP request for MobileAppInstallStatus
func (r *MobileAppInstallStatusRequest) Do(method, path string, reqObj interface{}) (resObj *MobileAppInstallStatus, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MobileAppInstallStatus
func (r *MobileAppInstallStatusRequest) Get() (*MobileAppInstallStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MobileAppInstallStatus
func (r *MobileAppInstallStatusRequest) Update(reqObj *MobileAppInstallStatus) (*MobileAppInstallStatus, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MobileAppInstallStatus
func (r *MobileAppInstallStatusRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// App is navigation property
func (b *MobileAppInstallStatusRequestBuilder) App() *MobileAppRequestBuilder {
	bb := &MobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/app"
	return bb
}