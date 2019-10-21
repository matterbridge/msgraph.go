// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PictureRequestBuilder is request builder for Picture
type PictureRequestBuilder struct{ BaseRequestBuilder }

// Request returns PictureRequest
func (b *PictureRequestBuilder) Request() *PictureRequest {
	return &PictureRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PictureRequest is request for Picture
type PictureRequest struct{ BaseRequest }

// Do performs HTTP request for Picture
func (r *PictureRequest) Do(method, path string, reqObj interface{}) (resObj *Picture, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Picture
func (r *PictureRequest) Get() (*Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Picture
func (r *PictureRequest) Update(reqObj *Picture) (*Picture, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Picture
func (r *PictureRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}