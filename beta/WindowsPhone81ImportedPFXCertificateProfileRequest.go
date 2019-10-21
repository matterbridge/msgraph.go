// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WindowsPhone81ImportedPFXCertificateProfileRequestBuilder is request builder for WindowsPhone81ImportedPFXCertificateProfile
type WindowsPhone81ImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsPhone81ImportedPFXCertificateProfileRequest
func (b *WindowsPhone81ImportedPFXCertificateProfileRequestBuilder) Request() *WindowsPhone81ImportedPFXCertificateProfileRequest {
	return &WindowsPhone81ImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsPhone81ImportedPFXCertificateProfileRequest is request for WindowsPhone81ImportedPFXCertificateProfile
type WindowsPhone81ImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsPhone81ImportedPFXCertificateProfile
func (r *WindowsPhone81ImportedPFXCertificateProfileRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsPhone81ImportedPFXCertificateProfile, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsPhone81ImportedPFXCertificateProfile
func (r *WindowsPhone81ImportedPFXCertificateProfileRequest) Get() (*WindowsPhone81ImportedPFXCertificateProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsPhone81ImportedPFXCertificateProfile
func (r *WindowsPhone81ImportedPFXCertificateProfileRequest) Update(reqObj *WindowsPhone81ImportedPFXCertificateProfile) (*WindowsPhone81ImportedPFXCertificateProfile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsPhone81ImportedPFXCertificateProfile
func (r *WindowsPhone81ImportedPFXCertificateProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *WindowsPhone81ImportedPFXCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceCertificateState collection
func (r *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
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
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
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
func (r *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *WindowsPhone81ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (*ManagedDeviceCertificateState, error) {
	return r.Do("POST", "", reqObj)
}
