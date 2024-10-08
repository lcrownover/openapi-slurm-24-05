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
	"bytes"
	"fmt"
)

// checks if the SlurmV0041GetLicenses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlurmV0041GetLicenses200Response{}

// SlurmV0041GetLicenses200Response struct for SlurmV0041GetLicenses200Response
type SlurmV0041GetLicenses200Response struct {
	// List of licenses
	Licenses []SlurmV0041GetLicenses200ResponseLicensesInner `json:"licenses"`
	LastUpdate SlurmV0041GetLicenses200ResponseLastUpdate `json:"last_update"`
	Meta *SlurmdbV0041DeleteSingleQos200ResponseMeta `json:"meta,omitempty"`
	// Query errors
	Errors []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _SlurmV0041GetLicenses200Response SlurmV0041GetLicenses200Response

// NewSlurmV0041GetLicenses200Response instantiates a new SlurmV0041GetLicenses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlurmV0041GetLicenses200Response(licenses []SlurmV0041GetLicenses200ResponseLicensesInner, lastUpdate SlurmV0041GetLicenses200ResponseLastUpdate) *SlurmV0041GetLicenses200Response {
	this := SlurmV0041GetLicenses200Response{}
	this.Licenses = licenses
	this.LastUpdate = lastUpdate
	return &this
}

// NewSlurmV0041GetLicenses200ResponseWithDefaults instantiates a new SlurmV0041GetLicenses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlurmV0041GetLicenses200ResponseWithDefaults() *SlurmV0041GetLicenses200Response {
	this := SlurmV0041GetLicenses200Response{}
	return &this
}

// GetLicenses returns the Licenses field value
func (o *SlurmV0041GetLicenses200Response) GetLicenses() []SlurmV0041GetLicenses200ResponseLicensesInner {
	if o == nil {
		var ret []SlurmV0041GetLicenses200ResponseLicensesInner
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200Response) GetLicensesOk() ([]SlurmV0041GetLicenses200ResponseLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *SlurmV0041GetLicenses200Response) SetLicenses(v []SlurmV0041GetLicenses200ResponseLicensesInner) {
	o.Licenses = v
}

// GetLastUpdate returns the LastUpdate field value
func (o *SlurmV0041GetLicenses200Response) GetLastUpdate() SlurmV0041GetLicenses200ResponseLastUpdate {
	if o == nil {
		var ret SlurmV0041GetLicenses200ResponseLastUpdate
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200Response) GetLastUpdateOk() (*SlurmV0041GetLicenses200ResponseLastUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *SlurmV0041GetLicenses200Response) SetLastUpdate(v SlurmV0041GetLicenses200ResponseLastUpdate) {
	o.LastUpdate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SlurmV0041GetLicenses200Response) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret SlurmdbV0041DeleteSingleQos200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200Response) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SlurmV0041GetLicenses200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SlurmdbV0041DeleteSingleQos200ResponseMeta and assigns it to the Meta field.
func (o *SlurmV0041GetLicenses200Response) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SlurmV0041GetLicenses200Response) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200Response) GetErrorsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SlurmV0041GetLicenses200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner and assigns it to the Errors field.
func (o *SlurmV0041GetLicenses200Response) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SlurmV0041GetLicenses200Response) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlurmV0041GetLicenses200Response) GetWarningsOk() ([]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SlurmV0041GetLicenses200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner and assigns it to the Warnings field.
func (o *SlurmV0041GetLicenses200Response) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner) {
	o.Warnings = v
}

func (o SlurmV0041GetLicenses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlurmV0041GetLicenses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["licenses"] = o.Licenses
	toSerialize["last_update"] = o.LastUpdate
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

func (o *SlurmV0041GetLicenses200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"licenses",
		"last_update",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSlurmV0041GetLicenses200Response := _SlurmV0041GetLicenses200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSlurmV0041GetLicenses200Response)

	if err != nil {
		return err
	}

	*o = SlurmV0041GetLicenses200Response(varSlurmV0041GetLicenses200Response)

	return err
}

type NullableSlurmV0041GetLicenses200Response struct {
	value *SlurmV0041GetLicenses200Response
	isSet bool
}

func (v NullableSlurmV0041GetLicenses200Response) Get() *SlurmV0041GetLicenses200Response {
	return v.value
}

func (v *NullableSlurmV0041GetLicenses200Response) Set(val *SlurmV0041GetLicenses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSlurmV0041GetLicenses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSlurmV0041GetLicenses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlurmV0041GetLicenses200Response(val *SlurmV0041GetLicenses200Response) *NullableSlurmV0041GetLicenses200Response {
	return &NullableSlurmV0041GetLicenses200Response{value: val, isSet: true}
}

func (v NullableSlurmV0041GetLicenses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlurmV0041GetLicenses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


