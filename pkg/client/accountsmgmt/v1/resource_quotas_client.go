/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// ResourceQuotasClient is the client of the 'resource_quotas' resource.
//
// Manages the collection of resource quotas for an organization.
type ResourceQuotasClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewResourceQuotasClient creates a new client for the 'resource_quotas'
// resource using the given transport to sned the requests and receive the
// responses.
func NewResourceQuotasClient(transport http.RoundTripper, path string, metric string) *ResourceQuotasClient {
	client := new(ResourceQuotasClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of resource quotas.
func (c *ResourceQuotasClient) List() *ResourceQuotasListRequest {
	request := new(ResourceQuotasListRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// Add creates a request for the 'add' method.
//
// Creates a new resource quota.
func (c *ResourceQuotasClient) Add() *ResourceQuotasAddRequest {
	request := new(ResourceQuotasAddRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// ResourceQuota returns the target 'resource_quota' resource for the given identifier.
//
// Reference to the service that manages an specific resource quota.
func (c *ResourceQuotasClient) ResourceQuota(id string) *ResourceQuotaClient {
	return NewResourceQuotaClient(
		c.transport,
		path.Join(c.path, id),
		path.Join(c.metric, "-"),
	)
}

// ResourceQuotasListRequest is the request for the 'list' method.
type ResourceQuotasListRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	page      *int
	size      *int
	total     *int
}

// Parameter adds a query parameter.
func (r *ResourceQuotasListRequest) Parameter(name string, value interface{}) *ResourceQuotasListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ResourceQuotasListRequest) Header(name string, value interface{}) *ResourceQuotasListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *ResourceQuotasListRequest) Page(value int) *ResourceQuotasListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *ResourceQuotasListRequest) Size(value int) *ResourceQuotasListRequest {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ResourceQuotasListRequest) Total(value int) *ResourceQuotasListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *ResourceQuotasListRequest) Send() (result *ResourceQuotasListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ResourceQuotasListRequest) SendContext(ctx context.Context) (result *ResourceQuotasListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	if r.total != nil {
		helpers.AddValue(&query, "total", *r.total)
	}
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(ResourceQuotasListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// ResourceQuotasListResponse is the response for the 'list' method.
type ResourceQuotasListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *ResourceQuotaList
}

// Status returns the response status code.
func (r *ResourceQuotasListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ResourceQuotasListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ResourceQuotasListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *ResourceQuotasListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *ResourceQuotasListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *ResourceQuotasListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *ResourceQuotasListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ResourceQuotasListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ResourceQuotasListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of resource quotas.
func (r *ResourceQuotasListResponse) Items() *ResourceQuotaList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of resource quotas.
func (r *ResourceQuotasListResponse) GetItems() (value *ResourceQuotaList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *ResourceQuotasListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(resourceQuotasListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	return err
}

// resourceQuotasListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type resourceQuotasListResponseData struct {
	Page  *int                  "json:\"page,omitempty\""
	Size  *int                  "json:\"size,omitempty\""
	Total *int                  "json:\"total,omitempty\""
	Items resourceQuotaListData "json:\"items,omitempty\""
}

// ResourceQuotasAddRequest is the request for the 'add' method.
type ResourceQuotasAddRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	body      *ResourceQuota
}

// Parameter adds a query parameter.
func (r *ResourceQuotasAddRequest) Parameter(name string, value interface{}) *ResourceQuotasAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *ResourceQuotasAddRequest) Header(name string, value interface{}) *ResourceQuotasAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Resource quota data.
func (r *ResourceQuotasAddRequest) Body(value *ResourceQuota) *ResourceQuotasAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *ResourceQuotasAddRequest) Send() (result *ResourceQuotasAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *ResourceQuotasAddRequest) SendContext(ctx context.Context) (result *ResourceQuotasAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodPost,
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(ResourceQuotasAddResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'add' method.
func (r *ResourceQuotasAddRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ResourceQuotasAddResponse is the response for the 'add' method.
type ResourceQuotasAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *ResourceQuota
}

// Status returns the response status code.
func (r *ResourceQuotasAddResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *ResourceQuotasAddResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *ResourceQuotasAddResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Resource quota data.
func (r *ResourceQuotasAddResponse) Body() *ResourceQuota {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Resource quota data.
func (r *ResourceQuotasAddResponse) GetBody() (value *ResourceQuota, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'add' method.
func (r *ResourceQuotasAddResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(resourceQuotaData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}
