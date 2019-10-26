// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// Windows10PkcsCertificateProfileRequestBuilder is request builder for Windows10PkcsCertificateProfile
type Windows10PkcsCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows10PkcsCertificateProfileRequest
func (b *Windows10PkcsCertificateProfileRequestBuilder) Request() *Windows10PkcsCertificateProfileRequest {
	return &Windows10PkcsCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows10PkcsCertificateProfileRequest is request for Windows10PkcsCertificateProfile
type Windows10PkcsCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for Windows10PkcsCertificateProfile
func (r *Windows10PkcsCertificateProfileRequest) Get() (resObj *Windows10PkcsCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows10PkcsCertificateProfile
func (r *Windows10PkcsCertificateProfileRequest) Update(reqObj *Windows10PkcsCertificateProfile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows10PkcsCertificateProfile
func (r *Windows10PkcsCertificateProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *Windows10PkcsCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDeviceCertificateState
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
			value  []ManagedDeviceCertificateState
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

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *Windows10PkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
