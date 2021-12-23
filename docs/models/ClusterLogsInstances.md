# ClusterLogsInstances



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the PostgreSQL instance. | [optional] |
|**Messages** | Pointer to [**[]ClusterLogsMessages**](ClusterLogsMessages.md) |  | [optional] |

## Methods


### GetName

`func (o *ClusterLogsInstances) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLogsInstances) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLogsInstances) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLogsInstances) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessages

`func (o *ClusterLogsInstances) GetMessages() []ClusterLogsMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ClusterLogsInstances) GetMessagesOk() (*[]ClusterLogsMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ClusterLogsInstances) SetMessages(v []ClusterLogsMessages)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ClusterLogsInstances) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


