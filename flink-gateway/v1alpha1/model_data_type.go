// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
	"fmt"
)

// DataType - struct for DataType
type DataType struct {
	ArrayType *ArrayType
	BigIntType *BigIntType
	BinaryType *BinaryType
	BooleanType *BooleanType
	CharType *CharType
	DateType *DateType
	DecimalType *DecimalType
	DoubleType *DoubleType
	FloatType *FloatType
	IntegerType *IntegerType
	IntervalDayTimeType *IntervalDayTimeType
	IntervalYearMonthType *IntervalYearMonthType
	MapType *MapType
	MultisetType *MultisetType
	RowType *RowType
	SmallIntType *SmallIntType
	TimeWithoutTimeZoneType *TimeWithoutTimeZoneType
	TimestampWithLocalTimeZoneType *TimestampWithLocalTimeZoneType
	TimestampWithTimeZoneType *TimestampWithTimeZoneType
	TimestampWithoutTimeZoneType *TimestampWithoutTimeZoneType
	TinyIntType *TinyIntType
	VarbinaryType *VarbinaryType
	VarcharType *VarcharType
}

// ArrayTypeAsDataType is a convenience function that returns ArrayType wrapped in DataType
func ArrayTypeAsDataType(v *ArrayType) DataType {
	return DataType{ ArrayType: v}
}

// BigIntTypeAsDataType is a convenience function that returns BigIntType wrapped in DataType
func BigIntTypeAsDataType(v *BigIntType) DataType {
	return DataType{ BigIntType: v}
}

// BinaryTypeAsDataType is a convenience function that returns BinaryType wrapped in DataType
func BinaryTypeAsDataType(v *BinaryType) DataType {
	return DataType{ BinaryType: v}
}

// BooleanTypeAsDataType is a convenience function that returns BooleanType wrapped in DataType
func BooleanTypeAsDataType(v *BooleanType) DataType {
	return DataType{ BooleanType: v}
}

// CharTypeAsDataType is a convenience function that returns CharType wrapped in DataType
func CharTypeAsDataType(v *CharType) DataType {
	return DataType{ CharType: v}
}

// DateTypeAsDataType is a convenience function that returns DateType wrapped in DataType
func DateTypeAsDataType(v *DateType) DataType {
	return DataType{ DateType: v}
}

// DecimalTypeAsDataType is a convenience function that returns DecimalType wrapped in DataType
func DecimalTypeAsDataType(v *DecimalType) DataType {
	return DataType{ DecimalType: v}
}

// DoubleTypeAsDataType is a convenience function that returns DoubleType wrapped in DataType
func DoubleTypeAsDataType(v *DoubleType) DataType {
	return DataType{ DoubleType: v}
}

// FloatTypeAsDataType is a convenience function that returns FloatType wrapped in DataType
func FloatTypeAsDataType(v *FloatType) DataType {
	return DataType{ FloatType: v}
}

// IntegerTypeAsDataType is a convenience function that returns IntegerType wrapped in DataType
func IntegerTypeAsDataType(v *IntegerType) DataType {
	return DataType{ IntegerType: v}
}

// IntervalDayTimeTypeAsDataType is a convenience function that returns IntervalDayTimeType wrapped in DataType
func IntervalDayTimeTypeAsDataType(v *IntervalDayTimeType) DataType {
	return DataType{ IntervalDayTimeType: v}
}

// IntervalYearMonthTypeAsDataType is a convenience function that returns IntervalYearMonthType wrapped in DataType
func IntervalYearMonthTypeAsDataType(v *IntervalYearMonthType) DataType {
	return DataType{ IntervalYearMonthType: v}
}

// MapTypeAsDataType is a convenience function that returns MapType wrapped in DataType
func MapTypeAsDataType(v *MapType) DataType {
	return DataType{ MapType: v}
}

// MultisetTypeAsDataType is a convenience function that returns MultisetType wrapped in DataType
func MultisetTypeAsDataType(v *MultisetType) DataType {
	return DataType{ MultisetType: v}
}

// RowTypeAsDataType is a convenience function that returns RowType wrapped in DataType
func RowTypeAsDataType(v *RowType) DataType {
	return DataType{ RowType: v}
}

