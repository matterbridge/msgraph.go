// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceRequestBuilder is request builder for Device
type DeviceRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceRequest
func (b *DeviceRequestBuilder) Request() *DeviceRequest {
	return &DeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceRequest is request for Device
type DeviceRequest struct{ BaseRequest }

// Do performs HTTP request for Device
func (r *DeviceRequest) Do(method, path string, reqObj interface{}) (resObj *Device, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Device
func (r *DeviceRequest) Get() (*Device, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Device
func (r *DeviceRequest) Update(reqObj *Device) (*Device, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Device
func (r *DeviceRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Commands returns request builder for Command collection
func (b *DeviceRequestBuilder) Commands() *DeviceCommandsCollectionRequestBuilder {
	bb := &DeviceCommandsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/commands"
	return bb
}

// DeviceCommandsCollectionRequestBuilder is request builder for Command collection
type DeviceCommandsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Command collection
func (b *DeviceCommandsCollectionRequestBuilder) Request() *DeviceCommandsCollectionRequest {
	return &DeviceCommandsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Command item
func (b *DeviceCommandsCollectionRequestBuilder) ID(id string) *CommandRequestBuilder {
	bb := &CommandRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceCommandsCollectionRequest is request for Command collection
type DeviceCommandsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Command collection
func (r *DeviceCommandsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Command, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Command collection
func (r *DeviceCommandsCollectionRequest) Paging(method, path string, obj interface{}) ([]Command, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Command
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Command
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

// Get performs GET request for Command collection
func (r *DeviceCommandsCollectionRequest) Get() ([]Command, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Command collection
func (r *DeviceCommandsCollectionRequest) Add(reqObj *Command) (*Command, error) {
	return r.Do("POST", "", reqObj)
}

// Extensions returns request builder for Extension collection
func (b *DeviceRequestBuilder) Extensions() *DeviceExtensionsCollectionRequestBuilder {
	bb := &DeviceExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// DeviceExtensionsCollectionRequestBuilder is request builder for Extension collection
type DeviceExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *DeviceExtensionsCollectionRequestBuilder) Request() *DeviceExtensionsCollectionRequest {
	return &DeviceExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *DeviceExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceExtensionsCollectionRequest is request for Extension collection
type DeviceExtensionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Extension collection
func (r *DeviceExtensionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Extension, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Extension collection
func (r *DeviceExtensionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Extension, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Extension
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *DeviceExtensionsCollectionRequest) Get() ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *DeviceExtensionsCollectionRequest) Add(reqObj *Extension) (*Extension, error) {
	return r.Do("POST", "", reqObj)
}

// MemberOf returns request builder for DirectoryObject collection
func (b *DeviceRequestBuilder) MemberOf() *DeviceMemberOfCollectionRequestBuilder {
	bb := &DeviceMemberOfCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/memberOf"
	return bb
}

// DeviceMemberOfCollectionRequestBuilder is request builder for DirectoryObject collection
type DeviceMemberOfCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DeviceMemberOfCollectionRequestBuilder) Request() *DeviceMemberOfCollectionRequest {
	return &DeviceMemberOfCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DeviceMemberOfCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceMemberOfCollectionRequest is request for DirectoryObject collection
type DeviceMemberOfCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *DeviceMemberOfCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *DeviceMemberOfCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *DeviceMemberOfCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DeviceMemberOfCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}

// RegisteredOwners returns request builder for DirectoryObject collection
func (b *DeviceRequestBuilder) RegisteredOwners() *DeviceRegisteredOwnersCollectionRequestBuilder {
	bb := &DeviceRegisteredOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/registeredOwners"
	return bb
}

// DeviceRegisteredOwnersCollectionRequestBuilder is request builder for DirectoryObject collection
type DeviceRegisteredOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DeviceRegisteredOwnersCollectionRequestBuilder) Request() *DeviceRegisteredOwnersCollectionRequest {
	return &DeviceRegisteredOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DeviceRegisteredOwnersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceRegisteredOwnersCollectionRequest is request for DirectoryObject collection
type DeviceRegisteredOwnersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *DeviceRegisteredOwnersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *DeviceRegisteredOwnersCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *DeviceRegisteredOwnersCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DeviceRegisteredOwnersCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}

// RegisteredUsers returns request builder for DirectoryObject collection
func (b *DeviceRequestBuilder) RegisteredUsers() *DeviceRegisteredUsersCollectionRequestBuilder {
	bb := &DeviceRegisteredUsersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/registeredUsers"
	return bb
}

// DeviceRegisteredUsersCollectionRequestBuilder is request builder for DirectoryObject collection
type DeviceRegisteredUsersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DeviceRegisteredUsersCollectionRequestBuilder) Request() *DeviceRegisteredUsersCollectionRequest {
	return &DeviceRegisteredUsersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DeviceRegisteredUsersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceRegisteredUsersCollectionRequest is request for DirectoryObject collection
type DeviceRegisteredUsersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *DeviceRegisteredUsersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *DeviceRegisteredUsersCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *DeviceRegisteredUsersCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DeviceRegisteredUsersCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}

// TransitiveMemberOf returns request builder for DirectoryObject collection
func (b *DeviceRequestBuilder) TransitiveMemberOf() *DeviceTransitiveMemberOfCollectionRequestBuilder {
	bb := &DeviceTransitiveMemberOfCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/transitiveMemberOf"
	return bb
}

// DeviceTransitiveMemberOfCollectionRequestBuilder is request builder for DirectoryObject collection
type DeviceTransitiveMemberOfCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DeviceTransitiveMemberOfCollectionRequestBuilder) Request() *DeviceTransitiveMemberOfCollectionRequest {
	return &DeviceTransitiveMemberOfCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DeviceTransitiveMemberOfCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceTransitiveMemberOfCollectionRequest is request for DirectoryObject collection
type DeviceTransitiveMemberOfCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *DeviceTransitiveMemberOfCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *DeviceTransitiveMemberOfCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *DeviceTransitiveMemberOfCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DeviceTransitiveMemberOfCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}
