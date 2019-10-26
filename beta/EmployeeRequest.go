// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EmployeeRequestBuilder is request builder for Employee
type EmployeeRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmployeeRequest
func (b *EmployeeRequestBuilder) Request() *EmployeeRequest {
	return &EmployeeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmployeeRequest is request for Employee
type EmployeeRequest struct{ BaseRequest }

// Get performs GET request for Employee
func (r *EmployeeRequest) Get() (resObj *Employee, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Employee
func (r *EmployeeRequest) Update(reqObj *Employee) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Employee
func (r *EmployeeRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Picture returns request builder for Picture collection
func (b *EmployeeRequestBuilder) Picture() *EmployeePictureCollectionRequestBuilder {
	bb := &EmployeePictureCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/picture"
	return bb
}

// EmployeePictureCollectionRequestBuilder is request builder for Picture collection
type EmployeePictureCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Picture collection
func (b *EmployeePictureCollectionRequestBuilder) Request() *EmployeePictureCollectionRequest {
	return &EmployeePictureCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Picture item
func (b *EmployeePictureCollectionRequestBuilder) ID(id string) *PictureRequestBuilder {
	bb := &PictureRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmployeePictureCollectionRequest is request for Picture collection
type EmployeePictureCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Picture collection
func (r *EmployeePictureCollectionRequest) Paging(method, path string, obj interface{}) ([]Picture, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Picture
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
			value  []Picture
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

// Get performs GET request for Picture collection
func (r *EmployeePictureCollectionRequest) Get() ([]Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Picture collection
func (r *EmployeePictureCollectionRequest) Add(reqObj *Picture) (resObj *Picture, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