// SmallIntTypeAsDataType is a convenience function that returns SmallIntType wrapped in DataType
func SmallIntTypeAsDataType(v *SmallIntType) DataType {
	return DataType{ SmallIntType: v}
}

// TimeWithoutTimeZoneTypeAsDataType is a convenience function that returns TimeWithoutTimeZoneType wrapped in DataType
func TimeWithoutTimeZoneTypeAsDataType(v *TimeWithoutTimeZoneType) DataType {
	return DataType{ TimeWithoutTimeZoneType: v}
}

// TimestampWithLocalTimeZoneTypeAsDataType is a convenience function that returns TimestampWithLocalTimeZoneType wrapped in DataType
func TimestampWithLocalTimeZoneTypeAsDataType(v *TimestampWithLocalTimeZoneType) DataType {
	return DataType{ TimestampWithLocalTimeZoneType: v}
}

// TimestampWithTimeZoneTypeAsDataType is a convenience function that returns TimestampWithTimeZoneType wrapped in DataType
func TimestampWithTimeZoneTypeAsDataType(v *TimestampWithTimeZoneType) DataType {
	return DataType{ TimestampWithTimeZoneType: v}
}

// TimestampWithoutTimeZoneTypeAsDataType is a convenience function that returns TimestampWithoutTimeZoneType wrapped in DataType
func TimestampWithoutTimeZoneTypeAsDataType(v *TimestampWithoutTimeZoneType) DataType {
	return DataType{ TimestampWithoutTimeZoneType: v}
}

// TinyIntTypeAsDataType is a convenience function that returns TinyIntType wrapped in DataType
func TinyIntTypeAsDataType(v *TinyIntType) DataType {
	return DataType{ TinyIntType: v}
}

// VarbinaryTypeAsDataType is a convenience function that returns VarbinaryType wrapped in DataType
func VarbinaryTypeAsDataType(v *VarbinaryType) DataType {
	return DataType{ VarbinaryType: v}
}

