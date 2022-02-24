/*
 * IONOS DBaaS REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ClusterListAllOf struct for ClusterListAllOf
type ClusterListAllOf struct {
	Type *ResourceType `json:"type,omitempty"`
	// The unique ID of the resource.
	Id    *string            `json:"id,omitempty"`
	Items *[]ClusterResponse `json:"items,omitempty"`
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for ResourceType will be returned
func (o *ClusterListAllOf) GetType() *ResourceType {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterListAllOf) GetTypeOk() (*ResourceType, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *ClusterListAllOf) SetType(v ResourceType) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *ClusterListAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClusterListAllOf) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *ClusterListAllOf) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *ClusterListAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []ClusterResponse will be returned
func (o *ClusterListAllOf) GetItems() *[]ClusterResponse {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterListAllOf) GetItemsOk() (*[]ClusterResponse, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *ClusterListAllOf) SetItems(v []ClusterResponse) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *ClusterListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o ClusterListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullableClusterListAllOf struct {
	value *ClusterListAllOf
	isSet bool
}

func (v NullableClusterListAllOf) Get() *ClusterListAllOf {
	return v.value
}

func (v *NullableClusterListAllOf) Set(val *ClusterListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterListAllOf(val *ClusterListAllOf) *NullableClusterListAllOf {
	return &NullableClusterListAllOf{value: val, isSet: true}
}

func (v NullableClusterListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
