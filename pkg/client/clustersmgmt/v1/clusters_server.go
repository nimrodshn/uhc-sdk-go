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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

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

// ClustersServer represents the interface the manages the 'clusters' resource.
type ClustersServer interface {

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of clusters.
	List(request *ClustersListServerRequest, response *ClustersListServerResponse) error

	// Add handles a request for the 'add' method.
	//
	// Provision a new cluster and add it to the collection of clusters.
	//
	// See the `register_cluster` method for adding an existing cluster.
	Add(request *ClustersAddServerRequest, response *ClustersAddServerResponse) error

	// Cluster returns the target 'cluster' server for the given identifier.
	//
	// Returns a reference to the service that manages an specific cluster.
	Cluster(id string) ClusterServer
}

// ClustersListServerRequest is the request for the 'list' method.
type ClustersListServerRequest struct {
	path   string
	query  url.Values
	ctx    context.Context
	page   *int
	size   *int
	search *string
	total  *int
}

// GetContext returns the request Context and
// a flag indicating if the parameter has a value.
func (r *ClustersListServerRequest) GetContext() (value context.Context, ok bool) {
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
func (r *ClustersListServerRequest) Page() int {
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
func (r *ClustersListServerRequest) GetPage() (value int, ok bool) {
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
func (r *ClustersListServerRequest) Size() int {
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
func (r *ClustersListServerRequest) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Search returns the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause
// of an SQL statement, but using the names of the attributes of the cluster
// instead of the names of the columns of a table. For example, in order to
// retrieve all the clusters with a name starting with `my` in the
// `us-east-1` region the value should be:
//
// [source,sql]
// ----
// name like 'my%' and region.id = 'us-east-1'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// clusters that the user has permission to see will be returned.
func (r *ClustersListServerRequest) Search() string {
	if r != nil && r.search != nil {
		return *r.search
	}
	return ""
}

// GetSearch returns the value of the 'search' parameter and
// a flag indicating if the parameter has a value.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause
// of an SQL statement, but using the names of the attributes of the cluster
// instead of the names of the columns of a table. For example, in order to
// retrieve all the clusters with a name starting with `my` in the
// `us-east-1` region the value should be:
//
// [source,sql]
// ----
// name like 'my%' and region.id = 'us-east-1'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the
// clusters that the user has permission to see will be returned.
func (r *ClustersListServerRequest) GetSearch() (value string, ok bool) {
	ok = r != nil && r.search != nil
	if ok {
		value = *r.search
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ClustersListServerRequest) Total() int {
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
func (r *ClustersListServerRequest) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// ClustersListServerResponse is the response for the 'list' method.
type ClustersListServerResponse struct {
	status int
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *ClusterList
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *ClustersListServerResponse) Page(value int) *ClustersListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *ClustersListServerResponse) Size(value int) *ClustersListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *ClustersListServerResponse) Total(value int) *ClustersListServerResponse {
	r.total = &value
	return r
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of clusters.
func (r *ClustersListServerResponse) Items(value *ClusterList) *ClustersListServerResponse {
	r.items = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClustersListServerResponse) SetStatusCode(status int) *ClustersListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *ClustersListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(clustersListServerResponseData)
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

// clustersListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type clustersListServerResponseData struct {
	Page  *int            "json:\"page,omitempty\""
	Size  *int            "json:\"size,omitempty\""
	Total *int            "json:\"total,omitempty\""
	Items clusterListData "json:\"items,omitempty\""
}

// ClustersAddServerRequest is the request for the 'add' method.
type ClustersAddServerRequest struct {
	path  string
	query url.Values
	ctx   context.Context
	body  *Cluster
}

// GetContext returns the request Context and
// a flag indicating if the parameter has a value.
func (r *ClustersAddServerRequest) GetContext() (value context.Context, ok bool) {
	ok = r != nil && r.ctx != nil
	if ok {
		value = r.ctx
	}
	return
}

// Body returns the value of the 'body' parameter.
//
// Description of the cluster.
func (r *ClustersAddServerRequest) Body() *Cluster {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Description of the cluster.
func (r *ClustersAddServerRequest) GetBody() (value *Cluster, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *ClustersAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(clusterData)
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

// ClustersAddServerResponse is the response for the 'add' method.
type ClustersAddServerResponse struct {
	status int
	err    *errors.Error
	body   *Cluster
}

// Body sets the value of the 'body' parameter.
//
// Description of the cluster.
func (r *ClustersAddServerResponse) Body(value *Cluster) *ClustersAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *ClustersAddServerResponse) SetStatusCode(status int) *ClustersAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *ClustersAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// ClustersServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type ClustersServerAdapter struct {
	server ClustersServer
	router *mux.Router
}

func NewClustersServerAdapter(server ClustersServer, router *mux.Router) *ClustersServerAdapter {
	adapter := new(ClustersServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}/").HandlerFunc(adapter.clusterHandler)
	adapter.router.HandleFunc("/", adapter.listHandler).Methods("GET")
	adapter.router.HandleFunc("/", adapter.addHandler).Methods("POST")
	return adapter
}
func (a *ClustersServerAdapter) clusterHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.Cluster(id)
	targetAdapter := NewClusterServerAdapter(target, a.router.PathPrefix("/{id}/").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *ClustersServerAdapter) readClustersListServerRequest(r *http.Request) (*ClustersListServerRequest, error) {
	result := new(ClustersListServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	result.ctx = r.Context()
	return result, nil
}
func (a *ClustersServerAdapter) writeClustersListServerResponse(w http.ResponseWriter, r *ClustersListServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClustersServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClustersListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClustersListServerResponse)
	err = a.server.List(req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClustersListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClustersServerAdapter) readClustersAddServerRequest(r *http.Request) (*ClustersAddServerRequest, error) {
	result := new(ClustersAddServerRequest)
	result.query = r.Form
	result.path = r.URL.Path
	result.ctx = r.Context()
	err := result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *ClustersServerAdapter) writeClustersAddServerResponse(w http.ResponseWriter, r *ClustersAddServerResponse) error {
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *ClustersServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readClustersAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(ClustersAddServerResponse)
	err = a.server.Add(req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeClustersAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *ClustersServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
