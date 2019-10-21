// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ItemActivityRequestBuilder is request builder for ItemActivity
type ItemActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ItemActivityRequest
func (b *ItemActivityRequestBuilder) Request() *ItemActivityRequest {
	return &ItemActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ItemActivityRequest is request for ItemActivity
type ItemActivityRequest struct{ BaseRequest }

// Do performs HTTP request for ItemActivity
func (r *ItemActivityRequest) Do(method, path string, reqObj interface{}) (resObj *ItemActivity, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ItemActivity
func (r *ItemActivityRequest) Get() (*ItemActivity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ItemActivity
func (r *ItemActivityRequest) Update(reqObj *ItemActivity) (*ItemActivity, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ItemActivity
func (r *ItemActivityRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DriveItem is navigation property
func (b *ItemActivityRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}