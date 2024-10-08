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

// checks if the SlurmdbV0041GetDiag200ResponseStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041GetDiag200ResponseStatistics{}

// SlurmdbV0041GetDiag200ResponseStatistics statistics
type SlurmdbV0041GetDiag200ResponseStatistics struct {
	// When data collection started (UNIX timestamp)
	TimeStart *int64 `json:"time_start,omitempty"`
	Rollups *SlurmdbV0041GetDiag200ResponseStatisticsRollups `json:"rollups,omitempty"`
	// List of RPCs sent to the slurmdbd
	RPCs []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner `json:"RPCs,omitempty"`
	// List of users that issued RPCs
	Users []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner `json:"users,omitempty"`
}

// NewSlurmdbV0041GetDiag200ResponseStatistics instantiates a new SlurmdbV0041GetDiag200ResponseStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041GetDiag200ResponseStatistics() *SlurmdbV0041GetDiag200ResponseStatistics {
	this := SlurmdbV0041GetDiag200ResponseStatistics{}
	return &this
}

// NewSlurmdbV0041GetDiag200ResponseStatisticsWithDefaults instantiates a new SlurmdbV0041GetDiag200ResponseStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041GetDiag200ResponseStatisticsWithDefaults() *SlurmdbV0041GetDiag200ResponseStatistics {
	this := SlurmdbV0041GetDiag200ResponseStatistics{}
	return &this
}

// GetTimeStart returns the TimeStart field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetTimeStart() int64 {
	if o == nil || IsNil(o.TimeStart) {
		var ret int64
		return ret
	}
	return *o.TimeStart
}

// GetTimeStartOk returns a tuple with the TimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeStart) {
		return nil, false
	}
	return o.TimeStart, true
}

// HasTimeStart returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasTimeStart() bool {
	if o != nil && !IsNil(o.TimeStart) {
		return true
	}

	return false
}

// SetTimeStart gets a reference to the given int64 and assigns it to the TimeStart field.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetTimeStart(v int64) {
	o.TimeStart = &v
}

// GetRollups returns the Rollups field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRollups() SlurmdbV0041GetDiag200ResponseStatisticsRollups {
	if o == nil || IsNil(o.Rollups) {
		var ret SlurmdbV0041GetDiag200ResponseStatisticsRollups
		return ret
	}
	return *o.Rollups
}

// GetRollupsOk returns a tuple with the Rollups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRollupsOk() (*SlurmdbV0041GetDiag200ResponseStatisticsRollups, bool) {
	if o == nil || IsNil(o.Rollups) {
		return nil, false
	}
	return o.Rollups, true
}

// HasRollups returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasRollups() bool {
	if o != nil && !IsNil(o.Rollups) {
		return true
	}

	return false
}

// SetRollups gets a reference to the given SlurmdbV0041GetDiag200ResponseStatisticsRollups and assigns it to the Rollups field.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetRollups(v SlurmdbV0041GetDiag200ResponseStatisticsRollups) {
	o.Rollups = &v
}

// GetRPCs returns the RPCs field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRPCs() []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner {
	if o == nil || IsNil(o.RPCs) {
		var ret []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner
		return ret
	}
	return o.RPCs
}

// GetRPCsOk returns a tuple with the RPCs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetRPCsOk() ([]SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner, bool) {
	if o == nil || IsNil(o.RPCs) {
		return nil, false
	}
	return o.RPCs, true
}

// HasRPCs returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasRPCs() bool {
	if o != nil && !IsNil(o.RPCs) {
		return true
	}

	return false
}

// SetRPCs gets a reference to the given []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner and assigns it to the RPCs field.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetRPCs(v []SlurmdbV0041GetDiag200ResponseStatisticsRPCsInner) {
	o.RPCs = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetUsers() []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner {
	if o == nil || IsNil(o.Users) {
		var ret []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) GetUsersOk() ([]SlurmdbV0041GetDiag200ResponseStatisticsUsersInner, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner and assigns it to the Users field.
func (o *SlurmdbV0041GetDiag200ResponseStatistics) SetUsers(v []SlurmdbV0041GetDiag200ResponseStatisticsUsersInner) {
	o.Users = v
}

func (o SlurmdbV0041GetDiag200ResponseStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041GetDiag200ResponseStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeStart) {
		toSerialize["time_start"] = o.TimeStart
	}
	if !IsNil(o.Rollups) {
		toSerialize["rollups"] = o.Rollups
	}
	if !IsNil(o.RPCs) {
		toSerialize["RPCs"] = o.RPCs
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableSlurmdbV0041GetDiag200ResponseStatistics struct {
	value *SlurmdbV0041GetDiag200ResponseStatistics
	isSet bool
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatistics) Get() *SlurmdbV0041GetDiag200ResponseStatistics {
	return v.value
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatistics) Set(val *SlurmdbV0041GetDiag200ResponseStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041GetDiag200ResponseStatistics(val *SlurmdbV0041GetDiag200ResponseStatistics) *NullableSlurmdbV0041GetDiag200ResponseStatistics {
	return &NullableSlurmdbV0041GetDiag200ResponseStatistics{value: val, isSet: true}
}

func (v NullableSlurmdbV0041GetDiag200ResponseStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041GetDiag200ResponseStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


