# SlurmV0041GetShares200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | [**SlurmV0041GetShares200ResponseShares**](SlurmV0041GetShares200ResponseShares.md) |  | 
**Meta** | Pointer to [**SlurmdbV0041DeleteSingleQos200ResponseMeta**](SlurmdbV0041DeleteSingleQos200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner**](SlurmdbV0041DeleteSingleQos200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner**](SlurmdbV0041DeleteSingleQos200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmV0041GetShares200Response

`func NewSlurmV0041GetShares200Response(shares SlurmV0041GetShares200ResponseShares, ) *SlurmV0041GetShares200Response`

NewSlurmV0041GetShares200Response instantiates a new SlurmV0041GetShares200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetShares200ResponseWithDefaults

`func NewSlurmV0041GetShares200ResponseWithDefaults() *SlurmV0041GetShares200Response`

NewSlurmV0041GetShares200ResponseWithDefaults instantiates a new SlurmV0041GetShares200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *SlurmV0041GetShares200Response) GetShares() SlurmV0041GetShares200ResponseShares`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *SlurmV0041GetShares200Response) GetSharesOk() (*SlurmV0041GetShares200ResponseShares, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *SlurmV0041GetShares200Response) SetShares(v SlurmV0041GetShares200ResponseShares)`

SetShares sets Shares field to given value.


### GetMeta

`func (o *SlurmV0041GetShares200Response) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmV0041GetShares200Response) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmV0041GetShares200Response) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmV0041GetShares200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmV0041GetShares200Response) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmV0041GetShares200Response) GetErrorsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmV0041GetShares200Response) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmV0041GetShares200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmV0041GetShares200Response) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmV0041GetShares200Response) GetWarningsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmV0041GetShares200Response) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmV0041GetShares200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


