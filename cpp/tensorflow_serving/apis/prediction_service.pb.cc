// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: tensorflow_serving/apis/prediction_service.proto

#include "tensorflow_serving/apis/prediction_service.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
namespace tensorflow {
namespace serving {
}  // namespace serving
}  // namespace tensorflow
static constexpr ::PROTOBUF_NAMESPACE_ID::Metadata* file_level_metadata_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::EnumDescriptor const** file_level_enum_descriptors_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::ServiceDescriptor const** file_level_service_descriptors_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto = nullptr;
const ::PROTOBUF_NAMESPACE_ID::uint32 TableStruct_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto::offsets[1] = {};
static constexpr ::PROTOBUF_NAMESPACE_ID::internal::MigrationSchema* schemas = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::Message* const* file_default_instances = nullptr;

const char descriptor_table_protodef_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n0tensorflow_serving/apis/prediction_ser"
  "vice.proto\022\022tensorflow.serving\032,tensorfl"
  "ow_serving/apis/classification.proto\0320te"
  "nsorflow_serving/apis/get_model_metadata"
  ".proto\032\'tensorflow_serving/apis/inferenc"
  "e.proto\032%tensorflow_serving/apis/predict"
  ".proto\032(tensorflow_serving/apis/regressi"
  "on.proto2\374\003\n\021PredictionService\022a\n\010Classi"
  "fy\022).tensorflow.serving.ClassificationRe"
  "quest\032*.tensorflow.serving.Classificatio"
  "nResponse\022X\n\007Regress\022%.tensorflow.servin"
  "g.RegressionRequest\032&.tensorflow.serving"
  ".RegressionResponse\022R\n\007Predict\022\".tensorf"
  "low.serving.PredictRequest\032#.tensorflow."
  "serving.PredictResponse\022g\n\016MultiInferenc"
  "e\022).tensorflow.serving.MultiInferenceReq"
  "uest\032*.tensorflow.serving.MultiInference"
  "Response\022m\n\020GetModelMetadata\022+.tensorflo"
  "w.serving.GetModelMetadataRequest\032,.tens"
  "orflow.serving.GetModelMetadataResponseB"
  "\003\370\001\001b\006proto3"
  ;
static const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable*const descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto_deps[5] = {
  &::descriptor_table_tensorflow_5fserving_2fapis_2fclassification_2eproto,
  &::descriptor_table_tensorflow_5fserving_2fapis_2fget_5fmodel_5fmetadata_2eproto,
  &::descriptor_table_tensorflow_5fserving_2fapis_2finference_2eproto,
  &::descriptor_table_tensorflow_5fserving_2fapis_2fpredict_2eproto,
  &::descriptor_table_tensorflow_5fserving_2fapis_2fregression_2eproto,
};
static ::PROTOBUF_NAMESPACE_ID::internal::SCCInfoBase*const descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto_sccs[1] = {
};
static ::PROTOBUF_NAMESPACE_ID::internal::once_flag descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto_once;
const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto = {
  false, false, descriptor_table_protodef_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto, "tensorflow_serving/apis/prediction_service.proto", 812,
  &descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto_once, descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto_sccs, descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto_deps, 0, 5,
  schemas, file_default_instances, TableStruct_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto::offsets,
  file_level_metadata_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto, 0, file_level_enum_descriptors_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto, file_level_service_descriptors_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto,
};

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto = (static_cast<void>(::PROTOBUF_NAMESPACE_ID::internal::AddDescriptors(&descriptor_table_tensorflow_5fserving_2fapis_2fprediction_5fservice_2eproto)), true);
namespace tensorflow {
namespace serving {

// @@protoc_insertion_point(namespace_scope)
}  // namespace serving
}  // namespace tensorflow
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
