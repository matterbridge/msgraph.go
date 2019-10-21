// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupLifecyclePolicyCollectionRenewGroupRequestParameter undocumented
type GroupLifecyclePolicyCollectionRenewGroupRequestParameter struct {
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
}

// GroupLifecyclePolicyAddGroupRequestParameter undocumented
type GroupLifecyclePolicyAddGroupRequestParameter struct {
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
}

// GroupLifecyclePolicyRemoveGroupRequestParameter undocumented
type GroupLifecyclePolicyRemoveGroupRequestParameter struct {
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
}

//
type GroupLifecyclePolicyAddGroupRequestBuilder struct{ BaseRequestBuilder }

// AddGroup action undocumented
func (b *GroupLifecyclePolicyRequestBuilder) AddGroup(reqObj *GroupLifecyclePolicyAddGroupRequestParameter) *GroupLifecyclePolicyAddGroupRequestBuilder {
	bb := &GroupLifecyclePolicyAddGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupLifecyclePolicyAddGroupRequest struct{ BaseRequest }

//
func (b *GroupLifecyclePolicyAddGroupRequestBuilder) Request() *GroupLifecyclePolicyAddGroupRequest {
	return &GroupLifecyclePolicyAddGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupLifecyclePolicyAddGroupRequest) Do(method, path string, reqObj interface{}) (resObj *bool, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *GroupLifecyclePolicyAddGroupRequest) Post() (*bool, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupLifecyclePolicyRemoveGroupRequestBuilder struct{ BaseRequestBuilder }

// RemoveGroup action undocumented
func (b *GroupLifecyclePolicyRequestBuilder) RemoveGroup(reqObj *GroupLifecyclePolicyRemoveGroupRequestParameter) *GroupLifecyclePolicyRemoveGroupRequestBuilder {
	bb := &GroupLifecyclePolicyRemoveGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupLifecyclePolicyRemoveGroupRequest struct{ BaseRequest }

//
func (b *GroupLifecyclePolicyRemoveGroupRequestBuilder) Request() *GroupLifecyclePolicyRemoveGroupRequest {
	return &GroupLifecyclePolicyRemoveGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupLifecyclePolicyRemoveGroupRequest) Do(method, path string, reqObj interface{}) (resObj *bool, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *GroupLifecyclePolicyRemoveGroupRequest) Post() (*bool, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupLifecyclePolicyCollectionRenewGroupRequestBuilder struct{ BaseRequestBuilder }

// RenewGroup action undocumented
func (b *GroupGroupLifecyclePoliciesCollectionRequestBuilder) RenewGroup(reqObj *GroupLifecyclePolicyCollectionRenewGroupRequestParameter) *GroupLifecyclePolicyCollectionRenewGroupRequestBuilder {
	bb := &GroupLifecyclePolicyCollectionRenewGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/renewGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupLifecyclePolicyCollectionRenewGroupRequest struct{ BaseRequest }

//
func (b *GroupLifecyclePolicyCollectionRenewGroupRequestBuilder) Request() *GroupLifecyclePolicyCollectionRenewGroupRequest {
	return &GroupLifecyclePolicyCollectionRenewGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupLifecyclePolicyCollectionRenewGroupRequest) Do(method, path string, reqObj interface{}) (resObj *bool, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *GroupLifecyclePolicyCollectionRenewGroupRequest) Post() (*bool, error) {
	return r.Do("POST", "", r.requestObject)
}
