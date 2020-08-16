# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tensorflow/core/protobuf/rewriter_config.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from tensorflow.core.framework import attr_value_pb2 as tensorflow_dot_core_dot_framework_dot_attr__value__pb2
from tensorflow.core.protobuf import verifier_config_pb2 as tensorflow_dot_core_dot_protobuf_dot_verifier__config__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='tensorflow/core/protobuf/rewriter_config.proto',
  package='tensorflow',
  syntax='proto3',
  serialized_options=b'\n\030org.tensorflow.frameworkB\024RewriterConfigProtosP\001ZHgithub.com/tensorflow/tensorflow/tensorflow/go/core/core_protos_go_proto\370\001\001',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n.tensorflow/core/protobuf/rewriter_config.proto\x12\ntensorflow\x1a*tensorflow/core/framework/attr_value.proto\x1a.tensorflow/core/protobuf/verifier_config.proto\";\n\x13\x41utoParallelOptions\x12\x0e\n\x06\x65nable\x18\x01 \x01(\x08\x12\x14\n\x0cnum_replicas\x18\x02 \x01(\x05\"+\n\x16ScopedAllocatorOptions\x12\x11\n\tenable_op\x18\x01 \x03(\t\"\x95\x11\n\x0eRewriterConfig\x12;\n\x10layout_optimizer\x18\x01 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12;\n\x10\x63onstant_folding\x18\x03 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12=\n\x12shape_optimization\x18\r \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x34\n\tremapping\x18\x0e \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x46\n\x1b\x63ommon_subgraph_elimination\x18\x18 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x42\n\x17\x61rithmetic_optimization\x18\x07 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x42\n\x17\x64\x65pendency_optimization\x18\x08 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12<\n\x11loop_optimization\x18\t \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12@\n\x15\x66unction_optimization\x18\n \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x39\n\x0e\x64\x65\x62ug_stripper\x18\x0b \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x1d\n\x15\x64isable_model_pruning\x18\x02 \x01(\x08\x12H\n\x1dscoped_allocator_optimization\x18\x0f \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x43\n\x18pin_to_host_optimization\x18\x12 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x42\n\x17implementation_selector\x18\x16 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12?\n\x14\x61uto_mixed_precision\x18\x17 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x43\n\x18\x61uto_mixed_precision_mkl\x18\x19 \x01(\x0e\x32!.tensorflow.RewriterConfig.Toggle\x12\x1e\n\x16\x64isable_meta_optimizer\x18\x13 \x01(\x08\x12O\n\x19meta_optimizer_iterations\x18\x0c \x01(\x0e\x32,.tensorflow.RewriterConfig.NumIterationsType\x12\x17\n\x0fmin_graph_nodes\x18\x11 \x01(\x05\x12\x42\n\x13memory_optimization\x18\x04 \x01(\x0e\x32%.tensorflow.RewriterConfig.MemOptType\x12/\n\'memory_optimizer_target_node_name_scope\x18\x06 \x01(\t\x12!\n\x19meta_optimizer_timeout_ms\x18\x14 \x01(\x03\x12\x36\n\rauto_parallel\x18\x05 \x01(\x0b\x32\x1f.tensorflow.AutoParallelOptions\x12 \n\x18\x66\x61il_on_optimizer_errors\x18\x15 \x01(\x08\x12\x41\n\x15scoped_allocator_opts\x18\x10 \x01(\x0b\x32\".tensorflow.ScopedAllocatorOptions\x12\x12\n\noptimizers\x18\x64 \x03(\t\x12K\n\x11\x63ustom_optimizers\x18\xc8\x01 \x03(\x0b\x32/.tensorflow.RewriterConfig.CustomGraphOptimizer\x12\x44\n\x1finter_optimizer_verifier_config\x18\xac\x02 \x01(\x0b\x32\x1a.tensorflow.VerifierConfig\x12\x46\n!post_optimization_verifier_config\x18\xad\x02 \x01(\x0b\x32\x1a.tensorflow.VerifierConfig\x1a\xca\x01\n\x14\x43ustomGraphOptimizer\x12\x0c\n\x04name\x18\x01 \x01(\t\x12X\n\rparameter_map\x18\x02 \x03(\x0b\x32\x41.tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry\x1aJ\n\x11ParameterMapEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12$\n\x05value\x18\x02 \x01(\x0b\x32\x15.tensorflow.AttrValue:\x02\x38\x01\"6\n\x06Toggle\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x06\n\x02ON\x10\x01\x12\x07\n\x03OFF\x10\x02\x12\x0e\n\nAGGRESSIVE\x10\x03\"<\n\x11NumIterationsType\x12\x15\n\x11\x44\x45\x46\x41ULT_NUM_ITERS\x10\x00\x12\x07\n\x03ONE\x10\x01\x12\x07\n\x03TWO\x10\x02\"\x9f\x01\n\nMemOptType\x12\x13\n\x0f\x44\x45\x46\x41ULT_MEM_OPT\x10\x00\x12\x0e\n\nNO_MEM_OPT\x10\x01\x12\n\n\x06MANUAL\x10\x02\x12\x17\n\x13SWAPPING_HEURISTICS\x10\x04\x12\x1c\n\x18RECOMPUTATION_HEURISTICS\x10\x05\x12\x19\n\x15SCHEDULING_HEURISTICS\x10\x06\x12\x0e\n\nHEURISTICS\x10\x03\x42\x7f\n\x18org.tensorflow.frameworkB\x14RewriterConfigProtosP\x01ZHgithub.com/tensorflow/tensorflow/tensorflow/go/core/core_protos_go_proto\xf8\x01\x01\x62\x06proto3'
  ,
  dependencies=[tensorflow_dot_core_dot_framework_dot_attr__value__pb2.DESCRIPTOR,tensorflow_dot_core_dot_protobuf_dot_verifier__config__pb2.DESCRIPTOR,])



