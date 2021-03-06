// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DataSharingConsentConsentToDataSharingRequestParameter undocumented
type DataSharingConsentConsentToDataSharingRequestParameter struct {
}

//
type DataSharingConsentConsentToDataSharingRequestBuilder struct{ BaseRequestBuilder }

// ConsentToDataSharing action undocumented
func (b *DataSharingConsentRequestBuilder) ConsentToDataSharing(reqObj *DataSharingConsentConsentToDataSharingRequestParameter) *DataSharingConsentConsentToDataSharingRequestBuilder {
	bb := &DataSharingConsentConsentToDataSharingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/consentToDataSharing"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DataSharingConsentConsentToDataSharingRequest struct{ BaseRequest }

//
func (b *DataSharingConsentConsentToDataSharingRequestBuilder) Request() *DataSharingConsentConsentToDataSharingRequest {
	return &DataSharingConsentConsentToDataSharingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DataSharingConsentConsentToDataSharingRequest) Post(ctx context.Context) (resObj *DataSharingConsent, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
