# BackupMetadata

Metadata of the backup resource.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] |
|**State** | Pointer to [**State**](State.md) |  | [optional] |

## Methods


### GetCreatedDate

`func (o *BackupMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BackupMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BackupMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BackupMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetState

`func (o *BackupMetadata) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupMetadata) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupMetadata) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *BackupMetadata) HasState() bool`

HasState returns a boolean if a field has been set.



