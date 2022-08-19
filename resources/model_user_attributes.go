/*
Boost API

This is a REST API of the Boost App.

API version: 0.0.1-beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
)

// UserAttributes struct for UserAttributes
type UserAttributes struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Username  *string    `json:"username,omitempty"`
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	Status    UserStatus `json:"status"`
	Role      []UserRole `json:"role"`
}

// NewUserAttributes instantiates a new UserAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAttributes(createdAt time.Time, status UserStatus, role []UserRole) *UserAttributes {
	this := UserAttributes{}
	this.CreatedAt = createdAt
	this.Status = status
	this.Role = role
	return &this
}

// NewUserAttributesWithDefaults instantiates a new UserAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAttributesWithDefaults() *UserAttributes {
	this := UserAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserAttributes) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserAttributes) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserAttributes) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserAttributes) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserAttributes) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserAttributes) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserAttributes) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserAttributes) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserAttributes) SetLastName(v string) {
	o.LastName = &v
}

// GetStatus returns the Status field value
func (o *UserAttributes) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserAttributes) SetStatus(v UserStatus) {
	o.Status = v
}

// GetRole returns the Role field value
func (o *UserAttributes) GetRole() []UserRole {
	if o == nil {
		var ret []UserRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetRoleOk() ([]UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role, true
}

// SetRole sets field value
func (o *UserAttributes) SetRole(v []UserRole) {
	o.Role = v
}

func (o UserAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableUserAttributes struct {
	value *UserAttributes
	isSet bool
}

func (v NullableUserAttributes) Get() *UserAttributes {
	return v.value
}

func (v *NullableUserAttributes) Set(val *UserAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAttributes(val *UserAttributes) *NullableUserAttributes {
	return &NullableUserAttributes{value: val, isSet: true}
}

func (v NullableUserAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}