_REWRITERCONFIG_TOGGLE = _descriptor.EnumDescriptor(
  name='Toggle',
  full_name='tensorflow.RewriterConfig.Toggle',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DEFAULT', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ON', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OFF', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='AGGRESSIVE', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2180,
  serialized_end=2234,
)
_sym_db.RegisterEnumDescriptor(_REWRITERCONFIG_TOGGLE)

_REWRITERCONFIG_NUMITERATIONSTYPE = _descriptor.EnumDescriptor(
  name='NumIterationsType',
  full_name='tensorflow.RewriterConfig.NumIterationsType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DEFAULT_NUM_ITERS', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ONE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TWO', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2236,
  serialized_end=2296,
)
_sym_db.RegisterEnumDescriptor(_REWRITERCONFIG_NUMITERATIONSTYPE)

_REWRITERCONFIG_MEMOPTTYPE = _descriptor.EnumDescriptor(
  name='MemOptType',
  full_name='tensorflow.RewriterConfig.MemOptType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DEFAULT_MEM_OPT', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NO_MEM_OPT', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='MANUAL', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SWAPPING_HEURISTICS', index=3, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RECOMPUTATION_HEURISTICS', index=4, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SCHEDULING_HEURISTICS', index=5, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HEURISTICS', index=6, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2299,
  serialized_end=2458,
)
_sym_db.RegisterEnumDescriptor(_REWRITERCONFIG_MEMOPTTYPE)


_AUTOPARALLELOPTIONS = _descriptor.Descriptor(
  name='AutoParallelOptions',
  full_name='tensorflow.AutoParallelOptions',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='enable', full_name='tensorflow.AutoParallelOptions.enable', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='num_replicas', full_name='tensorflow.AutoParallelOptions.num_replicas', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=154,
  serialized_end=213,
)


_SCOPEDALLOCATOROPTIONS = _descriptor.Descriptor(
  name='ScopedAllocatorOptions',
  full_name='tensorflow.ScopedAllocatorOptions',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='enable_op', full_name='tensorflow.ScopedAllocatorOptions.enable_op', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=215,
  serialized_end=258,
)


_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY = _descriptor.Descriptor(
  name='ParameterMapEntry',
  full_name='tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2104,
  serialized_end=2178,
)

_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER = _descriptor.Descriptor(
  name='CustomGraphOptimizer',
  full_name='tensorflow.RewriterConfig.CustomGraphOptimizer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='tensorflow.RewriterConfig.CustomGraphOptimizer.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='parameter_map', full_name='tensorflow.RewriterConfig.CustomGraphOptimizer.parameter_map', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1976,
  serialized_end=2178,
)

