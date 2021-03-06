// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// CalendarSharingMessageRequestBuilder is request builder for CalendarSharingMessage
type CalendarSharingMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarSharingMessageRequest
func (b *CalendarSharingMessageRequestBuilder) Request() *CalendarSharingMessageRequest {
	return &CalendarSharingMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarSharingMessageRequest is request for CalendarSharingMessage
type CalendarSharingMessageRequest struct{ BaseRequest }

// Get performs GET request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Get(ctx context.Context) (resObj *CalendarSharingMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Update(ctx context.Context, reqObj *CalendarSharingMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
