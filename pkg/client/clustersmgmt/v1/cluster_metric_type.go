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
	time "time"
)

// ClusterMetric represents the values of the 'cluster_metric' type.
//
// Metric describing the total and used amount of some resource (like RAM, CPU and storage) in
// a cluster.
type ClusterMetric struct {
	updatedTimestamp *time.Time
	total            *Value
	used             *Value
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterMetric) Empty() bool {
	return o == nil || (o.updatedTimestamp == nil &&
		o.total == nil &&
		o.used == nil &&
		true)
}

// UpdatedTimestamp returns the value of the 'updated_timestamp' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Collection timestamp of the metric.
func (o *ClusterMetric) UpdatedTimestamp() time.Time {
	if o != nil && o.updatedTimestamp != nil {
		return *o.updatedTimestamp
	}
	return time.Time{}
}

// GetUpdatedTimestamp returns the value of the 'updated_timestamp' attribute and
// a flag indicating if the attribute has a value.
//
// Collection timestamp of the metric.
func (o *ClusterMetric) GetUpdatedTimestamp() (value time.Time, ok bool) {
	ok = o != nil && o.updatedTimestamp != nil
	if ok {
		value = *o.updatedTimestamp
	}
	return
}

// Total returns the value of the 'total' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Total amount of the resource that exists in the cluster. For example the total amount
// of RAM.
func (o *ClusterMetric) Total() *Value {
	if o == nil {
		return nil
	}
	return o.total
}

// GetTotal returns the value of the 'total' attribute and
// a flag indicating if the attribute has a value.
//
// Total amount of the resource that exists in the cluster. For example the total amount
// of RAM.
func (o *ClusterMetric) GetTotal() (value *Value, ok bool) {
	ok = o != nil && o.total != nil
	if ok {
		value = o.total
	}
	return
}

// Used returns the value of the 'used' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Amount of the resource that is currently in use in the cluster. Fore example the amount
// of RAM in use.
func (o *ClusterMetric) Used() *Value {
	if o == nil {
		return nil
	}
	return o.used
}

// GetUsed returns the value of the 'used' attribute and
// a flag indicating if the attribute has a value.
//
// Amount of the resource that is currently in use in the cluster. Fore example the amount
// of RAM in use.
func (o *ClusterMetric) GetUsed() (value *Value, ok bool) {
	ok = o != nil && o.used != nil
	if ok {
		value = o.used
	}
	return
}

// ClusterMetricList is a list of values of the 'cluster_metric' type.
type ClusterMetricList struct {
	items []*ClusterMetric
}

// Len returns the length of the list.
func (l *ClusterMetricList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterMetricList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterMetricList) Get(i int) *ClusterMetric {
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
func (l *ClusterMetricList) Slice() []*ClusterMetric {
	var slice []*ClusterMetric
	if l == nil {
		slice = make([]*ClusterMetric, 0)
	} else {
		slice = make([]*ClusterMetric, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterMetricList) Each(f func(item *ClusterMetric) bool) {
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
func (l *ClusterMetricList) Range(f func(index int, item *ClusterMetric) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
