// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// VendorRequestBuilder is request builder for Vendor
type VendorRequestBuilder struct{ BaseRequestBuilder }

// Request returns VendorRequest
func (b *VendorRequestBuilder) Request() *VendorRequest {
	return &VendorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// VendorRequest is request for Vendor
type VendorRequest struct{ BaseRequest }

// Get performs GET request for Vendor
func (r *VendorRequest) Get() (resObj *Vendor, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Vendor
func (r *VendorRequest) Update(reqObj *Vendor) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Vendor
func (r *VendorRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Currency is navigation property
func (b *VendorRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// PaymentMethod is navigation property
func (b *VendorRequestBuilder) PaymentMethod() *PaymentMethodRequestBuilder {
	bb := &PaymentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentMethod"
	return bb
}

// PaymentTerm is navigation property
func (b *VendorRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// Picture returns request builder for Picture collection
func (b *VendorRequestBuilder) Picture() *VendorPictureCollectionRequestBuilder {
	bb := &VendorPictureCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/picture"
	return bb
}

// VendorPictureCollectionRequestBuilder is request builder for Picture collection
type VendorPictureCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Picture collection
func (b *VendorPictureCollectionRequestBuilder) Request() *VendorPictureCollectionRequest {
	return &VendorPictureCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Picture item
func (b *VendorPictureCollectionRequestBuilder) ID(id string) *PictureRequestBuilder {
	bb := &PictureRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// VendorPictureCollectionRequest is request for Picture collection
type VendorPictureCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Picture collection
func (r *VendorPictureCollectionRequest) Paging(method, path string, obj interface{}) ([]Picture, error) {
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
func (r *VendorPictureCollectionRequest) Get() ([]Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Picture collection
func (r *VendorPictureCollectionRequest) Add(reqObj *Picture) (resObj *Picture, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
