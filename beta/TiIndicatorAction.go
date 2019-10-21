// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TiIndicatorCollectionSubmitTiIndicatorsRequestParameter undocumented
type TiIndicatorCollectionSubmitTiIndicatorsRequestParameter struct {
	// Value undocumented
	Value []TiIndicator `json:"value,omitempty"`
}

// TiIndicatorCollectionUpdateTiIndicatorsRequestParameter undocumented
type TiIndicatorCollectionUpdateTiIndicatorsRequestParameter struct {
	// Value undocumented
	Value []TiIndicator `json:"value,omitempty"`
}

// TiIndicatorCollectionDeleteTiIndicatorsRequestParameter undocumented
type TiIndicatorCollectionDeleteTiIndicatorsRequestParameter struct {
	// Value undocumented
	Value []string `json:"value,omitempty"`
}

// TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestParameter undocumented
type TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestParameter struct {
	// Value undocumented
	Value []string `json:"value,omitempty"`
}

//
type TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder struct{ BaseRequestBuilder }

// SubmitTiIndicators action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) SubmitTiIndicators(reqObj *TiIndicatorCollectionSubmitTiIndicatorsRequestParameter) *TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder {
	bb := &TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/submitTiIndicators"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionSubmitTiIndicatorsRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionSubmitTiIndicatorsRequestBuilder) Request() *TiIndicatorCollectionSubmitTiIndicatorsRequest {
	return &TiIndicatorCollectionSubmitTiIndicatorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionSubmitTiIndicatorsRequest) Do(method, path string, reqObj interface{}) (resObj *[]TiIndicator, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *TiIndicatorCollectionSubmitTiIndicatorsRequest) Paging(method, path string, obj interface{}) ([][]TiIndicator, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]TiIndicator
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]TiIndicator
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
func (r *TiIndicatorCollectionSubmitTiIndicatorsRequest) Get() ([][]TiIndicator, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder struct{ BaseRequestBuilder }

// UpdateTiIndicators action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) UpdateTiIndicators(reqObj *TiIndicatorCollectionUpdateTiIndicatorsRequestParameter) *TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder {
	bb := &TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateTiIndicators"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionUpdateTiIndicatorsRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionUpdateTiIndicatorsRequestBuilder) Request() *TiIndicatorCollectionUpdateTiIndicatorsRequest {
	return &TiIndicatorCollectionUpdateTiIndicatorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionUpdateTiIndicatorsRequest) Do(method, path string, reqObj interface{}) (resObj *[]TiIndicator, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *TiIndicatorCollectionUpdateTiIndicatorsRequest) Paging(method, path string, obj interface{}) ([][]TiIndicator, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]TiIndicator
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]TiIndicator
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
func (r *TiIndicatorCollectionUpdateTiIndicatorsRequest) Get() ([][]TiIndicator, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder struct{ BaseRequestBuilder }

// DeleteTiIndicators action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) DeleteTiIndicators(reqObj *TiIndicatorCollectionDeleteTiIndicatorsRequestParameter) *TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder {
	bb := &TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/deleteTiIndicators"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionDeleteTiIndicatorsRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionDeleteTiIndicatorsRequestBuilder) Request() *TiIndicatorCollectionDeleteTiIndicatorsRequest {
	return &TiIndicatorCollectionDeleteTiIndicatorsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsRequest) Do(method, path string, reqObj interface{}) (resObj *[]ResultInfo, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsRequest) Paging(method, path string, obj interface{}) ([][]ResultInfo, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]ResultInfo
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]ResultInfo
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
func (r *TiIndicatorCollectionDeleteTiIndicatorsRequest) Get() ([][]ResultInfo, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder struct{ BaseRequestBuilder }

// DeleteTiIndicatorsByExternalID action undocumented
func (b *SecurityTiIndicatorsCollectionRequestBuilder) DeleteTiIndicatorsByExternalID(reqObj *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestParameter) *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder {
	bb := &TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/deleteTiIndicatorsByExternalId"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest struct{ BaseRequest }

//
func (b *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequestBuilder) Request() *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest {
	return &TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest) Do(method, path string, reqObj interface{}) (resObj *[]ResultInfo, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest) Paging(method, path string, obj interface{}) ([][]ResultInfo, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]ResultInfo
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]ResultInfo
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
func (r *TiIndicatorCollectionDeleteTiIndicatorsByExternalIDRequest) Get() ([][]ResultInfo, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