// VarcharTypeAsDataType is a convenience function that returns VarcharType wrapped in DataType
func VarcharTypeAsDataType(v *VarcharType) DataType {
	return DataType{ VarcharType: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataType) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ARRAY'
	if jsonDict["type"] == "ARRAY" {
		// try to unmarshal JSON data into ArrayType
		err = json.Unmarshal(data, &dst.ArrayType)
		if err == nil {
			return nil // data stored in dst.ArrayType, return on the first match
		} else {
			dst.ArrayType = nil
			return fmt.Errorf("Failed to unmarshal DataType as ArrayType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ArrayType'
	if jsonDict["type"] == "ArrayType" {
		// try to unmarshal JSON data into ArrayType
		err = json.Unmarshal(data, &dst.ArrayType)
		if err == nil {
			return nil // data stored in dst.ArrayType, return on the first match
		} else {
			dst.ArrayType = nil
			return fmt.Errorf("Failed to unmarshal DataType as ArrayType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BIGINT'
	if jsonDict["type"] == "BIGINT" {
		// try to unmarshal JSON data into BigIntType
		err = json.Unmarshal(data, &dst.BigIntType)
		if err == nil {
			return nil // data stored in dst.BigIntType, return on the first match
		} else {
			dst.BigIntType = nil
			return fmt.Errorf("Failed to unmarshal DataType as BigIntType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BINARY'
	if jsonDict["type"] == "BINARY" {
		// try to unmarshal JSON data into BinaryType
		err = json.Unmarshal(data, &dst.BinaryType)
		if err == nil {
			return nil // data stored in dst.BinaryType, return on the first match
		} else {
			dst.BinaryType = nil
			return fmt.Errorf("Failed to unmarshal DataType as BinaryType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BOOLEAN'
	if jsonDict["type"] == "BOOLEAN" {
		// try to unmarshal JSON data into BooleanType
		err = json.Unmarshal(data, &dst.BooleanType)
		if err == nil {
			return nil // data stored in dst.BooleanType, return on the first match
		} else {
			dst.BooleanType = nil
			return fmt.Errorf("Failed to unmarshal DataType as BooleanType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BigIntType'
	if jsonDict["type"] == "BigIntType" {
		// try to unmarshal JSON data into BigIntType
		err = json.Unmarshal(data, &dst.BigIntType)
		if err == nil {
			return nil // data stored in dst.BigIntType, return on the first match
		} else {
			dst.BigIntType = nil
			return fmt.Errorf("Failed to unmarshal DataType as BigIntType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BinaryType'
	if jsonDict["type"] == "BinaryType" {
		// try to unmarshal JSON data into BinaryType
		err = json.Unmarshal(data, &dst.BinaryType)
		if err == nil {
			return nil // data stored in dst.BinaryType, return on the first match
		} else {
			dst.BinaryType = nil
			return fmt.Errorf("Failed to unmarshal DataType as BinaryType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BooleanType'
	if jsonDict["type"] == "BooleanType" {
		// try to unmarshal JSON data into BooleanType
		err = json.Unmarshal(data, &dst.BooleanType)
		if err == nil {
			return nil // data stored in dst.BooleanType, return on the first match
		} else {
			dst.BooleanType = nil
			return fmt.Errorf("Failed to unmarshal DataType as BooleanType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CHAR'
	if jsonDict["type"] == "CHAR" {
		// try to unmarshal JSON data into CharType
		err = json.Unmarshal(data, &dst.CharType)
		if err == nil {
			return nil // data stored in dst.CharType, return on the first match
		} else {
			dst.CharType = nil
			return fmt.Errorf("Failed to unmarshal DataType as CharType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CharType'
	if jsonDict["type"] == "CharType" {
		// try to unmarshal JSON data into CharType
		err = json.Unmarshal(data, &dst.CharType)
		if err == nil {
			return nil // data stored in dst.CharType, return on the first match
		} else {
			dst.CharType = nil
			return fmt.Errorf("Failed to unmarshal DataType as CharType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DATE'
	if jsonDict["type"] == "DATE" {
		// try to unmarshal JSON data into DateType
		err = json.Unmarshal(data, &dst.DateType)
		if err == nil {
			return nil // data stored in dst.DateType, return on the first match
		} else {
			dst.DateType = nil
			return fmt.Errorf("Failed to unmarshal DataType as DateType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DECIMAL'
	if jsonDict["type"] == "DECIMAL" {
		// try to unmarshal JSON data into DecimalType
		err = json.Unmarshal(data, &dst.DecimalType)
		if err == nil {
			return nil // data stored in dst.DecimalType, return on the first match
		} else {
			dst.DecimalType = nil
			return fmt.Errorf("Failed to unmarshal DataType as DecimalType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DOUBLE'
	if jsonDict["type"] == "DOUBLE" {
		// try to unmarshal JSON data into DoubleType
		err = json.Unmarshal(data, &dst.DoubleType)
		if err == nil {
			return nil // data stored in dst.DoubleType, return on the first match
		} else {
			dst.DoubleType = nil
			return fmt.Errorf("Failed to unmarshal DataType as DoubleType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DateType'
	if jsonDict["type"] == "DateType" {
		// try to unmarshal JSON data into DateType
		err = json.Unmarshal(data, &dst.DateType)
		if err == nil {
			return nil // data stored in dst.DateType, return on the first match
		} else {
			dst.DateType = nil
			return fmt.Errorf("Failed to unmarshal DataType as DateType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DecimalType'
	if jsonDict["type"] == "DecimalType" {
		// try to unmarshal JSON data into DecimalType
		err = json.Unmarshal(data, &dst.DecimalType)
		if err == nil {
			return nil // data stored in dst.DecimalType, return on the first match
		} else {
			dst.DecimalType = nil
			return fmt.Errorf("Failed to unmarshal DataType as DecimalType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DoubleType'
	if jsonDict["type"] == "DoubleType" {
		// try to unmarshal JSON data into DoubleType
		err = json.Unmarshal(data, &dst.DoubleType)
		if err == nil {
			return nil // data stored in dst.DoubleType, return on the first match
		} else {
			dst.DoubleType = nil
			return fmt.Errorf("Failed to unmarshal DataType as DoubleType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FLOAT'
	if jsonDict["type"] == "FLOAT" {
		// try to unmarshal JSON data into FloatType
		err = json.Unmarshal(data, &dst.FloatType)
		if err == nil {
			return nil // data stored in dst.FloatType, return on the first match
		} else {
			dst.FloatType = nil
			return fmt.Errorf("Failed to unmarshal DataType as FloatType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FloatType'
	if jsonDict["type"] == "FloatType" {
		// try to unmarshal JSON data into FloatType
		err = json.Unmarshal(data, &dst.FloatType)
		if err == nil {
			return nil // data stored in dst.FloatType, return on the first match
		} else {
			dst.FloatType = nil
			return fmt.Errorf("Failed to unmarshal DataType as FloatType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'INTEGER'
	if jsonDict["type"] == "INTEGER" {
		// try to unmarshal JSON data into IntegerType
		err = json.Unmarshal(data, &dst.IntegerType)
		if err == nil {
			return nil // data stored in dst.IntegerType, return on the first match
		} else {
			dst.IntegerType = nil
			return fmt.Errorf("Failed to unmarshal DataType as IntegerType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'INTERVAL_DAY_TIME'
	if jsonDict["type"] == "INTERVAL_DAY_TIME" {
		// try to unmarshal JSON data into IntervalDayTimeType
		err = json.Unmarshal(data, &dst.IntervalDayTimeType)
		if err == nil {
			return nil // data stored in dst.IntervalDayTimeType, return on the first match
		} else {
			dst.IntervalDayTimeType = nil
			return fmt.Errorf("Failed to unmarshal DataType as IntervalDayTimeType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'INTERVAL_YEAR_MONTH'
	if jsonDict["type"] == "INTERVAL_YEAR_MONTH" {
		// try to unmarshal JSON data into IntervalYearMonthType
		err = json.Unmarshal(data, &dst.IntervalYearMonthType)
		if err == nil {
			return nil // data stored in dst.IntervalYearMonthType, return on the first match
		} else {
			dst.IntervalYearMonthType = nil
			return fmt.Errorf("Failed to unmarshal DataType as IntervalYearMonthType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IntegerType'
	if jsonDict["type"] == "IntegerType" {
		// try to unmarshal JSON data into IntegerType
		err = json.Unmarshal(data, &dst.IntegerType)
		if err == nil {
			return nil // data stored in dst.IntegerType, return on the first match
		} else {
			dst.IntegerType = nil
			return fmt.Errorf("Failed to unmarshal DataType as IntegerType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IntervalDayTimeType'
	if jsonDict["type"] == "IntervalDayTimeType" {
		// try to unmarshal JSON data into IntervalDayTimeType
		err = json.Unmarshal(data, &dst.IntervalDayTimeType)
		if err == nil {
			return nil // data stored in dst.IntervalDayTimeType, return on the first match
		} else {
			dst.IntervalDayTimeType = nil
			return fmt.Errorf("Failed to unmarshal DataType as IntervalDayTimeType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IntervalYearMonthType'
	if jsonDict["type"] == "IntervalYearMonthType" {
		// try to unmarshal JSON data into IntervalYearMonthType
		err = json.Unmarshal(data, &dst.IntervalYearMonthType)
		if err == nil {
			return nil // data stored in dst.IntervalYearMonthType, return on the first match
		} else {
			dst.IntervalYearMonthType = nil
			return fmt.Errorf("Failed to unmarshal DataType as IntervalYearMonthType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MAP'
	if jsonDict["type"] == "MAP" {
		// try to unmarshal JSON data into MapType
		err = json.Unmarshal(data, &dst.MapType)
		if err == nil {
			return nil // data stored in dst.MapType, return on the first match
		} else {
			dst.MapType = nil
			return fmt.Errorf("Failed to unmarshal DataType as MapType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MULTISET'
	if jsonDict["type"] == "MULTISET" {
		// try to unmarshal JSON data into MultisetType
		err = json.Unmarshal(data, &dst.MultisetType)
		if err == nil {
			return nil // data stored in dst.MultisetType, return on the first match
		} else {
			dst.MultisetType = nil
			return fmt.Errorf("Failed to unmarshal DataType as MultisetType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MapType'
	if jsonDict["type"] == "MapType" {
		// try to unmarshal JSON data into MapType
		err = json.Unmarshal(data, &dst.MapType)
		if err == nil {
			return nil // data stored in dst.MapType, return on the first match
		} else {
			dst.MapType = nil
			return fmt.Errorf("Failed to unmarshal DataType as MapType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MultisetType'
	if jsonDict["type"] == "MultisetType" {
		// try to unmarshal JSON data into MultisetType
		err = json.Unmarshal(data, &dst.MultisetType)
		if err == nil {
			return nil // data stored in dst.MultisetType, return on the first match
		} else {
			dst.MultisetType = nil
			return fmt.Errorf("Failed to unmarshal DataType as MultisetType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ROW'
	if jsonDict["type"] == "ROW" {
		// try to unmarshal JSON data into RowType
		err = json.Unmarshal(data, &dst.RowType)
		if err == nil {
			return nil // data stored in dst.RowType, return on the first match
		} else {
			dst.RowType = nil
			return fmt.Errorf("Failed to unmarshal DataType as RowType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RowType'
	if jsonDict["type"] == "RowType" {
		// try to unmarshal JSON data into RowType
		err = json.Unmarshal(data, &dst.RowType)
		if err == nil {
			return nil // data stored in dst.RowType, return on the first match
		} else {
			dst.RowType = nil
			return fmt.Errorf("Failed to unmarshal DataType as RowType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SMALLINT'
	if jsonDict["type"] == "SMALLINT" {
		// try to unmarshal JSON data into SmallIntType
		err = json.Unmarshal(data, &dst.SmallIntType)
		if err == nil {
			return nil // data stored in dst.SmallIntType, return on the first match
		} else {
			dst.SmallIntType = nil
			return fmt.Errorf("Failed to unmarshal DataType as SmallIntType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SmallIntType'
	if jsonDict["type"] == "SmallIntType" {
		// try to unmarshal JSON data into SmallIntType
		err = json.Unmarshal(data, &dst.SmallIntType)
		if err == nil {
			return nil // data stored in dst.SmallIntType, return on the first match
		} else {
			dst.SmallIntType = nil
			return fmt.Errorf("Failed to unmarshal DataType as SmallIntType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TIMESTAMP_WITHOUT_TIME_ZONE'
	if jsonDict["type"] == "TIMESTAMP_WITHOUT_TIME_ZONE" {
		// try to unmarshal JSON data into TimestampWithoutTimeZoneType
		err = json.Unmarshal(data, &dst.TimestampWithoutTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimestampWithoutTimeZoneType, return on the first match
		} else {
			dst.TimestampWithoutTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimestampWithoutTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TIMESTAMP_WITH_LOCAL_TIME_ZONE'
	if jsonDict["type"] == "TIMESTAMP_WITH_LOCAL_TIME_ZONE" {
		// try to unmarshal JSON data into TimestampWithLocalTimeZoneType
		err = json.Unmarshal(data, &dst.TimestampWithLocalTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimestampWithLocalTimeZoneType, return on the first match
		} else {
			dst.TimestampWithLocalTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimestampWithLocalTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TIMESTAMP_WITH_TIME_ZONE'
	if jsonDict["type"] == "TIMESTAMP_WITH_TIME_ZONE" {
		// try to unmarshal JSON data into TimestampWithTimeZoneType
		err = json.Unmarshal(data, &dst.TimestampWithTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimestampWithTimeZoneType, return on the first match
		} else {
			dst.TimestampWithTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimestampWithTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TIME_WITHOUT_TIME_ZONE'
	if jsonDict["type"] == "TIME_WITHOUT_TIME_ZONE" {
		// try to unmarshal JSON data into TimeWithoutTimeZoneType
		err = json.Unmarshal(data, &dst.TimeWithoutTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimeWithoutTimeZoneType, return on the first match
		} else {
			dst.TimeWithoutTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimeWithoutTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TINYINT'
	if jsonDict["type"] == "TINYINT" {
		// try to unmarshal JSON data into TinyIntType
		err = json.Unmarshal(data, &dst.TinyIntType)
		if err == nil {
			return nil // data stored in dst.TinyIntType, return on the first match
		} else {
			dst.TinyIntType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TinyIntType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TimeWithoutTimeZoneType'
	if jsonDict["type"] == "TimeWithoutTimeZoneType" {
		// try to unmarshal JSON data into TimeWithoutTimeZoneType
		err = json.Unmarshal(data, &dst.TimeWithoutTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimeWithoutTimeZoneType, return on the first match
		} else {
			dst.TimeWithoutTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimeWithoutTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TimestampWithLocalTimeZoneType'
	if jsonDict["type"] == "TimestampWithLocalTimeZoneType" {
		// try to unmarshal JSON data into TimestampWithLocalTimeZoneType
		err = json.Unmarshal(data, &dst.TimestampWithLocalTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimestampWithLocalTimeZoneType, return on the first match
		} else {
			dst.TimestampWithLocalTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimestampWithLocalTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TimestampWithTimeZoneType'
	if jsonDict["type"] == "TimestampWithTimeZoneType" {
		// try to unmarshal JSON data into TimestampWithTimeZoneType
		err = json.Unmarshal(data, &dst.TimestampWithTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimestampWithTimeZoneType, return on the first match
		} else {
			dst.TimestampWithTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimestampWithTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TimestampWithoutTimeZoneType'
	if jsonDict["type"] == "TimestampWithoutTimeZoneType" {
		// try to unmarshal JSON data into TimestampWithoutTimeZoneType
		err = json.Unmarshal(data, &dst.TimestampWithoutTimeZoneType)
		if err == nil {
			return nil // data stored in dst.TimestampWithoutTimeZoneType, return on the first match
		} else {
			dst.TimestampWithoutTimeZoneType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TimestampWithoutTimeZoneType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TinyIntType'
	if jsonDict["type"] == "TinyIntType" {
		// try to unmarshal JSON data into TinyIntType
		err = json.Unmarshal(data, &dst.TinyIntType)
		if err == nil {
			return nil // data stored in dst.TinyIntType, return on the first match
		} else {
			dst.TinyIntType = nil
			return fmt.Errorf("Failed to unmarshal DataType as TinyIntType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VARBINARY'
	if jsonDict["type"] == "VARBINARY" {
		// try to unmarshal JSON data into VarbinaryType
		err = json.Unmarshal(data, &dst.VarbinaryType)
		if err == nil {
			return nil // data stored in dst.VarbinaryType, return on the first match
		} else {
			dst.VarbinaryType = nil
			return fmt.Errorf("Failed to unmarshal DataType as VarbinaryType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VARCHAR'
	if jsonDict["type"] == "VARCHAR" {
		// try to unmarshal JSON data into VarcharType
		err = json.Unmarshal(data, &dst.VarcharType)
		if err == nil {
			return nil // data stored in dst.VarcharType, return on the first match
		} else {
			dst.VarcharType = nil
			return fmt.Errorf("Failed to unmarshal DataType as VarcharType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VarbinaryType'
	if jsonDict["type"] == "VarbinaryType" {
		// try to unmarshal JSON data into VarbinaryType
		err = json.Unmarshal(data, &dst.VarbinaryType)
		if err == nil {
			return nil // data stored in dst.VarbinaryType, return on the first match
		} else {
			dst.VarbinaryType = nil
			return fmt.Errorf("Failed to unmarshal DataType as VarbinaryType: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VarcharType'
	if jsonDict["type"] == "VarcharType" {
		// try to unmarshal JSON data into VarcharType
		err = json.Unmarshal(data, &dst.VarcharType)
		if err == nil {
			return nil // data stored in dst.VarcharType, return on the first match
		} else {
			dst.VarcharType = nil
			return fmt.Errorf("Failed to unmarshal DataType as VarcharType: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataType) MarshalJSON() ([]byte, error) {
	if src.ArrayType != nil {
		return json.Marshal(&src.ArrayType)
	}

	if src.BigIntType != nil {
		return json.Marshal(&src.BigIntType)
	}

	if src.BinaryType != nil {
		return json.Marshal(&src.BinaryType)
	}

	if src.BooleanType != nil {
		return json.Marshal(&src.BooleanType)
	}

	if src.CharType != nil {
		return json.Marshal(&src.CharType)
	}

	if src.DateType != nil {
		return json.Marshal(&src.DateType)
	}

	if src.DecimalType != nil {
		return json.Marshal(&src.DecimalType)
	}

	if src.DoubleType != nil {
		return json.Marshal(&src.DoubleType)
	}

	if src.FloatType != nil {
		return json.Marshal(&src.FloatType)
	}

	if src.IntegerType != nil {
		return json.Marshal(&src.IntegerType)
	}

	if src.IntervalDayTimeType != nil {
		return json.Marshal(&src.IntervalDayTimeType)
	}

	if src.IntervalYearMonthType != nil {
		return json.Marshal(&src.IntervalYearMonthType)
	}

	if src.MapType != nil {
		return json.Marshal(&src.MapType)
	}

	if src.MultisetType != nil {
		return json.Marshal(&src.MultisetType)
	}

	if src.RowType != nil {
		return json.Marshal(&src.RowType)
	}

	if src.SmallIntType != nil {
		return json.Marshal(&src.SmallIntType)
	}

	if src.TimeWithoutTimeZoneType != nil {
		return json.Marshal(&src.TimeWithoutTimeZoneType)
	}

	if src.TimestampWithLocalTimeZoneType != nil {
		return json.Marshal(&src.TimestampWithLocalTimeZoneType)
	}

	if src.TimestampWithTimeZoneType != nil {
		return json.Marshal(&src.TimestampWithTimeZoneType)
	}

	if src.TimestampWithoutTimeZoneType != nil {
		return json.Marshal(&src.TimestampWithoutTimeZoneType)
	}

	if src.TinyIntType != nil {
		return json.Marshal(&src.TinyIntType)
	}

	if src.VarbinaryType != nil {
		return json.Marshal(&src.VarbinaryType)
	}

	if src.VarcharType != nil {
		return json.Marshal(&src.VarcharType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataType) GetActualInstance() (interface{}) {
	if obj.ArrayType != nil {
		return obj.ArrayType
	}

	if obj.BigIntType != nil {
		return obj.BigIntType
	}

	if obj.BinaryType != nil {
		return obj.BinaryType
	}

	if obj.BooleanType != nil {
		return obj.BooleanType
	}

	if obj.CharType != nil {
		return obj.CharType
	}

	if obj.DateType != nil {
		return obj.DateType
	}

	if obj.DecimalType != nil {
		return obj.DecimalType
	}

	if obj.DoubleType != nil {
		return obj.DoubleType
	}

	if obj.FloatType != nil {
		return obj.FloatType
	}

	if obj.IntegerType != nil {
		return obj.IntegerType
	}

	if obj.IntervalDayTimeType != nil {
		return obj.IntervalDayTimeType
	}

	if obj.IntervalYearMonthType != nil {
		return obj.IntervalYearMonthType
	}

	if obj.MapType != nil {
		return obj.MapType
	}

	if obj.MultisetType != nil {
		return obj.MultisetType
	}

	if obj.RowType != nil {
		return obj.RowType
	}

	if obj.SmallIntType != nil {
		return obj.SmallIntType
	}

	if obj.TimeWithoutTimeZoneType != nil {
		return obj.TimeWithoutTimeZoneType
	}

	if obj.TimestampWithLocalTimeZoneType != nil {
		return obj.TimestampWithLocalTimeZoneType
	}

	if obj.TimestampWithTimeZoneType != nil {
		return obj.TimestampWithTimeZoneType
	}

	if obj.TimestampWithoutTimeZoneType != nil {
		return obj.TimestampWithoutTimeZoneType
	}

	if obj.TinyIntType != nil {
		return obj.TinyIntType
	}

	if obj.VarbinaryType != nil {
		return obj.VarbinaryType
	}

	if obj.VarcharType != nil {
		return obj.VarcharType
	}

	// all schemas are nil
	return nil
}

type NullableDataType struct {
	value *DataType
	isSet bool
}

func (v NullableDataType) Get() *DataType {
	return v.value
}

func (v *NullableDataType) Set(val *DataType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataType(val *DataType) *NullableDataType {
	return &NullableDataType{value: val, isSet: true}
}

func (v NullableDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

