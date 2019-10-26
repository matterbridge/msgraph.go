// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder is request builder for AdvancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdvancedThreatProtectionOnboardingStateSummaryRequest
func (b *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Request() *AdvancedThreatProtectionOnboardingStateSummaryRequest {
	return &AdvancedThreatProtectionOnboardingStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AdvancedThreatProtectionOnboardingStateSummaryRequest is request for AdvancedThreatProtectionOnboardingStateSummary
type AdvancedThreatProtectionOnboardingStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Get() (resObj *AdvancedThreatProtectionOnboardingStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Update(reqObj *AdvancedThreatProtectionOnboardingStateSummary) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AdvancedThreatProtectionOnboardingStateSummary
func (r *AdvancedThreatProtectionOnboardingStateSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AdvancedThreatProtectionOnboardingDeviceSettingStates returns request builder for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (b *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates() *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder {
	bb := &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/advancedThreatProtectionOnboardingDeviceSettingStates"
	return bb
}

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder is request builder for AdvancedThreatProtectionOnboardingDeviceSettingState collection
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (b *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder) Request() *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest {
	return &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AdvancedThreatProtectionOnboardingDeviceSettingState item
func (b *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder) ID(id string) *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder {
	bb := &AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest is request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AdvancedThreatProtectionOnboardingDeviceSettingState
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
			value  []AdvancedThreatProtectionOnboardingDeviceSettingState
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

// Get performs GET request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Get() ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Add(reqObj *AdvancedThreatProtectionOnboardingDeviceSettingState) (resObj *AdvancedThreatProtectionOnboardingDeviceSettingState, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
