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

// checks if the SlurmdbV0041PostAccountsAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmdbV0041PostAccountsAssociationRequest{}

// SlurmdbV0041PostAccountsAssociationRequest struct for SlurmdbV0041PostAccountsAssociationRequest
type SlurmdbV0041PostAccountsAssociationRequest struct {
	AssociationCondition *SlurmdbV0041PostAccountsAssociationRequestAssociationCondition `json:"association_condition,omitempty"`
	Account *SlurmdbV0041PostAccountsAssociationRequestAccount `json:"account,omitempty"`
	Meta *SlurmdbV0041DeleteSingleQos200ResponseMeta `json:"meta,omitempty"`
	// Query errors
	Errors []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner `json:"warnings,omitempty"`
}

// NewSlurmdbV0041PostAccountsAssociationRequest instantiates a new SlurmdbV0041PostAccountsAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmdbV0041PostAccountsAssociationRequest() *SlurmdbV0041PostAccountsAssociationRequest {
	this := SlurmdbV0041PostAccountsAssociationRequest{}
	return &this
}

// NewSlurmdbV0041PostAccountsAssociationRequestWithDefaults instantiates a new SlurmdbV0041PostAccountsAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmdbV0041PostAccountsAssociationRequestWithDefaults() *SlurmdbV0041PostAccountsAssociationRequest {
	this := SlurmdbV0041PostAccountsAssociationRequest{}
	return &this
}

// GetAssociationCondition returns the AssociationCondition field value if set, zero value otherwise.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAssociationCondition() SlurmdbV0041PostAccountsAssociationRequestAssociationCondition {
	if o == nil || IsNil(o.AssociationCondition) {
		var ret SlurmdbV0041PostAccountsAssociationRequestAssociationCondition
		return ret
	}
	return *o.AssociationCondition
}

// GetAssociationConditionOk returns a tuple with the AssociationCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAssociationConditionOk() (*SlurmdbV0041PostAccountsAssociationRequestAssociationCondition, bool) {
	if o == nil || IsNil(o.AssociationCondition) {
		return nil, false
	}
	return o.AssociationCondition, true
}

// HasAssociationCondition returns a boolean if a field has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) HasAssociationCondition() bool {
	if o != nil && !IsNil(o.AssociationCondition) {
		return true
	}

	return false
}

// SetAssociationCondition gets a reference to the given SlurmdbV0041PostAccountsAssociationRequestAssociationCondition and assigns it to the AssociationCondition field.
func (o *SlurmdbV0041PostAccountsAssociationRequest) SetAssociationCondition(v SlurmdbV0041PostAccountsAssociationRequestAssociationCondition) {
	o.AssociationCondition = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAccount() SlurmdbV0041PostAccountsAssociationRequestAccount {
	if o == nil || IsNil(o.Account) {
		var ret SlurmdbV0041PostAccountsAssociationRequestAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetAccountOk() (*SlurmdbV0041PostAccountsAssociationRequestAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SlurmdbV0041PostAccountsAssociationRequestAccount and assigns it to the Account field.
func (o *SlurmdbV0041PostAccountsAssociationRequest) SetAccount(v SlurmdbV0041PostAccountsAssociationRequestAccount) {
	o.Account = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMeta and assigns it to the Meta field.
func (o *SlurmdbV0041PostAccountsAssociationRequest) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetErrorsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner and assigns it to the Errors field.
func (o *SlurmdbV0041PostAccountsAssociationRequest) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) GetWarningsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SlurmdbV0041PostAccountsAssociationRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner and assigns it to the Warnings field.
func (o *SlurmdbV0041PostAccountsAssociationRequest) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	o.Warnings = v
}

func (o SlurmdbV0041PostAccountsAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmdbV0041PostAccountsAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociationCondition) {
		toSerialize["association_condition"] = o.AssociationCondition
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
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

type NullableSlurmdbV0041PostAccountsAssociationRequest struct {
	value *SlurmdbV0041PostAccountsAssociationRequest
	isSet bool
}

func (v NullableSlurmdbV0041PostAccountsAssociationRequest) Get() *SlurmdbV0041PostAccountsAssociationRequest {
	return v.value
}

func (v *NullableSlurmdbV0041PostAccountsAssociationRequest) Set(val *SlurmdbV0041PostAccountsAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmdbV0041PostAccountsAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmdbV0041PostAccountsAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmdbV0041PostAccountsAssociationRequest(val *SlurmdbV0041PostAccountsAssociationRequest) *NullableSlurmdbV0041PostAccountsAssociationRequest {
	return &NullableSlurmdbV0041PostAccountsAssociationRequest{value: val, isSet: true}
}

func (v NullableSlurmdbV0041PostAccountsAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmdbV0041PostAccountsAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


