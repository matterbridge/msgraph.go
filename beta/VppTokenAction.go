// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VppTokenSyncLicensesRequestParameter undocumented
type VppTokenSyncLicensesRequestParameter struct {
}

// VppTokenRevokeLicensesRequestParameter undocumented
type VppTokenRevokeLicensesRequestParameter struct {
	// NotifyManagedDevices undocumented
	NotifyManagedDevices *bool `json:"notifyManagedDevices,omitempty"`
}

//
type VppTokenSyncLicensesRequestBuilder struct{ BaseRequestBuilder }

// SyncLicenses action undocumented
func (b *VppTokenRequestBuilder) SyncLicenses(reqObj *VppTokenSyncLicensesRequestParameter) *VppTokenSyncLicensesRequestBuilder {
	bb := &VppTokenSyncLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/syncLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type VppTokenSyncLicensesRequest struct{ BaseRequest }

//
func (b *VppTokenSyncLicensesRequestBuilder) Request() *VppTokenSyncLicensesRequest {
	return &VppTokenSyncLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *VppTokenSyncLicensesRequest) Do(method, path string, reqObj interface{}) (resObj *VppToken, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *VppTokenSyncLicensesRequest) Post() (*VppToken, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type VppTokenRevokeLicensesRequestBuilder struct{ BaseRequestBuilder }

// RevokeLicenses action undocumented
func (b *VppTokenRequestBuilder) RevokeLicenses(reqObj *VppTokenRevokeLicensesRequestParameter) *VppTokenRevokeLicensesRequestBuilder {
	bb := &VppTokenRevokeLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type VppTokenRevokeLicensesRequest struct{ BaseRequest }

//
func (b *VppTokenRevokeLicensesRequestBuilder) Request() *VppTokenRevokeLicensesRequest {
	return &VppTokenRevokeLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *VppTokenRevokeLicensesRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *VppTokenRevokeLicensesRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}