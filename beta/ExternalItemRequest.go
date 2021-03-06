// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ExternalItemRequestBuilder is request builder for ExternalItem
type ExternalItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalItemRequest
func (b *ExternalItemRequestBuilder) Request() *ExternalItemRequest {
	return &ExternalItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExternalItemRequest is request for ExternalItem
type ExternalItemRequest struct{ BaseRequest }

// Get performs GET request for ExternalItem
func (r *ExternalItemRequest) Get(ctx context.Context) (resObj *ExternalItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExternalItem
func (r *ExternalItemRequest) Update(ctx context.Context, reqObj *ExternalItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExternalItem
func (r *ExternalItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
