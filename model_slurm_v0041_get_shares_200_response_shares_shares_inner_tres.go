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

// checks if the SlurmV0041GetShares200ResponseSharesSharesInnerTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041GetShares200ResponseSharesSharesInnerTres{}

// SlurmV0041GetShares200ResponseSharesSharesInnerTres struct for SlurmV0041GetShares200ResponseSharesSharesInnerTres
type SlurmV0041GetShares200ResponseSharesSharesInnerTres struct {
	// Currently running tres-secs = grp_used_tres_run_secs
	RunSeconds []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner `json:"run_seconds,omitempty"`
	// TRES-minute limit
	GroupMinutes []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner `json:"group_minutes,omitempty"`
	// Measure of each TRES usage
	Usage []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner `json:"usage,omitempty"`
}

// NewSlurmV0041GetShares200ResponseSharesSharesInnerTres instantiates a new SlurmV0041GetShares200ResponseSharesSharesInnerTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041GetShares200ResponseSharesSharesInnerTres() *SlurmV0041GetShares200ResponseSharesSharesInnerTres {
	this := SlurmV0041GetShares200ResponseSharesSharesInnerTres{}
	return &this
}

// NewSlurmV0041GetShares200ResponseSharesSharesInnerTresWithDefaults instantiates a new SlurmV0041GetShares200ResponseSharesSharesInnerTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041GetShares200ResponseSharesSharesInnerTresWithDefaults() *SlurmV0041GetShares200ResponseSharesSharesInnerTres {
	this := SlurmV0041GetShares200ResponseSharesSharesInnerTres{}
	return &this
}

// GetRunSeconds returns the RunSeconds field value if set, zero value otherwise.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetRunSeconds() []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner {
	if o == nil || IsNil(o.RunSeconds) {
		var ret []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner
		return ret
	}
	return o.RunSeconds
}

// GetRunSecondsOk returns a tuple with the RunSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetRunSecondsOk() ([]SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner, bool) {
	if o == nil || IsNil(o.RunSeconds) {
		return nil, false
	}
	return o.RunSeconds, true
}

// HasRunSeconds returns a boolean if a field has been set.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) HasRunSeconds() bool {
	if o != nil && !IsNil(o.RunSeconds) {
		return true
	}

	return false
}

// SetRunSeconds gets a reference to the given []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner and assigns it to the RunSeconds field.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) SetRunSeconds(v []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner) {
	o.RunSeconds = v
}

// GetGroupMinutes returns the GroupMinutes field value if set, zero value otherwise.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetGroupMinutes() []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner {
	if o == nil || IsNil(o.GroupMinutes) {
		var ret []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner
		return ret
	}
	return o.GroupMinutes
}

// GetGroupMinutesOk returns a tuple with the GroupMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetGroupMinutesOk() ([]SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner, bool) {
	if o == nil || IsNil(o.GroupMinutes) {
		return nil, false
	}
	return o.GroupMinutes, true
}

// HasGroupMinutes returns a boolean if a field has been set.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) HasGroupMinutes() bool {
	if o != nil && !IsNil(o.GroupMinutes) {
		return true
	}

	return false
}

// SetGroupMinutes gets a reference to the given []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner and assigns it to the GroupMinutes field.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) SetGroupMinutes(v []SlurmV0041GetShares200ResponseSharesSharesInnerTresRunSecondsInner) {
	o.GroupMinutes = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetUsage() []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner {
	if o == nil || IsNil(o.Usage) {
		var ret []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) GetUsageOk() ([]SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner and assigns it to the Usage field.
func (o *SlurmV0041GetShares200ResponseSharesSharesInnerTres) SetUsage(v []SlurmV0041GetShares200ResponseSharesSharesInnerTresUsageInner) {
	o.Usage = v
}

func (o SlurmV0041GetShares200ResponseSharesSharesInnerTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041GetShares200ResponseSharesSharesInnerTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RunSeconds) {
		toSerialize["run_seconds"] = o.RunSeconds
	}
	if !IsNil(o.GroupMinutes) {
		toSerialize["group_minutes"] = o.GroupMinutes
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres struct {
	value *SlurmV0041GetShares200ResponseSharesSharesInnerTres
	isSet bool
}

func (v NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres) Get() *SlurmV0041GetShares200ResponseSharesSharesInnerTres {
	return v.value
}

func (v *NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres) Set(val *SlurmV0041GetShares200ResponseSharesSharesInnerTres) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041GetShares200ResponseSharesSharesInnerTres(val *SlurmV0041GetShares200ResponseSharesSharesInnerTres) *NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres {
	return &NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres{value: val, isSet: true}
}

func (v NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041GetShares200ResponseSharesSharesInnerTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


