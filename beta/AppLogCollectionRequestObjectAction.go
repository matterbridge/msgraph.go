// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppLogCollectionRequestObjectCreateDownloadURLRequestParameter undocumented
type AppLogCollectionRequestObjectCreateDownloadURLRequestParameter struct {
}

//
type AppLogCollectionRequestObjectCreateDownloadURLRequestBuilder struct{ BaseRequestBuilder }

// CreateDownloadURL action undocumented
func (b *AppLogCollectionRequestObjectRequestBuilder) CreateDownloadURL(reqObj *AppLogCollectionRequestObjectCreateDownloadURLRequestParameter) *AppLogCollectionRequestObjectCreateDownloadURLRequestBuilder {
	bb := &AppLogCollectionRequestObjectCreateDownloadURLRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createDownloadUrl"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AppLogCollectionRequestObjectCreateDownloadURLRequest struct{ BaseRequest }

//
func (b *AppLogCollectionRequestObjectCreateDownloadURLRequestBuilder) Request() *AppLogCollectionRequestObjectCreateDownloadURLRequest {
	return &AppLogCollectionRequestObjectCreateDownloadURLRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AppLogCollectionRequestObjectCreateDownloadURLRequest) Do(method, path string, reqObj interface{}) (resObj *AppLogCollectionDownloadDetails, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *AppLogCollectionRequestObjectCreateDownloadURLRequest) Post() (*AppLogCollectionDownloadDetails, error) {
	return r.Do("POST", "", r.requestObject)
}