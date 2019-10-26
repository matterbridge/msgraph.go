// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ListItemRequestBuilder is request builder for ListItem
type ListItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ListItemRequest
func (b *ListItemRequestBuilder) Request() *ListItemRequest {
	return &ListItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ListItemRequest is request for ListItem
type ListItemRequest struct{ BaseRequest }

// Get performs GET request for ListItem
func (r *ListItemRequest) Get() (resObj *ListItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ListItem
func (r *ListItemRequest) Update(reqObj *ListItem) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ListItem
func (r *ListItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Activities returns request builder for ItemActivityOLD collection
func (b *ListItemRequestBuilder) Activities() *ListItemActivitiesCollectionRequestBuilder {
	bb := &ListItemActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// ListItemActivitiesCollectionRequestBuilder is request builder for ItemActivityOLD collection
type ListItemActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityOLD collection
func (b *ListItemActivitiesCollectionRequestBuilder) Request() *ListItemActivitiesCollectionRequest {
	return &ListItemActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityOLD item
func (b *ListItemActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityOLDRequestBuilder {
	bb := &ItemActivityOLDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemActivitiesCollectionRequest is request for ItemActivityOLD collection
type ListItemActivitiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ItemActivityOLD collection
func (r *ListItemActivitiesCollectionRequest) Paging(method, path string, obj interface{}) ([]ItemActivityOLD, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ItemActivityOLD
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
			value  []ItemActivityOLD
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
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ItemActivityOLD collection
func (r *ListItemActivitiesCollectionRequest) Get() ([]ItemActivityOLD, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ItemActivityOLD collection
func (r *ListItemActivitiesCollectionRequest) Add(reqObj *ItemActivityOLD) (resObj *ItemActivityOLD, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Analytics is navigation property
func (b *ListItemRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// DriveItem is navigation property
func (b *ListItemRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// Fields is navigation property
func (b *ListItemRequestBuilder) Fields() *FieldValueSetRequestBuilder {
	bb := &FieldValueSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fields"
	return bb
}

// Versions returns request builder for ListItemVersion collection
func (b *ListItemRequestBuilder) Versions() *ListItemVersionsCollectionRequestBuilder {
	bb := &ListItemVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/versions"
	return bb
}

// ListItemVersionsCollectionRequestBuilder is request builder for ListItemVersion collection
type ListItemVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItemVersion collection
func (b *ListItemVersionsCollectionRequestBuilder) Request() *ListItemVersionsCollectionRequest {
	return &ListItemVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItemVersion item
func (b *ListItemVersionsCollectionRequestBuilder) ID(id string) *ListItemVersionRequestBuilder {
	bb := &ListItemVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemVersionsCollectionRequest is request for ListItemVersion collection
type ListItemVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Paging(method, path string, obj interface{}) ([]ListItemVersion, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ListItemVersion
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
			value  []ListItemVersion
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
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Get() ([]ListItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Add(reqObj *ListItemVersion) (resObj *ListItemVersion, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
