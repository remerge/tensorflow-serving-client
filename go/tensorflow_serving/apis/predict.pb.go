// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/predict.proto

package tensorflow_serving

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tensor_go_proto "github.com/figroc/tensorflow-serving-client/v2/go/tensorflow/core/framework/tensor_go_proto"
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

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'inputs' field.
	Inputs map[string]*tensor_go_proto.TensorProto `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'outputs' field.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter         []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter,proto3" json:"output_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0799bd5d5992a93, []int{0}
}

func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictRequest) GetInputs() map[string]*tensor_go_proto.TensorProto {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PredictRequest) GetOutputFilter() []string {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	// Effective Model Specification used to process PredictRequest.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Output tensors.
	Outputs              map[string]*tensor_go_proto.TensorProto `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *PredictResponse) Reset()         { *m = PredictResponse{} }
func (m *PredictResponse) String() string { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()    {}
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0799bd5d5992a93, []int{1}
}

func (m *PredictResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictResponse.Unmarshal(m, b)
}
func (m *PredictResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictResponse.Marshal(b, m, deterministic)
}
func (m *PredictResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictResponse.Merge(m, src)
}
func (m *PredictResponse) XXX_Size() int {
	return xxx_messageInfo_PredictResponse.Size(m)
}
func (m *PredictResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictResponse proto.InternalMessageInfo

func (m *PredictResponse) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictResponse) GetOutputs() map[string]*tensor_go_proto.TensorProto {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "tensorflow.serving.PredictRequest")
	proto.RegisterMapType((map[string]*tensor_go_proto.TensorProto)(nil), "tensorflow.serving.PredictRequest.InputsEntry")
	proto.RegisterType((*PredictResponse)(nil), "tensorflow.serving.PredictResponse")
	proto.RegisterMapType((map[string]*tensor_go_proto.TensorProto)(nil), "tensorflow.serving.PredictResponse.OutputsEntry")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/predict.proto", fileDescriptor_d0799bd5d5992a93)
}

var fileDescriptor_d0799bd5d5992a93 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x64, 0x13, 0x7e, 0xfd, 0xd1, 0xd7, 0xfa, 0x87, 0xbd, 0x18, 0x02, 0x42, 0x69, 0x51, 0x7a,
	0x71, 0x23, 0xf5, 0x22, 0xe2, 0x49, 0xb0, 0xa0, 0x20, 0x96, 0xad, 0xf7, 0x52, 0xd3, 0x57, 0x09,
	0x4d, 0xb2, 0xeb, 0xee, 0xa6, 0xa5, 0x9f, 0xc2, 0xaf, 0xea, 0xd1, 0xa3, 0x64, 0xb7, 0xb5, 0xf1,
	0x2f, 0x08, 0xde, 0x1e, 0xf3, 0x66, 0x76, 0xe6, 0x0d, 0x0b, 0x07, 0x06, 0x73, 0x2d, 0xd4, 0x34,
	0x15, 0x8b, 0x91, 0x46, 0x35, 0x4f, 0xf2, 0x87, 0x68, 0x2c, 0x13, 0x1d, 0x49, 0x85, 0x93, 0x24,
	0x36, 0x4c, 0x2a, 0x61, 0x04, 0xa5, 0x1b, 0x1a, 0x5b, 0xd1, 0xc2, 0xc3, 0x0d, 0x16, 0xc5, 0x42,
	0x61, 0x34, 0x55, 0xe3, 0x0c, 0x17, 0x42, 0xcd, 0x22, 0xb7, 0x71, 0xda, 0xb0, 0xf3, 0x9d, 0x45,
	0x26, 0x26, 0x98, 0x3a, 0x52, 0xfb, 0xc9, 0x83, 0xed, 0x81, 0xb3, 0xe4, 0xf8, 0x58, 0xa0, 0x36,
	0xf4, 0x1c, 0xc0, 0x32, 0x46, 0x5a, 0x62, 0x1c, 0x90, 0x16, 0xe9, 0x36, 0x7a, 0xfb, 0xec, 0x73,
	0x10, 0x76, 0x53, 0xb2, 0x86, 0x12, 0x63, 0x5e, 0xcf, 0xd6, 0x23, 0xed, 0x43, 0x2d, 0xc9, 0x65,
	0x61, 0x74, 0xe0, 0xb5, 0xfc, 0x6e, 0xa3, 0xc7, 0xbe, 0x52, 0xbe, 0x77, 0x64, 0x57, 0x56, 0x70,
	0x99, 0x1b, 0xb5, 0xe4, 0x2b, 0x35, 0xed, 0xc0, 0x96, 0x28, 0x8c, 0x2c, 0xcc, 0x68, 0x9a, 0xa4,
	0x06, 0x55, 0xe0, 0xb7, 0xfc, 0x6e, 0x9d, 0x37, 0x1d, 0xd8, 0xb7, 0x58, 0xc8, 0xa1, 0x51, 0xd1,
	0xd2, 0x5d, 0xf0, 0x67, 0xb8, 0xb4, 0x91, 0xeb, 0xbc, 0x1c, 0xe9, 0x11, 0xfc, 0x9b, 0x8f, 0xd3,
	0x02, 0x03, 0xcf, 0x9e, 0xb1, 0x57, 0x0d, 0x73, 0x67, 0xc7, 0x41, 0x59, 0x03, 0x77, 0xac, 0x33,
	0xef, 0x94, 0xb4, 0x9f, 0x09, 0xec, 0xbc, 0xe5, 0xd3, 0x52, 0xe4, 0x1a, 0x3f, 0x54, 0xe2, 0xfd,
	0xb2, 0x92, 0x6b, 0xf8, 0xef, 0x52, 0xeb, 0x80, 0xd8, 0x4e, 0x8e, 0x7f, 0xec, 0xc4, 0x79, 0xb2,
	0x5b, 0x27, 0x71, 0xad, 0xac, 0x1f, 0x08, 0x87, 0xd0, 0xac, 0x2e, 0xfe, 0xe4, 0xe4, 0x0b, 0xff,
	0x85, 0x90, 0xfb, 0x9a, 0xfd, 0x10, 0x27, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x57, 0xf4,
	0x30, 0x9a, 0x02, 0x00, 0x00,
}
