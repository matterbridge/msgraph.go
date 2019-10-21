// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SynchronizationJobPauseRequestParameter undocumented
type SynchronizationJobPauseRequestParameter struct {
}

// SynchronizationJobStartRequestParameter undocumented
type SynchronizationJobStartRequestParameter struct {
}

// SynchronizationJobStopRequestParameter undocumented
type SynchronizationJobStopRequestParameter struct {
}

// SynchronizationJobApplyRequestParameter undocumented
type SynchronizationJobApplyRequestParameter struct {
	// ObjectID undocumented
	ObjectID *string `json:"objectId,omitempty"`
	// TypeName undocumented
	TypeName *string `json:"typeName,omitempty"`
	// RuleID undocumented
	RuleID *string `json:"ruleId,omitempty"`
}

// SynchronizationJobRestartRequestParameter undocumented
type SynchronizationJobRestartRequestParameter struct {
	// Criteria undocumented
	Criteria *SynchronizationJobRestartCriteria `json:"criteria,omitempty"`
}

// SynchronizationJobValidateCredentialsRequestParameter undocumented
type SynchronizationJobValidateCredentialsRequestParameter struct {
	// ApplicationIdentifier undocumented
	ApplicationIdentifier *string `json:"applicationIdentifier,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// UseSavedCredentials undocumented
	UseSavedCredentials *bool `json:"useSavedCredentials,omitempty"`
	// Credentials undocumented
	Credentials []SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}

// SynchronizationJobCollectionValidateCredentialsRequestParameter undocumented
type SynchronizationJobCollectionValidateCredentialsRequestParameter struct {
	// ApplicationIdentifier undocumented
	ApplicationIdentifier *string `json:"applicationIdentifier,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// UseSavedCredentials undocumented
	UseSavedCredentials *bool `json:"useSavedCredentials,omitempty"`
	// Credentials undocumented
	Credentials []SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}

//
type SynchronizationJobCollectionValidateCredentialsRequestBuilder struct{ BaseRequestBuilder }

// ValidateCredentials action undocumented
func (b *SynchronizationJobsCollectionRequestBuilder) ValidateCredentials(reqObj *SynchronizationJobCollectionValidateCredentialsRequestParameter) *SynchronizationJobCollectionValidateCredentialsRequestBuilder {
	bb := &SynchronizationJobCollectionValidateCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobCollectionValidateCredentialsRequest struct{ BaseRequest }

//
func (b *SynchronizationJobCollectionValidateCredentialsRequestBuilder) Request() *SynchronizationJobCollectionValidateCredentialsRequest {
	return &SynchronizationJobCollectionValidateCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobCollectionValidateCredentialsRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobCollectionValidateCredentialsRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type SynchronizationJobPauseRequestBuilder struct{ BaseRequestBuilder }

// Pause action undocumented
func (b *SynchronizationJobRequestBuilder) Pause(reqObj *SynchronizationJobPauseRequestParameter) *SynchronizationJobPauseRequestBuilder {
	bb := &SynchronizationJobPauseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/pause"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobPauseRequest struct{ BaseRequest }

//
func (b *SynchronizationJobPauseRequestBuilder) Request() *SynchronizationJobPauseRequest {
	return &SynchronizationJobPauseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobPauseRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobPauseRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type SynchronizationJobStartRequestBuilder struct{ BaseRequestBuilder }

// Start action undocumented
func (b *SynchronizationJobRequestBuilder) Start(reqObj *SynchronizationJobStartRequestParameter) *SynchronizationJobStartRequestBuilder {
	bb := &SynchronizationJobStartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/start"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobStartRequest struct{ BaseRequest }

//
func (b *SynchronizationJobStartRequestBuilder) Request() *SynchronizationJobStartRequest {
	return &SynchronizationJobStartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobStartRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobStartRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type SynchronizationJobStopRequestBuilder struct{ BaseRequestBuilder }

// Stop action undocumented
func (b *SynchronizationJobRequestBuilder) Stop(reqObj *SynchronizationJobStopRequestParameter) *SynchronizationJobStopRequestBuilder {
	bb := &SynchronizationJobStopRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/stop"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobStopRequest struct{ BaseRequest }

//
func (b *SynchronizationJobStopRequestBuilder) Request() *SynchronizationJobStopRequest {
	return &SynchronizationJobStopRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobStopRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobStopRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type SynchronizationJobApplyRequestBuilder struct{ BaseRequestBuilder }

// Apply action undocumented
func (b *SynchronizationJobRequestBuilder) Apply(reqObj *SynchronizationJobApplyRequestParameter) *SynchronizationJobApplyRequestBuilder {
	bb := &SynchronizationJobApplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/apply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobApplyRequest struct{ BaseRequest }

//
func (b *SynchronizationJobApplyRequestBuilder) Request() *SynchronizationJobApplyRequest {
	return &SynchronizationJobApplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobApplyRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobApplyRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type SynchronizationJobRestartRequestBuilder struct{ BaseRequestBuilder }

// Restart action undocumented
func (b *SynchronizationJobRequestBuilder) Restart(reqObj *SynchronizationJobRestartRequestParameter) *SynchronizationJobRestartRequestBuilder {
	bb := &SynchronizationJobRestartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/restart"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobRestartRequest struct{ BaseRequest }

//
func (b *SynchronizationJobRestartRequestBuilder) Request() *SynchronizationJobRestartRequest {
	return &SynchronizationJobRestartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobRestartRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobRestartRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type SynchronizationJobValidateCredentialsRequestBuilder struct{ BaseRequestBuilder }

// ValidateCredentials action undocumented
func (b *SynchronizationJobRequestBuilder) ValidateCredentials(reqObj *SynchronizationJobValidateCredentialsRequestParameter) *SynchronizationJobValidateCredentialsRequestBuilder {
	bb := &SynchronizationJobValidateCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationJobValidateCredentialsRequest struct{ BaseRequest }

//
func (b *SynchronizationJobValidateCredentialsRequestBuilder) Request() *SynchronizationJobValidateCredentialsRequest {
	return &SynchronizationJobValidateCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationJobValidateCredentialsRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *SynchronizationJobValidateCredentialsRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
