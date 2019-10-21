// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter undocumented
type DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// DeviceEnrollmentConfigurationSetPriorityRequestParameter undocumented
type DeviceEnrollmentConfigurationSetPriorityRequestParameter struct {
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
}

// DeviceEnrollmentConfigurationAssignRequestParameter undocumented
type DeviceEnrollmentConfigurationAssignRequestParameter struct {
	// EnrollmentConfigurationAssignments undocumented
	EnrollmentConfigurationAssignments []EnrollmentConfigurationAssignment `json:"enrollmentConfigurationAssignments,omitempty"`
}

//
type DeviceEnrollmentConfigurationSetPriorityRequestBuilder struct{ BaseRequestBuilder }

// SetPriority action undocumented
func (b *DeviceEnrollmentConfigurationRequestBuilder) SetPriority(reqObj *DeviceEnrollmentConfigurationSetPriorityRequestParameter) *DeviceEnrollmentConfigurationSetPriorityRequestBuilder {
	bb := &DeviceEnrollmentConfigurationSetPriorityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setPriority"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceEnrollmentConfigurationSetPriorityRequest struct{ BaseRequest }

//
func (b *DeviceEnrollmentConfigurationSetPriorityRequestBuilder) Request() *DeviceEnrollmentConfigurationSetPriorityRequest {
	return &DeviceEnrollmentConfigurationSetPriorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceEnrollmentConfigurationSetPriorityRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *DeviceEnrollmentConfigurationSetPriorityRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type DeviceEnrollmentConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceEnrollmentConfigurationRequestBuilder) Assign(reqObj *DeviceEnrollmentConfigurationAssignRequestParameter) *DeviceEnrollmentConfigurationAssignRequestBuilder {
	bb := &DeviceEnrollmentConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceEnrollmentConfigurationAssignRequest struct{ BaseRequest }

//
func (b *DeviceEnrollmentConfigurationAssignRequestBuilder) Request() *DeviceEnrollmentConfigurationAssignRequest {
	return &DeviceEnrollmentConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceEnrollmentConfigurationAssignRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *DeviceEnrollmentConfigurationAssignRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceManagementDeviceEnrollmentConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter) *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// HasPayloadLinks action undocumented
func (b *UserDeviceEnrollmentConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter) *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder) Request() *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest {
	return &DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest) Do(method, path string, reqObj interface{}) (resObj *[]HasPayloadLinkResultItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest) Paging(method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]HasPayloadLinkResultItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]HasPayloadLinkResultItem
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

//
func (r *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest) Get() ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
