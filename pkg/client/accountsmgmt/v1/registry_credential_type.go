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

// RegistryCredentialKind is the name of the type used to represent objects
// of type 'registry_credential'.
const RegistryCredentialKind = "RegistryCredential"

// RegistryCredentialLinkKind is the name of the type used to represent links
// to objects of type 'registry_credential'.
const RegistryCredentialLinkKind = "RegistryCredentialLink"

// RegistryCredentialNilKind is the name of the type used to nil references
// to objects of type 'registry_credential'.
const RegistryCredentialNilKind = "RegistryCredentialNil"

// RegistryCredential represents the values of the 'registry_credential' type.
//
//
type RegistryCredential struct {
	id       *string
	href     *string
	link     bool
	username *string
	token    *string
	registry *Registry
	account  *Account
}

// Kind returns the name of the type of the object.
func (o *RegistryCredential) Kind() string {
	if o == nil {
		return RegistryCredentialNilKind
	}
	if o.link {
		return RegistryCredentialLinkKind
	}
	return RegistryCredentialKind
}

// ID returns the identifier of the object.
func (o *RegistryCredential) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *RegistryCredential) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *RegistryCredential) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *RegistryCredential) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *RegistryCredential) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *RegistryCredential) Empty() bool {
	return o == nil || (o.id == nil &&
		o.username == nil &&
		o.token == nil &&
		o.registry == nil &&
		o.account == nil &&
		true)
}

// Username returns the value of the 'username' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *RegistryCredential) Username() string {
	if o != nil && o.username != nil {
		return *o.username
	}
	return ""
}

// GetUsername returns the value of the 'username' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *RegistryCredential) GetUsername() (value string, ok bool) {
	ok = o != nil && o.username != nil
	if ok {
		value = *o.username
	}
	return
}

// Token returns the value of the 'token' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *RegistryCredential) Token() string {
	if o != nil && o.token != nil {
		return *o.token
	}
	return ""
}

// GetToken returns the value of the 'token' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *RegistryCredential) GetToken() (value string, ok bool) {
	ok = o != nil && o.token != nil
	if ok {
		value = *o.token
	}
	return
}

// Registry returns the value of the 'registry' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *RegistryCredential) Registry() *Registry {
	if o == nil {
		return nil
	}
	return o.registry
}

// GetRegistry returns the value of the 'registry' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *RegistryCredential) GetRegistry() (value *Registry, ok bool) {
	ok = o != nil && o.registry != nil
	if ok {
		value = o.registry
	}
	return
}

// Account returns the value of the 'account' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *RegistryCredential) Account() *Account {
	if o == nil {
		return nil
	}
	return o.account
}

// GetAccount returns the value of the 'account' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *RegistryCredential) GetAccount() (value *Account, ok bool) {
	ok = o != nil && o.account != nil
	if ok {
		value = o.account
	}
	return
}

// RegistryCredentialListKind is the name of the type used to represent list of
// objects of type 'registry_credential'.
const RegistryCredentialListKind = "RegistryCredentialList"

// RegistryCredentialListLinkKind is the name of the type used to represent links
// to list of objects of type 'registry_credential'.
const RegistryCredentialListLinkKind = "RegistryCredentialListLink"

// RegistryCredentialNilKind is the name of the type used to nil lists of
// objects of type 'registry_credential'.
const RegistryCredentialListNilKind = "RegistryCredentialListNil"

// RegistryCredentialList is a list of values of the 'registry_credential' type.
type RegistryCredentialList struct {
	href  *string
	link  bool
	items []*RegistryCredential
}

// Kind returns the name of the type of the object.
func (l *RegistryCredentialList) Kind() string {
	if l == nil {
		return RegistryCredentialListNilKind
	}
	if l.link {
		return RegistryCredentialListLinkKind
	}
	return RegistryCredentialListKind
}

// Link returns true iif this is a link.
func (l *RegistryCredentialList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *RegistryCredentialList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *RegistryCredentialList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *RegistryCredentialList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *RegistryCredentialList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *RegistryCredentialList) Get(i int) *RegistryCredential {
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
func (l *RegistryCredentialList) Slice() []*RegistryCredential {
	var slice []*RegistryCredential
	if l == nil {
		slice = make([]*RegistryCredential, 0)
	} else {
		slice = make([]*RegistryCredential, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *RegistryCredentialList) Each(f func(item *RegistryCredential) bool) {
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
func (l *RegistryCredentialList) Range(f func(index int, item *RegistryCredential) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
