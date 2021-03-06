// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessReviewDecisionRequestBuilder is request builder for AccessReviewDecision
type AccessReviewDecisionRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessReviewDecisionRequest
func (b *AccessReviewDecisionRequestBuilder) Request() *AccessReviewDecisionRequest {
	return &AccessReviewDecisionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessReviewDecisionRequest is request for AccessReviewDecision
type AccessReviewDecisionRequest struct{ BaseRequest }

// Get performs GET request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Get(ctx context.Context) (resObj *AccessReviewDecision, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Update(ctx context.Context, reqObj *AccessReviewDecision) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessReviewDecision
func (r *AccessReviewDecisionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
