// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/cost_graph.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CostGraphDef struct {
	Node                 []*CostGraphDef_Node           `protobuf:"bytes,1,rep,name=node,proto3" json:"node,omitempty"`
	Cost                 []*CostGraphDef_AggregatedCost `protobuf:"bytes,2,rep,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CostGraphDef) Reset()         { *m = CostGraphDef{} }
func (m *CostGraphDef) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef) ProtoMessage()    {}
func (*CostGraphDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0}
}

func (m *CostGraphDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef.Unmarshal(m, b)
}
func (m *CostGraphDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef.Marshal(b, m, deterministic)
}
func (m *CostGraphDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef.Merge(m, src)
}
func (m *CostGraphDef) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef.Size(m)
}
func (m *CostGraphDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef proto.InternalMessageInfo

func (m *CostGraphDef) GetNode() []*CostGraphDef_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CostGraphDef) GetCost() []*CostGraphDef_AggregatedCost {
	if m != nil {
		return m.Cost
	}
	return nil
}

type CostGraphDef_Node struct {
	// The name of the node. Names are globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The device of the node. Can be empty if the node is mapped to the
	// default partition or partitioning hasn't been run yet.
	Device string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	// The id of the node. Node ids are only unique inside a partition.
	Id         int32                           `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	InputInfo  []*CostGraphDef_Node_InputInfo  `protobuf:"bytes,4,rep,name=input_info,json=inputInfo,proto3" json:"input_info,omitempty"`
	OutputInfo []*CostGraphDef_Node_OutputInfo `protobuf:"bytes,5,rep,name=output_info,json=outputInfo,proto3" json:"output_info,omitempty"`
	// Temporary memory used by this node.
	TemporaryMemorySize int64 `protobuf:"varint,6,opt,name=temporary_memory_size,json=temporaryMemorySize,proto3" json:"temporary_memory_size,omitempty"`
	// Persistent memory used by this node.
	PersistentMemorySize       int64 `protobuf:"varint,12,opt,name=persistent_memory_size,json=persistentMemorySize,proto3" json:"persistent_memory_size,omitempty"`
	HostTempMemorySize         int64 `protobuf:"varint,10,opt,name=host_temp_memory_size,json=hostTempMemorySize,proto3" json:"host_temp_memory_size,omitempty"`                         // Deprecated: Do not use.
	DeviceTempMemorySize       int64 `protobuf:"varint,11,opt,name=device_temp_memory_size,json=deviceTempMemorySize,proto3" json:"device_temp_memory_size,omitempty"`                   // Deprecated: Do not use.
	DevicePersistentMemorySize int64 `protobuf:"varint,16,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize,proto3" json:"device_persistent_memory_size,omitempty"` // Deprecated: Do not use.
	// Estimate of the computational cost of this node, in microseconds.
	ComputeCost int64 `protobuf:"varint,9,opt,name=compute_cost,json=computeCost,proto3" json:"compute_cost,omitempty"`
	// Analytical estimate of the computational cost of this node, in
	// microseconds.
	ComputeTime int64 `protobuf:"varint,14,opt,name=compute_time,json=computeTime,proto3" json:"compute_time,omitempty"`
	// Analytical estimate of the memory access cost of this node, in
	// microseconds.
	MemoryTime int64 `protobuf:"varint,15,opt,name=memory_time,json=memoryTime,proto3" json:"memory_time,omitempty"`
	// If true, the output is permanent: it can't be discarded, because this
	// node is part of the "final output". Nodes may depend on final nodes.
	IsFinal bool `protobuf:"varint,7,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	// Ids of the control inputs for this node.
	ControlInput []int32 `protobuf:"varint,8,rep,packed,name=control_input,json=controlInput,proto3" json:"control_input,omitempty"`
	// Are the costs inaccurate?
	Inaccurate           bool     `protobuf:"varint,17,opt,name=inaccurate,proto3" json:"inaccurate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostGraphDef_Node) Reset()         { *m = CostGraphDef_Node{} }
func (m *CostGraphDef_Node) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node) ProtoMessage()    {}
func (*CostGraphDef_Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 0}
}

