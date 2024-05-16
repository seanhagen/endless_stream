//
//  Generated code. Do not modify.
//  source: test/test.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import 'test_types.pbjson.dart' as $5;

const $core.Map<$core.String, $core.dynamic> TestServiceBase$json = {
  '1': 'Test',
  '2': [
    {'1': 'Ping', '2': '.endless.PingReq', '3': '.endless.PongResp', '4': {}},
    {'1': 'ClientStream', '2': '.endless.TestStreamRequest', '3': '.endless.TestResponse', '5': true},
    {'1': 'ServerStream', '2': '.endless.TestRequest', '3': '.endless.TestStreamResponse', '6': true},
    {'1': 'BiDiStream', '2': '.endless.TestStreamRequest', '3': '.endless.TestStreamResponse', '5': true},
  ],
};

@$core.Deprecated('Use testServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> TestServiceBase$messageJson = {
  '.endless.PingReq': $5.PingReq$json,
  '.endless.PongResp': $5.PongResp$json,
  '.endless.TestStreamRequest': $5.TestStreamRequest$json,
  '.endless.TestResponse': $5.TestResponse$json,
  '.endless.TestRequest': $5.TestRequest$json,
  '.endless.TestStreamResponse': $5.TestStreamResponse$json,
};

/// Descriptor for `Test`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List testServiceDescriptor = $convert.base64Decode(
    'CgRUZXN0EkAKBFBpbmcSEC5lbmRsZXNzLlBpbmdSZXEaES5lbmRsZXNzLlBvbmdSZXNwIhOC0+'
    'STAg06ASoiCC92MS9waW5nEkMKDENsaWVudFN0cmVhbRIaLmVuZGxlc3MuVGVzdFN0cmVhbVJl'
    'cXVlc3QaFS5lbmRsZXNzLlRlc3RSZXNwb25zZSgBEkMKDFNlcnZlclN0cmVhbRIULmVuZGxlc3'
    'MuVGVzdFJlcXVlc3QaGy5lbmRsZXNzLlRlc3RTdHJlYW1SZXNwb25zZTABEkcKCkJpRGlTdHJl'
    'YW0SGi5lbmRsZXNzLlRlc3RTdHJlYW1SZXF1ZXN0GhsuZW5kbGVzcy5UZXN0U3RyZWFtUmVzcG'
    '9uc2UoAQ==');
