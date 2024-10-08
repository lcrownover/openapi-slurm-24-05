/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.3&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_slurm_24_05

import (
	"encoding/json"
)

// checks if the V0041OpenapiPartitionRespPartitionsInnerGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiPartitionRespPartitionsInnerGroups{}

// V0041OpenapiPartitionRespPartitionsInnerGroups struct for V0041OpenapiPartitionRespPartitionsInnerGroups
type V0041OpenapiPartitionRespPartitionsInnerGroups struct {
	// AllowGroups
	Allowed *string `json:"allowed,omitempty"`
}

// NewV0041OpenapiPartitionRespPartitionsInnerGroups instantiates a new V0041OpenapiPartitionRespPartitionsInnerGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiPartitionRespPartitionsInnerGroups() *V0041OpenapiPartitionRespPartitionsInnerGroups {
	this := V0041OpenapiPartitionRespPartitionsInnerGroups{}
	return &this
}

// NewV0041OpenapiPartitionRespPartitionsInnerGroupsWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiPartitionRespPartitionsInnerGroupsWithDefaults() *V0041OpenapiPartitionRespPartitionsInnerGroups {
	this := V0041OpenapiPartitionRespPartitionsInnerGroups{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerGroups) GetAllowed() string {
	if o == nil || IsNil(o.Allowed) {
		var ret string
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerGroups) GetAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerGroups) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given string and assigns it to the Allowed field.
func (o *V0041OpenapiPartitionRespPartitionsInnerGroups) SetAllowed(v string) {
	o.Allowed = &v
}

func (o V0041OpenapiPartitionRespPartitionsInnerGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiPartitionRespPartitionsInnerGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableV0041OpenapiPartitionRespPartitionsInnerGroups struct {
	value *V0041OpenapiPartitionRespPartitionsInnerGroups
	isSet bool
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerGroups) Get() *V0041OpenapiPartitionRespPartitionsInnerGroups {
	return v.value
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerGroups) Set(val *V0041OpenapiPartitionRespPartitionsInnerGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiPartitionRespPartitionsInnerGroups(val *V0041OpenapiPartitionRespPartitionsInnerGroups) *NullableV0041OpenapiPartitionRespPartitionsInnerGroups {
	return &NullableV0041OpenapiPartitionRespPartitionsInnerGroups{value: val, isSet: true}
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


