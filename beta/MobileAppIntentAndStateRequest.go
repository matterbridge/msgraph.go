// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppIntentAndStateRequestBuilder is request builder for MobileAppIntentAndState
type MobileAppIntentAndStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppIntentAndStateRequest
func (b *MobileAppIntentAndStateRequestBuilder) Request() *MobileAppIntentAndStateRequest {
	return &MobileAppIntentAndStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppIntentAndStateRequest is request for MobileAppIntentAndState
type MobileAppIntentAndStateRequest struct{ BaseRequest }

// Do performs HTTP request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Do(method, path string, reqObj interface{}) (resObj *MobileAppIntentAndState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Get() (*MobileAppIntentAndState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Update(reqObj *MobileAppIntentAndState) (*MobileAppIntentAndState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MobileAppIntentAndState
func (r *MobileAppIntentAndStateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}