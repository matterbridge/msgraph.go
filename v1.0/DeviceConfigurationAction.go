// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceConfigurationAssignRequestParameter undocumented
type DeviceConfigurationAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []DeviceConfigurationAssignment `json:"assignments,omitempty"`
}

//
type DeviceConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceConfigurationRequestBuilder) Assign(reqObj *DeviceConfigurationAssignRequestParameter) *DeviceConfigurationAssignRequestBuilder {
	bb := &DeviceConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationAssignRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationAssignRequestBuilder) Request() *DeviceConfigurationAssignRequest {
	return &DeviceConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationAssignRequest) Do(method, path string, reqObj interface{}) (resObj *[]DeviceConfigurationAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *DeviceConfigurationAssignRequest) Paging(method, path string, obj interface{}) ([][]DeviceConfigurationAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]DeviceConfigurationAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]DeviceConfigurationAssignment
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
func (r *DeviceConfigurationAssignRequest) Get() ([][]DeviceConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
