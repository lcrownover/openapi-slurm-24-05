# SlurmdbV0041DeleteWckey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedWckeys** | **[]string** | deleted wckeys | 
**Meta** | Pointer to [**SlurmdbV0041DeleteSingleQos200ResponseMeta**](SlurmdbV0041DeleteSingleQos200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner**](SlurmdbV0041DeleteSingleQos200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner**](SlurmdbV0041DeleteSingleQos200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmdbV0041DeleteWckey200Response

`func NewSlurmdbV0041DeleteWckey200Response(deletedWckeys []string, ) *SlurmdbV0041DeleteWckey200Response`

NewSlurmdbV0041DeleteWckey200Response instantiates a new SlurmdbV0041DeleteWckey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmdbV0041DeleteWckey200ResponseWithDefaults

`func NewSlurmdbV0041DeleteWckey200ResponseWithDefaults() *SlurmdbV0041DeleteWckey200Response`

NewSlurmdbV0041DeleteWckey200ResponseWithDefaults instantiates a new SlurmdbV0041DeleteWckey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedWckeys

`func (o *SlurmdbV0041DeleteWckey200Response) GetDeletedWckeys() []string`

GetDeletedWckeys returns the DeletedWckeys field if non-nil, zero value otherwise.

### GetDeletedWckeysOk

`func (o *SlurmdbV0041DeleteWckey200Response) GetDeletedWckeysOk() (*[]string, bool)`

GetDeletedWckeysOk returns a tuple with the DeletedWckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedWckeys

`func (o *SlurmdbV0041DeleteWckey200Response) SetDeletedWckeys(v []string)`

SetDeletedWckeys sets DeletedWckeys field to given value.


### GetMeta

`func (o *SlurmdbV0041DeleteWckey200Response) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmdbV0041DeleteWckey200Response) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmdbV0041DeleteWckey200Response) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmdbV0041DeleteWckey200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmdbV0041DeleteWckey200Response) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmdbV0041DeleteWckey200Response) GetErrorsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmdbV0041DeleteWckey200Response) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmdbV0041DeleteWckey200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmdbV0041DeleteWckey200Response) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmdbV0041DeleteWckey200Response) GetWarningsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmdbV0041DeleteWckey200Response) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmdbV0041DeleteWckey200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


