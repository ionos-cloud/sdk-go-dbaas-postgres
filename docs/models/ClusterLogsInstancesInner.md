# ClusterLogsInstancesInner

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the PostgreSQL instance. | [optional] |
|**Messages** | Pointer to [**[]ClusterLogsInstancesInnerMessagesInner**](ClusterLogsInstancesInnerMessagesInner.md) |  | [optional] |

## Methods

### NewClusterLogsInstancesInner

`func NewClusterLogsInstancesInner() *ClusterLogsInstancesInner`

NewClusterLogsInstancesInner instantiates a new ClusterLogsInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsInstancesInnerWithDefaults

`func NewClusterLogsInstancesInnerWithDefaults() *ClusterLogsInstancesInner`

NewClusterLogsInstancesInnerWithDefaults instantiates a new ClusterLogsInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterLogsInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLogsInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLogsInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLogsInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessages

`func (o *ClusterLogsInstancesInner) GetMessages() []ClusterLogsInstancesInnerMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ClusterLogsInstancesInner) GetMessagesOk() (*[]ClusterLogsInstancesInnerMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ClusterLogsInstancesInner) SetMessages(v []ClusterLogsInstancesInnerMessagesInner)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ClusterLogsInstancesInner) HasMessages() bool`

HasMessages returns a boolean if a field has been set.



