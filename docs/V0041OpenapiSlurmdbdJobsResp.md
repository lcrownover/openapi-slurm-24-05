# V0041OpenapiSlurmdbdJobsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]V0041OpenapiSlurmdbdJobsRespJobsInner**](V0041OpenapiSlurmdbdJobsRespJobsInner.md) | jobs | 
**Meta** | Pointer to [**SlurmdbV0041DeleteSingleQos200ResponseMeta**](SlurmdbV0041DeleteSingleQos200ResponseMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner**](SlurmdbV0041DeleteSingleQos200ResponseErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner**](SlurmdbV0041DeleteSingleQos200ResponseWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsResp

`func NewV0041OpenapiSlurmdbdJobsResp(jobs []V0041OpenapiSlurmdbdJobsRespJobsInner, ) *V0041OpenapiSlurmdbdJobsResp`

NewV0041OpenapiSlurmdbdJobsResp instantiates a new V0041OpenapiSlurmdbdJobsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespWithDefaults() *V0041OpenapiSlurmdbdJobsResp`

NewV0041OpenapiSlurmdbdJobsRespWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *V0041OpenapiSlurmdbdJobsResp) GetJobs() []V0041OpenapiSlurmdbdJobsRespJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0041OpenapiSlurmdbdJobsResp) GetJobsOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0041OpenapiSlurmdbdJobsResp) SetJobs(v []V0041OpenapiSlurmdbdJobsRespJobsInner)`

SetJobs sets Jobs field to given value.


### GetMeta

`func (o *V0041OpenapiSlurmdbdJobsResp) GetMeta() SlurmdbV0041DeleteSingleQos200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiSlurmdbdJobsResp) GetMetaOk() (*SlurmdbV0041DeleteSingleQos200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiSlurmdbdJobsResp) SetMeta(v SlurmdbV0041DeleteSingleQos200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiSlurmdbdJobsResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiSlurmdbdJobsResp) GetErrors() []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiSlurmdbdJobsResp) GetErrorsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiSlurmdbdJobsResp) SetErrors(v []SlurmdbV0041DeleteSingleQos200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiSlurmdbdJobsResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiSlurmdbdJobsResp) GetWarnings() []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiSlurmdbdJobsResp) GetWarningsOk() (*[]SlurmdbV0041DeleteSingleQos200ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiSlurmdbdJobsResp) SetWarnings(v []SlurmdbV0041DeleteSingleQos200ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiSlurmdbdJobsResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


