# V0040AcctGatherEnergy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageWatts** | Pointer to **int32** | Average power consumption, in watts | [optional] 
**BaseConsumedEnergy** | Pointer to **int64** | The energy consumed between when the node was powered on and the last time it was registered by slurmd, in joules | [optional] 
**ConsumedEnergy** | Pointer to **int64** | The energy consumed between the last time the node was registered by the slurmd daemon and the last node energy accounting sample, in joules | [optional] 
**CurrentWatts** | Pointer to [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | [optional] 
**PreviousConsumedEnergy** | Pointer to **int64** | Previous value of consumed_energy | [optional] 
**LastCollected** | Pointer to **int64** | Time when energy data was last retrieved (UNIX timestamp) | [optional] 

## Methods

### NewV0040AcctGatherEnergy

`func NewV0040AcctGatherEnergy() *V0040AcctGatherEnergy`

NewV0040AcctGatherEnergy instantiates a new V0040AcctGatherEnergy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040AcctGatherEnergyWithDefaults

`func NewV0040AcctGatherEnergyWithDefaults() *V0040AcctGatherEnergy`

NewV0040AcctGatherEnergyWithDefaults instantiates a new V0040AcctGatherEnergy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageWatts

`func (o *V0040AcctGatherEnergy) GetAverageWatts() int32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *V0040AcctGatherEnergy) GetAverageWattsOk() (*int32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *V0040AcctGatherEnergy) SetAverageWatts(v int32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *V0040AcctGatherEnergy) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetBaseConsumedEnergy

`func (o *V0040AcctGatherEnergy) GetBaseConsumedEnergy() int64`

GetBaseConsumedEnergy returns the BaseConsumedEnergy field if non-nil, zero value otherwise.

### GetBaseConsumedEnergyOk

`func (o *V0040AcctGatherEnergy) GetBaseConsumedEnergyOk() (*int64, bool)`

GetBaseConsumedEnergyOk returns a tuple with the BaseConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseConsumedEnergy

`func (o *V0040AcctGatherEnergy) SetBaseConsumedEnergy(v int64)`

SetBaseConsumedEnergy sets BaseConsumedEnergy field to given value.

### HasBaseConsumedEnergy

`func (o *V0040AcctGatherEnergy) HasBaseConsumedEnergy() bool`

HasBaseConsumedEnergy returns a boolean if a field has been set.

### GetConsumedEnergy

`func (o *V0040AcctGatherEnergy) GetConsumedEnergy() int64`

GetConsumedEnergy returns the ConsumedEnergy field if non-nil, zero value otherwise.

### GetConsumedEnergyOk

`func (o *V0040AcctGatherEnergy) GetConsumedEnergyOk() (*int64, bool)`

GetConsumedEnergyOk returns a tuple with the ConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedEnergy

`func (o *V0040AcctGatherEnergy) SetConsumedEnergy(v int64)`

SetConsumedEnergy sets ConsumedEnergy field to given value.

### HasConsumedEnergy

`func (o *V0040AcctGatherEnergy) HasConsumedEnergy() bool`

HasConsumedEnergy returns a boolean if a field has been set.

### GetCurrentWatts

`func (o *V0040AcctGatherEnergy) GetCurrentWatts() V0040Uint32NoVal`

GetCurrentWatts returns the CurrentWatts field if non-nil, zero value otherwise.

### GetCurrentWattsOk

`func (o *V0040AcctGatherEnergy) GetCurrentWattsOk() (*V0040Uint32NoVal, bool)`

GetCurrentWattsOk returns a tuple with the CurrentWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWatts

`func (o *V0040AcctGatherEnergy) SetCurrentWatts(v V0040Uint32NoVal)`

SetCurrentWatts sets CurrentWatts field to given value.

### HasCurrentWatts

`func (o *V0040AcctGatherEnergy) HasCurrentWatts() bool`

HasCurrentWatts returns a boolean if a field has been set.

### GetPreviousConsumedEnergy

`func (o *V0040AcctGatherEnergy) GetPreviousConsumedEnergy() int64`

GetPreviousConsumedEnergy returns the PreviousConsumedEnergy field if non-nil, zero value otherwise.

### GetPreviousConsumedEnergyOk

`func (o *V0040AcctGatherEnergy) GetPreviousConsumedEnergyOk() (*int64, bool)`

GetPreviousConsumedEnergyOk returns a tuple with the PreviousConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousConsumedEnergy

`func (o *V0040AcctGatherEnergy) SetPreviousConsumedEnergy(v int64)`

SetPreviousConsumedEnergy sets PreviousConsumedEnergy field to given value.

### HasPreviousConsumedEnergy

`func (o *V0040AcctGatherEnergy) HasPreviousConsumedEnergy() bool`

HasPreviousConsumedEnergy returns a boolean if a field has been set.

### GetLastCollected

`func (o *V0040AcctGatherEnergy) GetLastCollected() int64`

GetLastCollected returns the LastCollected field if non-nil, zero value otherwise.

### GetLastCollectedOk

`func (o *V0040AcctGatherEnergy) GetLastCollectedOk() (*int64, bool)`

GetLastCollectedOk returns a tuple with the LastCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCollected

`func (o *V0040AcctGatherEnergy) SetLastCollected(v int64)`

SetLastCollected sets LastCollected field to given value.

### HasLastCollected

`func (o *V0040AcctGatherEnergy) HasLastCollected() bool`

HasLastCollected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


