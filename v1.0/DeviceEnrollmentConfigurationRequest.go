// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// DeviceEnrollmentConfigurationRequestBuilder is request builder for DeviceEnrollmentConfiguration
type DeviceEnrollmentConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceEnrollmentConfigurationRequest
func (b *DeviceEnrollmentConfigurationRequestBuilder) Request() *DeviceEnrollmentConfigurationRequest {
	return &DeviceEnrollmentConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceEnrollmentConfigurationRequest is request for DeviceEnrollmentConfiguration
type DeviceEnrollmentConfigurationRequest struct{ BaseRequest }

// Get performs GET request for DeviceEnrollmentConfiguration
func (r *DeviceEnrollmentConfigurationRequest) Get(ctx context.Context) (resObj *DeviceEnrollmentConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceEnrollmentConfiguration
func (r *DeviceEnrollmentConfigurationRequest) Update(ctx context.Context, reqObj *DeviceEnrollmentConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceEnrollmentConfiguration
func (r *DeviceEnrollmentConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for EnrollmentConfigurationAssignment collection
func (b *DeviceEnrollmentConfigurationRequestBuilder) Assignments() *DeviceEnrollmentConfigurationAssignmentsCollectionRequestBuilder {
	bb := &DeviceEnrollmentConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// DeviceEnrollmentConfigurationAssignmentsCollectionRequestBuilder is request builder for EnrollmentConfigurationAssignment collection
type DeviceEnrollmentConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EnrollmentConfigurationAssignment collection
func (b *DeviceEnrollmentConfigurationAssignmentsCollectionRequestBuilder) Request() *DeviceEnrollmentConfigurationAssignmentsCollectionRequest {
	return &DeviceEnrollmentConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EnrollmentConfigurationAssignment item
func (b *DeviceEnrollmentConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *EnrollmentConfigurationAssignmentRequestBuilder {
	bb := &EnrollmentConfigurationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceEnrollmentConfigurationAssignmentsCollectionRequest is request for EnrollmentConfigurationAssignment collection
type DeviceEnrollmentConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EnrollmentConfigurationAssignment collection
func (r *DeviceEnrollmentConfigurationAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]EnrollmentConfigurationAssignment, error) {
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
	var values []EnrollmentConfigurationAssignment
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
			value  []EnrollmentConfigurationAssignment
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

// Get performs GET request for EnrollmentConfigurationAssignment collection
func (r *DeviceEnrollmentConfigurationAssignmentsCollectionRequest) Get(ctx context.Context) ([]EnrollmentConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for EnrollmentConfigurationAssignment collection
func (r *DeviceEnrollmentConfigurationAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *EnrollmentConfigurationAssignment) (resObj *EnrollmentConfigurationAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
