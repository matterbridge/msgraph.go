// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GroupCollectionEvaluateDynamicMembershipRequestParameter undocumented
type GroupCollectionEvaluateDynamicMembershipRequestParameter struct {
	// MemberID undocumented
	MemberID *string `json:"memberId,omitempty"`
	// MembershipRule undocumented
	MembershipRule *string `json:"membershipRule,omitempty"`
}

// GroupValidatePropertiesRequestParameter undocumented
type GroupValidatePropertiesRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnBehalfOfUserID undocumented
	OnBehalfOfUserID *UUID `json:"onBehalfOfUserId,omitempty"`
}

// GroupCheckGrantedPermissionsForAppRequestParameter undocumented
type GroupCheckGrantedPermissionsForAppRequestParameter struct {
}

// GroupAssignLicenseRequestParameter undocumented
type GroupAssignLicenseRequestParameter struct {
	// AddLicenses undocumented
	AddLicenses []AssignedLicense `json:"addLicenses,omitempty"`
	// RemoveLicenses undocumented
	RemoveLicenses []UUID `json:"removeLicenses,omitempty"`
}

// GroupSubscribeByMailRequestParameter undocumented
type GroupSubscribeByMailRequestParameter struct {
}

// GroupUnsubscribeByMailRequestParameter undocumented
type GroupUnsubscribeByMailRequestParameter struct {
}

// GroupAddFavoriteRequestParameter undocumented
type GroupAddFavoriteRequestParameter struct {
}

// GroupRemoveFavoriteRequestParameter undocumented
type GroupRemoveFavoriteRequestParameter struct {
}

// GroupResetUnseenCountRequestParameter undocumented
type GroupResetUnseenCountRequestParameter struct {
}

// GroupRenewRequestParameter undocumented
type GroupRenewRequestParameter struct {
}

// GroupEvaluateDynamicMembershipRequestParameter undocumented
type GroupEvaluateDynamicMembershipRequestParameter struct {
	// MemberID undocumented
	MemberID *string `json:"memberId,omitempty"`
}

//
type GroupCollectionEvaluateDynamicMembershipRequestBuilder struct{ BaseRequestBuilder }

