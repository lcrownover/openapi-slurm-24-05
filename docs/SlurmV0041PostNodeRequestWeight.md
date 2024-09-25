# SlurmV0041PostNodeRequestWeight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | True if number has been set; False if number is unset | [optional] 
**Infinite** | Pointer to **bool** | True if number has been set to infinite; \&quot;set\&quot; and \&quot;number\&quot; will be ignored | [optional] 
**Number** | Pointer to **int32** | If \&quot;set\&quot; is True the number will be set with value; otherwise ignore number contents | [optional] 

## Methods

### NewSlurmV0041PostNodeRequestWeight

`func NewSlurmV0041PostNodeRequestWeight() *SlurmV0041PostNodeRequestWeight`

NewSlurmV0041PostNodeRequestWeight instantiates a new SlurmV0041PostNodeRequestWeight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostNodeRequestWeightWithDefaults

`func NewSlurmV0041PostNodeRequestWeightWithDefaults() *SlurmV0041PostNodeRequestWeight`

NewSlurmV0041PostNodeRequestWeightWithDefaults instantiates a new SlurmV0041PostNodeRequestWeight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *SlurmV0041PostNodeRequestWeight) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *SlurmV0041PostNodeRequestWeight) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *SlurmV0041PostNodeRequestWeight) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *SlurmV0041PostNodeRequestWeight) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetInfinite

`func (o *SlurmV0041PostNodeRequestWeight) GetInfinite() bool`

GetInfinite returns the Infinite field if non-nil, zero value otherwise.

### GetInfiniteOk

`func (o *SlurmV0041PostNodeRequestWeight) GetInfiniteOk() (*bool, bool)`

GetInfiniteOk returns a tuple with the Infinite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfinite

`func (o *SlurmV0041PostNodeRequestWeight) SetInfinite(v bool)`

SetInfinite sets Infinite field to given value.

### HasInfinite

`func (o *SlurmV0041PostNodeRequestWeight) HasInfinite() bool`

HasInfinite returns a boolean if a field has been set.

### GetNumber

`func (o *SlurmV0041PostNodeRequestWeight) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SlurmV0041PostNodeRequestWeight) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SlurmV0041PostNodeRequestWeight) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SlurmV0041PostNodeRequestWeight) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


