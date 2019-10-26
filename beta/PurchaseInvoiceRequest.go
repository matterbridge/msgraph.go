// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PurchaseInvoiceRequestBuilder is request builder for PurchaseInvoice
type PurchaseInvoiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PurchaseInvoiceRequest
func (b *PurchaseInvoiceRequestBuilder) Request() *PurchaseInvoiceRequest {
	return &PurchaseInvoiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PurchaseInvoiceRequest is request for PurchaseInvoice
type PurchaseInvoiceRequest struct{ BaseRequest }

// Get performs GET request for PurchaseInvoice
func (r *PurchaseInvoiceRequest) Get() (resObj *PurchaseInvoice, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PurchaseInvoice
func (r *PurchaseInvoiceRequest) Update(reqObj *PurchaseInvoice) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PurchaseInvoice
func (r *PurchaseInvoiceRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Currency is navigation property
func (b *PurchaseInvoiceRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// PurchaseInvoiceLines returns request builder for PurchaseInvoiceLine collection
func (b *PurchaseInvoiceRequestBuilder) PurchaseInvoiceLines() *PurchaseInvoicePurchaseInvoiceLinesCollectionRequestBuilder {
	bb := &PurchaseInvoicePurchaseInvoiceLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/purchaseInvoiceLines"
	return bb
}

// PurchaseInvoicePurchaseInvoiceLinesCollectionRequestBuilder is request builder for PurchaseInvoiceLine collection
type PurchaseInvoicePurchaseInvoiceLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PurchaseInvoiceLine collection
func (b *PurchaseInvoicePurchaseInvoiceLinesCollectionRequestBuilder) Request() *PurchaseInvoicePurchaseInvoiceLinesCollectionRequest {
	return &PurchaseInvoicePurchaseInvoiceLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PurchaseInvoiceLine item
func (b *PurchaseInvoicePurchaseInvoiceLinesCollectionRequestBuilder) ID(id string) *PurchaseInvoiceLineRequestBuilder {
	bb := &PurchaseInvoiceLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PurchaseInvoicePurchaseInvoiceLinesCollectionRequest is request for PurchaseInvoiceLine collection
type PurchaseInvoicePurchaseInvoiceLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PurchaseInvoiceLine collection
func (r *PurchaseInvoicePurchaseInvoiceLinesCollectionRequest) Paging(method, path string, obj interface{}) ([]PurchaseInvoiceLine, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PurchaseInvoiceLine
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
			value  []PurchaseInvoiceLine
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

// Get performs GET request for PurchaseInvoiceLine collection
func (r *PurchaseInvoicePurchaseInvoiceLinesCollectionRequest) Get() ([]PurchaseInvoiceLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PurchaseInvoiceLine collection
func (r *PurchaseInvoicePurchaseInvoiceLinesCollectionRequest) Add(reqObj *PurchaseInvoiceLine) (resObj *PurchaseInvoiceLine, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Vendor is navigation property
func (b *PurchaseInvoiceRequestBuilder) Vendor() *VendorRequestBuilder {
	bb := &VendorRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/vendor"
	return bb
}
