# V0041OpenapiAssocsRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemovedAssociations** | **[]string** | removed_associations | 
**Meta** | Pointer to [**SlurmdbV0041DeleteSingleQos200ResponseMeta**](SlurmdbV0041DeleteSingleQos200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner**](SlurmdbV0041DeleteSingleQos200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner**](SlurmdbV0041DeleteSingleQos200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiAssocsRemovedResp

`func NewV0041OpenapiAssocsRemovedResp(removedAssociations []string, ) *V0041OpenapiAssocsRemovedResp`

NewV0041OpenapiAssocsRemovedResp instantiates a new V0041OpenapiAssocsRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiAssocsRemovedRespWithDefaults

`func NewV0041OpenapiAssocsRemovedRespWithDefaults() *V0041OpenapiAssocsRemovedResp`

NewV0041OpenapiAssocsRemovedRespWithDefaults instantiates a new V0041OpenapiAssocsRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemovedAssociations

`func (o *V0041OpenapiAssocsRemovedResp) GetRemovedAssociations() []string`

GetRemovedAssociations returns the RemovedAssociations field if non-nil, zero value otherwise.

### GetRemovedAssociationsOk

`func (o *V0041OpenapiAssocsRemovedResp) GetRemovedAssociationsOk() (*[]string, bool)`

GetRemovedAssociationsOk returns a tuple with the RemovedAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedAssociations

`func (o *V0041OpenapiAssocsRemovedResp) SetRemovedAssociations(v []string)`

SetRemovedAssociations sets RemovedAssociations field to given value.


### GetMeta

`func (o *V0041OpenapiAssocsRemovedResp) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiAssocsRemovedResp) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiAssocsRemovedResp) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiAssocsRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiAssocsRemovedResp) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiAssocsRemovedResp) GetErrorsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiAssocsRemovedResp) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiAssocsRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiAssocsRemovedResp) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiAssocsRemovedResp) GetWarningsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiAssocsRemovedResp) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiAssocsRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


