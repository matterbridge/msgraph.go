// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// NotebookRequestBuilder is request builder for Notebook
type NotebookRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotebookRequest
func (b *NotebookRequestBuilder) Request() *NotebookRequest {
	return &NotebookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NotebookRequest is request for Notebook
type NotebookRequest struct{ BaseRequest }

// Get performs GET request for Notebook
func (r *NotebookRequest) Get(ctx context.Context) (resObj *Notebook, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Notebook
func (r *NotebookRequest) Update(ctx context.Context, reqObj *Notebook) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Notebook
func (r *NotebookRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SectionGroups returns request builder for SectionGroup collection
func (b *NotebookRequestBuilder) SectionGroups() *NotebookSectionGroupsCollectionRequestBuilder {
	bb := &NotebookSectionGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sectionGroups"
	return bb
}

// NotebookSectionGroupsCollectionRequestBuilder is request builder for SectionGroup collection
type NotebookSectionGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SectionGroup collection
func (b *NotebookSectionGroupsCollectionRequestBuilder) Request() *NotebookSectionGroupsCollectionRequest {
	return &NotebookSectionGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SectionGroup item
func (b *NotebookSectionGroupsCollectionRequestBuilder) ID(id string) *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// NotebookSectionGroupsCollectionRequest is request for SectionGroup collection
type NotebookSectionGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SectionGroup, error) {
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
	var values []SectionGroup
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
			value  []SectionGroup
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

// Get performs GET request for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Get(ctx context.Context) ([]SectionGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Add(ctx context.Context, reqObj *SectionGroup) (resObj *SectionGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sections returns request builder for OnenoteSection collection
func (b *NotebookRequestBuilder) Sections() *NotebookSectionsCollectionRequestBuilder {
	bb := &NotebookSectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sections"
	return bb
}

// NotebookSectionsCollectionRequestBuilder is request builder for OnenoteSection collection
type NotebookSectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteSection collection
func (b *NotebookSectionsCollectionRequestBuilder) Request() *NotebookSectionsCollectionRequest {
	return &NotebookSectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteSection item
func (b *NotebookSectionsCollectionRequestBuilder) ID(id string) *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// NotebookSectionsCollectionRequest is request for OnenoteSection collection
type NotebookSectionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnenoteSection, error) {
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
	var values []OnenoteSection
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
			value  []OnenoteSection
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

// Get performs GET request for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Get(ctx context.Context) ([]OnenoteSection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Add(ctx context.Context, reqObj *OnenoteSection) (resObj *OnenoteSection, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