func (m *CostGraphDef_Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_Node.Unmarshal(m, b)
}
func (m *CostGraphDef_Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_Node.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_Node.Merge(m, src)
}
func (m *CostGraphDef_Node) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_Node.Size(m)
}
func (m *CostGraphDef_Node) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_Node.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_Node proto.InternalMessageInfo

func (m *CostGraphDef_Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CostGraphDef_Node) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *CostGraphDef_Node) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CostGraphDef_Node) GetInputInfo() []*CostGraphDef_Node_InputInfo {
	if m != nil {
		return m.InputInfo
	}
	return nil
}

func (m *CostGraphDef_Node) GetOutputInfo() []*CostGraphDef_Node_OutputInfo {
	if m != nil {
		return m.OutputInfo
	}
	return nil
}

func (m *CostGraphDef_Node) GetTemporaryMemorySize() int64 {
	if m != nil {
		return m.TemporaryMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetPersistentMemorySize() int64 {
	if m != nil {
		return m.PersistentMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *CostGraphDef_Node) GetHostTempMemorySize() int64 {
	if m != nil {
		return m.HostTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *CostGraphDef_Node) GetDeviceTempMemorySize() int64 {
	if m != nil {
		return m.DeviceTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *CostGraphDef_Node) GetDevicePersistentMemorySize() int64 {
	if m != nil {
		return m.DevicePersistentMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetComputeCost() int64 {
	if m != nil {
		return m.ComputeCost
	}
	return 0
}

func (m *CostGraphDef_Node) GetComputeTime() int64 {
	if m != nil {
		return m.ComputeTime
	}
	return 0
}

func (m *CostGraphDef_Node) GetMemoryTime() int64 {
	if m != nil {
		return m.MemoryTime
	}
	return 0
}

func (m *CostGraphDef_Node) GetIsFinal() bool {
	if m != nil {
		return m.IsFinal
	}
	return false
}

func (m *CostGraphDef_Node) GetControlInput() []int32 {
	if m != nil {
		return m.ControlInput
	}
	return nil
}

func (m *CostGraphDef_Node) GetInaccurate() bool {
	if m != nil {
		return m.Inaccurate
	}
	return false
}

// Inputs of this node. They must be executed before this node can be
// executed. An input is a particular output of another node, specified
// by the node id and the output index.
type CostGraphDef_Node_InputInfo struct {
	PrecedingNode        int32    `protobuf:"varint,1,opt,name=preceding_node,json=precedingNode,proto3" json:"preceding_node,omitempty"`
	PrecedingPort        int32    `protobuf:"varint,2,opt,name=preceding_port,json=precedingPort,proto3" json:"preceding_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostGraphDef_Node_InputInfo) Reset()         { *m = CostGraphDef_Node_InputInfo{} }
func (m *CostGraphDef_Node_InputInfo) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node_InputInfo) ProtoMessage()    {}
func (*CostGraphDef_Node_InputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 0, 0}
}

func (m *CostGraphDef_Node_InputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_Node_InputInfo.Unmarshal(m, b)
}
func (m *CostGraphDef_Node_InputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_Node_InputInfo.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_Node_InputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_Node_InputInfo.Merge(m, src)
}
func (m *CostGraphDef_Node_InputInfo) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_Node_InputInfo.Size(m)
}
func (m *CostGraphDef_Node_InputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_Node_InputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_Node_InputInfo proto.InternalMessageInfo

func (m *CostGraphDef_Node_InputInfo) GetPrecedingNode() int32 {
	if m != nil {
		return m.PrecedingNode
	}
	return 0
}

func (m *CostGraphDef_Node_InputInfo) GetPrecedingPort() int32 {
	if m != nil {
		return m.PrecedingPort
	}
	return 0
}

// Outputs of this node.
type CostGraphDef_Node_OutputInfo struct {
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// If >= 0, the output is an alias of an input. Note that an alias input
	// may itself be an alias. The algorithm will therefore need to follow
	// those pointers.
	AliasInputPort       int64             `protobuf:"varint,2,opt,name=alias_input_port,json=aliasInputPort,proto3" json:"alias_input_port,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,3,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype                DataType          `protobuf:"varint,4,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CostGraphDef_Node_OutputInfo) Reset()         { *m = CostGraphDef_Node_OutputInfo{} }
func (m *CostGraphDef_Node_OutputInfo) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node_OutputInfo) ProtoMessage()    {}
func (*CostGraphDef_Node_OutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 0, 1}
}

func (m *CostGraphDef_Node_OutputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_Node_OutputInfo.Unmarshal(m, b)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_Node_OutputInfo.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_Node_OutputInfo.Merge(m, src)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_Node_OutputInfo.Size(m)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_Node_OutputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_Node_OutputInfo proto.InternalMessageInfo

func (m *CostGraphDef_Node_OutputInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CostGraphDef_Node_OutputInfo) GetAliasInputPort() int64 {
	if m != nil {
		return m.AliasInputPort
	}
	return 0
}

func (m *CostGraphDef_Node_OutputInfo) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *CostGraphDef_Node_OutputInfo) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

// Total cost of this graph, typically used for balancing decisions.
type CostGraphDef_AggregatedCost struct {
	// Aggregated cost value.
	Cost float32 `protobuf:"fixed32,1,opt,name=cost,proto3" json:"cost,omitempty"`
	// Aggregated cost dimension (e.g. 'memory', 'compute', 'network').
	Dimension            string   `protobuf:"bytes,2,opt,name=dimension,proto3" json:"dimension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostGraphDef_AggregatedCost) Reset()         { *m = CostGraphDef_AggregatedCost{} }
func (m *CostGraphDef_AggregatedCost) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_AggregatedCost) ProtoMessage()    {}
func (*CostGraphDef_AggregatedCost) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 1}
}

func (m *CostGraphDef_AggregatedCost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_AggregatedCost.Unmarshal(m, b)
}
func (m *CostGraphDef_AggregatedCost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_AggregatedCost.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_AggregatedCost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_AggregatedCost.Merge(m, src)
}
func (m *CostGraphDef_AggregatedCost) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_AggregatedCost.Size(m)
}
func (m *CostGraphDef_AggregatedCost) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_AggregatedCost.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_AggregatedCost proto.InternalMessageInfo

func (m *CostGraphDef_AggregatedCost) GetCost() float32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *CostGraphDef_AggregatedCost) GetDimension() string {
	if m != nil {
		return m.Dimension
	}
	return ""
}

func init() {
	proto.RegisterType((*CostGraphDef)(nil), "tensorflow.CostGraphDef")
	proto.RegisterType((*CostGraphDef_Node)(nil), "tensorflow.CostGraphDef.Node")
	proto.RegisterType((*CostGraphDef_Node_InputInfo)(nil), "tensorflow.CostGraphDef.Node.InputInfo")
	proto.RegisterType((*CostGraphDef_Node_OutputInfo)(nil), "tensorflow.CostGraphDef.Node.OutputInfo")
	proto.RegisterType((*CostGraphDef_AggregatedCost)(nil), "tensorflow.CostGraphDef.AggregatedCost")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/cost_graph.proto", fileDescriptor_5f8948141565ace8)
}

var fileDescriptor_5f8948141565ace8 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0xf3, 0xa7, 0x6d, 0x26, 0x69, 0xda, 0xdf, 0xfe, 0xda, 0x62, 0xa2, 0x16, 0x02, 0xa8,
	0xc2, 0xaa, 0x50, 0x22, 0x02, 0x1c, 0x10, 0xe2, 0x40, 0x28, 0x45, 0x3d, 0x00, 0x91, 0x9b, 0x0b,
	0x5c, 0x2c, 0xd7, 0x9e, 0x38, 0x2b, 0x62, 0xaf, 0xb5, 0xbb, 0xa1, 0x4a, 0x9f, 0x81, 0x27, 0xe1,
	0x45, 0x78, 0x25, 0x8e, 0x68, 0xc7, 0xc6, 0x71, 0x4a, 0x5b, 0x6e, 0xe3, 0xd9, 0xef, 0xfb, 0x66,
	0x76, 0xfd, 0xcd, 0xc0, 0x91, 0xc6, 0x44, 0x09, 0x39, 0x99, 0x89, 0x8b, 0x7e, 0x20, 0x24, 0xf6,
	0x27, 0xd2, 0x8f, 0xf1, 0x42, 0xc8, 0xaf, 0xfd, 0x40, 0x28, 0xed, 0x45, 0xd2, 0x4f, 0xa7, 0xbd,
	0x54, 0x0a, 0x2d, 0x18, 0x2c, 0xb1, 0x9d, 0x27, 0x37, 0xf3, 0xb2, 0x13, 0x4f, 0x4d, 0xfd, 0x14,
	0x33, 0x66, 0xe7, 0xf0, 0x16, 0xf4, 0x22, 0x45, 0x95, 0xc1, 0x1e, 0x7e, 0x6f, 0x40, 0xeb, 0xad,
	0x50, 0xfa, 0xbd, 0x29, 0x7a, 0x8c, 0x13, 0xf6, 0x14, 0x6a, 0x89, 0x08, 0xd1, 0xb6, 0xba, 0x55,
	0xa7, 0x39, 0x38, 0xe8, 0x2d, 0x65, 0x7a, 0x65, 0x5c, 0xef, 0xa3, 0x08, 0xd1, 0x25, 0x28, 0x7b,
	0x05, 0x35, 0xd3, 0xb8, 0x5d, 0x21, 0xca, 0xe3, 0x1b, 0x29, 0x6f, 0xa2, 0x48, 0x62, 0xe4, 0x6b,
	0x0c, 0x4d, 0xda, 0x25, 0x52, 0xe7, 0xe7, 0x3a, 0xd4, 0x8c, 0x16, 0x63, 0x50, 0x4b, 0xfc, 0xd8,
	0x14, 0xb6, 0x9c, 0x86, 0x4b, 0x31, 0xdb, 0x83, 0xb5, 0x10, 0xbf, 0xf1, 0x00, 0xed, 0x0a, 0x65,
	0xf3, 0x2f, 0xd6, 0x86, 0x0a, 0x0f, 0xed, 0x6a, 0xd7, 0x72, 0xea, 0x6e, 0x85, 0x87, 0xec, 0x04,
	0x80, 0x27, 0xe9, 0x5c, 0x7b, 0x3c, 0x99, 0x08, 0xbb, 0xf6, 0x8f, 0x3e, 0x4c, 0xb9, 0xde, 0xa9,
	0xc1, 0x9f, 0x26, 0x13, 0xe1, 0x36, 0xf8, 0x9f, 0x90, 0x9d, 0x42, 0x53, 0xcc, 0x75, 0x21, 0x54,
	0x27, 0x21, 0xe7, 0x76, 0xa1, 0x4f, 0x44, 0x20, 0x25, 0x10, 0x45, 0xcc, 0x06, 0xb0, 0xab, 0x31,
	0x4e, 0x85, 0xf4, 0xe5, 0xc2, 0x8b, 0x31, 0x16, 0x72, 0xe1, 0x29, 0x7e, 0x89, 0xf6, 0x5a, 0xd7,
	0x72, 0xaa, 0xee, 0xff, 0xc5, 0xe1, 0x07, 0x3a, 0x3b, 0xe3, 0x97, 0xc8, 0x9e, 0xc3, 0x5e, 0x8a,
	0x52, 0x71, 0xa5, 0x31, 0xd1, 0x2b, 0xa4, 0x16, 0x91, 0x76, 0x96, 0xa7, 0x25, 0xd6, 0x0b, 0xd8,
	0x9d, 0x1a, 0xdf, 0x18, 0xc5, 0x15, 0x12, 0x18, 0xd2, 0xb0, 0x62, 0x5b, 0x2e, 0x33, 0x80, 0x31,
	0xc6, 0x69, 0x89, 0xf6, 0x12, 0xee, 0x64, 0xaf, 0xf9, 0x37, 0xb1, 0x59, 0x10, 0x77, 0x32, 0xc8,
	0x15, 0xea, 0x3b, 0x38, 0xc8, 0xa9, 0x37, 0xb4, 0xbb, 0x5d, 0x08, 0x74, 0x32, 0xe0, 0xe8, 0xba,
	0xc6, 0x1f, 0x40, 0x2b, 0x10, 0x71, 0x3a, 0xd7, 0xe8, 0x91, 0x7f, 0x1a, 0x74, 0xc9, 0x66, 0x9e,
	0x33, 0x2f, 0x5d, 0x86, 0x68, 0x1e, 0xa3, 0xdd, 0x5e, 0x81, 0x8c, 0x79, 0x8c, 0xec, 0x3e, 0x34,
	0xf3, 0xd2, 0x84, 0xd8, 0x22, 0x04, 0x64, 0x29, 0x02, 0xdc, 0x85, 0x0d, 0xae, 0xbc, 0x09, 0x4f,
	0xfc, 0x99, 0xbd, 0xde, 0xb5, 0x9c, 0x0d, 0x77, 0x9d, 0xab, 0x13, 0xf3, 0xc9, 0x1e, 0xc1, 0x66,
	0x20, 0x12, 0x2d, 0xc5, 0xcc, 0x23, 0x13, 0xd8, 0x1b, 0xdd, 0xaa, 0x53, 0x77, 0x5b, 0x79, 0x92,
	0x3c, 0xc2, 0xee, 0x19, 0x73, 0xf9, 0x41, 0x30, 0x97, 0xbe, 0x46, 0xfb, 0x3f, 0x52, 0x28, 0x65,
	0x3a, 0x9f, 0xa1, 0x51, 0x98, 0x89, 0x1d, 0x42, 0x3b, 0x95, 0x18, 0x60, 0xc8, 0x93, 0xc8, 0xcb,
	0x07, 0xc9, 0xb8, 0x74, 0xb3, 0xc8, 0x92, 0xd9, 0x57, 0x60, 0xa9, 0x90, 0x9a, 0x0c, 0x5e, 0x86,
	0x8d, 0x84, 0xd4, 0x9d, 0x1f, 0x16, 0xc0, 0xd2, 0x5f, 0x66, 0x44, 0xe8, 0x79, 0x2d, 0xba, 0x23,
	0xc5, 0xcc, 0x81, 0x6d, 0x7f, 0xc6, 0x7d, 0x95, 0x5d, 0x60, 0xa9, 0x55, 0x75, 0xdb, 0x94, 0xa7,
	0xd6, 0x8c, 0x18, 0x1b, 0x40, 0x9d, 0x16, 0x04, 0xcd, 0x4d, 0x73, 0xb0, 0x5f, 0xb6, 0xf5, 0x98,
	0xc2, 0x33, 0x73, 0x3c, 0x32, 0x7b, 0xc1, 0xcd, 0xa0, 0xec, 0x08, 0xea, 0xa1, 0x59, 0x17, 0x76,
	0xad, 0x6b, 0x39, 0xed, 0xc1, 0x4e, 0x99, 0x73, 0xec, 0x6b, 0x7f, 0xbc, 0x48, 0xd1, 0xcd, 0x20,
	0x9d, 0x21, 0xb4, 0x57, 0x27, 0xdc, 0xf4, 0x4b, 0x3f, 0xd6, 0xf4, 0x5b, 0xc9, 0xe6, 0x9d, 0xed,
	0x43, 0x23, 0xe4, 0x31, 0x26, 0x8a, 0x8b, 0x24, 0x9f, 0xea, 0x65, 0x62, 0x28, 0xc0, 0x16, 0x32,
	0x2a, 0x57, 0x29, 0xd6, 0xd6, 0x70, 0xab, 0x98, 0x3d, 0x6a, 0x51, 0x8d, 0xac, 0x2f, 0xaf, 0x23,
	0xae, 0xa7, 0xf3, 0xf3, 0x5e, 0x20, 0xe2, 0x7e, 0x69, 0xdf, 0x5d, 0x1f, 0x46, 0xe2, 0xca, 0x22,
	0xfc, 0x65, 0x59, 0xe7, 0x6b, 0xb4, 0x06, 0x9f, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x8c,
	0xdb, 0x50, 0x95, 0x05, 0x00, 0x00,
}