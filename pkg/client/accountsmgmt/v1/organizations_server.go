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
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
)

// OrganizationsServer represents the interface the manages the 'organizations' resource.
type OrganizationsServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves a list of organizations.
	List(request *OrganizationsListServerRequest, response *OrganizationsListServerResponse) error

	// Add handles a request for the 'add' method.
	//
	// Creates a new organization.
	Add(request *OrganizationsAddServerRequest, response *OrganizationsAddServerResponse) error

	// Organization returns the target 'organization' server for the given identifier.
	//
	// Reference to the service that manages a specific organization.
	Organization(id string) OrganizationServer
}

// OrganizationsListServerRequest is the request for the 'list' method.
type OrganizationsListServerRequest struct {
	path  string
	query url.Values
	ctx   context.Context
	page  *int
	size  *int
	total *int
}

// GetContext returns the request Context and
// a flag indicating if the parameter has a value.
func (r *OrganizationsListServerRequest) GetContext() (value context.Context, ok bool) {
	ok = r != nil && r.ctx != nil
	if ok {
		value = r.ctx
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *OrganizationsListServerRequest) Page() int {
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
func (r *OrganizationsListServerRequest) GetPage() (value int, ok bool) {
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
func (r *OrganizationsListServerRequest) Size() int {
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
func (r *OrganizationsListServerRequest) GetSize() (value int, ok bool) {
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
func (r *OrganizationsListServerRequest) Total() int {
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
func (r *OrganizationsListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// OrganizationsListServerResponse is the response for the 'list' method.
type OrganizationsListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *OrganizationList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *OrganizationsListServerResponse) Page(value int) *OrganizationsListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *OrganizationsListServerResponse) Size(value int) *OrganizationsListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *OrganizationsListServerResponse) Total(value int) *OrganizationsListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of organizations.
func (r *OrganizationsListServerResponse) Items(value *OrganizationList) *OrganizationsListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *OrganizationsListServerResponse) SetStatusCode(status int) *OrganizationsListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *OrganizationsListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(organizationsListServerResponseData)
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// organizationsListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type organizationsListServerResponseData struct {
	Page  *int                 "json:\"page,omitempty\""
	Size  *int                 "json:\"size,omitempty\""
	Total *int                 "json:\"total,omitempty\""
	Items organizationListData "json:\"items,omitempty\""
}

// OrganizationsAddServerRequest is the request for the 'add' method.
type OrganizationsAddServerRequest struct {
	path  string
	query url.Values
	ctx   context.Context
	body  *Organization
}

// GetContext returns the request Context and
// a flag indicating if the parameter has a value.
func (r *OrganizationsAddServerRequest) GetContext() (value context.Context, ok bool) {
	ok = r != nil && r.ctx != nil
	if ok {
		value = r.ctx
	}
	return
}

// Body returns the value of the 'body' parameter.
//
// Organization data.
func (r *OrganizationsAddServerRequest) Body() *Organization {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Organization data.
func (r *OrganizationsAddServerRequest) GetBody() (value *Organization, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *OrganizationsAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(organizationData)
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

// OrganizationsAddServerResponse is the response for the 'add' method.
type OrganizationsAddServerResponse struct {
	status int
	err    *errors.Error
	body   *Organization
}

// Body sets the value of the 'body' parameter.
//
// Organization data.
func (r *OrganizationsAddServerResponse) Body(value *Organization) *OrganizationsAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *OrganizationsAddServerResponse) SetStatusCode(status int) *OrganizationsAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *OrganizationsAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// OrganizationsServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type OrganizationsServerAdapter struct {
	server OrganizationsServer
	router *mux.Router
}

func NewOrganizationsServerAdapter(server OrganizationsServer, router *mux.Router) *OrganizationsServerAdapter {
	adapter := new(OrganizationsServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}/").HandlerFunc(adapter.organizationHandler)
	adapter.router.HandleFunc("/", adapter.listHandler).Methods("GET")
	adapter.router.HandleFunc("/", adapter.addHandler).Methods("POST")
	return adapter
}
func (a *OrganizationsServerAdapter) organizationHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Organization(id)
	targetAdapter := NewOrganizationServerAdapter(target, a.router.PathPrefix("/{id}/").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *OrganizationsServerAdapter) readOrganizationsListServerRequest(r *http.Request) (*OrganizationsListServerRequest, error) {
	result := new(OrganizationsListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	result.ctx = r.Context()
	return result, nil
}
func (a *OrganizationsServerAdapter) writeOrganizationsListServerResponse(w http.ResponseWriter, r *OrganizationsListServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *OrganizationsServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readOrganizationsListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(OrganizationsListServerResponse)
	err = a.server.List(req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeOrganizationsListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *OrganizationsServerAdapter) readOrganizationsAddServerRequest(r *http.Request) (*OrganizationsAddServerRequest, error) {
	result := new(OrganizationsAddServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	result.ctx = r.Context()
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *OrganizationsServerAdapter) writeOrganizationsAddServerResponse(w http.ResponseWriter, r *OrganizationsAddServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *OrganizationsServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readOrganizationsAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(OrganizationsAddServerResponse)
	err = a.server.Add(req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeOrganizationsAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *OrganizationsServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
