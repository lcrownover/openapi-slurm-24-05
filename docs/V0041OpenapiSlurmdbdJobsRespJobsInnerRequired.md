# V0041OpenapiSlurmdbdJobsRespJobsInnerRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CPUs** | Pointer to **int32** | Minimum number of CPUs required | [optional] 
**MemoryPerCpu** | Pointer to [**SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu**](SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerNode**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerNode.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerRequired

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerRequired() *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerRequired instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerRequiredWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerRequiredWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerRequiredWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCPUs

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) GetCPUs() int32`

GetCPUs returns the CPUs field if non-nil, zero value otherwise.

### GetCPUsOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) GetCPUsOk() (*int32, bool)`

GetCPUsOk returns a tuple with the CPUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPUs

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) SetCPUs(v int32)`

SetCPUs sets CPUs field to given value.

### HasCPUs

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) HasCPUs() bool`

HasCPUs returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) GetMemoryPerCpu() SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) GetMemoryPerCpuOk() (*SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) SetMemoryPerCpu(v SlurmV0041PostJobSubmitRequestJobsInnerMemoryPerCpu)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) GetMemoryPerNode() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerNode`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) GetMemoryPerNodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerNode, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) SetMemoryPerNode(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerNode)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerRequired) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


