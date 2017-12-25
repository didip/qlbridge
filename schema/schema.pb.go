// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	schema.proto

It has these top-level messages:
	Command
	TablePartition
	Partition
	TablePb
	FieldPb
	Index
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Command_Operation int32

const (
	Command_Unknown   Command_Operation = 0
	Command_AddUpdate Command_Operation = 1
	Command_Drop      Command_Operation = 2
)

var Command_Operation_name = map[int32]string{
	0: "Unknown",
	1: "AddUpdate",
	2: "Drop",
}
var Command_Operation_value = map[string]int32{
	"Unknown":   0,
	"AddUpdate": 1,
	"Drop":      2,
}

func (x Command_Operation) String() string {
	return proto.EnumName(Command_Operation_name, int32(x))
}
func (Command_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Command defines a Schema Replication command/message such as
// Drop, Alter, Create-Schema, etc.  Used to replicate schema changes
// across servers.
type Command struct {
	Op     Command_Operation `protobuf:"varint,1,opt,name=op,enum=schema.Command_Operation" json:"op,omitempty"`
	Origin string            `protobuf:"bytes,2,opt,name=origin" json:"origin,omitempty"`
	Schema string            `protobuf:"bytes,3,opt,name=schema" json:"schema,omitempty"`
	Index  uint64            `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Ts     int64             `protobuf:"varint,5,opt,name=ts" json:"ts,omitempty"`
	Msg    []byte            `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetOp() Command_Operation {
	if m != nil {
		return m.Op
	}
	return Command_Unknown
}

func (m *Command) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Command) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Command) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Command) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *Command) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// Partition describes a range of data (in a Table).
// left-key is contained in this partition
// right key is not contained in this partition, in the next partition.
// So any value >= left-key, and < right-key is contained herein.
type TablePartition struct {
	Table      string       `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Keys       []string     `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
	Partitions []*Partition `protobuf:"bytes,3,rep,name=partitions" json:"partitions,omitempty"`
}

func (m *TablePartition) Reset()                    { *m = TablePartition{} }
func (m *TablePartition) String() string            { return proto.CompactTextString(m) }
func (*TablePartition) ProtoMessage()               {}
func (*TablePartition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TablePartition) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *TablePartition) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TablePartition) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Partition describes a range of data
// the left-key is contained in this partition
// the right key is not contained in this partition, in the next one
type Partition struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Left  string `protobuf:"bytes,2,opt,name=left" json:"left,omitempty"`
	Right string `protobuf:"bytes,3,opt,name=right" json:"right,omitempty"`
}

func (m *Partition) Reset()                    { *m = Partition{} }
func (m *Partition) String() string            { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()               {}
func (*Partition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Partition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Partition) GetLeft() string {
	if m != nil {
		return m.Left
	}
	return ""
}

func (m *Partition) GetRight() string {
	if m != nil {
		return m.Right
	}
	return ""
}

// TablePb defines the fields that define table attributes, and
// can be serialized.
type TablePb struct {
	// Name of table lowercased
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Name of table (not lowercased)
	NameOriginal string `protobuf:"bytes,2,opt,name=nameOriginal" json:"nameOriginal,omitempty"`
	// some dbs are more hiearchical (table-column-family)
	Parent string `protobuf:"bytes,3,opt,name=parent" json:"parent,omitempty"`
	// Character set, default = utf8
	Charset uint32 `protobuf:"varint,4,opt,name=Charset" json:"Charset,omitempty"`
	// Partitions in this table, optional may be empty
	Partition *TablePartition `protobuf:"bytes,5,opt,name=partition" json:"partition,omitempty"`
	// Partition Count
	PartitionCt uint32 `protobuf:"varint,6,opt,name=PartitionCt" json:"PartitionCt,omitempty"`
	// List of indexes for this table
	Indexes []*Index `protobuf:"bytes,7,rep,name=indexes" json:"indexes,omitempty"`
	// context json bytes
	ContextJson []byte `protobuf:"bytes,8,opt,name=contextJson,proto3" json:"contextJson,omitempty"`
	// List of Fields, in order
	Fieldpbs []*FieldPb `protobuf:"bytes,9,rep,name=fieldpbs" json:"fieldpbs,omitempty"`
}

func (m *TablePb) Reset()                    { *m = TablePb{} }
func (m *TablePb) String() string            { return proto.CompactTextString(m) }
func (*TablePb) ProtoMessage()               {}
func (*TablePb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TablePb) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TablePb) GetNameOriginal() string {
	if m != nil {
		return m.NameOriginal
	}
	return ""
}

func (m *TablePb) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *TablePb) GetCharset() uint32 {
	if m != nil {
		return m.Charset
	}
	return 0
}

