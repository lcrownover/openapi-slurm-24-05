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

// checks if the SlurmV0041PostJobAllocate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041PostJobAllocate200Response{}

// SlurmV0041PostJobAllocate200Response struct for SlurmV0041PostJobAllocate200Response
type SlurmV0041PostJobAllocate200Response struct {
	// Submitted Job ID
	JobId *int32 `json:"job_id,omitempty"`
	// Job submission user message
	JobSubmitUserMsg *string `json:"job_submit_user_msg,omitempty"`
	Meta *SlurmdbV0041DeleteSingleQos200ResponseMeta `json:"meta,omitempty"`
	// Query errors
	Errors []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner `json:"warnings,omitempty"`
}

// NewSlurmV0041PostJobAllocate200Response instantiates a new SlurmV0041PostJobAllocate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041PostJobAllocate200Response() *SlurmV0041PostJobAllocate200Response {
	this := SlurmV0041PostJobAllocate200Response{}
	return &this
}

// NewSlurmV0041PostJobAllocate200ResponseWithDefaults instantiates a new SlurmV0041PostJobAllocate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041PostJobAllocate200ResponseWithDefaults() *SlurmV0041PostJobAllocate200Response {
	this := SlurmV0041PostJobAllocate200Response{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *SlurmV0041PostJobAllocate200Response) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobAllocate200Response) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *SlurmV0041PostJobAllocate200Response) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *SlurmV0041PostJobAllocate200Response) SetJobId(v int32) {
	o.JobId = &v
}

// GetJobSubmitUserMsg returns the JobSubmitUserMsg field value if set, zero value otherwise.
func (o *SlurmV0041PostJobAllocate200Response) GetJobSubmitUserMsg() string {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		var ret string
		return ret
	}
	return *o.JobSubmitUserMsg
}

// GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobAllocate200Response) GetJobSubmitUserMsgOk() (*string, bool) {
	if o == nil || IsNil(o.JobSubmitUserMsg) {
		return nil, false
	}
	return o.JobSubmitUserMsg, true
}

// HasJobSubmitUserMsg returns a boolean if a field has been set.
func (o *SlurmV0041PostJobAllocate200Response) HasJobSubmitUserMsg() bool {
	if o != nil && !IsNil(o.JobSubmitUserMsg) {
		return true
	}

	return false
}

// SetJobSubmitUserMsg gets a reference to the given string and assigns it to the JobSubmitUserMsg field.
func (o *SlurmV0041PostJobAllocate200Response) SetJobSubmitUserMsg(v string) {
	o.JobSubmitUserMsg = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SlurmV0041PostJobAllocate200Response) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobAllocate200Response) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SlurmV0041PostJobAllocate200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMeta and assigns it to the Meta field.
func (o *SlurmV0041PostJobAllocate200Response) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SlurmV0041PostJobAllocate200Response) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobAllocate200Response) GetErrorsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SlurmV0041PostJobAllocate200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner and assigns it to the Errors field.
func (o *SlurmV0041PostJobAllocate200Response) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SlurmV0041PostJobAllocate200Response) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041PostJobAllocate200Response) GetWarningsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SlurmV0041PostJobAllocate200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner and assigns it to the Warnings field.
func (o *SlurmV0041PostJobAllocate200Response) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	o.Warnings = v
}

func (o SlurmV0041PostJobAllocate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041PostJobAllocate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobSubmitUserMsg) {
		toSerialize["job_submit_user_msg"] = o.JobSubmitUserMsg
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableSlurmV0041PostJobAllocate200Response struct {
	value *SlurmV0041PostJobAllocate200Response
	isSet bool
}

func (v NullableSlurmV0041PostJobAllocate200Response) Get() *SlurmV0041PostJobAllocate200Response {
	return v.value
}

func (v *NullableSlurmV0041PostJobAllocate200Response) Set(val *SlurmV0041PostJobAllocate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041PostJobAllocate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041PostJobAllocate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041PostJobAllocate200Response(val *SlurmV0041PostJobAllocate200Response) *NullableSlurmV0041PostJobAllocate200Response {
	return &NullableSlurmV0041PostJobAllocate200Response{value: val, isSet: true}
}

func (v NullableSlurmV0041PostJobAllocate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041PostJobAllocate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


