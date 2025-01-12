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
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// PermissionClient is the client of the 'permission' resource.
//
// Manages a specific permission.
type PermissionClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewPermissionClient creates a new client for the 'permission'
// resource using the given transport to sned the requests and receive the
// responses.
func NewPermissionClient(transport http.RoundTripper, path string, metric string) *PermissionClient {
	client := new(PermissionClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the permission.
func (c *PermissionClient) Get() *PermissionGetRequest {
	request := new(PermissionGetRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// Delete creates a request for the 'delete' method.
//
// Deletes the permission.
func (c *PermissionClient) Delete() *PermissionDeleteRequest {
	request := new(PermissionDeleteRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// PermissionGetRequest is the request for the 'get' method.
type PermissionGetRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *PermissionGetRequest) Parameter(name string, value interface{}) *PermissionGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *PermissionGetRequest) Header(name string, value interface{}) *PermissionGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *PermissionGetRequest) Send() (result *PermissionGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *PermissionGetRequest) SendContext(ctx context.Context) (result *PermissionGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
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
	result = new(PermissionGetResponse)
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

// PermissionGetResponse is the response for the 'get' method.
type PermissionGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Permission
}

// Status returns the response status code.
func (r *PermissionGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *PermissionGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *PermissionGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *PermissionGetResponse) Body() *Permission {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *PermissionGetResponse) GetBody() (value *Permission, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *PermissionGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(permissionData)
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

// PermissionDeleteRequest is the request for the 'delete' method.
type PermissionDeleteRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *PermissionDeleteRequest) Parameter(name string, value interface{}) *PermissionDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *PermissionDeleteRequest) Header(name string, value interface{}) *PermissionDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *PermissionDeleteRequest) Send() (result *PermissionDeleteResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *PermissionDeleteRequest) SendContext(ctx context.Context) (result *PermissionDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodDelete,
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
	result = new(PermissionDeleteResponse)
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
	return
}

// PermissionDeleteResponse is the response for the 'delete' method.
type PermissionDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *PermissionDeleteResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *PermissionDeleteResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *PermissionDeleteResponse) Error() *errors.Error {
	return r.err
}
