// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// UserAssignLicenseRequestParameter undocumented
type UserAssignLicenseRequestParameter struct {
	// AddLicenses undocumented
	AddLicenses []AssignedLicense `json:"addLicenses,omitempty"`
	// RemoveLicenses undocumented
	RemoveLicenses []UUID `json:"removeLicenses,omitempty"`
}

// UserChangePasswordRequestParameter undocumented
type UserChangePasswordRequestParameter struct {
	// CurrentPassword undocumented
	CurrentPassword *string `json:"currentPassword,omitempty"`
	// NewPassword undocumented
	NewPassword *string `json:"newPassword,omitempty"`
}

// UserInvalidateAllRefreshTokensRequestParameter undocumented
type UserInvalidateAllRefreshTokensRequestParameter struct {
}

// UserRevokeSignInSessionsRequestParameter undocumented
type UserRevokeSignInSessionsRequestParameter struct {
}

// UserReprocessLicenseAssignmentRequestParameter undocumented
type UserReprocessLicenseAssignmentRequestParameter struct {
}

// UserFindMeetingTimesRequestParameter undocumented
type UserFindMeetingTimesRequestParameter struct {
	// Attendees undocumented
	Attendees []AttendeeBase `json:"attendees,omitempty"`
	// LocationConstraint undocumented
	LocationConstraint *LocationConstraint `json:"locationConstraint,omitempty"`
	// TimeConstraint undocumented
	TimeConstraint *TimeConstraint `json:"timeConstraint,omitempty"`
	// MeetingDuration undocumented
	MeetingDuration *time.Duration `json:"meetingDuration,omitempty"`
	// MaxCandidates undocumented
	MaxCandidates *int `json:"maxCandidates,omitempty"`
	// IsOrganizerOptional undocumented
	IsOrganizerOptional *bool `json:"isOrganizerOptional,omitempty"`
	// ReturnSuggestionReasons undocumented
	ReturnSuggestionReasons *bool `json:"returnSuggestionReasons,omitempty"`
	// MinimumAttendeePercentage undocumented
	MinimumAttendeePercentage *float64 `json:"minimumAttendeePercentage,omitempty"`
}

// UserSendMailRequestParameter undocumented
type UserSendMailRequestParameter struct {
	// Message undocumented
	Message *Message `json:"Message,omitempty"`
	// SaveToSentItems undocumented
	SaveToSentItems *bool `json:"SaveToSentItems,omitempty"`
}

// UserGetMailTipsRequestParameter undocumented
type UserGetMailTipsRequestParameter struct {
	// EmailAddresses undocumented
	EmailAddresses []string `json:"EmailAddresses,omitempty"`
	// MailTipsOptions undocumented
	MailTipsOptions *MailTipsType `json:"MailTipsOptions,omitempty"`
}

// UserTranslateExchangeIDsRequestParameter undocumented
type UserTranslateExchangeIDsRequestParameter struct {
	// InputIDs undocumented
	InputIDs []string `json:"InputIds,omitempty"`
	// TargetIDType undocumented
	TargetIDType *ExchangeIDFormat `json:"TargetIdType,omitempty"`
	// SourceIDType undocumented
	SourceIDType *ExchangeIDFormat `json:"SourceIdType,omitempty"`
}

// UserRemoveAllDevicesFromManagementRequestParameter undocumented
type UserRemoveAllDevicesFromManagementRequestParameter struct {
}

// UserWipeManagedAppRegistrationByDeviceTagRequestParameter undocumented
type UserWipeManagedAppRegistrationByDeviceTagRequestParameter struct {
	// DeviceTag undocumented
	DeviceTag *string `json:"deviceTag,omitempty"`
}

// UserWipeManagedAppRegistrationsByDeviceTagRequestParameter undocumented
type UserWipeManagedAppRegistrationsByDeviceTagRequestParameter struct {
	// DeviceTag undocumented
	DeviceTag *string `json:"deviceTag,omitempty"`
}

// UserExportPersonalDataRequestParameter undocumented
type UserExportPersonalDataRequestParameter struct {
	// StorageLocation undocumented
	StorageLocation *string `json:"storageLocation,omitempty"`
}

