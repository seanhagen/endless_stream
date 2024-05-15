//
//  Generated code. Do not modify.
//  source: test/test_types.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pingReqDescriptor instead')
const PingReq$json = {
  '1': 'PingReq',
  '2': [
    {'1': 'msg', '3': 1, '4': 1, '5': 9, '10': 'msg'},
  ],
};

/// Descriptor for `PingReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pingReqDescriptor = $convert.base64Decode(
    'CgdQaW5nUmVxEhAKA21zZxgBIAEoCVIDbXNn');

@$core.Deprecated('Use pongRespDescriptor instead')
const PongResp$json = {
  '1': 'PongResp',
  '2': [
    {'1': 'gsm', '3': 1, '4': 1, '5': 9, '10': 'gsm'},
  ],
};

/// Descriptor for `PongResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pongRespDescriptor = $convert.base64Decode(
    'CghQb25nUmVzcBIQCgNnc20YASABKAlSA2dzbQ==');

@$core.Deprecated('Use testStreamRequestDescriptor instead')
const TestStreamRequest$json = {
  '1': 'TestStreamRequest',
  '2': [
    {'1': 'chunk_id', '3': 1, '4': 1, '5': 5, '10': 'chunkId'},
    {'1': 'msg', '3': 2, '4': 1, '5': 9, '10': 'msg'},
  ],
};

/// Descriptor for `TestStreamRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List testStreamRequestDescriptor = $convert.base64Decode(
    'ChFUZXN0U3RyZWFtUmVxdWVzdBIZCghjaHVua19pZBgBIAEoBVIHY2h1bmtJZBIQCgNtc2cYAi'
    'ABKAlSA21zZw==');

@$core.Deprecated('Use testStreamResponseDescriptor instead')
const TestStreamResponse$json = {
  '1': 'TestStreamResponse',
  '2': [
    {'1': 'resp_id', '3': 1, '4': 1, '5': 5, '10': 'respId'},
    {'1': 'gsm', '3': 2, '4': 1, '5': 9, '10': 'gsm'},
  ],
};

/// Descriptor for `TestStreamResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List testStreamResponseDescriptor = $convert.base64Decode(
    'ChJUZXN0U3RyZWFtUmVzcG9uc2USFwoHcmVzcF9pZBgBIAEoBVIGcmVzcElkEhAKA2dzbRgCIA'
    'EoCVIDZ3Nt');

@$core.Deprecated('Use testRequestDescriptor instead')
const TestRequest$json = {
  '1': 'TestRequest',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `TestRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List testRequestDescriptor = $convert.base64Decode(
    'CgtUZXN0UmVxdWVzdBISCgRuYW1lGAEgASgJUgRuYW1l');

@$core.Deprecated('Use testResponseDescriptor instead')
const TestResponse$json = {
  '1': 'TestResponse',
  '2': [
    {'1': 'resp', '3': 1, '4': 1, '5': 9, '10': 'resp'},
  ],
};

/// Descriptor for `TestResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List testResponseDescriptor = $convert.base64Decode(
    'CgxUZXN0UmVzcG9uc2USEgoEcmVzcBgBIAEoCVIEcmVzcA==');