// EvaluateDynamicMembership action undocumented
func (b *UserJoinedGroupsCollectionRequestBuilder) EvaluateDynamicMembership(reqObj *GroupCollectionEvaluateDynamicMembershipRequestParameter) *GroupCollectionEvaluateDynamicMembershipRequestBuilder {
	bb := &GroupCollectionEvaluateDynamicMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateDynamicMembership"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupCollectionEvaluateDynamicMembershipRequest struct{ BaseRequest }

//
func (b *GroupCollectionEvaluateDynamicMembershipRequestBuilder) Request() *GroupCollectionEvaluateDynamicMembershipRequest {
	return &GroupCollectionEvaluateDynamicMembershipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupCollectionEvaluateDynamicMembershipRequest) Post(ctx context.Context) (resObj *EvaluateDynamicMembershipResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type GroupValidatePropertiesRequestBuilder struct{ BaseRequestBuilder }

// ValidateProperties action undocumented
func (b *GroupRequestBuilder) ValidateProperties(reqObj *GroupValidatePropertiesRequestParameter) *GroupValidatePropertiesRequestBuilder {
	bb := &GroupValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupValidatePropertiesRequest struct{ BaseRequest }

//
func (b *GroupValidatePropertiesRequestBuilder) Request() *GroupValidatePropertiesRequest {
	return &GroupValidatePropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupValidatePropertiesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupCheckGrantedPermissionsForAppRequestBuilder struct{ BaseRequestBuilder }

// CheckGrantedPermissionsForApp action undocumented
func (b *GroupRequestBuilder) CheckGrantedPermissionsForApp(reqObj *GroupCheckGrantedPermissionsForAppRequestParameter) *GroupCheckGrantedPermissionsForAppRequestBuilder {
	bb := &GroupCheckGrantedPermissionsForAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/checkGrantedPermissionsForApp"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupCheckGrantedPermissionsForAppRequest struct{ BaseRequest }

//
func (b *GroupCheckGrantedPermissionsForAppRequestBuilder) Request() *GroupCheckGrantedPermissionsForAppRequest {
	return &GroupCheckGrantedPermissionsForAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupCheckGrantedPermissionsForAppRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]ResourceSpecificPermissionGrant, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]ResourceSpecificPermissionGrant
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
			value  [][]ResourceSpecificPermissionGrant
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
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *GroupCheckGrantedPermissionsForAppRequest) Get(ctx context.Context) ([][]ResourceSpecificPermissionGrant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

//
type GroupAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumented
func (b *GroupRequestBuilder) AssignLicense(reqObj *GroupAssignLicenseRequestParameter) *GroupAssignLicenseRequestBuilder {
	bb := &GroupAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupAssignLicenseRequest struct{ BaseRequest }

//
func (b *GroupAssignLicenseRequestBuilder) Request() *GroupAssignLicenseRequest {
	return &GroupAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupAssignLicenseRequest) Post(ctx context.Context) (resObj *Group, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type GroupSubscribeByMailRequestBuilder struct{ BaseRequestBuilder }

// SubscribeByMail action undocumented
func (b *GroupRequestBuilder) SubscribeByMail(reqObj *GroupSubscribeByMailRequestParameter) *GroupSubscribeByMailRequestBuilder {
	bb := &GroupSubscribeByMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/subscribeByMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupSubscribeByMailRequest struct{ BaseRequest }

//
func (b *GroupSubscribeByMailRequestBuilder) Request() *GroupSubscribeByMailRequest {
	return &GroupSubscribeByMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupSubscribeByMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupUnsubscribeByMailRequestBuilder struct{ BaseRequestBuilder }

// UnsubscribeByMail action undocumented
func (b *GroupRequestBuilder) UnsubscribeByMail(reqObj *GroupUnsubscribeByMailRequestParameter) *GroupUnsubscribeByMailRequestBuilder {
	bb := &GroupUnsubscribeByMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unsubscribeByMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupUnsubscribeByMailRequest struct{ BaseRequest }

//
func (b *GroupUnsubscribeByMailRequestBuilder) Request() *GroupUnsubscribeByMailRequest {
	return &GroupUnsubscribeByMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupUnsubscribeByMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupAddFavoriteRequestBuilder struct{ BaseRequestBuilder }

// AddFavorite action undocumented
func (b *GroupRequestBuilder) AddFavorite(reqObj *GroupAddFavoriteRequestParameter) *GroupAddFavoriteRequestBuilder {
	bb := &GroupAddFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addFavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupAddFavoriteRequest struct{ BaseRequest }

//
func (b *GroupAddFavoriteRequestBuilder) Request() *GroupAddFavoriteRequest {
	return &GroupAddFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupAddFavoriteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupRemoveFavoriteRequestBuilder struct{ BaseRequestBuilder }

// RemoveFavorite action undocumented
func (b *GroupRequestBuilder) RemoveFavorite(reqObj *GroupRemoveFavoriteRequestParameter) *GroupRemoveFavoriteRequestBuilder {
	bb := &GroupRemoveFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeFavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupRemoveFavoriteRequest struct{ BaseRequest }

//
func (b *GroupRemoveFavoriteRequestBuilder) Request() *GroupRemoveFavoriteRequest {
	return &GroupRemoveFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupRemoveFavoriteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupResetUnseenCountRequestBuilder struct{ BaseRequestBuilder }

// ResetUnseenCount action undocumented
func (b *GroupRequestBuilder) ResetUnseenCount(reqObj *GroupResetUnseenCountRequestParameter) *GroupResetUnseenCountRequestBuilder {
	bb := &GroupResetUnseenCountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/resetUnseenCount"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupResetUnseenCountRequest struct{ BaseRequest }

//
func (b *GroupResetUnseenCountRequestBuilder) Request() *GroupResetUnseenCountRequest {
	return &GroupResetUnseenCountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupResetUnseenCountRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupRenewRequestBuilder struct{ BaseRequestBuilder }

// Renew action undocumented
func (b *GroupRequestBuilder) Renew(reqObj *GroupRenewRequestParameter) *GroupRenewRequestBuilder {
	bb := &GroupRenewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/renew"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupRenewRequest struct{ BaseRequest }

//
func (b *GroupRenewRequestBuilder) Request() *GroupRenewRequest {
	return &GroupRenewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupRenewRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type GroupEvaluateDynamicMembershipRequestBuilder struct{ BaseRequestBuilder }

// EvaluateDynamicMembership action undocumented
func (b *GroupRequestBuilder) EvaluateDynamicMembership(reqObj *GroupEvaluateDynamicMembershipRequestParameter) *GroupEvaluateDynamicMembershipRequestBuilder {
	bb := &GroupEvaluateDynamicMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateDynamicMembership"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupEvaluateDynamicMembershipRequest struct{ BaseRequest }

//
func (b *GroupEvaluateDynamicMembershipRequestBuilder) Request() *GroupEvaluateDynamicMembershipRequest {
	return &GroupEvaluateDynamicMembershipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupEvaluateDynamicMembershipRequest) Post(ctx context.Context) (resObj *EvaluateDynamicMembershipResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
