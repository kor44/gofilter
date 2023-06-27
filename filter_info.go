package gofilter

import (
	"errors"
)

type fieldId int
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

type Context struct {
	idToFieldNameMap map[fieldId]string
	fieldNameToIdMap map[string]fieldId
	idToFieldTypeMap map[fieldId]ftenum
}

var ErrFieldExist = errors.New("gofilter: field is already registered")

// RegisterField adds field with name and f_type to known fields.
// When try to register field with name which was already registered
// return ErrFieldExist.
func (ctx *Context) RegisterField(name string, f_type ftenum) error {
	if exists := ctx.fieldNameToIdMap[name]; exists != 0 {
		return ErrFieldExist
	}
	// field id
	id := fieldId(len(ctx.idToFieldNameMap))
	ctx.idToFieldNameMap[id] = name
	ctx.fieldNameToIdMap[name] = id
	ctx.idToFieldTypeMap[id] = f_type

	return nil
}

func CreateContext() Context {
	return Context{
		idToFieldNameMap: make(map[fieldId]string),
		fieldNameToIdMap: make(map[string]fieldId),
		idToFieldTypeMap: make(map[fieldId]ftenum),
	}
}

func (ctx Context) nameToId(name string) fieldId {
	return ctx.fieldNameToIdMap[name]
}

func (ctx Context) nameToFieldType(name string) ftenum {
	return ctx.idToFieldType(ctx.nameToId(name))
}

func (ctx Context) idToFieldType(id fieldId) ftenum {
	return ctx.idToFieldTypeMap[id]
}
