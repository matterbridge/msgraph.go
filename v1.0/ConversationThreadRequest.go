// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// ConversationThreadRequestBuilder is request builder for ConversationThread
type ConversationThreadRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConversationThreadRequest
func (b *ConversationThreadRequestBuilder) Request() *ConversationThreadRequest {
	return &ConversationThreadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConversationThreadRequest is request for ConversationThread
type ConversationThreadRequest struct{ BaseRequest }

// Get performs GET request for ConversationThread
func (r *ConversationThreadRequest) Get(ctx context.Context) (resObj *ConversationThread, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConversationThread
func (r *ConversationThreadRequest) Update(ctx context.Context, reqObj *ConversationThread) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConversationThread
func (r *ConversationThreadRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Posts returns request builder for Post collection
func (b *ConversationThreadRequestBuilder) Posts() *ConversationThreadPostsCollectionRequestBuilder {
	bb := &ConversationThreadPostsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/posts"
	return bb
}

// ConversationThreadPostsCollectionRequestBuilder is request builder for Post collection
type ConversationThreadPostsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Post collection
func (b *ConversationThreadPostsCollectionRequestBuilder) Request() *ConversationThreadPostsCollectionRequest {
	return &ConversationThreadPostsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Post item
func (b *ConversationThreadPostsCollectionRequestBuilder) ID(id string) *PostRequestBuilder {
	bb := &PostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConversationThreadPostsCollectionRequest is request for Post collection
type ConversationThreadPostsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Post collection
func (r *ConversationThreadPostsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Post, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Post
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Post
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for Post collection
func (r *ConversationThreadPostsCollectionRequest) Get(ctx context.Context) ([]Post, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Post collection
func (r *ConversationThreadPostsCollectionRequest) Add(ctx context.Context, reqObj *Post) (resObj *Post, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