_REWRITERCONFIG = _descriptor.Descriptor(
  name='RewriterConfig',
  full_name='tensorflow.RewriterConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='layout_optimizer', full_name='tensorflow.RewriterConfig.layout_optimizer', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='constant_folding', full_name='tensorflow.RewriterConfig.constant_folding', index=1,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='shape_optimization', full_name='tensorflow.RewriterConfig.shape_optimization', index=2,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='remapping', full_name='tensorflow.RewriterConfig.remapping', index=3,
      number=14, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='common_subgraph_elimination', full_name='tensorflow.RewriterConfig.common_subgraph_elimination', index=4,
      number=24, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='arithmetic_optimization', full_name='tensorflow.RewriterConfig.arithmetic_optimization', index=5,
      number=7, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='dependency_optimization', full_name='tensorflow.RewriterConfig.dependency_optimization', index=6,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='loop_optimization', full_name='tensorflow.RewriterConfig.loop_optimization', index=7,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='function_optimization', full_name='tensorflow.RewriterConfig.function_optimization', index=8,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='debug_stripper', full_name='tensorflow.RewriterConfig.debug_stripper', index=9,
      number=11, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='disable_model_pruning', full_name='tensorflow.RewriterConfig.disable_model_pruning', index=10,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='scoped_allocator_optimization', full_name='tensorflow.RewriterConfig.scoped_allocator_optimization', index=11,
      number=15, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pin_to_host_optimization', full_name='tensorflow.RewriterConfig.pin_to_host_optimization', index=12,
      number=18, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='implementation_selector', full_name='tensorflow.RewriterConfig.implementation_selector', index=13,
      number=22, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='auto_mixed_precision', full_name='tensorflow.RewriterConfig.auto_mixed_precision', index=14,
      number=23, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='auto_mixed_precision_mkl', full_name='tensorflow.RewriterConfig.auto_mixed_precision_mkl', index=15,
      number=25, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='disable_meta_optimizer', full_name='tensorflow.RewriterConfig.disable_meta_optimizer', index=16,
      number=19, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='meta_optimizer_iterations', full_name='tensorflow.RewriterConfig.meta_optimizer_iterations', index=17,
      number=12, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='min_graph_nodes', full_name='tensorflow.RewriterConfig.min_graph_nodes', index=18,
      number=17, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='memory_optimization', full_name='tensorflow.RewriterConfig.memory_optimization', index=19,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='memory_optimizer_target_node_name_scope', full_name='tensorflow.RewriterConfig.memory_optimizer_target_node_name_scope', index=20,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='meta_optimizer_timeout_ms', full_name='tensorflow.RewriterConfig.meta_optimizer_timeout_ms', index=21,
      number=20, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='auto_parallel', full_name='tensorflow.RewriterConfig.auto_parallel', index=22,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='fail_on_optimizer_errors', full_name='tensorflow.RewriterConfig.fail_on_optimizer_errors', index=23,
      number=21, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='scoped_allocator_opts', full_name='tensorflow.RewriterConfig.scoped_allocator_opts', index=24,
      number=16, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='optimizers', full_name='tensorflow.RewriterConfig.optimizers', index=25,
      number=100, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='custom_optimizers', full_name='tensorflow.RewriterConfig.custom_optimizers', index=26,
      number=200, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='inter_optimizer_verifier_config', full_name='tensorflow.RewriterConfig.inter_optimizer_verifier_config', index=27,
      number=300, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='post_optimization_verifier_config', full_name='tensorflow.RewriterConfig.post_optimization_verifier_config', index=28,
      number=301, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER, ],
  enum_types=[
    _REWRITERCONFIG_TOGGLE,
    _REWRITERCONFIG_NUMITERATIONSTYPE,
    _REWRITERCONFIG_MEMOPTTYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=261,
  serialized_end=2458,
)

