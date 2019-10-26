// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DataClassificationServiceRequestBuilder is request builder for DataClassificationService
type DataClassificationServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataClassificationServiceRequest
func (b *DataClassificationServiceRequestBuilder) Request() *DataClassificationServiceRequest {
	return &DataClassificationServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataClassificationServiceRequest is request for DataClassificationService
type DataClassificationServiceRequest struct{ BaseRequest }

// Get performs GET request for DataClassificationService
func (r *DataClassificationServiceRequest) Get() (resObj *DataClassificationService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataClassificationService
func (r *DataClassificationServiceRequest) Update(reqObj *DataClassificationService) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataClassificationService
func (r *DataClassificationServiceRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ClassifyFile returns request builder for FileClassificationRequestObject collection
func (b *DataClassificationServiceRequestBuilder) ClassifyFile() *DataClassificationServiceClassifyFileCollectionRequestBuilder {
	bb := &DataClassificationServiceClassifyFileCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classifyFile"
	return bb
}

// DataClassificationServiceClassifyFileCollectionRequestBuilder is request builder for FileClassificationRequestObject collection
type DataClassificationServiceClassifyFileCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for FileClassificationRequestObject collection
func (b *DataClassificationServiceClassifyFileCollectionRequestBuilder) Request() *DataClassificationServiceClassifyFileCollectionRequest {
	return &DataClassificationServiceClassifyFileCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for FileClassificationRequestObject item
func (b *DataClassificationServiceClassifyFileCollectionRequestBuilder) ID(id string) *FileClassificationRequestObjectRequestBuilder {
	bb := &FileClassificationRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceClassifyFileCollectionRequest is request for FileClassificationRequestObject collection
type DataClassificationServiceClassifyFileCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for FileClassificationRequestObject collection
func (r *DataClassificationServiceClassifyFileCollectionRequest) Paging(method, path string, obj interface{}) ([]FileClassificationRequestObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []FileClassificationRequestObject
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
			value  []FileClassificationRequestObject
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

// Get performs GET request for FileClassificationRequestObject collection
func (r *DataClassificationServiceClassifyFileCollectionRequest) Get() ([]FileClassificationRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for FileClassificationRequestObject collection
func (r *DataClassificationServiceClassifyFileCollectionRequest) Add(reqObj *FileClassificationRequestObject) (resObj *FileClassificationRequestObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ClassifyFileJobs returns request builder for JobResponseBase collection
func (b *DataClassificationServiceRequestBuilder) ClassifyFileJobs() *DataClassificationServiceClassifyFileJobsCollectionRequestBuilder {
	bb := &DataClassificationServiceClassifyFileJobsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classifyFileJobs"
	return bb
}

// DataClassificationServiceClassifyFileJobsCollectionRequestBuilder is request builder for JobResponseBase collection
type DataClassificationServiceClassifyFileJobsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for JobResponseBase collection
func (b *DataClassificationServiceClassifyFileJobsCollectionRequestBuilder) Request() *DataClassificationServiceClassifyFileJobsCollectionRequest {
	return &DataClassificationServiceClassifyFileJobsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for JobResponseBase item
func (b *DataClassificationServiceClassifyFileJobsCollectionRequestBuilder) ID(id string) *JobResponseBaseRequestBuilder {
	bb := &JobResponseBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceClassifyFileJobsCollectionRequest is request for JobResponseBase collection
type DataClassificationServiceClassifyFileJobsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for JobResponseBase collection
func (r *DataClassificationServiceClassifyFileJobsCollectionRequest) Paging(method, path string, obj interface{}) ([]JobResponseBase, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []JobResponseBase
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
			value  []JobResponseBase
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

// Get performs GET request for JobResponseBase collection
func (r *DataClassificationServiceClassifyFileJobsCollectionRequest) Get() ([]JobResponseBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for JobResponseBase collection
func (r *DataClassificationServiceClassifyFileJobsCollectionRequest) Add(reqObj *JobResponseBase) (resObj *JobResponseBase, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ClassifyText returns request builder for TextClassificationRequestObject collection
func (b *DataClassificationServiceRequestBuilder) ClassifyText() *DataClassificationServiceClassifyTextCollectionRequestBuilder {
	bb := &DataClassificationServiceClassifyTextCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classifyText"
	return bb
}

// DataClassificationServiceClassifyTextCollectionRequestBuilder is request builder for TextClassificationRequestObject collection
type DataClassificationServiceClassifyTextCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TextClassificationRequestObject collection
func (b *DataClassificationServiceClassifyTextCollectionRequestBuilder) Request() *DataClassificationServiceClassifyTextCollectionRequest {
	return &DataClassificationServiceClassifyTextCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TextClassificationRequestObject item
func (b *DataClassificationServiceClassifyTextCollectionRequestBuilder) ID(id string) *TextClassificationRequestObjectRequestBuilder {
	bb := &TextClassificationRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceClassifyTextCollectionRequest is request for TextClassificationRequestObject collection
type DataClassificationServiceClassifyTextCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TextClassificationRequestObject collection
func (r *DataClassificationServiceClassifyTextCollectionRequest) Paging(method, path string, obj interface{}) ([]TextClassificationRequestObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TextClassificationRequestObject
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
			value  []TextClassificationRequestObject
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

// Get performs GET request for TextClassificationRequestObject collection
func (r *DataClassificationServiceClassifyTextCollectionRequest) Get() ([]TextClassificationRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TextClassificationRequestObject collection
func (r *DataClassificationServiceClassifyTextCollectionRequest) Add(reqObj *TextClassificationRequestObject) (resObj *TextClassificationRequestObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ClassifyTextJobs returns request builder for JobResponseBase collection
func (b *DataClassificationServiceRequestBuilder) ClassifyTextJobs() *DataClassificationServiceClassifyTextJobsCollectionRequestBuilder {
	bb := &DataClassificationServiceClassifyTextJobsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classifyTextJobs"
	return bb
}

// DataClassificationServiceClassifyTextJobsCollectionRequestBuilder is request builder for JobResponseBase collection
type DataClassificationServiceClassifyTextJobsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for JobResponseBase collection
func (b *DataClassificationServiceClassifyTextJobsCollectionRequestBuilder) Request() *DataClassificationServiceClassifyTextJobsCollectionRequest {
	return &DataClassificationServiceClassifyTextJobsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for JobResponseBase item
func (b *DataClassificationServiceClassifyTextJobsCollectionRequestBuilder) ID(id string) *JobResponseBaseRequestBuilder {
	bb := &JobResponseBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceClassifyTextJobsCollectionRequest is request for JobResponseBase collection
type DataClassificationServiceClassifyTextJobsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for JobResponseBase collection
func (r *DataClassificationServiceClassifyTextJobsCollectionRequest) Paging(method, path string, obj interface{}) ([]JobResponseBase, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []JobResponseBase
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
			value  []JobResponseBase
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

// Get performs GET request for JobResponseBase collection
func (r *DataClassificationServiceClassifyTextJobsCollectionRequest) Get() ([]JobResponseBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for JobResponseBase collection
func (r *DataClassificationServiceClassifyTextJobsCollectionRequest) Add(reqObj *JobResponseBase) (resObj *JobResponseBase, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// EvaluateLabelJobs returns request builder for JobResponseBase collection
func (b *DataClassificationServiceRequestBuilder) EvaluateLabelJobs() *DataClassificationServiceEvaluateLabelJobsCollectionRequestBuilder {
	bb := &DataClassificationServiceEvaluateLabelJobsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/evaluateLabelJobs"
	return bb
}

// DataClassificationServiceEvaluateLabelJobsCollectionRequestBuilder is request builder for JobResponseBase collection
type DataClassificationServiceEvaluateLabelJobsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for JobResponseBase collection
func (b *DataClassificationServiceEvaluateLabelJobsCollectionRequestBuilder) Request() *DataClassificationServiceEvaluateLabelJobsCollectionRequest {
	return &DataClassificationServiceEvaluateLabelJobsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for JobResponseBase item
func (b *DataClassificationServiceEvaluateLabelJobsCollectionRequestBuilder) ID(id string) *JobResponseBaseRequestBuilder {
	bb := &JobResponseBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceEvaluateLabelJobsCollectionRequest is request for JobResponseBase collection
type DataClassificationServiceEvaluateLabelJobsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for JobResponseBase collection
func (r *DataClassificationServiceEvaluateLabelJobsCollectionRequest) Paging(method, path string, obj interface{}) ([]JobResponseBase, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []JobResponseBase
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
			value  []JobResponseBase
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

// Get performs GET request for JobResponseBase collection
func (r *DataClassificationServiceEvaluateLabelJobsCollectionRequest) Get() ([]JobResponseBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for JobResponseBase collection
func (r *DataClassificationServiceEvaluateLabelJobsCollectionRequest) Add(reqObj *JobResponseBase) (resObj *JobResponseBase, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ExactMatchDataStores returns request builder for ExactMatchDataStore collection
func (b *DataClassificationServiceRequestBuilder) ExactMatchDataStores() *DataClassificationServiceExactMatchDataStoresCollectionRequestBuilder {
	bb := &DataClassificationServiceExactMatchDataStoresCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/exactMatchDataStores"
	return bb
}

// DataClassificationServiceExactMatchDataStoresCollectionRequestBuilder is request builder for ExactMatchDataStore collection
type DataClassificationServiceExactMatchDataStoresCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExactMatchDataStore collection
func (b *DataClassificationServiceExactMatchDataStoresCollectionRequestBuilder) Request() *DataClassificationServiceExactMatchDataStoresCollectionRequest {
	return &DataClassificationServiceExactMatchDataStoresCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExactMatchDataStore item
func (b *DataClassificationServiceExactMatchDataStoresCollectionRequestBuilder) ID(id string) *ExactMatchDataStoreRequestBuilder {
	bb := &ExactMatchDataStoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceExactMatchDataStoresCollectionRequest is request for ExactMatchDataStore collection
type DataClassificationServiceExactMatchDataStoresCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExactMatchDataStore collection
func (r *DataClassificationServiceExactMatchDataStoresCollectionRequest) Paging(method, path string, obj interface{}) ([]ExactMatchDataStore, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ExactMatchDataStore
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
			value  []ExactMatchDataStore
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

// Get performs GET request for ExactMatchDataStore collection
func (r *DataClassificationServiceExactMatchDataStoresCollectionRequest) Get() ([]ExactMatchDataStore, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ExactMatchDataStore collection
func (r *DataClassificationServiceExactMatchDataStoresCollectionRequest) Add(reqObj *ExactMatchDataStore) (resObj *ExactMatchDataStore, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ExactMatchUploadAgents returns request builder for ExactMatchUploadAgent collection
func (b *DataClassificationServiceRequestBuilder) ExactMatchUploadAgents() *DataClassificationServiceExactMatchUploadAgentsCollectionRequestBuilder {
	bb := &DataClassificationServiceExactMatchUploadAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/exactMatchUploadAgents"
	return bb
}

// DataClassificationServiceExactMatchUploadAgentsCollectionRequestBuilder is request builder for ExactMatchUploadAgent collection
type DataClassificationServiceExactMatchUploadAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExactMatchUploadAgent collection
func (b *DataClassificationServiceExactMatchUploadAgentsCollectionRequestBuilder) Request() *DataClassificationServiceExactMatchUploadAgentsCollectionRequest {
	return &DataClassificationServiceExactMatchUploadAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExactMatchUploadAgent item
func (b *DataClassificationServiceExactMatchUploadAgentsCollectionRequestBuilder) ID(id string) *ExactMatchUploadAgentRequestBuilder {
	bb := &ExactMatchUploadAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceExactMatchUploadAgentsCollectionRequest is request for ExactMatchUploadAgent collection
type DataClassificationServiceExactMatchUploadAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExactMatchUploadAgent collection
func (r *DataClassificationServiceExactMatchUploadAgentsCollectionRequest) Paging(method, path string, obj interface{}) ([]ExactMatchUploadAgent, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ExactMatchUploadAgent
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
			value  []ExactMatchUploadAgent
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

// Get performs GET request for ExactMatchUploadAgent collection
func (r *DataClassificationServiceExactMatchUploadAgentsCollectionRequest) Get() ([]ExactMatchUploadAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ExactMatchUploadAgent collection
func (r *DataClassificationServiceExactMatchUploadAgentsCollectionRequest) Add(reqObj *ExactMatchUploadAgent) (resObj *ExactMatchUploadAgent, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Jobs returns request builder for JobResponseBase collection
func (b *DataClassificationServiceRequestBuilder) Jobs() *DataClassificationServiceJobsCollectionRequestBuilder {
	bb := &DataClassificationServiceJobsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/jobs"
	return bb
}

// DataClassificationServiceJobsCollectionRequestBuilder is request builder for JobResponseBase collection
type DataClassificationServiceJobsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for JobResponseBase collection
func (b *DataClassificationServiceJobsCollectionRequestBuilder) Request() *DataClassificationServiceJobsCollectionRequest {
	return &DataClassificationServiceJobsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for JobResponseBase item
func (b *DataClassificationServiceJobsCollectionRequestBuilder) ID(id string) *JobResponseBaseRequestBuilder {
	bb := &JobResponseBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceJobsCollectionRequest is request for JobResponseBase collection
type DataClassificationServiceJobsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for JobResponseBase collection
func (r *DataClassificationServiceJobsCollectionRequest) Paging(method, path string, obj interface{}) ([]JobResponseBase, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []JobResponseBase
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
			value  []JobResponseBase
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

// Get performs GET request for JobResponseBase collection
func (r *DataClassificationServiceJobsCollectionRequest) Get() ([]JobResponseBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for JobResponseBase collection
func (r *DataClassificationServiceJobsCollectionRequest) Add(reqObj *JobResponseBase) (resObj *JobResponseBase, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// SensitiveTypes returns request builder for SensitiveType collection
func (b *DataClassificationServiceRequestBuilder) SensitiveTypes() *DataClassificationServiceSensitiveTypesCollectionRequestBuilder {
	bb := &DataClassificationServiceSensitiveTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sensitiveTypes"
	return bb
}

// DataClassificationServiceSensitiveTypesCollectionRequestBuilder is request builder for SensitiveType collection
type DataClassificationServiceSensitiveTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SensitiveType collection
func (b *DataClassificationServiceSensitiveTypesCollectionRequestBuilder) Request() *DataClassificationServiceSensitiveTypesCollectionRequest {
	return &DataClassificationServiceSensitiveTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SensitiveType item
func (b *DataClassificationServiceSensitiveTypesCollectionRequestBuilder) ID(id string) *SensitiveTypeRequestBuilder {
	bb := &SensitiveTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceSensitiveTypesCollectionRequest is request for SensitiveType collection
type DataClassificationServiceSensitiveTypesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SensitiveType collection
func (r *DataClassificationServiceSensitiveTypesCollectionRequest) Paging(method, path string, obj interface{}) ([]SensitiveType, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SensitiveType
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
			value  []SensitiveType
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

// Get performs GET request for SensitiveType collection
func (r *DataClassificationServiceSensitiveTypesCollectionRequest) Get() ([]SensitiveType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SensitiveType collection
func (r *DataClassificationServiceSensitiveTypesCollectionRequest) Add(reqObj *SensitiveType) (resObj *SensitiveType, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// SensitivityLabels returns request builder for SensitivityLabel collection
func (b *DataClassificationServiceRequestBuilder) SensitivityLabels() *DataClassificationServiceSensitivityLabelsCollectionRequestBuilder {
	bb := &DataClassificationServiceSensitivityLabelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sensitivityLabels"
	return bb
}

// DataClassificationServiceSensitivityLabelsCollectionRequestBuilder is request builder for SensitivityLabel collection
type DataClassificationServiceSensitivityLabelsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SensitivityLabel collection
func (b *DataClassificationServiceSensitivityLabelsCollectionRequestBuilder) Request() *DataClassificationServiceSensitivityLabelsCollectionRequest {
	return &DataClassificationServiceSensitivityLabelsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SensitivityLabel item
func (b *DataClassificationServiceSensitivityLabelsCollectionRequestBuilder) ID(id string) *SensitivityLabelRequestBuilder {
	bb := &SensitivityLabelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceSensitivityLabelsCollectionRequest is request for SensitivityLabel collection
type DataClassificationServiceSensitivityLabelsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SensitivityLabel collection
func (r *DataClassificationServiceSensitivityLabelsCollectionRequest) Paging(method, path string, obj interface{}) ([]SensitivityLabel, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SensitivityLabel
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
			value  []SensitivityLabel
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

// Get performs GET request for SensitivityLabel collection
func (r *DataClassificationServiceSensitivityLabelsCollectionRequest) Get() ([]SensitivityLabel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SensitivityLabel collection
func (r *DataClassificationServiceSensitivityLabelsCollectionRequest) Add(reqObj *SensitivityLabel) (resObj *SensitivityLabel, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