func (m *TablePb) GetPartition() *TablePartition {
	if m != nil {
		return m.Partition
	}
	return nil
}

func (m *TablePb) GetPartitionCt() uint32 {
	if m != nil {
		return m.PartitionCt
	}
	return 0
}

func (m *TablePb) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *TablePb) GetContextJson() []byte {
	if m != nil {
		return m.ContextJson
	}
	return nil
}

func (m *TablePb) GetFieldpbs() []*FieldPb {
	if m != nil {
		return m.Fieldpbs
	}
	return nil
}

// FieldPb defines attributes of a field/column that can
// be serialized and transported.
type FieldPb struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Key         string   `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Extra       string   `protobuf:"bytes,4,opt,name=extra" json:"extra,omitempty"`
	Data        string   `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	Length      uint32   `protobuf:"varint,6,opt,name=length" json:"length,omitempty"`
	Type        uint32   `protobuf:"varint,7,opt,name=type" json:"type,omitempty"`
	NativeType  uint32   `protobuf:"varint,8,opt,name=nativeType" json:"nativeType,omitempty"`
	DefLength   uint64   `protobuf:"varint,9,opt,name=defLength" json:"defLength,omitempty"`
	DefVal      []byte   `protobuf:"bytes,11,opt,name=defVal,proto3" json:"defVal,omitempty"`
	Indexed     bool     `protobuf:"varint,13,opt,name=indexed" json:"indexed,omitempty"`
	NoNulls     bool     `protobuf:"varint,14,opt,name=noNulls" json:"noNulls,omitempty"`
	Collation   string   `protobuf:"bytes,15,opt,name=collation" json:"collation,omitempty"`
	Roles       []string `protobuf:"bytes,16,rep,name=roles" json:"roles,omitempty"`
	Indexes     []*Index `protobuf:"bytes,17,rep,name=indexes" json:"indexes,omitempty"`
	ContextJson []byte   `protobuf:"bytes,18,opt,name=contextJson,proto3" json:"contextJson,omitempty"`
	Position    uint64   `protobuf:"varint,19,opt,name=position" json:"position,omitempty"`
}

func (m *FieldPb) Reset()                    { *m = FieldPb{} }
func (m *FieldPb) String() string            { return proto.CompactTextString(m) }
func (*FieldPb) ProtoMessage()               {}
func (*FieldPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FieldPb) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldPb) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FieldPb) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FieldPb) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *FieldPb) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FieldPb) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *FieldPb) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FieldPb) GetNativeType() uint32 {
	if m != nil {
		return m.NativeType
	}
	return 0
}

func (m *FieldPb) GetDefLength() uint64 {
	if m != nil {
		return m.DefLength
	}
	return 0
}

func (m *FieldPb) GetDefVal() []byte {
	if m != nil {
		return m.DefVal
	}
	return nil
}

func (m *FieldPb) GetIndexed() bool {
	if m != nil {
		return m.Indexed
	}
	return false
}

func (m *FieldPb) GetNoNulls() bool {
	if m != nil {
		return m.NoNulls
	}
	return false
}

func (m *FieldPb) GetCollation() string {
	if m != nil {
		return m.Collation
	}
	return ""
}

func (m *FieldPb) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *FieldPb) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *FieldPb) GetContextJson() []byte {
	if m != nil {
		return m.ContextJson
	}
	return nil
}

func (m *FieldPb) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

