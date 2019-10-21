// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OutlookTaskCompleteRequestParameter undocumented
type OutlookTaskCompleteRequestParameter struct {
}

//
type OutlookTaskCompleteRequestBuilder struct{ BaseRequestBuilder }

// Complete action undocumented
func (b *OutlookTaskRequestBuilder) Complete(reqObj *OutlookTaskCompleteRequestParameter) *OutlookTaskCompleteRequestBuilder {
	bb := &OutlookTaskCompleteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/complete"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OutlookTaskCompleteRequest struct{ BaseRequest }

//
func (b *OutlookTaskCompleteRequestBuilder) Request() *OutlookTaskCompleteRequest {
	return &OutlookTaskCompleteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OutlookTaskCompleteRequest) Do(method, path string, reqObj interface{}) (resObj *[]OutlookTask, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *OutlookTaskCompleteRequest) Paging(method, path string, obj interface{}) ([][]OutlookTask, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]OutlookTask
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]OutlookTask
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
func (r *OutlookTaskCompleteRequest) Get() ([][]OutlookTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
