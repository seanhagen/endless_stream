//
//  Generated code. Do not modify.
//  source: common/logs.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use logLevelDescriptor instead')
const LogLevel$json = {
  '1': 'LogLevel',
  '2': [
    {'1': 'Info', '2': 0},
    {'1': 'Debug', '2': -1},
    {'1': 'Warn', '2': 2},
    {'1': 'Error', '2': 3},
    {'1': 'Fatal', '2': 4},
  ],
};

/// Descriptor for `LogLevel`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List logLevelDescriptor = $convert.base64Decode(
    'CghMb2dMZXZlbBIICgRJbmZvEAASEgoFRGVidWcQ////////////ARIICgRXYXJuEAISCQoFRX'
    'Jyb3IQAxIJCgVGYXRhbBAE');

@$core.Deprecated('Use logSourceDescriptor instead')
const LogSource$json = {
  '1': 'LogSource',
  '2': [
    {'1': 'Server', '2': 0},
    {'1': 'Player', '2': 100},
    {'1': 'Game', '2': 200},
  ],
};

/// Descriptor for `LogSource`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List logSourceDescriptor = $convert.base64Decode(
    'CglMb2dTb3VyY2USCgoGU2VydmVyEAASCgoGUGxheWVyEGQSCQoER2FtZRDIAQ==');

@$core.Deprecated('Use logDescriptor instead')
const Log$json = {
  '1': 'Log',
  '2': [
    {'1': 'msg', '3': 1, '4': 1, '5': 9, '10': 'msg'},
    {'1': 'at', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'at'},
    {'1': 'level', '3': 3, '4': 1, '5': 14, '6': '.endless.LogLevel', '10': 'level'},
    {'1': 'source', '3': 4, '4': 1, '5': 14, '6': '.endless.LogSource', '10': 'source'},
  ],
};

/// Descriptor for `Log`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logDescriptor = $convert.base64Decode(
    'CgNMb2cSEAoDbXNnGAEgASgJUgNtc2cSKgoCYXQYAiABKAsyGi5nb29nbGUucHJvdG9idWYuVG'
    'ltZXN0YW1wUgJhdBInCgVsZXZlbBgDIAEoDjIRLmVuZGxlc3MuTG9nTGV2ZWxSBWxldmVsEioK'
    'BnNvdXJjZRgEIAEoDjISLmVuZGxlc3MuTG9nU291cmNlUgZzb3VyY2U=');
