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

// AccountKind is the name of the type used to represent objects
// of type 'account'.
const AccountKind = "Account"

// AccountLinkKind is the name of the type used to represent links
// to objects of type 'account'.
const AccountLinkKind = "AccountLink"

// AccountNilKind is the name of the type used to nil references
// to objects of type 'account'.
const AccountNilKind = "AccountNil"

// Account represents the values of the 'account' type.
//
//
type Account struct {
	id             *string
	href           *string
	link           bool
	name           *string
	username       *string
	email          *string
	firstName      *string
	lastName       *string
	banned         *bool
	banDescription *string
	organization   *Organization
}

// Kind returns the name of the type of the object.
func (o *Account) Kind() string {
	if o == nil {
		return AccountNilKind
	}
	if o.link {
		return AccountLinkKind
	}
	return AccountKind
}

// ID returns the identifier of the object.
func (o *Account) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Account) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Account) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Account) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Account) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Account) Empty() bool {
	return o == nil || (o.id == nil &&
		o.name == nil &&
		o.username == nil &&
		o.email == nil &&
		o.firstName == nil &&
		o.lastName == nil &&
		o.banned == nil &&
		o.banDescription == nil &&
		o.organization == nil &&
		true)
}

// Name returns the value of the 'name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) Name() string {
	if o != nil && o.name != nil {
		return *o.name
	}
	return ""
}

// GetName returns the value of the 'name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetName() (value string, ok bool) {
	ok = o != nil && o.name != nil
	if ok {
		value = *o.name
	}
	return
}

// Username returns the value of the 'username' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) Username() string {
	if o != nil && o.username != nil {
		return *o.username
	}
	return ""
}

// GetUsername returns the value of the 'username' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetUsername() (value string, ok bool) {
	ok = o != nil && o.username != nil
	if ok {
		value = *o.username
	}
	return
}

// Email returns the value of the 'email' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) Email() string {
	if o != nil && o.email != nil {
		return *o.email
	}
	return ""
}

// GetEmail returns the value of the 'email' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetEmail() (value string, ok bool) {
	ok = o != nil && o.email != nil
	if ok {
		value = *o.email
	}
	return
}

// FirstName returns the value of the 'first_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) FirstName() string {
	if o != nil && o.firstName != nil {
		return *o.firstName
	}
	return ""
}

// GetFirstName returns the value of the 'first_name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetFirstName() (value string, ok bool) {
	ok = o != nil && o.firstName != nil
	if ok {
		value = *o.firstName
	}
	return
}

// LastName returns the value of the 'last_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) LastName() string {
	if o != nil && o.lastName != nil {
		return *o.lastName
	}
	return ""
}

// GetLastName returns the value of the 'last_name' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetLastName() (value string, ok bool) {
	ok = o != nil && o.lastName != nil
	if ok {
		value = *o.lastName
	}
	return
}

// Banned returns the value of the 'banned' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) Banned() bool {
	if o != nil && o.banned != nil {
		return *o.banned
	}
	return false
}

// GetBanned returns the value of the 'banned' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetBanned() (value bool, ok bool) {
	ok = o != nil && o.banned != nil
	if ok {
		value = *o.banned
	}
	return
}

// BanDescription returns the value of the 'ban_description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) BanDescription() string {
	if o != nil && o.banDescription != nil {
		return *o.banDescription
	}
	return ""
}

// GetBanDescription returns the value of the 'ban_description' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetBanDescription() (value string, ok bool) {
	ok = o != nil && o.banDescription != nil
	if ok {
		value = *o.banDescription
	}
	return
}

// Organization returns the value of the 'organization' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Account) Organization() *Organization {
	if o == nil {
		return nil
	}
	return o.organization
}

// GetOrganization returns the value of the 'organization' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Account) GetOrganization() (value *Organization, ok bool) {
	ok = o != nil && o.organization != nil
	if ok {
		value = o.organization
	}
	return
}

// AccountListKind is the name of the type used to represent list of
// objects of type 'account'.
const AccountListKind = "AccountList"

// AccountListLinkKind is the name of the type used to represent links
// to list of objects of type 'account'.
const AccountListLinkKind = "AccountListLink"

// AccountNilKind is the name of the type used to nil lists of
// objects of type 'account'.
const AccountListNilKind = "AccountListNil"

// AccountList is a list of values of the 'account' type.
type AccountList struct {
	href  *string
	link  bool
	items []*Account
}

// Kind returns the name of the type of the object.
func (l *AccountList) Kind() string {
	if l == nil {
		return AccountListNilKind
	}
	if l.link {
		return AccountListLinkKind
	}
	return AccountListKind
}

// Link returns true iif this is a link.
func (l *AccountList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *AccountList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *AccountList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *AccountList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *AccountList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AccountList) Get(i int) *Account {
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
func (l *AccountList) Slice() []*Account {
	var slice []*Account
	if l == nil {
		slice = make([]*Account, 0)
	} else {
		slice = make([]*Account, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AccountList) Each(f func(item *Account) bool) {
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
func (l *AccountList) Range(f func(index int, item *Account) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