_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY.fields_by_name['value'].message_type = tensorflow_dot_core_dot_framework_dot_attr__value__pb2._ATTRVALUE
_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY.containing_type = _REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER
_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER.fields_by_name['parameter_map'].message_type = _REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY
_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER.containing_type = _REWRITERCONFIG
_REWRITERCONFIG.fields_by_name['layout_optimizer'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['constant_folding'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['shape_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['remapping'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['common_subgraph_elimination'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['arithmetic_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['dependency_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['loop_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['function_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['debug_stripper'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['scoped_allocator_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['pin_to_host_optimization'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['implementation_selector'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['auto_mixed_precision'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['auto_mixed_precision_mkl'].enum_type = _REWRITERCONFIG_TOGGLE
_REWRITERCONFIG.fields_by_name['meta_optimizer_iterations'].enum_type = _REWRITERCONFIG_NUMITERATIONSTYPE
_REWRITERCONFIG.fields_by_name['memory_optimization'].enum_type = _REWRITERCONFIG_MEMOPTTYPE
_REWRITERCONFIG.fields_by_name['auto_parallel'].message_type = _AUTOPARALLELOPTIONS
_REWRITERCONFIG.fields_by_name['scoped_allocator_opts'].message_type = _SCOPEDALLOCATOROPTIONS
_REWRITERCONFIG.fields_by_name['custom_optimizers'].message_type = _REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER
_REWRITERCONFIG.fields_by_name['inter_optimizer_verifier_config'].message_type = tensorflow_dot_core_dot_protobuf_dot_verifier__config__pb2._VERIFIERCONFIG
_REWRITERCONFIG.fields_by_name['post_optimization_verifier_config'].message_type = tensorflow_dot_core_dot_protobuf_dot_verifier__config__pb2._VERIFIERCONFIG
_REWRITERCONFIG_TOGGLE.containing_type = _REWRITERCONFIG
_REWRITERCONFIG_NUMITERATIONSTYPE.containing_type = _REWRITERCONFIG
_REWRITERCONFIG_MEMOPTTYPE.containing_type = _REWRITERCONFIG
DESCRIPTOR.message_types_by_name['AutoParallelOptions'] = _AUTOPARALLELOPTIONS
DESCRIPTOR.message_types_by_name['ScopedAllocatorOptions'] = _SCOPEDALLOCATOROPTIONS
DESCRIPTOR.message_types_by_name['RewriterConfig'] = _REWRITERCONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AutoParallelOptions = _reflection.GeneratedProtocolMessageType('AutoParallelOptions', (_message.Message,), {
  'DESCRIPTOR' : _AUTOPARALLELOPTIONS,
  '__module__' : 'tensorflow.core.protobuf.rewriter_config_pb2'
  # @@protoc_insertion_point(class_scope:tensorflow.AutoParallelOptions)
  })
_sym_db.RegisterMessage(AutoParallelOptions)

ScopedAllocatorOptions = _reflection.GeneratedProtocolMessageType('ScopedAllocatorOptions', (_message.Message,), {
  'DESCRIPTOR' : _SCOPEDALLOCATOROPTIONS,
  '__module__' : 'tensorflow.core.protobuf.rewriter_config_pb2'
  # @@protoc_insertion_point(class_scope:tensorflow.ScopedAllocatorOptions)
  })
_sym_db.RegisterMessage(ScopedAllocatorOptions)

RewriterConfig = _reflection.GeneratedProtocolMessageType('RewriterConfig', (_message.Message,), {

  'CustomGraphOptimizer' : _reflection.GeneratedProtocolMessageType('CustomGraphOptimizer', (_message.Message,), {

    'ParameterMapEntry' : _reflection.GeneratedProtocolMessageType('ParameterMapEntry', (_message.Message,), {
      'DESCRIPTOR' : _REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY,
      '__module__' : 'tensorflow.core.protobuf.rewriter_config_pb2'
      # @@protoc_insertion_point(class_scope:tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry)
      })
    ,
    'DESCRIPTOR' : _REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER,
    '__module__' : 'tensorflow.core.protobuf.rewriter_config_pb2'
    # @@protoc_insertion_point(class_scope:tensorflow.RewriterConfig.CustomGraphOptimizer)
    })
  ,
  'DESCRIPTOR' : _REWRITERCONFIG,
  '__module__' : 'tensorflow.core.protobuf.rewriter_config_pb2'
  # @@protoc_insertion_point(class_scope:tensorflow.RewriterConfig)
  })
_sym_db.RegisterMessage(RewriterConfig)
_sym_db.RegisterMessage(RewriterConfig.CustomGraphOptimizer)
_sym_db.RegisterMessage(RewriterConfig.CustomGraphOptimizer.ParameterMapEntry)


DESCRIPTOR._options = None
_REWRITERCONFIG_CUSTOMGRAPHOPTIMIZER_PARAMETERMAPENTRY._options = None
# @@protoc_insertion_point(module_scope)
