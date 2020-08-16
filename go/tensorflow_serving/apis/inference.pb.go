// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/inference.proto

package tensorflow_serving

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

// Inference request such as classification, regression, etc...
type InferenceTask struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	// All ModelSpecs in a MultiInferenceRequest must access the same model name.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Signature's method_name. Should be one of the method names defined in
	// third_party/tensorflow/python/saved_model/signature_constants.py.
	// e.g. "tensorflow/serving/classify".
	MethodName           string   `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferenceTask) Reset()         { *m = InferenceTask{} }
func (m *InferenceTask) String() string { return proto.CompactTextString(m) }
func (*InferenceTask) ProtoMessage()    {}
func (*InferenceTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{0}
}

func (m *InferenceTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferenceTask.Unmarshal(m, b)
}
func (m *InferenceTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferenceTask.Marshal(b, m, deterministic)
}
func (m *InferenceTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferenceTask.Merge(m, src)
}
func (m *InferenceTask) XXX_Size() int {
	return xxx_messageInfo_InferenceTask.Size(m)
}
func (m *InferenceTask) XXX_DiscardUnknown() {
	xxx_messageInfo_InferenceTask.DiscardUnknown(m)
}

var xxx_messageInfo_InferenceTask proto.InternalMessageInfo

func (m *InferenceTask) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *InferenceTask) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

// Inference result, matches the type of request or is an error.
type InferenceResult struct {
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*InferenceResult_ClassificationResult
	//	*InferenceResult_RegressionResult
	Result               isInferenceResult_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *InferenceResult) Reset()         { *m = InferenceResult{} }
func (m *InferenceResult) String() string { return proto.CompactTextString(m) }
func (*InferenceResult) ProtoMessage()    {}
func (*InferenceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{1}
}

func (m *InferenceResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferenceResult.Unmarshal(m, b)
}
func (m *InferenceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferenceResult.Marshal(b, m, deterministic)
}
func (m *InferenceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferenceResult.Merge(m, src)
}
func (m *InferenceResult) XXX_Size() int {
	return xxx_messageInfo_InferenceResult.Size(m)
}
func (m *InferenceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InferenceResult.DiscardUnknown(m)
}

var xxx_messageInfo_InferenceResult proto.InternalMessageInfo

func (m *InferenceResult) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

type isInferenceResult_Result interface {
	isInferenceResult_Result()
}

type InferenceResult_ClassificationResult struct {
	ClassificationResult *ClassificationResult `protobuf:"bytes,2,opt,name=classification_result,json=classificationResult,proto3,oneof"`
}

type InferenceResult_RegressionResult struct {
	RegressionResult *RegressionResult `protobuf:"bytes,3,opt,name=regression_result,json=regressionResult,proto3,oneof"`
}

func (*InferenceResult_ClassificationResult) isInferenceResult_Result() {}

func (*InferenceResult_RegressionResult) isInferenceResult_Result() {}

func (m *InferenceResult) GetResult() isInferenceResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *InferenceResult) GetClassificationResult() *ClassificationResult {
	if x, ok := m.GetResult().(*InferenceResult_ClassificationResult); ok {
		return x.ClassificationResult
	}
	return nil
}

func (m *InferenceResult) GetRegressionResult() *RegressionResult {
	if x, ok := m.GetResult().(*InferenceResult_RegressionResult); ok {
		return x.RegressionResult
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InferenceResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InferenceResult_ClassificationResult)(nil),
		(*InferenceResult_RegressionResult)(nil),
	}
}

// Inference request containing one or more requests.
type MultiInferenceRequest struct {
	// Inference tasks.
	Tasks []*InferenceTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// Input data.
	Input                *Input   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiInferenceRequest) Reset()         { *m = MultiInferenceRequest{} }
func (m *MultiInferenceRequest) String() string { return proto.CompactTextString(m) }
func (*MultiInferenceRequest) ProtoMessage()    {}
func (*MultiInferenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{2}
}

func (m *MultiInferenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiInferenceRequest.Unmarshal(m, b)
}
func (m *MultiInferenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiInferenceRequest.Marshal(b, m, deterministic)
}
func (m *MultiInferenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiInferenceRequest.Merge(m, src)
}
func (m *MultiInferenceRequest) XXX_Size() int {
	return xxx_messageInfo_MultiInferenceRequest.Size(m)
}
func (m *MultiInferenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiInferenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiInferenceRequest proto.InternalMessageInfo

func (m *MultiInferenceRequest) GetTasks() []*InferenceTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *MultiInferenceRequest) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

// Inference request containing one or more responses.
type MultiInferenceResponse struct {
	// List of results; one for each InferenceTask in the request, returned in the
	// same order as the request.
	Results              []*InferenceResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MultiInferenceResponse) Reset()         { *m = MultiInferenceResponse{} }
func (m *MultiInferenceResponse) String() string { return proto.CompactTextString(m) }
func (*MultiInferenceResponse) ProtoMessage()    {}
func (*MultiInferenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4842740eef0fb1a6, []int{3}
}

func (m *MultiInferenceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiInferenceResponse.Unmarshal(m, b)
}
func (m *MultiInferenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiInferenceResponse.Marshal(b, m, deterministic)
}
func (m *MultiInferenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiInferenceResponse.Merge(m, src)
}
func (m *MultiInferenceResponse) XXX_Size() int {
	return xxx_messageInfo_MultiInferenceResponse.Size(m)
}
func (m *MultiInferenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiInferenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultiInferenceResponse proto.InternalMessageInfo

func (m *MultiInferenceResponse) GetResults() []*InferenceResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*InferenceTask)(nil), "tensorflow.serving.InferenceTask")
	proto.RegisterType((*InferenceResult)(nil), "tensorflow.serving.InferenceResult")
	proto.RegisterType((*MultiInferenceRequest)(nil), "tensorflow.serving.MultiInferenceRequest")
	proto.RegisterType((*MultiInferenceResponse)(nil), "tensorflow.serving.MultiInferenceResponse")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/inference.proto", fileDescriptor_4842740eef0fb1a6)
}

var fileDescriptor_4842740eef0fb1a6 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x4e, 0xe3, 0x40,
	0x10, 0x86, 0xcf, 0x89, 0x92, 0xbb, 0x8c, 0x75, 0xba, 0xbb, 0xd5, 0xe5, 0x94, 0x8b, 0x84, 0x08,
	0x0e, 0x12, 0x2e, 0x90, 0x23, 0x85, 0x82, 0x06, 0x9a, 0xd0, 0x40, 0x11, 0x8a, 0x0d, 0x12, 0xa5,
	0x65, 0x9c, 0x49, 0xb0, 0x62, 0xef, 0x9a, 0x9d, 0x35, 0xd4, 0x3c, 0x01, 0xaf, 0x4b, 0x89, 0xf0,
	0xc6, 0x0e, 0x09, 0x0e, 0x14, 0x74, 0xd6, 0xe8, 0x9b, 0xef, 0x9f, 0x19, 0x2f, 0x1c, 0x68, 0x14,
	0x24, 0xd5, 0x2c, 0x96, 0x0f, 0x3e, 0xa1, 0xba, 0x8f, 0xc4, 0x7c, 0x10, 0xa4, 0x11, 0x0d, 0x22,
	0x31, 0x43, 0x85, 0x22, 0x44, 0x2f, 0x55, 0x52, 0x4b, 0xc6, 0x56, 0xa0, 0xb7, 0x04, 0xbb, 0x87,
	0xdb, 0x9a, 0xc3, 0x38, 0x20, 0x8a, 0x66, 0x51, 0x18, 0xe8, 0x48, 0x0a, 0x63, 0xe8, 0xf6, 0xb7,
	0x47, 0xa5, 0x99, 0xfe, 0x0c, 0x4a, 0xe4, 0x14, 0xe3, 0x25, 0xe4, 0x6e, 0x83, 0x14, 0xce, 0x15,
	0x12, 0x95, 0x99, 0x8e, 0x80, 0x9f, 0x17, 0xc5, 0x22, 0x57, 0x01, 0x2d, 0xd8, 0x09, 0x40, 0x6e,
	0xf2, 0x29, 0xc5, 0xb0, 0x63, 0xf5, 0x2c, 0xd7, 0x1e, 0xee, 0x78, 0xef, 0x77, 0xf3, 0xc6, 0xaf,
	0xd4, 0x24, 0xc5, 0x90, 0xb7, 0x92, 0xe2, 0x93, 0xed, 0x82, 0x9d, 0xa0, 0xbe, 0x95, 0x53, 0x5f,
	0x04, 0x09, 0x76, 0x6a, 0x3d, 0xcb, 0x6d, 0x71, 0x30, 0xa5, 0xcb, 0x20, 0x41, 0xe7, 0xa9, 0x06,
	0xbf, 0xca, 0x40, 0x8e, 0x94, 0xc5, 0xfa, 0x8b, 0x91, 0x3e, 0xb4, 0xd7, 0xaf, 0xe9, 0xab, 0x5c,
	0x9b, 0x87, 0xdb, 0x43, 0xb7, 0x4a, 0x74, 0xb6, 0xd6, 0x60, 0xc6, 0x38, 0xff, 0xc6, 0xff, 0x86,
	0x15, 0x75, 0x36, 0x81, 0x3f, 0xab, 0xb3, 0x15, 0xf2, 0x7a, 0x2e, 0xdf, 0xaf, 0x92, 0xf3, 0x12,
	0x2e, 0xc5, 0xbf, 0xd5, 0x46, 0x6d, 0xf4, 0x03, 0x9a, 0xc6, 0xe4, 0x3c, 0x5a, 0xd0, 0x1e, 0x67,
	0xb1, 0x8e, 0xde, 0x9c, 0xe5, 0x2e, 0x43, 0xd2, 0xec, 0x18, 0x1a, 0x3a, 0xa0, 0x05, 0x75, 0xac,
	0x5e, 0xdd, 0xb5, 0x87, 0x7b, 0x55, 0x61, 0x6b, 0x3f, 0x8f, 0x1b, 0x9e, 0x0d, 0xa0, 0x91, 0x3f,
	0x99, 0xe5, 0x09, 0xfe, 0x57, 0x37, 0xa6, 0x99, 0xe6, 0x86, 0x73, 0xae, 0xe1, 0xdf, 0xe6, 0x08,
	0x94, 0x4a, 0x41, 0xc8, 0x4e, 0xe1, 0xbb, 0x99, 0xb3, 0x98, 0xa2, 0xff, 0xe1, 0x14, 0x66, 0x3b,
	0x5e, 0xf4, 0x8c, 0xea, 0xcf, 0x96, 0x75, 0xd3, 0xcc, 0x9f, 0xda, 0xd1, 0x4b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdf, 0x63, 0x5f, 0x61, 0x4b, 0x03, 0x00, 0x00,
}
