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

// checks if the SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration{}

// SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration struct for SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration
type SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration struct {
	// Total time spent doing last daily rollup (seconds)
	Last *int64 `json:"last,omitempty"`
	// Longest hourly rollup time (seconds)
	Max *int64 `json:"max,omitempty"`
	// Total time spent doing hourly rollups (seconds)
	Time *int64 `json:"time,omitempty"`
}

// NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration() *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration {
	this := SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration{}
	return &this
}

// NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDurationWithDefaults instantiates a new SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDurationWithDefaults() *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration {
	this := SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration{}
	return &this
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) GetLast() int64 {
	if o == nil || IsNil(o.Last) {
		var ret int64
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) GetLastOk() (*int64, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given int64 and assigns it to the Last field.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) SetLast(v int64) {
	o.Last = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) GetMax() int64 {
	if o == nil || IsNil(o.Max) {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) GetMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) SetMax(v int64) {
	o.Max = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) SetTime(v int64) {
	o.Time = &v
}

func (o SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration struct {
	value *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration
	isSet bool
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) Get() *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration {
	return v.value
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) Set(val *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration(val *SlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) *NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration {
	return &NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration{value: val, isSet: true}
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatisticsRollupsHourlyDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


