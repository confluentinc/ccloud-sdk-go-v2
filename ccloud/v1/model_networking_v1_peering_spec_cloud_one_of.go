/*
 * Network API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1-alpha1
 * Contact: cire-traffic@confluent.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccloud

import (
	"encoding/json"
	"fmt"
)

// NetworkingV1PeeringSpecCloudOneOf - struct for NetworkingV1PeeringSpecCloudOneOf
type NetworkingV1PeeringSpecCloudOneOf struct {
	NetworkingV1AwsPeering *NetworkingV1AwsPeering
	NetworkingV1AzurePeering *NetworkingV1AzurePeering
	NetworkingV1GcpPeering *NetworkingV1GcpPeering
}

// NetworkingV1AwsPeeringAsNetworkingV1PeeringSpecCloudOneOf is a convenience function that returns NetworkingV1AwsPeering wrapped in NetworkingV1PeeringSpecCloudOneOf
func NetworkingV1AwsPeeringAsNetworkingV1PeeringSpecCloudOneOf(v *NetworkingV1AwsPeering) NetworkingV1PeeringSpecCloudOneOf {
	return NetworkingV1PeeringSpecCloudOneOf{ NetworkingV1AwsPeering: v}
}

// NetworkingV1AzurePeeringAsNetworkingV1PeeringSpecCloudOneOf is a convenience function that returns NetworkingV1AzurePeering wrapped in NetworkingV1PeeringSpecCloudOneOf
func NetworkingV1AzurePeeringAsNetworkingV1PeeringSpecCloudOneOf(v *NetworkingV1AzurePeering) NetworkingV1PeeringSpecCloudOneOf {
	return NetworkingV1PeeringSpecCloudOneOf{ NetworkingV1AzurePeering: v}
}

// NetworkingV1GcpPeeringAsNetworkingV1PeeringSpecCloudOneOf is a convenience function that returns NetworkingV1GcpPeering wrapped in NetworkingV1PeeringSpecCloudOneOf
func NetworkingV1GcpPeeringAsNetworkingV1PeeringSpecCloudOneOf(v *NetworkingV1GcpPeering) NetworkingV1PeeringSpecCloudOneOf {
	return NetworkingV1PeeringSpecCloudOneOf{ NetworkingV1GcpPeering: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkingV1PeeringSpecCloudOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'networking.v1.AwsPeering'
	if jsonDict["kind"] == "networking.v1.AwsPeering" {
		// try to unmarshal JSON data into NetworkingV1AwsPeering
		err = json.Unmarshal(data, &dst.NetworkingV1AwsPeering)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsPeering, return on the first match
		} else {
			dst.NetworkingV1AwsPeering = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1PeeringSpecCloudOneOf as NetworkingV1AwsPeering: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AzurePeering'
	if jsonDict["kind"] == "networking.v1.AzurePeering" {
		// try to unmarshal JSON data into NetworkingV1AzurePeering
		err = json.Unmarshal(data, &dst.NetworkingV1AzurePeering)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AzurePeering, return on the first match
		} else {
			dst.NetworkingV1AzurePeering = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1PeeringSpecCloudOneOf as NetworkingV1AzurePeering: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.GcpPeering'
	if jsonDict["kind"] == "networking.v1.GcpPeering" {
		// try to unmarshal JSON data into NetworkingV1GcpPeering
		err = json.Unmarshal(data, &dst.NetworkingV1GcpPeering)
		if err == nil {
			return nil // data stored in dst.NetworkingV1GcpPeering, return on the first match
		} else {
			dst.NetworkingV1GcpPeering = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1PeeringSpecCloudOneOf as NetworkingV1GcpPeering: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkingV1PeeringSpecCloudOneOf) MarshalJSON() ([]byte, error) {
	if src.NetworkingV1AwsPeering != nil {
		return json.Marshal(&src.NetworkingV1AwsPeering)
	}

	if src.NetworkingV1AzurePeering != nil {
		return json.Marshal(&src.NetworkingV1AzurePeering)
	}

	if src.NetworkingV1GcpPeering != nil {
		return json.Marshal(&src.NetworkingV1GcpPeering)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NetworkingV1PeeringSpecCloudOneOf) GetActualInstance() (interface{}) {
	if obj.NetworkingV1AwsPeering != nil {
		return obj.NetworkingV1AwsPeering
	}

	if obj.NetworkingV1AzurePeering != nil {
		return obj.NetworkingV1AzurePeering
	}

	if obj.NetworkingV1GcpPeering != nil {
		return obj.NetworkingV1GcpPeering
	}

	// all schemas are nil
	return nil
}

type NullableNetworkingV1PeeringSpecCloudOneOf struct {
	value *NetworkingV1PeeringSpecCloudOneOf
	isSet bool
}

func (v NullableNetworkingV1PeeringSpecCloudOneOf) Get() *NetworkingV1PeeringSpecCloudOneOf {
	return v.value
}

func (v *NullableNetworkingV1PeeringSpecCloudOneOf) Set(val *NetworkingV1PeeringSpecCloudOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PeeringSpecCloudOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PeeringSpecCloudOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PeeringSpecCloudOneOf(val *NetworkingV1PeeringSpecCloudOneOf) *NullableNetworkingV1PeeringSpecCloudOneOf {
	return &NullableNetworkingV1PeeringSpecCloudOneOf{value: val, isSet: true}
}

func (v NullableNetworkingV1PeeringSpecCloudOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1PeeringSpecCloudOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

