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

// checks if the SlurmV0041PostJobSubmitRequestJobsInnerBeginTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041PostJobSubmitRequestJobsInnerBeginTime{}

// SlurmV0041PostJobSubmitRequestJobsInnerBeginTime Defer the allocation of the job until the specified time (UNIX timestamp)
type SlurmV0041PostJobSubmitRequestJobsInnerBeginTime struct {
	// True if number has been set; False if number is unset
	Set *bool `json:"set,omitempty"`
	// True if number has been set to infinite; \"set\" and \"number\" will be ignored
	Infinite *bool `json:"infinite,omitempty"`
	// If \"set\" is True the number will be set with value; otherwise ignore number contents
	Number *int64 `json:"number,omitempty"`
}

// NewSlurmV0041PostJobSubmitRequestJobsInnerBeginTime instantiates a new SlurmV0041PostJobSubmitRequestJobsInnerBeginTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041PostJobSubmitRequestJobsInnerBeginTime() *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime {
	this := SlurmV0041PostJobSubmitRequestJobsInnerBeginTime{}
	return &this
}

// NewSlurmV0041PostJobSubmitRequestJobsInnerBeginTimeWithDefaults instantiates a new SlurmV0041PostJobSubmitRequestJobsInnerBeginTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041PostJobSubmitRequestJobsInnerBeginTimeWithDefaults() *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime {
	this := SlurmV0041PostJobSubmitRequestJobsInnerBeginTime{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) SetSet(v bool) {
	o.Set = &v
}

// GetInfinite returns the Infinite field value if set, zero value otherwise.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) GetInfinite() bool {
	if o == nil || IsNil(o.Infinite) {
		var ret bool
		return ret
	}
	return *o.Infinite
}

// GetInfiniteOk returns a tuple with the Infinite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) GetInfiniteOk() (*bool, bool) {
	if o == nil || IsNil(o.Infinite) {
		return nil, false
	}
	return o.Infinite, true
}

// HasInfinite returns a boolean if a field has been set.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) HasInfinite() bool {
	if o != nil && !IsNil(o.Infinite) {
		return true
	}

	return false
}

// SetInfinite gets a reference to the given bool and assigns it to the Infinite field.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) SetInfinite(v bool) {
	o.Infinite = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) GetNumber() int64 {
	if o == nil || IsNil(o.Number) {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) GetNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) SetNumber(v int64) {
	o.Number = &v
}

func (o SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.Infinite) {
		toSerialize["infinite"] = o.Infinite
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime struct {
	value *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime
	isSet bool
}

func (v NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime) Get() *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime {
	return v.value
}

func (v *NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime) Set(val *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime(val *SlurmV0041PostJobSubmitRequestJobsInnerBeginTime) *NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime {
	return &NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime{value: val, isSet: true}
}

func (v NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041PostJobSubmitRequestJobsInnerBeginTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


