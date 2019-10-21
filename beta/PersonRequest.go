// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PersonRequestBuilder is request builder for Person
type PersonRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonRequest
func (b *PersonRequestBuilder) Request() *PersonRequest {
	return &PersonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonRequest is request for Person
type PersonRequest struct{ BaseRequest }

// Do performs HTTP request for Person
func (r *PersonRequest) Do(method, path string, reqObj interface{}) (resObj *Person, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Person
func (r *PersonRequest) Get() (*Person, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Person
func (r *PersonRequest) Update(reqObj *Person) (*Person, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Person
func (r *PersonRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}