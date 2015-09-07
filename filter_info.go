package gofilter

import (
	"errors"
)

type fieldId int

var idToFieldNameMap map[fieldId]string
var fieldNameToIdMap map[string]fieldId
var idToFieldTypeMap map[fieldId]ftenum

type ftenum int

// Lis of available field types
const (
	_ = iota

	FT_BOOL   ftenum = iota // bool
	FT_STRING ftenum = iota // string
	FT_INT    ftenum = iota // int
	FT_UINT   ftenum = iota // uint

	FT_UINT8  ftenum = iota // uint8
	FT_UINT16 ftenum = iota // uint16
	FT_UINT24 ftenum = iota // uint24
	FT_UINT32 ftenum = iota // uint32
	FT_UINT64 ftenum = iota // uint64
	FT_INT8   ftenum = iota // int8
	FT_INT16  ftenum = iota // int16
	FT_INT24  ftenum = iota // int24
	FT_INT32  ftenum = iota // int32
	FT_INT64  ftenum = iota // int64

	FT_FLOAT32 ftenum = iota // float32
	FT_FLOAT64 ftenum = iota // float64

	FT_BYTES ftenum = iota // []byte

	FT_IP  ftenum = iota // net.IP
	FT_MAC ftenum = iota // net.HardwareAddr

)

var ErrFieldExist = errors.New("Field is already registered")

var id fieldId

// RegisterField adds field with name and f_type to known fields.
// When try to register field with name which was already registered
// return ErrFieldExist.
func RegisterField(name string, f_type ftenum) error {
	if exists := fieldNameToIdMap[name]; exists != 0 {
		return ErrFieldExist
	}
	// field id
	id++
	idToFieldNameMap[id] = name
	fieldNameToIdMap[name] = id
	idToFieldTypeMap[id] = f_type

	return nil
}
func init() {
	idToFieldNameMap = make(map[fieldId]string)
	fieldNameToIdMap = make(map[string]fieldId)
	idToFieldTypeMap = make(map[fieldId]ftenum)
}

func nameToId(name string) fieldId {
	return fieldNameToIdMap[name]
}

func nameToFieldType(name string) ftenum {
	return idToFieldType(nameToId(name))
}

func idToFieldType(id fieldId) ftenum {
	return idToFieldTypeMap[id]
}
