// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PlannerBucketTaskBoardTaskFormatRequestBuilder is request builder for PlannerBucketTaskBoardTaskFormat
type PlannerBucketTaskBoardTaskFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerBucketTaskBoardTaskFormatRequest
func (b *PlannerBucketTaskBoardTaskFormatRequestBuilder) Request() *PlannerBucketTaskBoardTaskFormatRequest {
	return &PlannerBucketTaskBoardTaskFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerBucketTaskBoardTaskFormatRequest is request for PlannerBucketTaskBoardTaskFormat
type PlannerBucketTaskBoardTaskFormatRequest struct{ BaseRequest }

// Get performs GET request for PlannerBucketTaskBoardTaskFormat
func (r *PlannerBucketTaskBoardTaskFormatRequest) Get(ctx context.Context) (resObj *PlannerBucketTaskBoardTaskFormat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PlannerBucketTaskBoardTaskFormat
func (r *PlannerBucketTaskBoardTaskFormatRequest) Update(ctx context.Context, reqObj *PlannerBucketTaskBoardTaskFormat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PlannerBucketTaskBoardTaskFormat
func (r *PlannerBucketTaskBoardTaskFormatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
