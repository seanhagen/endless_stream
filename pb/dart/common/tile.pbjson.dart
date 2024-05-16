//
//  Generated code. Do not modify.
//  source: common/tile.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use typeDescriptor instead')
const Type$json = {
  '1': 'Type',
  '2': [
    {'1': 'Empty', '2': 0},
    {'1': 'Floor', '2': 1},
    {'1': 'Wall', '2': 2},
  ],
};

/// Descriptor for `Type`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List typeDescriptor = $convert.base64Decode(
    'CgRUeXBlEgkKBUVtcHR5EAASCQoFRmxvb3IQARIICgRXYWxsEAI=');

@$core.Deprecated('Use tilesetDescriptor instead')
const Tileset$json = {
  '1': 'Tileset',
  '2': [
    {'1': 'Dungeon', '2': 0},
    {'1': 'Woods', '2': 1},
  ],
};

/// Descriptor for `Tileset`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List tilesetDescriptor = $convert.base64Decode(
    'CgdUaWxlc2V0EgsKB0R1bmdlb24QABIJCgVXb29kcxAB');

@$core.Deprecated('Use coordinateDescriptor instead')
const Coordinate$json = {
  '1': 'Coordinate',
  '2': [
    {'1': 'x', '3': 1, '4': 1, '5': 5, '10': 'x'},
    {'1': 'y', '3': 2, '4': 1, '5': 5, '10': 'y'},
  ],
};

/// Descriptor for `Coordinate`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List coordinateDescriptor = $convert.base64Decode(
    'CgpDb29yZGluYXRlEgwKAXgYASABKAVSAXgSDAoBeRgCIAEoBVIBeQ==');

@$core.Deprecated('Use tileDescriptor instead')
const Tile$json = {
  '1': 'Tile',
  '2': [
    {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.endless.Type', '10': 'type'},
    {'1': 'coords', '3': 2, '4': 1, '5': 11, '6': '.endless.Coordinate', '10': 'coords'},
  ],
};

/// Descriptor for `Tile`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tileDescriptor = $convert.base64Decode(
    'CgRUaWxlEiEKBHR5cGUYASABKA4yDS5lbmRsZXNzLlR5cGVSBHR5cGUSKwoGY29vcmRzGAIgAS'
    'gLMhMuZW5kbGVzcy5Db29yZGluYXRlUgZjb29yZHM=');
