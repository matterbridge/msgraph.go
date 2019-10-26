// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DriveItemCheckinRequestParameter undocumented
type DriveItemCheckinRequestParameter struct {
	// CheckInAs undocumented
	CheckInAs *string `json:"checkInAs,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
}

// DriveItemCheckoutRequestParameter undocumented
type DriveItemCheckoutRequestParameter struct {
}

// DriveItemCopyRequestParameter undocumented
type DriveItemCopyRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ParentReference undocumented
	ParentReference *ItemReference `json:"parentReference,omitempty"`
}

// DriveItemCreateLinkRequestParameter undocumented
type DriveItemCreateLinkRequestParameter struct {
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Recipients undocumented
	Recipients []DriveRecipient `json:"recipients,omitempty"`
}

// DriveItemCreateUploadSessionRequestParameter undocumented
type DriveItemCreateUploadSessionRequestParameter struct {
	// Item undocumented
	Item *DriveItemUploadableProperties `json:"item,omitempty"`
	// DeferCommit undocumented
	DeferCommit *bool `json:"deferCommit,omitempty"`
}

// DriveItemFollowRequestParameter undocumented
type DriveItemFollowRequestParameter struct {
}

// DriveItemUnfollowRequestParameter undocumented
type DriveItemUnfollowRequestParameter struct {
}

// DriveItemInviteRequestParameter undocumented
type DriveItemInviteRequestParameter struct {
	// RequireSignIn undocumented
	RequireSignIn *bool `json:"requireSignIn,omitempty"`
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// SendInvitation undocumented
	SendInvitation *bool `json:"sendInvitation,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Recipients undocumented
	Recipients []DriveRecipient `json:"recipients,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
}

// DriveItemPreviewRequestParameter undocumented
type DriveItemPreviewRequestParameter struct {
	// Viewer undocumented
	Viewer *string `json:"viewer,omitempty"`
	// Chromeless undocumented
	Chromeless *bool `json:"chromeless,omitempty"`
	// AllowEdit undocumented
	AllowEdit *bool `json:"allowEdit,omitempty"`
	// Page undocumented
	Page *string `json:"page,omitempty"`
	// Zoom undocumented
	Zoom *float64 `json:"zoom,omitempty"`
}

// DriveItemRestoreRequestParameter undocumented
type DriveItemRestoreRequestParameter struct {
	// ParentReference undocumented
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// DriveItemValidatePermissionRequestParameter undocumented
type DriveItemValidatePermissionRequestParameter struct {
	// ChallengeToken undocumented
	ChallengeToken *string `json:"challengeToken,omitempty"`
	// Password undocumented
	Password *string `json:"password,omitempty"`
}

//
type DriveItemCheckinRequestBuilder struct{ BaseRequestBuilder }

// Checkin action undocumented
func (b *DriveItemRequestBuilder) Checkin(reqObj *DriveItemCheckinRequestParameter) *DriveItemCheckinRequestBuilder {
	bb := &DriveItemCheckinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/checkin"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemCheckinRequest struct{ BaseRequest }

//
func (b *DriveItemCheckinRequestBuilder) Request() *DriveItemCheckinRequest {
	return &DriveItemCheckinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemCheckinRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type DriveItemCheckoutRequestBuilder struct{ BaseRequestBuilder }

// Checkout action undocumented
func (b *DriveItemRequestBuilder) Checkout(reqObj *DriveItemCheckoutRequestParameter) *DriveItemCheckoutRequestBuilder {
	bb := &DriveItemCheckoutRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/checkout"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemCheckoutRequest struct{ BaseRequest }

//
func (b *DriveItemCheckoutRequestBuilder) Request() *DriveItemCheckoutRequest {
	return &DriveItemCheckoutRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemCheckoutRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type DriveItemCopyRequestBuilder struct{ BaseRequestBuilder }

// Copy action undocumented
func (b *DriveItemRequestBuilder) Copy(reqObj *DriveItemCopyRequestParameter) *DriveItemCopyRequestBuilder {
	bb := &DriveItemCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemCopyRequest struct{ BaseRequest }

//
func (b *DriveItemCopyRequestBuilder) Request() *DriveItemCopyRequest {
	return &DriveItemCopyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemCopyRequest) Post() (resObj *DriveItem, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type DriveItemCreateLinkRequestBuilder struct{ BaseRequestBuilder }

// CreateLink action undocumented
func (b *DriveItemRequestBuilder) CreateLink(reqObj *DriveItemCreateLinkRequestParameter) *DriveItemCreateLinkRequestBuilder {
	bb := &DriveItemCreateLinkRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createLink"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemCreateLinkRequest struct{ BaseRequest }

//
func (b *DriveItemCreateLinkRequestBuilder) Request() *DriveItemCreateLinkRequest {
	return &DriveItemCreateLinkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemCreateLinkRequest) Post() (resObj *Permission, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type DriveItemCreateUploadSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateUploadSession action undocumented
func (b *DriveItemRequestBuilder) CreateUploadSession(reqObj *DriveItemCreateUploadSessionRequestParameter) *DriveItemCreateUploadSessionRequestBuilder {
	bb := &DriveItemCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemCreateUploadSessionRequest struct{ BaseRequest }

//
func (b *DriveItemCreateUploadSessionRequestBuilder) Request() *DriveItemCreateUploadSessionRequest {
	return &DriveItemCreateUploadSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemCreateUploadSessionRequest) Post() (resObj *UploadSession, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type DriveItemFollowRequestBuilder struct{ BaseRequestBuilder }

// Follow action undocumented
func (b *DriveItemRequestBuilder) Follow(reqObj *DriveItemFollowRequestParameter) *DriveItemFollowRequestBuilder {
	bb := &DriveItemFollowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/follow"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemFollowRequest struct{ BaseRequest }

//
func (b *DriveItemFollowRequestBuilder) Request() *DriveItemFollowRequest {
	return &DriveItemFollowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemFollowRequest) Post() (resObj *DriveItem, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type DriveItemUnfollowRequestBuilder struct{ BaseRequestBuilder }

// Unfollow action undocumented
func (b *DriveItemRequestBuilder) Unfollow(reqObj *DriveItemUnfollowRequestParameter) *DriveItemUnfollowRequestBuilder {
	bb := &DriveItemUnfollowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unfollow"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemUnfollowRequest struct{ BaseRequest }

//
func (b *DriveItemUnfollowRequestBuilder) Request() *DriveItemUnfollowRequest {
	return &DriveItemUnfollowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemUnfollowRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type DriveItemInviteRequestBuilder struct{ BaseRequestBuilder }

// Invite action undocumented
func (b *DriveItemRequestBuilder) Invite(reqObj *DriveItemInviteRequestParameter) *DriveItemInviteRequestBuilder {
	bb := &DriveItemInviteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/invite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemInviteRequest struct{ BaseRequest }

//
func (b *DriveItemInviteRequestBuilder) Request() *DriveItemInviteRequest {
	return &DriveItemInviteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemInviteRequest) Paging(method, path string, obj interface{}) ([][]Permission, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]Permission
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
			value  [][]Permission
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
func (r *DriveItemInviteRequest) Get() ([][]Permission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type DriveItemPreviewRequestBuilder struct{ BaseRequestBuilder }

// Preview action undocumented
func (b *DriveItemRequestBuilder) Preview(reqObj *DriveItemPreviewRequestParameter) *DriveItemPreviewRequestBuilder {
	bb := &DriveItemPreviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/preview"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemPreviewRequest struct{ BaseRequest }

//
func (b *DriveItemPreviewRequestBuilder) Request() *DriveItemPreviewRequest {
	return &DriveItemPreviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemPreviewRequest) Post() (resObj *ItemPreviewInfo, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type DriveItemRestoreRequestBuilder struct{ BaseRequestBuilder }

// Restore action undocumented
func (b *DriveItemRequestBuilder) Restore(reqObj *DriveItemRestoreRequestParameter) *DriveItemRestoreRequestBuilder {
	bb := &DriveItemRestoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/restore"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemRestoreRequest struct{ BaseRequest }

//
func (b *DriveItemRestoreRequestBuilder) Request() *DriveItemRestoreRequest {
	return &DriveItemRestoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemRestoreRequest) Post() (resObj *DriveItem, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type DriveItemValidatePermissionRequestBuilder struct{ BaseRequestBuilder }

// ValidatePermission action undocumented
func (b *DriveItemRequestBuilder) ValidatePermission(reqObj *DriveItemValidatePermissionRequestParameter) *DriveItemValidatePermissionRequestBuilder {
	bb := &DriveItemValidatePermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validatePermission"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemValidatePermissionRequest struct{ BaseRequest }

//
func (b *DriveItemValidatePermissionRequestBuilder) Request() *DriveItemValidatePermissionRequest {
	return &DriveItemValidatePermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemValidatePermissionRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
