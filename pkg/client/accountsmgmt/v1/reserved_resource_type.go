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

// ReservedResource represents the values of the 'reserved_resource' type.
//
//
type ReservedResource struct {
	resourceName         *string
	resourceType         *string
	byoc                 *bool
	availabilityZoneType *string
	count                *int
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ReservedResource) Empty() bool {
	return o == nil || (o.resourceName == nil &&
		o.resourceType == nil &&
		o.byoc == nil &&
		o.availabilityZoneType == nil &&
		o.count == nil &&
		true)
}

// ResourceName returns the value of the 'resource_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ReservedResource) ResourceName() string {
	if o != nil && o.resourceName != nil {
		return *o.resourceName
	}
	return ""
}

// GetResourceName returns the value of the 'resource_name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ReservedResource) GetResourceName() (value string, ok bool) {
	ok = o != nil && o.resourceName != nil
	if ok {
		value = *o.resourceName
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ReservedResource) ResourceType() string {
	if o != nil && o.resourceType != nil {
		return *o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ReservedResource) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.resourceType != nil
	if ok {
		value = *o.resourceType
	}
	return
}

// BYOC returns the value of the 'BYOC' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ReservedResource) BYOC() bool {
	if o != nil && o.byoc != nil {
		return *o.byoc
	}
	return false
}

// GetBYOC returns the value of the 'BYOC' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ReservedResource) GetBYOC() (value bool, ok bool) {
	ok = o != nil && o.byoc != nil
	if ok {
		value = *o.byoc
	}
	return
}

// AvailabilityZoneType returns the value of the 'availability_zone_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ReservedResource) AvailabilityZoneType() string {
	if o != nil && o.availabilityZoneType != nil {
		return *o.availabilityZoneType
	}
	return ""
}

// GetAvailabilityZoneType returns the value of the 'availability_zone_type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ReservedResource) GetAvailabilityZoneType() (value string, ok bool) {
	ok = o != nil && o.availabilityZoneType != nil
	if ok {
		value = *o.availabilityZoneType
	}
	return
}

// Count returns the value of the 'count' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *ReservedResource) Count() int {
	if o != nil && o.count != nil {
		return *o.count
	}
	return 0
}

// GetCount returns the value of the 'count' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *ReservedResource) GetCount() (value int, ok bool) {
	ok = o != nil && o.count != nil
	if ok {
		value = *o.count
	}
	return
}

// ReservedResourceList is a list of values of the 'reserved_resource' type.
type ReservedResourceList struct {
	items []*ReservedResource
}

// Len returns the length of the list.
func (l *ReservedResourceList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ReservedResourceList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ReservedResourceList) Get(i int) *ReservedResource {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ReservedResourceList) Slice() []*ReservedResource {
	var slice []*ReservedResource
	if l == nil {
		slice = make([]*ReservedResource, 0)
	} else {
		slice = make([]*ReservedResource, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ReservedResourceList) Each(f func(item *ReservedResource) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ReservedResourceList) Range(f func(index int, item *ReservedResource) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
