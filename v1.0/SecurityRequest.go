// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SecurityRequestBuilder is request builder for Security
type SecurityRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityRequest
func (b *SecurityRequestBuilder) Request() *SecurityRequest {
	return &SecurityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityRequest is request for Security
type SecurityRequest struct{ BaseRequest }

// Get performs GET request for Security
func (r *SecurityRequest) Get() (resObj *Security, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Security
func (r *SecurityRequest) Update(reqObj *Security) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Security
func (r *SecurityRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Alerts returns request builder for Alert collection
func (b *SecurityRequestBuilder) Alerts() *SecurityAlertsCollectionRequestBuilder {
	bb := &SecurityAlertsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/alerts"
	return bb
}

// SecurityAlertsCollectionRequestBuilder is request builder for Alert collection
type SecurityAlertsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Alert collection
func (b *SecurityAlertsCollectionRequestBuilder) Request() *SecurityAlertsCollectionRequest {
	return &SecurityAlertsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Alert item
func (b *SecurityAlertsCollectionRequestBuilder) ID(id string) *AlertRequestBuilder {
	bb := &AlertRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityAlertsCollectionRequest is request for Alert collection
type SecurityAlertsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Alert collection
func (r *SecurityAlertsCollectionRequest) Paging(method, path string, obj interface{}) ([]Alert, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Alert
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
			value  []Alert
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

// Get performs GET request for Alert collection
func (r *SecurityAlertsCollectionRequest) Get() ([]Alert, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Alert collection
func (r *SecurityAlertsCollectionRequest) Add(reqObj *Alert) (resObj *Alert, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// SecureScoreControlProfiles returns request builder for SecureScoreControlProfile collection
func (b *SecurityRequestBuilder) SecureScoreControlProfiles() *SecuritySecureScoreControlProfilesCollectionRequestBuilder {
	bb := &SecuritySecureScoreControlProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/secureScoreControlProfiles"
	return bb
}

// SecuritySecureScoreControlProfilesCollectionRequestBuilder is request builder for SecureScoreControlProfile collection
type SecuritySecureScoreControlProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecureScoreControlProfile collection
func (b *SecuritySecureScoreControlProfilesCollectionRequestBuilder) Request() *SecuritySecureScoreControlProfilesCollectionRequest {
	return &SecuritySecureScoreControlProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecureScoreControlProfile item
func (b *SecuritySecureScoreControlProfilesCollectionRequestBuilder) ID(id string) *SecureScoreControlProfileRequestBuilder {
	bb := &SecureScoreControlProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecureScoreControlProfilesCollectionRequest is request for SecureScoreControlProfile collection
type SecuritySecureScoreControlProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Paging(method, path string, obj interface{}) ([]SecureScoreControlProfile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SecureScoreControlProfile
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
			value  []SecureScoreControlProfile
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

// Get performs GET request for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Get() ([]SecureScoreControlProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Add(reqObj *SecureScoreControlProfile) (resObj *SecureScoreControlProfile, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// SecureScores returns request builder for SecureScore collection
func (b *SecurityRequestBuilder) SecureScores() *SecuritySecureScoresCollectionRequestBuilder {
	bb := &SecuritySecureScoresCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/secureScores"
	return bb
}

// SecuritySecureScoresCollectionRequestBuilder is request builder for SecureScore collection
type SecuritySecureScoresCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecureScore collection
func (b *SecuritySecureScoresCollectionRequestBuilder) Request() *SecuritySecureScoresCollectionRequest {
	return &SecuritySecureScoresCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecureScore item
func (b *SecuritySecureScoresCollectionRequestBuilder) ID(id string) *SecureScoreRequestBuilder {
	bb := &SecureScoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecureScoresCollectionRequest is request for SecureScore collection
type SecuritySecureScoresCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Paging(method, path string, obj interface{}) ([]SecureScore, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SecureScore
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
			value  []SecureScore
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

// Get performs GET request for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Get() ([]SecureScore, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Add(reqObj *SecureScore) (resObj *SecureScore, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
