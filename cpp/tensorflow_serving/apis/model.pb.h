// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: tensorflow_serving/apis/model.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_tensorflow_5fserving_2fapis_2fmodel_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_tensorflow_5fserving_2fapis_2fmodel_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3012000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3012000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include <google/protobuf/wrappers.pb.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_tensorflow_5fserving_2fapis_2fmodel_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_tensorflow_5fserving_2fapis_2fmodel_2eproto {
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::FieldMetadata field_metadata[];
  static const ::PROTOBUF_NAMESPACE_ID::internal::SerializationTable serialization_table[];
  static const ::PROTOBUF_NAMESPACE_ID::uint32 offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_tensorflow_5fserving_2fapis_2fmodel_2eproto;
namespace tensorflow {
namespace serving {
class ModelSpec;
class ModelSpecDefaultTypeInternal;
extern ModelSpecDefaultTypeInternal _ModelSpec_default_instance_;
}  // namespace serving
}  // namespace tensorflow
PROTOBUF_NAMESPACE_OPEN
template<> ::tensorflow::serving::ModelSpec* Arena::CreateMaybeMessage<::tensorflow::serving::ModelSpec>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace tensorflow {
namespace serving {

// ===================================================================

class ModelSpec PROTOBUF_FINAL :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:tensorflow.serving.ModelSpec) */ {
 public:
  inline ModelSpec() : ModelSpec(nullptr) {};
  virtual ~ModelSpec();

  ModelSpec(const ModelSpec& from);
  ModelSpec(ModelSpec&& from) noexcept
    : ModelSpec() {
    *this = ::std::move(from);
  }

  inline ModelSpec& operator=(const ModelSpec& from) {
    CopyFrom(from);
    return *this;
  }
  inline ModelSpec& operator=(ModelSpec&& from) noexcept {
    if (GetArena() == from.GetArena()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return GetMetadataStatic().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return GetMetadataStatic().reflection;
  }
  static const ModelSpec& default_instance();

  enum VersionChoiceCase {
    kVersion = 2,
    kVersionLabel = 4,
    VERSION_CHOICE_NOT_SET = 0,
  };

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const ModelSpec* internal_default_instance() {
    return reinterpret_cast<const ModelSpec*>(
               &_ModelSpec_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(ModelSpec& a, ModelSpec& b) {
    a.Swap(&b);
  }
  inline void Swap(ModelSpec* other) {
    if (other == this) return;
    if (GetArena() == other->GetArena()) {
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(ModelSpec* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  inline ModelSpec* New() const final {
    return CreateMaybeMessage<ModelSpec>(nullptr);
  }

  ModelSpec* New(::PROTOBUF_NAMESPACE_ID::Arena* arena) const final {
    return CreateMaybeMessage<ModelSpec>(arena);
  }
  void CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) final;
  void CopyFrom(const ModelSpec& from);
  void MergeFrom(const ModelSpec& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  ::PROTOBUF_NAMESPACE_ID::uint8* _InternalSerialize(
      ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  inline void SharedCtor();
  inline void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(ModelSpec* other);
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "tensorflow.serving.ModelSpec";
  }
  protected:
  explicit ModelSpec(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;
  private:
  static ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadataStatic() {
    ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&::descriptor_table_tensorflow_5fserving_2fapis_2fmodel_2eproto);
    return ::descriptor_table_tensorflow_5fserving_2fapis_2fmodel_2eproto.file_level_metadata[kIndexInFileMessages];
  }

  public:

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kNameFieldNumber = 1,
    kSignatureNameFieldNumber = 3,
    kVersionFieldNumber = 2,
    kVersionLabelFieldNumber = 4,
  };
  // string name = 1;
  void clear_name();
  const std::string& name() const;
  void set_name(const std::string& value);
  void set_name(std::string&& value);
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  std::string* mutable_name();
  std::string* release_name();
  void set_allocated_name(std::string* name);
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  std::string* unsafe_arena_release_name();
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  void unsafe_arena_set_allocated_name(
      std::string* name);
  private:
  const std::string& _internal_name() const;
  void _internal_set_name(const std::string& value);
  std::string* _internal_mutable_name();
  public:

  // string signature_name = 3;
  void clear_signature_name();
  const std::string& signature_name() const;
  void set_signature_name(const std::string& value);
  void set_signature_name(std::string&& value);
  void set_signature_name(const char* value);
  void set_signature_name(const char* value, size_t size);
  std::string* mutable_signature_name();
  std::string* release_signature_name();
  void set_allocated_signature_name(std::string* signature_name);
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  std::string* unsafe_arena_release_signature_name();
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  void unsafe_arena_set_allocated_signature_name(
      std::string* signature_name);
  private:
  const std::string& _internal_signature_name() const;
  void _internal_set_signature_name(const std::string& value);
  std::string* _internal_mutable_signature_name();
  public:

  // .google.protobuf.Int64Value version = 2;
  bool has_version() const;
  private:
  bool _internal_has_version() const;
  public:
  void clear_version();
  const PROTOBUF_NAMESPACE_ID::Int64Value& version() const;
  PROTOBUF_NAMESPACE_ID::Int64Value* release_version();
  PROTOBUF_NAMESPACE_ID::Int64Value* mutable_version();
  void set_allocated_version(PROTOBUF_NAMESPACE_ID::Int64Value* version);
  private:
  const PROTOBUF_NAMESPACE_ID::Int64Value& _internal_version() const;
  PROTOBUF_NAMESPACE_ID::Int64Value* _internal_mutable_version();
  public:
  void unsafe_arena_set_allocated_version(
      PROTOBUF_NAMESPACE_ID::Int64Value* version);
  PROTOBUF_NAMESPACE_ID::Int64Value* unsafe_arena_release_version();

  // string version_label = 4;
  private:
  bool _internal_has_version_label() const;
  public:
  void clear_version_label();
  const std::string& version_label() const;
  void set_version_label(const std::string& value);
  void set_version_label(std::string&& value);
  void set_version_label(const char* value);
  void set_version_label(const char* value, size_t size);
  std::string* mutable_version_label();
  std::string* release_version_label();
  void set_allocated_version_label(std::string* version_label);
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  std::string* unsafe_arena_release_version_label();
  GOOGLE_PROTOBUF_RUNTIME_DEPRECATED("The unsafe_arena_ accessors for"
  "    string fields are deprecated and will be removed in a"
  "    future release.")
  void unsafe_arena_set_allocated_version_label(
      std::string* version_label);
  private:
  const std::string& _internal_version_label() const;
  void _internal_set_version_label(const std::string& value);
  std::string* _internal_mutable_version_label();
  public:

  void clear_version_choice();
  VersionChoiceCase version_choice_case() const;
  // @@protoc_insertion_point(class_scope:tensorflow.serving.ModelSpec)
 private:
  class _Internal;
  void set_has_version();
  void set_has_version_label();

  inline bool has_version_choice() const;
  inline void clear_has_version_choice();

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr name_;
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr signature_name_;
  union VersionChoiceUnion {
    VersionChoiceUnion() {}
    PROTOBUF_NAMESPACE_ID::Int64Value* version_;
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr version_label_;
  } version_choice_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  ::PROTOBUF_NAMESPACE_ID::uint32 _oneof_case_[1];

  friend struct ::TableStruct_tensorflow_5fserving_2fapis_2fmodel_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// ModelSpec

// string name = 1;
inline void ModelSpec::clear_name() {
  name_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline const std::string& ModelSpec::name() const {
  // @@protoc_insertion_point(field_get:tensorflow.serving.ModelSpec.name)
  return _internal_name();
}
inline void ModelSpec::set_name(const std::string& value) {
  _internal_set_name(value);
  // @@protoc_insertion_point(field_set:tensorflow.serving.ModelSpec.name)
}
inline std::string* ModelSpec::mutable_name() {
  // @@protoc_insertion_point(field_mutable:tensorflow.serving.ModelSpec.name)
  return _internal_mutable_name();
}
inline const std::string& ModelSpec::_internal_name() const {
  return name_.Get();
}
inline void ModelSpec::_internal_set_name(const std::string& value) {
  
  name_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value, GetArena());
}
inline void ModelSpec::set_name(std::string&& value) {
  
  name_.Set(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:tensorflow.serving.ModelSpec.name)
}
inline void ModelSpec::set_name(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  name_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value),
              GetArena());
  // @@protoc_insertion_point(field_set_char:tensorflow.serving.ModelSpec.name)
}
inline void ModelSpec::set_name(const char* value,
    size_t size) {
  
  name_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size), GetArena());
  // @@protoc_insertion_point(field_set_pointer:tensorflow.serving.ModelSpec.name)
}
inline std::string* ModelSpec::_internal_mutable_name() {
  
  return name_.Mutable(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline std::string* ModelSpec::release_name() {
  // @@protoc_insertion_point(field_release:tensorflow.serving.ModelSpec.name)
  return name_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline void ModelSpec::set_allocated_name(std::string* name) {
  if (name != nullptr) {
    
  } else {
    
  }
  name_.SetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), name,
      GetArena());
  // @@protoc_insertion_point(field_set_allocated:tensorflow.serving.ModelSpec.name)
}
inline std::string* ModelSpec::unsafe_arena_release_name() {
  // @@protoc_insertion_point(field_unsafe_arena_release:tensorflow.serving.ModelSpec.name)
  GOOGLE_DCHECK(GetArena() != nullptr);
  
  return name_.UnsafeArenaRelease(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      GetArena());
}
inline void ModelSpec::unsafe_arena_set_allocated_name(
    std::string* name) {
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (name != nullptr) {
    
  } else {
    
  }
  name_.UnsafeArenaSetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      name, GetArena());
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:tensorflow.serving.ModelSpec.name)
}

// .google.protobuf.Int64Value version = 2;
inline bool ModelSpec::_internal_has_version() const {
  return version_choice_case() == kVersion;
}
inline bool ModelSpec::has_version() const {
  return _internal_has_version();
}
inline void ModelSpec::set_has_version() {
  _oneof_case_[0] = kVersion;
}
inline PROTOBUF_NAMESPACE_ID::Int64Value* ModelSpec::release_version() {
  // @@protoc_insertion_point(field_release:tensorflow.serving.ModelSpec.version)
  if (_internal_has_version()) {
    clear_has_version_choice();
      PROTOBUF_NAMESPACE_ID::Int64Value* temp = version_choice_.version_;
    if (GetArena() != nullptr) {
      temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
    }
    version_choice_.version_ = nullptr;
    return temp;
  } else {
    return nullptr;
  }
}
inline const PROTOBUF_NAMESPACE_ID::Int64Value& ModelSpec::_internal_version() const {
  return _internal_has_version()
      ? *version_choice_.version_
      : *reinterpret_cast< PROTOBUF_NAMESPACE_ID::Int64Value*>(&PROTOBUF_NAMESPACE_ID::_Int64Value_default_instance_);
}
inline const PROTOBUF_NAMESPACE_ID::Int64Value& ModelSpec::version() const {
  // @@protoc_insertion_point(field_get:tensorflow.serving.ModelSpec.version)
  return _internal_version();
}
inline PROTOBUF_NAMESPACE_ID::Int64Value* ModelSpec::unsafe_arena_release_version() {
  // @@protoc_insertion_point(field_unsafe_arena_release:tensorflow.serving.ModelSpec.version)
  if (_internal_has_version()) {
    clear_has_version_choice();
    PROTOBUF_NAMESPACE_ID::Int64Value* temp = version_choice_.version_;
    version_choice_.version_ = nullptr;
    return temp;
  } else {
    return nullptr;
  }
}
inline void ModelSpec::unsafe_arena_set_allocated_version(PROTOBUF_NAMESPACE_ID::Int64Value* version) {
  clear_version_choice();
  if (version) {
    set_has_version();
    version_choice_.version_ = version;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:tensorflow.serving.ModelSpec.version)
}
inline PROTOBUF_NAMESPACE_ID::Int64Value* ModelSpec::_internal_mutable_version() {
  if (!_internal_has_version()) {
    clear_version_choice();
    set_has_version();
    version_choice_.version_ = CreateMaybeMessage< PROTOBUF_NAMESPACE_ID::Int64Value >(GetArena());
  }
  return version_choice_.version_;
}
inline PROTOBUF_NAMESPACE_ID::Int64Value* ModelSpec::mutable_version() {
  // @@protoc_insertion_point(field_mutable:tensorflow.serving.ModelSpec.version)
  return _internal_mutable_version();
}

// string version_label = 4;
inline bool ModelSpec::_internal_has_version_label() const {
  return version_choice_case() == kVersionLabel;
}
inline void ModelSpec::set_has_version_label() {
  _oneof_case_[0] = kVersionLabel;
}
inline void ModelSpec::clear_version_label() {
  if (_internal_has_version_label()) {
    version_choice_.version_label_.Destroy(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
    clear_has_version_choice();
  }
}
inline const std::string& ModelSpec::version_label() const {
  // @@protoc_insertion_point(field_get:tensorflow.serving.ModelSpec.version_label)
  return _internal_version_label();
}
inline void ModelSpec::set_version_label(const std::string& value) {
  _internal_set_version_label(value);
  // @@protoc_insertion_point(field_set:tensorflow.serving.ModelSpec.version_label)
}
inline std::string* ModelSpec::mutable_version_label() {
  // @@protoc_insertion_point(field_mutable:tensorflow.serving.ModelSpec.version_label)
  return _internal_mutable_version_label();
}
inline const std::string& ModelSpec::_internal_version_label() const {
  if (_internal_has_version_label()) {
    return version_choice_.version_label_.Get();
  }
  return *&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited();
}
inline void ModelSpec::_internal_set_version_label(const std::string& value) {
  if (!_internal_has_version_label()) {
    clear_version_choice();
    set_has_version_label();
    version_choice_.version_label_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  }
  version_choice_.version_label_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value, GetArena());
}
inline void ModelSpec::set_version_label(std::string&& value) {
  // @@protoc_insertion_point(field_set:tensorflow.serving.ModelSpec.version_label)
  if (!_internal_has_version_label()) {
    clear_version_choice();
    set_has_version_label();
    version_choice_.version_label_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  }
  version_choice_.version_label_.Set(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:tensorflow.serving.ModelSpec.version_label)
}
inline void ModelSpec::set_version_label(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  if (!_internal_has_version_label()) {
    clear_version_choice();
    set_has_version_label();
    version_choice_.version_label_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  }
  version_choice_.version_label_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      ::std::string(value), GetArena());
  // @@protoc_insertion_point(field_set_char:tensorflow.serving.ModelSpec.version_label)
}
inline void ModelSpec::set_version_label(const char* value,
                             size_t size) {
  if (!_internal_has_version_label()) {
    clear_version_choice();
    set_has_version_label();
    version_choice_.version_label_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  }
  version_choice_.version_label_.Set(
      &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size),
      GetArena());
  // @@protoc_insertion_point(field_set_pointer:tensorflow.serving.ModelSpec.version_label)
}
inline std::string* ModelSpec::_internal_mutable_version_label() {
  if (!_internal_has_version_label()) {
    clear_version_choice();
    set_has_version_label();
    version_choice_.version_label_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  }
  return version_choice_.version_label_.Mutable(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline std::string* ModelSpec::release_version_label() {
  // @@protoc_insertion_point(field_release:tensorflow.serving.ModelSpec.version_label)
  if (_internal_has_version_label()) {
    clear_has_version_choice();
    return version_choice_.version_label_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  } else {
    return nullptr;
  }
}
inline void ModelSpec::set_allocated_version_label(std::string* version_label) {
  if (has_version_choice()) {
    clear_version_choice();
  }
  if (version_label != nullptr) {
    set_has_version_label();
    version_choice_.version_label_.UnsafeSetDefault(version_label);
  }
  // @@protoc_insertion_point(field_set_allocated:tensorflow.serving.ModelSpec.version_label)
}
inline std::string* ModelSpec::unsafe_arena_release_version_label() {
  // @@protoc_insertion_point(field_unsafe_arena_release:tensorflow.serving.ModelSpec.version_label)
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (_internal_has_version_label()) {
    clear_has_version_choice();
    return version_choice_.version_label_.UnsafeArenaRelease(
        &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  } else {
    return nullptr;
  }
}
inline void ModelSpec::unsafe_arena_set_allocated_version_label(std::string* version_label) {
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (!_internal_has_version_label()) {
    version_choice_.version_label_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  }
  clear_version_choice();
  if (version_label) {
    set_has_version_label();
    version_choice_.version_label_.UnsafeArenaSetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), version_label, GetArena());
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:tensorflow.serving.ModelSpec.version_label)
}

// string signature_name = 3;
inline void ModelSpec::clear_signature_name() {
  signature_name_.ClearToEmpty(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline const std::string& ModelSpec::signature_name() const {
  // @@protoc_insertion_point(field_get:tensorflow.serving.ModelSpec.signature_name)
  return _internal_signature_name();
}
inline void ModelSpec::set_signature_name(const std::string& value) {
  _internal_set_signature_name(value);
  // @@protoc_insertion_point(field_set:tensorflow.serving.ModelSpec.signature_name)
}
inline std::string* ModelSpec::mutable_signature_name() {
  // @@protoc_insertion_point(field_mutable:tensorflow.serving.ModelSpec.signature_name)
  return _internal_mutable_signature_name();
}
inline const std::string& ModelSpec::_internal_signature_name() const {
  return signature_name_.Get();
}
inline void ModelSpec::_internal_set_signature_name(const std::string& value) {
  
  signature_name_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), value, GetArena());
}
inline void ModelSpec::set_signature_name(std::string&& value) {
  
  signature_name_.Set(
    &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::move(value), GetArena());
  // @@protoc_insertion_point(field_set_rvalue:tensorflow.serving.ModelSpec.signature_name)
}
inline void ModelSpec::set_signature_name(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  signature_name_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(value),
              GetArena());
  // @@protoc_insertion_point(field_set_char:tensorflow.serving.ModelSpec.signature_name)
}
inline void ModelSpec::set_signature_name(const char* value,
    size_t size) {
  
  signature_name_.Set(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), ::std::string(
      reinterpret_cast<const char*>(value), size), GetArena());
  // @@protoc_insertion_point(field_set_pointer:tensorflow.serving.ModelSpec.signature_name)
}
inline std::string* ModelSpec::_internal_mutable_signature_name() {
  
  return signature_name_.Mutable(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline std::string* ModelSpec::release_signature_name() {
  // @@protoc_insertion_point(field_release:tensorflow.serving.ModelSpec.signature_name)
  return signature_name_.Release(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}
inline void ModelSpec::set_allocated_signature_name(std::string* signature_name) {
  if (signature_name != nullptr) {
    
  } else {
    
  }
  signature_name_.SetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), signature_name,
      GetArena());
  // @@protoc_insertion_point(field_set_allocated:tensorflow.serving.ModelSpec.signature_name)
}
inline std::string* ModelSpec::unsafe_arena_release_signature_name() {
  // @@protoc_insertion_point(field_unsafe_arena_release:tensorflow.serving.ModelSpec.signature_name)
  GOOGLE_DCHECK(GetArena() != nullptr);
  
  return signature_name_.UnsafeArenaRelease(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      GetArena());
}
inline void ModelSpec::unsafe_arena_set_allocated_signature_name(
    std::string* signature_name) {
  GOOGLE_DCHECK(GetArena() != nullptr);
  if (signature_name != nullptr) {
    
  } else {
    
  }
  signature_name_.UnsafeArenaSetAllocated(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(),
      signature_name, GetArena());
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:tensorflow.serving.ModelSpec.signature_name)
}

inline bool ModelSpec::has_version_choice() const {
  return version_choice_case() != VERSION_CHOICE_NOT_SET;
}
inline void ModelSpec::clear_has_version_choice() {
  _oneof_case_[0] = VERSION_CHOICE_NOT_SET;
}
inline ModelSpec::VersionChoiceCase ModelSpec::version_choice_case() const {
  return ModelSpec::VersionChoiceCase(_oneof_case_[0]);
}
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace serving
}  // namespace tensorflow

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_tensorflow_5fserving_2fapis_2fmodel_2eproto
