/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
	"fmt"
)

// StorageType The storage type used in your cluster. (Value \"SSD\" is deprecated. Use the equivalent \"SSD Premium\" instead)
type StorageType string

// List of StorageType
const (
	STORAGETYPE_HDD          StorageType = "HDD"
	STORAGETYPE_SSD          StorageType = "SSD"
	STORAGETYPE_SSD_STANDARD StorageType = "SSD Standard"
	STORAGETYPE_SSD_PREMIUM  StorageType = "SSD Premium"
)

func (v *StorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageType(value)
	for _, existing := range []StorageType{"HDD", "SSD", "SSD Standard", "SSD Premium"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageType", value)
}

// Ptr returns reference to StorageType value
func (v StorageType) Ptr() *StorageType {
	return &v
}

type NullableStorageType struct {
	value *StorageType
	isSet bool
}

func (v NullableStorageType) Get() *StorageType {
	return v.value
}

func (v *NullableStorageType) Set(val *StorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageType(val *StorageType) *NullableStorageType {
	return &NullableStorageType{value: val, isSet: true}
}

func (v NullableStorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
