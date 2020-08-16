// This file is generated by rust-protobuf 2.14.0. Do not edit
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![cfg_attr(rustfmt, rustfmt_skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unsafe_code)]
#![allow(unused_imports)]
#![allow(unused_results)]
//! Generated file from `tensorflow_serving/apis/model_service.proto`

use protobuf::Message as Message_imported_for_functions;
use protobuf::ProtobufEnum as ProtobufEnum_imported_for_functions;

/// Generated files are compatible only with the same version
/// of protobuf runtime.
// const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_14_0;

static file_descriptor_proto_data: &'static [u8] = b"\
    \n+tensorflow_serving/apis/model_service.proto\x12\x12tensorflow.serving\
    \x1a.tensorflow_serving/apis/get_model_status.proto\x1a.tensorflow_servi\
    ng/apis/model_management.proto2\xe7\x01\n\x0cModelService\x12g\n\x0eGetM\
    odelStatus\x12).tensorflow.serving.GetModelStatusRequest\x1a*.tensorflow\
    .serving.GetModelStatusResponse\x12n\n\x19HandleReloadConfigRequest\x12'\
    .tensorflow.serving.ReloadConfigRequest\x1a(.tensorflow.serving.ReloadCo\
    nfigResponseB\x03\xf8\x01\x01J\xf2\x05\n\x06\x12\x04\0\0\x17\x01\n\x08\n\
    \x01\x0c\x12\x03\0\0\x12\n\x08\n\x01\x08\x12\x03\x02\0\x1f\n\t\n\x02\x08\
    \x1f\x12\x03\x02\0\x1f\n\t\n\x02\x03\0\x12\x03\x04\08\n\t\n\x02\x03\x01\
    \x12\x03\x05\08\n\x08\n\x01\x02\x12\x03\x07\0\x1b\n\x86\x01\n\x02\x06\0\
    \x12\x04\x0b\0\x17\x01\x1az\x20ModelService\x20provides\x20methods\x20to\
    \x20query\x20and\x20update\x20the\x20state\x20of\x20the\x20server,\n\x20\
    e.g.\x20which\x20models/versions\x20are\x20being\x20served.\n\n\n\n\x03\
    \x06\0\x01\x12\x03\x0b\x08\x14\n\x8a\x02\n\x04\x06\0\x02\0\x12\x03\x10\
    \x02M\x1a\xfc\x01\x20Gets\x20status\x20of\x20model.\x20If\x20the\x20Mode\
    lSpec\x20in\x20the\x20request\x20does\x20not\x20specify\n\x20version,\
    \x20information\x20about\x20all\x20versions\x20of\x20the\x20model\x20wil\
    l\x20be\x20returned.\x20If\n\x20the\x20ModelSpec\x20in\x20the\x20request\
    \x20does\x20specify\x20a\x20version,\x20the\x20status\x20of\x20only\n\
    \x20that\x20version\x20will\x20be\x20returned.\n\n\x0c\n\x05\x06\0\x02\0\
    \x01\x12\x03\x10\x06\x14\n\x0c\n\x05\x06\0\x02\0\x02\x12\x03\x10\x15*\n\
    \x0c\n\x05\x06\0\x02\0\x03\x12\x03\x105K\n\xb2\x01\n\x04\x06\0\x02\x01\
    \x12\x04\x15\x02\x16%\x1a\xa3\x01\x20Reloads\x20the\x20set\x20of\x20serv\
    ed\x20models.\x20The\x20new\x20config\x20supersedes\x20the\x20old\x20one\
    ,\n\x20so\x20if\x20a\x20model\x20is\x20omitted\x20from\x20the\x20new\x20\
    config\x20it\x20will\x20be\x20unloaded\x20and\x20no\n\x20longer\x20serve\
    d.\n\n\x0c\n\x05\x06\0\x02\x01\x01\x12\x03\x15\x06\x1f\n\x0c\n\x05\x06\0\
    \x02\x01\x02\x12\x03\x15\x203\n\x0c\n\x05\x06\0\x02\x01\x03\x12\x03\x16\
    \x0f#b\x06proto3\
";

static mut file_descriptor_proto_lazy: ::protobuf::lazy::Lazy<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::lazy::Lazy::INIT;

fn parse_descriptor_proto() -> ::protobuf::descriptor::FileDescriptorProto {
    ::protobuf::parse_from_bytes(file_descriptor_proto_data).unwrap()
}

pub fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    unsafe {
        file_descriptor_proto_lazy.get(|| {
            parse_descriptor_proto()
        })
    }
}