// Index a description of how field(s) should be indexed for a table.
type Index struct {
	Name          string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Fields        []string `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	PrimaryKey    bool     `protobuf:"varint,3,opt,name=primaryKey" json:"primaryKey,omitempty"`
	HashPartition []string `protobuf:"bytes,4,rep,name=hashPartition" json:"hashPartition,omitempty"`
	PartitionSize int32    `protobuf:"varint,5,opt,name=partitionSize" json:"partitionSize,omitempty"`
}

func (m *Index) Reset()                    { *m = Index{} }
func (m *Index) String() string            { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()               {}
func (*Index) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Index) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Index) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Index) GetPrimaryKey() bool {
	if m != nil {
		return m.PrimaryKey
	}
	return false
}

func (m *Index) GetHashPartition() []string {
	if m != nil {
		return m.HashPartition
	}
	return nil
}

func (m *Index) GetPartitionSize() int32 {
	if m != nil {
		return m.PartitionSize
	}
	return 0
}

func init() {
	proto.RegisterType((*Command)(nil), "schema.Command")
	proto.RegisterType((*TablePartition)(nil), "schema.TablePartition")
	proto.RegisterType((*Partition)(nil), "schema.Partition")
	proto.RegisterType((*TablePb)(nil), "schema.TablePb")
	proto.RegisterType((*FieldPb)(nil), "schema.FieldPb")
	proto.RegisterType((*Index)(nil), "schema.Index")
	proto.RegisterEnum("schema.Command_Operation", Command_Operation_name, Command_Operation_value)
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x76, 0x7e, 0x1c, 0x4f, 0x9a, 0x34, 0xdd, 0x53, 0x55, 0x7b, 0x8e, 0x10, 0xb2, 0x2c,
	0x24, 0x8c, 0x90, 0x2a, 0xb5, 0xf0, 0x02, 0xa8, 0x80, 0xc4, 0x8f, 0x68, 0xb5, 0xb4, 0xdc, 0x6f,
	0xe2, 0x4d, 0x62, 0xd5, 0xd9, 0xb5, 0xbc, 0x0b, 0x34, 0xbc, 0x0c, 0x0f, 0xc3, 0x2d, 0x12, 0xaf,
	0x84, 0x76, 0x76, 0xed, 0x38, 0xa2, 0x37, 0x5c, 0x65, 0xbe, 0x6f, 0x66, 0x67, 0x3c, 0xdf, 0xb7,
	0x1b, 0x38, 0xd0, 0x8b, 0xb5, 0xd8, 0xf0, 0xd3, 0xaa, 0x56, 0x46, 0x91, 0xa1, 0x43, 0xe9, 0xaf,
	0x00, 0xa2, 0x0b, 0xb5, 0xd9, 0x70, 0x99, 0x93, 0x27, 0x10, 0xaa, 0x8a, 0x06, 0x49, 0x90, 0x4d,
	0xcf, 0xff, 0x3b, 0xf5, 0xe5, 0x3e, 0x79, 0x7a, 0x59, 0x89, 0x9a, 0x9b, 0x42, 0x49, 0x16, 0xaa,
	0x8a, 0x9c, 0xc0, 0x50, 0xd5, 0xc5, 0xaa, 0x90, 0x34, 0x4c, 0x82, 0x2c, 0x66, 0x1e, 0x59, 0xde,
	0x9d, 0xa3, 0x3d, 0xc7, 0x3b, 0x44, 0x8e, 0x61, 0x50, 0xc8, 0x5c, 0xdc, 0xd1, 0x7e, 0x12, 0x64,
	0x7d, 0xe6, 0x00, 0x99, 0x42, 0x68, 0x34, 0x1d, 0x24, 0x41, 0xd6, 0x63, 0xa1, 0xd1, 0x64, 0x06,
	0xbd, 0x8d, 0x5e, 0xd1, 0x61, 0x12, 0x64, 0x07, 0xcc, 0x86, 0xe9, 0x19, 0xc4, 0xed, 0x60, 0x32,
	0x86, 0xe8, 0x46, 0xde, 0x4a, 0xf5, 0x55, 0xce, 0xfe, 0x21, 0x13, 0x88, 0x5f, 0xe4, 0xf9, 0x4d,
	0x95, 0x73, 0x23, 0x66, 0x01, 0x19, 0x41, 0xff, 0x65, 0xad, 0xaa, 0x59, 0x98, 0x6e, 0x60, 0x7a,
	0xcd, 0xe7, 0xa5, 0xb8, 0xe2, 0xb5, 0x29, 0xf0, 0xdc, 0x31, 0x0c, 0x8c, 0x65, 0x70, 0xb5, 0x98,
	0x39, 0x40, 0x08, 0xf4, 0x6f, 0xc5, 0x56, 0xd3, 0x30, 0xe9, 0x65, 0x31, 0xc3, 0x98, 0x9c, 0x01,
	0x54, 0xcd, 0x31, 0x4d, 0x7b, 0x49, 0x2f, 0x1b, 0x9f, 0x1f, 0x35, 0x4a, 0xb4, 0x0d, 0x59, 0xa7,
	0x28, 0x7d, 0x05, 0xf1, 0x6e, 0xd2, 0x14, 0xc2, 0x22, 0xf7, 0x63, 0xc2, 0x22, 0xb7, 0x33, 0x4a,
	0xb1, 0x34, 0x5e, 0x24, 0x8c, 0xed, 0xd7, 0xd4, 0xc5, 0x6a, 0x6d, 0xbc, 0x42, 0x0e, 0xa4, 0x3f,
	0x42, 0x88, 0xdc, 0x67, 0xcf, 0xed, 0x29, 0xc9, 0x37, 0xcd, 0xe7, 0x62, 0x4c, 0x52, 0x38, 0xb0,
	0xbf, 0x97, 0x28, 0x33, 0x2f, 0x7d, 0xc7, 0x3d, 0xce, 0x8a, 0x5f, 0xf1, 0x5a, 0xc8, 0xa6, 0xb5,
	0x47, 0x84, 0x42, 0x74, 0xb1, 0xe6, 0xb5, 0x16, 0x06, 0xe5, 0x9f, 0xb0, 0x06, 0x92, 0xe7, 0x10,
	0xb7, 0xab, 0xa0, 0x0f, 0xe3, 0xf3, 0x93, 0x66, 0xdd, 0x7d, 0x11, 0xd9, 0xae, 0x90, 0x24, 0x30,
	0x6e, 0xf9, 0x0b, 0x83, 0x76, 0x4d, 0x58, 0x97, 0x22, 0x8f, 0x21, 0x42, 0x87, 0x85, 0xa6, 0x11,
	0x8a, 0x38, 0x69, 0xba, 0xbe, 0xb1, 0x34, 0x6b, 0xb2, 0xb6, 0xd5, 0x42, 0x49, 0x23, 0xee, 0xcc,
	0x5b, 0xad, 0x24, 0x1d, 0xa1, 0xf3, 0x5d, 0x8a, 0x3c, 0x85, 0xd1, 0xb2, 0x10, 0x65, 0x5e, 0xcd,
	0x35, 0x8d, 0xb1, 0xd7, 0x61, 0xd3, 0xeb, 0xb5, 0xe5, 0xaf, 0xe6, 0xac, 0x2d, 0x48, 0x7f, 0xf6,
	0x20, 0xf2, 0xec, 0xbd, 0x2a, 0x26, 0x30, 0xce, 0x85, 0x5e, 0xd4, 0x45, 0x85, 0x1b, 0x3b, 0x11,
	0xbb, 0x94, 0xbd, 0x82, 0xb7, 0x62, 0xeb, 0x05, 0xb4, 0xa1, 0xf5, 0x4b, 0xdc, 0x99, 0x9a, 0xa3,
	0x76, 0x31, 0x73, 0xc0, 0x76, 0xcf, 0xb9, 0xe1, 0x28, 0x5a, 0xcc, 0x30, 0xb6, 0xfa, 0x97, 0x42,
	0xae, 0xcc, 0xda, 0x4b, 0xe2, 0x91, 0xad, 0x35, 0xdb, 0x4a, 0xd0, 0x08, 0x59, 0x8c, 0xc9, 0x43,
	0x00, 0xc9, 0x4d, 0xf1, 0x45, 0x5c, 0xdb, 0xcc, 0x08, 0x33, 0x1d, 0x86, 0x3c, 0x80, 0x38, 0x17,
	0xcb, 0xf7, 0xae, 0x5d, 0x8c, 0x8f, 0x66, 0x47, 0xd8, 0x49, 0xb9, 0x58, 0x7e, 0xe2, 0x25, 0x1d,
	0xa3, 0x62, 0x1e, 0x59, 0xa7, 0x9d, 0xb2, 0x39, 0x9d, 0x24, 0x41, 0x36, 0x6a, 0x84, 0xce, 0x6d,
	0x46, 0xaa, 0x0f, 0x9f, 0xcb, 0x52, 0xd3, 0xa9, 0xcb, 0x78, 0x68, 0x27, 0x2d, 0x54, 0x59, 0xe2,
	0x13, 0xa3, 0x87, 0xb8, 0xce, 0x8e, 0xc0, 0xdb, 0xaa, 0x4a, 0xa1, 0xe9, 0x0c, 0x9f, 0x89, 0x03,
	0x5d, 0x7f, 0x8f, 0xfe, 0xc6, 0x5f, 0xf2, 0xa7, 0xbf, 0xff, 0xc3, 0xa8, 0x52, 0xda, 0xdd, 0xc0,
	0x7f, 0x71, 0xcf, 0x16, 0xa7, 0xdf, 0x03, 0x18, 0x60, 0xc3, 0x7b, 0xcd, 0x3c, 0x81, 0x21, 0x1a,
	0xdf, 0x3c, 0x61, 0x8f, 0xac, 0xb4, 0x55, 0x5d, 0x6c, 0x78, 0xbd, 0x7d, 0xe7, 0x9d, 0x1c, 0xb1,
	0x0e, 0x43, 0x1e, 0xc1, 0x64, 0xcd, 0xf5, 0xba, 0xbd, 0xaf, 0xb4, 0x8f, 0xc7, 0xf7, 0x49, 0x5b,
	0xd5, 0xde, 0xf8, 0x8f, 0xc5, 0x37, 0x81, 0x4e, 0x0f, 0xd8, 0x3e, 0x39, 0x1f, 0xe2, 0xbf, 0xe9,
	0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x99, 0x99, 0xd3, 0x5d, 0x05, 0x00, 0x00,
}
