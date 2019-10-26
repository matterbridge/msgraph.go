// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceManagementSettingCategoryRequestBuilder is request builder for DeviceManagementSettingCategory
type DeviceManagementSettingCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementSettingCategoryRequest
func (b *DeviceManagementSettingCategoryRequestBuilder) Request() *DeviceManagementSettingCategoryRequest {
	return &DeviceManagementSettingCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementSettingCategoryRequest is request for DeviceManagementSettingCategory
type DeviceManagementSettingCategoryRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementSettingCategory
func (r *DeviceManagementSettingCategoryRequest) Get() (resObj *DeviceManagementSettingCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementSettingCategory
func (r *DeviceManagementSettingCategoryRequest) Update(reqObj *DeviceManagementSettingCategory) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementSettingCategory
func (r *DeviceManagementSettingCategoryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// SettingDefinitions returns request builder for DeviceManagementSettingDefinition collection
func (b *DeviceManagementSettingCategoryRequestBuilder) SettingDefinitions() *DeviceManagementSettingCategorySettingDefinitionsCollectionRequestBuilder {
	bb := &DeviceManagementSettingCategorySettingDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/settingDefinitions"
	return bb
}

// DeviceManagementSettingCategorySettingDefinitionsCollectionRequestBuilder is request builder for DeviceManagementSettingDefinition collection
type DeviceManagementSettingCategorySettingDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementSettingDefinition collection
func (b *DeviceManagementSettingCategorySettingDefinitionsCollectionRequestBuilder) Request() *DeviceManagementSettingCategorySettingDefinitionsCollectionRequest {
	return &DeviceManagementSettingCategorySettingDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementSettingDefinition item
func (b *DeviceManagementSettingCategorySettingDefinitionsCollectionRequestBuilder) ID(id string) *DeviceManagementSettingDefinitionRequestBuilder {
	bb := &DeviceManagementSettingDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementSettingCategorySettingDefinitionsCollectionRequest is request for DeviceManagementSettingDefinition collection
type DeviceManagementSettingCategorySettingDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementSettingDefinition collection
func (r *DeviceManagementSettingCategorySettingDefinitionsCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementSettingDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementSettingDefinition
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
			value  []DeviceManagementSettingDefinition
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

// Get performs GET request for DeviceManagementSettingDefinition collection
func (r *DeviceManagementSettingCategorySettingDefinitionsCollectionRequest) Get() ([]DeviceManagementSettingDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementSettingDefinition collection
func (r *DeviceManagementSettingCategorySettingDefinitionsCollectionRequest) Add(reqObj *DeviceManagementSettingDefinition) (resObj *DeviceManagementSettingDefinition, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
