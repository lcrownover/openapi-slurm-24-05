# SlurmV0041GetLicenses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | [**[]SlurmV0041GetLicenses200ResponseLicensesInner**](SlurmV0041GetLicenses200ResponseLicensesInner.md) | List of licenses | 
**LastUpdate** | [**SlurmV0041GetLicenses200ResponseLastUpdate**](SlurmV0041GetLicenses200ResponseLastUpdate.md) |  | 
**Meta** | Pointer to [**SlurmdbV0041DeleteSingleQos200ResponseMeta**](SlurmdbV0041DeleteSingleQos200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner**](SlurmdbV0041DeleteSingleQos200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner**](SlurmdbV0041DeleteSingleQos200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewSlurmV0041GetLicenses200Response

`func NewSlurmV0041GetLicenses200Response(licenses []SlurmV0041GetLicenses200ResponseLicensesInner, lastUpdate SlurmV0041GetLicenses200ResponseLastUpdate, ) *SlurmV0041GetLicenses200Response`

NewSlurmV0041GetLicenses200Response instantiates a new SlurmV0041GetLicenses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetLicenses200ResponseWithDefaults

`func NewSlurmV0041GetLicenses200ResponseWithDefaults() *SlurmV0041GetLicenses200Response`

NewSlurmV0041GetLicenses200ResponseWithDefaults instantiates a new SlurmV0041GetLicenses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenses

`func (o *SlurmV0041GetLicenses200Response) GetLicenses() []SlurmV0041GetLicenses200ResponseLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *SlurmV0041GetLicenses200Response) GetLicensesOk() (*[]SlurmV0041GetLicenses200ResponseLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *SlurmV0041GetLicenses200Response) SetLicenses(v []SlurmV0041GetLicenses200ResponseLicensesInner)`

SetLicenses sets Licenses field to given value.


### GetLastUpdate

`func (o *SlurmV0041GetLicenses200Response) GetLastUpdate() SlurmV0041GetLicenses200ResponseLastUpdate`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *SlurmV0041GetLicenses200Response) GetLastUpdateOk() (*SlurmV0041GetLicenses200ResponseLastUpdate, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *SlurmV0041GetLicenses200Response) SetLastUpdate(v SlurmV0041GetLicenses200ResponseLastUpdate)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *SlurmV0041GetLicenses200Response) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SlurmV0041GetLicenses200Response) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SlurmV0041GetLicenses200Response) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SlurmV0041GetLicenses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *SlurmV0041GetLicenses200Response) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SlurmV0041GetLicenses200Response) GetErrorsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SlurmV0041GetLicenses200Response) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SlurmV0041GetLicenses200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *SlurmV0041GetLicenses200Response) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SlurmV0041GetLicenses200Response) GetWarningsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SlurmV0041GetLicenses200Response) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SlurmV0041GetLicenses200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


