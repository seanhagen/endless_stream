//
//  Generated code. Do not modify.
//  source: proto/hex_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import '../google/protobuf/timestamp.pbjson.dart' as $0;
import 'hex.pbjson.dart' as $1;

const $core.Map<$core.String, $core.dynamic> HexServiceBase$json = {
  '1': 'Hex',
  '2': [
    {'1': 'Info', '2': '.endless.InfoRequest', '3': '.endless.InfoResponse'},
    {'1': 'Game', '2': '.endless.GameRequest', '3': '.endless.GameResponse', '5': true, '6': true},
  ],
};

@$core.Deprecated('Use hexServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> HexServiceBase$messageJson = {
  '.endless.InfoRequest': $1.InfoRequest$json,
  '.endless.InfoResponse': $1.InfoResponse$json,
  '.endless.GameRequest': $1.GameRequest$json,
  '.endless.GetLevel': $1.GetLevel$json,
  '.endless.GameResponse': $1.GameResponse$json,
  '.endless.Log': $1.Log$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
  '.endless.Level': $1.Level$json,
  '.endless.Tile': $1.Tile$json,
  '.endless.Coordinate': $1.Coordinate$json,
};

/// Descriptor for `Hex`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List hexServiceDescriptor = $convert.base64Decode(
    'CgNIZXgSMwoESW5mbxIULmVuZGxlc3MuSW5mb1JlcXVlc3QaFS5lbmRsZXNzLkluZm9SZXNwb2'
    '5zZRI3CgRHYW1lEhQuZW5kbGVzcy5HYW1lUmVxdWVzdBoVLmVuZGxlc3MuR2FtZVJlc3BvbnNl'
    'KAEwAQ==');

const $core.Map<$core.String, $core.dynamic> AdminServiceBase$json = {
  '1': 'Admin',
  '2': [
    {'1': 'Manage', '2': '.endless.AdminRequest', '3': '.endless.AdminResponse', '5': true, '6': true},
  ],
};

@$core.Deprecated('Use adminServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AdminServiceBase$messageJson = {
  '.endless.AdminRequest': $1.AdminRequest$json,
  '.endless.AddTile': $1.AddTile$json,
  '.endless.Tile': $1.Tile$json,
  '.endless.Coordinate': $1.Coordinate$json,
  '.endless.RemoveTile': $1.RemoveTile$json,
  '.endless.AdminResponse': $1.AdminResponse$json,
  '.endless.Log': $1.Log$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
};

/// Descriptor for `Admin`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List adminServiceDescriptor = $convert.base64Decode(
    'CgVBZG1pbhI7CgZNYW5hZ2USFS5lbmRsZXNzLkFkbWluUmVxdWVzdBoWLmVuZGxlc3MuQWRtaW'
    '5SZXNwb25zZSgBMAE=');
