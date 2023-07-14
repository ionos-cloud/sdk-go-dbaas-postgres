# ClusterLogsInstancesInnerMessagesInner

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Time** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**Message** | Pointer to **string** |  | [optional] |

## Methods

### NewClusterLogsInstancesInnerMessagesInner

`func NewClusterLogsInstancesInnerMessagesInner() *ClusterLogsInstancesInnerMessagesInner`

NewClusterLogsInstancesInnerMessagesInner instantiates a new ClusterLogsInstancesInnerMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsInstancesInnerMessagesInnerWithDefaults

`func NewClusterLogsInstancesInnerMessagesInnerWithDefaults() *ClusterLogsInstancesInnerMessagesInner`

NewClusterLogsInstancesInnerMessagesInnerWithDefaults instantiates a new ClusterLogsInstancesInnerMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ClusterLogsInstancesInnerMessagesInner) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClusterLogsInstancesInnerMessagesInner) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClusterLogsInstancesInnerMessagesInner) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClusterLogsInstancesInnerMessagesInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterLogsInstancesInnerMessagesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterLogsInstancesInnerMessagesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterLogsInstancesInnerMessagesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterLogsInstancesInnerMessagesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.



