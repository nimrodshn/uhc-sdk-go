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

// ClusterAuthorizationRequestBuilder contains the data and logic needed to build 'cluster_authorization_request' objects.
//
//
type ClusterAuthorizationRequestBuilder struct {
	clusterID        *string
	accountUsername  *string
	managed          *bool
	reserve          *bool
	byoc             *bool
	availabilityZone *string
	resources        []*ReservedResourceBuilder
}

// NewClusterAuthorizationRequest creates a new builder of 'cluster_authorization_request' objects.
func NewClusterAuthorizationRequest() *ClusterAuthorizationRequestBuilder {
	return new(ClusterAuthorizationRequestBuilder)
}

// ClusterID sets the value of the 'cluster_ID' attribute
// to the given value.
//
//
func (b *ClusterAuthorizationRequestBuilder) ClusterID(value string) *ClusterAuthorizationRequestBuilder {
	b.clusterID = &value
	return b
}

// AccountUsername sets the value of the 'account_username' attribute
// to the given value.
//
//
func (b *ClusterAuthorizationRequestBuilder) AccountUsername(value string) *ClusterAuthorizationRequestBuilder {
	b.accountUsername = &value
	return b
}

// Managed sets the value of the 'managed' attribute
// to the given value.
//
//
func (b *ClusterAuthorizationRequestBuilder) Managed(value bool) *ClusterAuthorizationRequestBuilder {
	b.managed = &value
	return b
}

// Reserve sets the value of the 'reserve' attribute
// to the given value.
//
//
func (b *ClusterAuthorizationRequestBuilder) Reserve(value bool) *ClusterAuthorizationRequestBuilder {
	b.reserve = &value
	return b
}

// BYOC sets the value of the 'BYOC' attribute
// to the given value.
//
//
func (b *ClusterAuthorizationRequestBuilder) BYOC(value bool) *ClusterAuthorizationRequestBuilder {
	b.byoc = &value
	return b
}

// AvailabilityZone sets the value of the 'availability_zone' attribute
// to the given value.
//
//
func (b *ClusterAuthorizationRequestBuilder) AvailabilityZone(value string) *ClusterAuthorizationRequestBuilder {
	b.availabilityZone = &value
	return b
}

// Resources sets the value of the 'resources' attribute
// to the given values.
//
//
func (b *ClusterAuthorizationRequestBuilder) Resources(values ...*ReservedResourceBuilder) *ClusterAuthorizationRequestBuilder {
	b.resources = make([]*ReservedResourceBuilder, len(values))
	copy(b.resources, values)
	return b
}

// Build creates a 'cluster_authorization_request' object using the configuration stored in the builder.
func (b *ClusterAuthorizationRequestBuilder) Build() (object *ClusterAuthorizationRequest, err error) {
	object = new(ClusterAuthorizationRequest)
	if b.clusterID != nil {
		object.clusterID = b.clusterID
	}
	if b.accountUsername != nil {
		object.accountUsername = b.accountUsername
	}
	if b.managed != nil {
		object.managed = b.managed
	}
	if b.reserve != nil {
		object.reserve = b.reserve
	}
	if b.byoc != nil {
		object.byoc = b.byoc
	}
	if b.availabilityZone != nil {
		object.availabilityZone = b.availabilityZone
	}
	if b.resources != nil {
		object.resources = new(ReservedResourceList)
		object.resources.items = make([]*ReservedResource, len(b.resources))
		for i, item := range b.resources {
			object.resources.items[i], err = item.Build()
			if err != nil {
				return
			}
		}
	}
	return
}
