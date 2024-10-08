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

// checks if the SlurmV0041PostJobSubmitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041PostJobSubmitRequest{}

// SlurmV0041PostJobSubmitRequest struct for SlurmV0041PostJobSubmitRequest
type SlurmV0041PostJobSubmitRequest struct {
	// Deprecated; Populate script field in jobs[0] or job
	// Deprecated
	Script *string `json:"script,omitempty"`
	// HetJob description
	Jobs []SlurmV0041PostJobSubmitRequestJobsInner `json:"jobs,omitempty"`
	Job *SlurmV0041PostJobSubmitRequestJob `json:"job,omitempty"`
}

// NewSlurmV0041PostJobSubmitRequest instantiates a new SlurmV0041PostJobSubmitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041PostJobSubmitRequest() *SlurmV0041PostJobSubmitRequest {
	this := SlurmV0041PostJobSubmitRequest{}
	return &this
}

// NewSlurmV0041PostJobSubmitRequestWithDefaults instantiates a new SlurmV0041PostJobSubmitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041PostJobSubmitRequestWithDefaults() *SlurmV0041PostJobSubmitRequest {
	this := SlurmV0041PostJobSubmitRequest{}
	return &this
}

// GetScript returns the Script field value if set, zero value otherwise.
// Deprecated
func (o *SlurmV0041PostJobSubmitRequest) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SlurmV0041PostJobSubmitRequest) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *SlurmV0041PostJobSubmitRequest) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
// Deprecated
func (o *SlurmV0041PostJobSubmitRequest) SetScript(v string) {
	o.Script = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *SlurmV0041PostJobSubmitRequest) GetJobs() []SlurmV0041PostJobSubmitRequestJobsInner {
	if o == nil || IsNil(o.Jobs) {
		var ret []SlurmV0041PostJobSubmitRequestJobsInner
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobSubmitRequest) GetJobsOk() ([]SlurmV0041PostJobSubmitRequestJobsInner, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *SlurmV0041PostJobSubmitRequest) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []SlurmV0041PostJobSubmitRequestJobsInner and assigns it to the Jobs field.
func (o *SlurmV0041PostJobSubmitRequest) SetJobs(v []SlurmV0041PostJobSubmitRequestJobsInner) {
	o.Jobs = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *SlurmV0041PostJobSubmitRequest) GetJob() SlurmV0041PostJobSubmitRequestJob {
	if o == nil || IsNil(o.Job) {
		var ret SlurmV0041PostJobSubmitRequestJob
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobSubmitRequest) GetJobOk() (*SlurmV0041PostJobSubmitRequestJob, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *SlurmV0041PostJobSubmitRequest) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given SlurmV0041PostJobSubmitRequestJob and assigns it to the Job field.
func (o *SlurmV0041PostJobSubmitRequest) SetJob(v SlurmV0041PostJobSubmitRequestJob) {
	o.Job = &v
}

func (o SlurmV0041PostJobSubmitRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041PostJobSubmitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableSlurmV0041PostJobSubmitRequest struct {
	value *SlurmV0041PostJobSubmitRequest
	isSet bool
}

func (v NullableSlurmV0041PostJobSubmitRequest) Get() *SlurmV0041PostJobSubmitRequest {
	return v.value
}

func (v *NullableSlurmV0041PostJobSubmitRequest) Set(val *SlurmV0041PostJobSubmitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041PostJobSubmitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041PostJobSubmitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041PostJobSubmitRequest(val *SlurmV0041PostJobSubmitRequest) *NullableSlurmV0041PostJobSubmitRequest {
	return &NullableSlurmV0041PostJobSubmitRequest{value: val, isSet: true}
}

func (v NullableSlurmV0041PostJobSubmitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041PostJobSubmitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


