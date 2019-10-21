// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceMobileAppConfigurationStateRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationState
type ManagedDeviceMobileAppConfigurationStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationStateRequest
func (b *ManagedDeviceMobileAppConfigurationStateRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationStateRequest {
	return &ManagedDeviceMobileAppConfigurationStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationStateRequest is request for ManagedDeviceMobileAppConfigurationState
type ManagedDeviceMobileAppConfigurationStateRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceMobileAppConfigurationState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Get() (*ManagedDeviceMobileAppConfigurationState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Update(reqObj *ManagedDeviceMobileAppConfigurationState) (*ManagedDeviceMobileAppConfigurationState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationState
func (r *ManagedDeviceMobileAppConfigurationStateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}