//
type UserAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumented
func (b *UserRequestBuilder) AssignLicense(reqObj *UserAssignLicenseRequestParameter) *UserAssignLicenseRequestBuilder {
	bb := &UserAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserAssignLicenseRequest struct{ BaseRequest }

//
func (b *UserAssignLicenseRequestBuilder) Request() *UserAssignLicenseRequest {
	return &UserAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserAssignLicenseRequest) Post() (resObj *User, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type UserChangePasswordRequestBuilder struct{ BaseRequestBuilder }

// ChangePassword action undocumented
func (b *UserRequestBuilder) ChangePassword(reqObj *UserChangePasswordRequestParameter) *UserChangePasswordRequestBuilder {
	bb := &UserChangePasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/changePassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserChangePasswordRequest struct{ BaseRequest }

//
func (b *UserChangePasswordRequestBuilder) Request() *UserChangePasswordRequest {
	return &UserChangePasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserChangePasswordRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type UserInvalidateAllRefreshTokensRequestBuilder struct{ BaseRequestBuilder }

// InvalidateAllRefreshTokens action undocumented
func (b *UserRequestBuilder) InvalidateAllRefreshTokens(reqObj *UserInvalidateAllRefreshTokensRequestParameter) *UserInvalidateAllRefreshTokensRequestBuilder {
	bb := &UserInvalidateAllRefreshTokensRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/invalidateAllRefreshTokens"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserInvalidateAllRefreshTokensRequest struct{ BaseRequest }

//
func (b *UserInvalidateAllRefreshTokensRequestBuilder) Request() *UserInvalidateAllRefreshTokensRequest {
	return &UserInvalidateAllRefreshTokensRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserInvalidateAllRefreshTokensRequest) Post() (resObj *bool, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type UserRevokeSignInSessionsRequestBuilder struct{ BaseRequestBuilder }

// RevokeSignInSessions action undocumented
func (b *UserRequestBuilder) RevokeSignInSessions(reqObj *UserRevokeSignInSessionsRequestParameter) *UserRevokeSignInSessionsRequestBuilder {
	bb := &UserRevokeSignInSessionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeSignInSessions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserRevokeSignInSessionsRequest struct{ BaseRequest }

//
func (b *UserRevokeSignInSessionsRequestBuilder) Request() *UserRevokeSignInSessionsRequest {
	return &UserRevokeSignInSessionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserRevokeSignInSessionsRequest) Post() (resObj *bool, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type UserReprocessLicenseAssignmentRequestBuilder struct{ BaseRequestBuilder }

// ReprocessLicenseAssignment action undocumented
func (b *UserRequestBuilder) ReprocessLicenseAssignment(reqObj *UserReprocessLicenseAssignmentRequestParameter) *UserReprocessLicenseAssignmentRequestBuilder {
	bb := &UserReprocessLicenseAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reprocessLicenseAssignment"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserReprocessLicenseAssignmentRequest struct{ BaseRequest }

//
func (b *UserReprocessLicenseAssignmentRequestBuilder) Request() *UserReprocessLicenseAssignmentRequest {
	return &UserReprocessLicenseAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserReprocessLicenseAssignmentRequest) Post() (resObj *User, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type UserFindMeetingTimesRequestBuilder struct{ BaseRequestBuilder }

// FindMeetingTimes action undocumented
func (b *UserRequestBuilder) FindMeetingTimes(reqObj *UserFindMeetingTimesRequestParameter) *UserFindMeetingTimesRequestBuilder {
	bb := &UserFindMeetingTimesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/findMeetingTimes"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserFindMeetingTimesRequest struct{ BaseRequest }

//
func (b *UserFindMeetingTimesRequestBuilder) Request() *UserFindMeetingTimesRequest {
	return &UserFindMeetingTimesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserFindMeetingTimesRequest) Post() (resObj *MeetingTimeSuggestionsResult, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type UserSendMailRequestBuilder struct{ BaseRequestBuilder }

// SendMail action undocumented
func (b *UserRequestBuilder) SendMail(reqObj *UserSendMailRequestParameter) *UserSendMailRequestBuilder {
	bb := &UserSendMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserSendMailRequest struct{ BaseRequest }

//
func (b *UserSendMailRequestBuilder) Request() *UserSendMailRequest {
	return &UserSendMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserSendMailRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type UserGetMailTipsRequestBuilder struct{ BaseRequestBuilder }

// GetMailTips action undocumented
func (b *UserRequestBuilder) GetMailTips(reqObj *UserGetMailTipsRequestParameter) *UserGetMailTipsRequestBuilder {
	bb := &UserGetMailTipsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getMailTips"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserGetMailTipsRequest struct{ BaseRequest }

//
func (b *UserGetMailTipsRequestBuilder) Request() *UserGetMailTipsRequest {
	return &UserGetMailTipsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserGetMailTipsRequest) Paging(method, path string, obj interface{}) ([][]MailTips, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]MailTips
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
			value  [][]MailTips
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
func (r *UserGetMailTipsRequest) Get() ([][]MailTips, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type UserTranslateExchangeIDsRequestBuilder struct{ BaseRequestBuilder }

// TranslateExchangeIDs action undocumented
func (b *UserRequestBuilder) TranslateExchangeIDs(reqObj *UserTranslateExchangeIDsRequestParameter) *UserTranslateExchangeIDsRequestBuilder {
	bb := &UserTranslateExchangeIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/translateExchangeIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserTranslateExchangeIDsRequest struct{ BaseRequest }

//
func (b *UserTranslateExchangeIDsRequestBuilder) Request() *UserTranslateExchangeIDsRequest {
	return &UserTranslateExchangeIDsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserTranslateExchangeIDsRequest) Paging(method, path string, obj interface{}) ([][]ConvertIDResult, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]ConvertIDResult
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
			value  [][]ConvertIDResult
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
func (r *UserTranslateExchangeIDsRequest) Get() ([][]ConvertIDResult, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type UserRemoveAllDevicesFromManagementRequestBuilder struct{ BaseRequestBuilder }

// RemoveAllDevicesFromManagement action undocumented
func (b *UserRequestBuilder) RemoveAllDevicesFromManagement(reqObj *UserRemoveAllDevicesFromManagementRequestParameter) *UserRemoveAllDevicesFromManagementRequestBuilder {
	bb := &UserRemoveAllDevicesFromManagementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeAllDevicesFromManagement"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserRemoveAllDevicesFromManagementRequest struct{ BaseRequest }

//
func (b *UserRemoveAllDevicesFromManagementRequestBuilder) Request() *UserRemoveAllDevicesFromManagementRequest {
	return &UserRemoveAllDevicesFromManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserRemoveAllDevicesFromManagementRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type UserWipeManagedAppRegistrationByDeviceTagRequestBuilder struct{ BaseRequestBuilder }

// WipeManagedAppRegistrationByDeviceTag action undocumented
func (b *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag(reqObj *UserWipeManagedAppRegistrationByDeviceTagRequestParameter) *UserWipeManagedAppRegistrationByDeviceTagRequestBuilder {
	bb := &UserWipeManagedAppRegistrationByDeviceTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/wipeManagedAppRegistrationByDeviceTag"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserWipeManagedAppRegistrationByDeviceTagRequest struct{ BaseRequest }

//
func (b *UserWipeManagedAppRegistrationByDeviceTagRequestBuilder) Request() *UserWipeManagedAppRegistrationByDeviceTagRequest {
	return &UserWipeManagedAppRegistrationByDeviceTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserWipeManagedAppRegistrationByDeviceTagRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct{ BaseRequestBuilder }

// WipeManagedAppRegistrationsByDeviceTag action undocumented
func (b *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag(reqObj *UserWipeManagedAppRegistrationsByDeviceTagRequestParameter) *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder {
	bb := &UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/wipeManagedAppRegistrationsByDeviceTag"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserWipeManagedAppRegistrationsByDeviceTagRequest struct{ BaseRequest }

//
func (b *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Request() *UserWipeManagedAppRegistrationsByDeviceTagRequest {
	return &UserWipeManagedAppRegistrationsByDeviceTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserWipeManagedAppRegistrationsByDeviceTagRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type UserExportPersonalDataRequestBuilder struct{ BaseRequestBuilder }

// ExportPersonalData action undocumented
func (b *UserRequestBuilder) ExportPersonalData(reqObj *UserExportPersonalDataRequestParameter) *UserExportPersonalDataRequestBuilder {
	bb := &UserExportPersonalDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/exportPersonalData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserExportPersonalDataRequest struct{ BaseRequest }

//
func (b *UserExportPersonalDataRequestBuilder) Request() *UserExportPersonalDataRequest {
	return &UserExportPersonalDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserExportPersonalDataRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
