// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CalendarRequestBuilder is request builder for Calendar
type CalendarRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarRequest
func (b *CalendarRequestBuilder) Request() *CalendarRequest {
	return &CalendarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarRequest is request for Calendar
type CalendarRequest struct{ BaseRequest }

// Get performs GET request for Calendar
func (r *CalendarRequest) Get() (resObj *Calendar, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Calendar
func (r *CalendarRequest) Update(reqObj *Calendar) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Calendar
func (r *CalendarRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// CalendarView returns request builder for Event collection
func (b *CalendarRequestBuilder) CalendarView() *CalendarCalendarViewCollectionRequestBuilder {
	bb := &CalendarCalendarViewCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendarView"
	return bb
}

// CalendarCalendarViewCollectionRequestBuilder is request builder for Event collection
type CalendarCalendarViewCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *CalendarCalendarViewCollectionRequestBuilder) Request() *CalendarCalendarViewCollectionRequest {
	return &CalendarCalendarViewCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *CalendarCalendarViewCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarCalendarViewCollectionRequest is request for Event collection
type CalendarCalendarViewCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *CalendarCalendarViewCollectionRequest) Paging(method, path string, obj interface{}) ([]Event, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Event
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
			value  []Event
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

// Get performs GET request for Event collection
func (r *CalendarCalendarViewCollectionRequest) Get() ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Event collection
func (r *CalendarCalendarViewCollectionRequest) Add(reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Events returns request builder for Event collection
func (b *CalendarRequestBuilder) Events() *CalendarEventsCollectionRequestBuilder {
	bb := &CalendarEventsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/events"
	return bb
}

// CalendarEventsCollectionRequestBuilder is request builder for Event collection
type CalendarEventsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *CalendarEventsCollectionRequestBuilder) Request() *CalendarEventsCollectionRequest {
	return &CalendarEventsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *CalendarEventsCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarEventsCollectionRequest is request for Event collection
type CalendarEventsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Event collection
func (r *CalendarEventsCollectionRequest) Paging(method, path string, obj interface{}) ([]Event, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Event
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
			value  []Event
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

// Get performs GET request for Event collection
func (r *CalendarEventsCollectionRequest) Get() ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Event collection
func (r *CalendarEventsCollectionRequest) Add(reqObj *Event) (resObj *Event, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *CalendarRequestBuilder) MultiValueExtendedProperties() *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &CalendarMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// CalendarMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type CalendarMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *CalendarMultiValueExtendedPropertiesCollectionRequest {
	return &CalendarMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *CalendarMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type CalendarMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MultiValueLegacyExtendedProperty
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
			value  []MultiValueLegacyExtendedProperty
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

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Get() ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *CalendarMultiValueExtendedPropertiesCollectionRequest) Add(reqObj *MultiValueLegacyExtendedProperty) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *CalendarRequestBuilder) SingleValueExtendedProperties() *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &CalendarSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// CalendarSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type CalendarSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *CalendarSingleValueExtendedPropertiesCollectionRequest {
	return &CalendarSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *CalendarSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CalendarSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type CalendarSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SingleValueLegacyExtendedProperty
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
			value  []SingleValueLegacyExtendedProperty
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

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Get() ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *CalendarSingleValueExtendedPropertiesCollectionRequest) Add(reqObj *SingleValueLegacyExtendedProperty) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
