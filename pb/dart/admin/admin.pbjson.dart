//
//  Generated code. Do not modify.
//  source: admin/admin.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import '../common/logs.pbjson.dart' as $2;
import '../common/tile.pbjson.dart' as $1;
import '../google/protobuf/timestamp.pbjson.dart' as $0;

@$core.Deprecated('Use resultDescriptor instead')
const Result$json = {
  '1': 'Result',
  '2': [
    {'1': 'Unknown', '2': 0},
    {'1': 'Failure', '2': 1},
    {'1': 'Success', '2': 2},
  ],
};

/// Descriptor for `Result`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List resultDescriptor = $convert.base64Decode(
    'CgZSZXN1bHQSCwoHVW5rbm93bhAAEgsKB0ZhaWx1cmUQARILCgdTdWNjZXNzEAI=');

@$core.Deprecated('Use addTileDescriptor instead')
const AddTile$json = {
  '1': 'AddTile',
  '2': [
    {'1': 'tile', '3': 1, '4': 1, '5': 11, '6': '.endless.Tile', '10': 'tile'},
  ],
};

/// Descriptor for `AddTile`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addTileDescriptor = $convert.base64Decode(
    'CgdBZGRUaWxlEiEKBHRpbGUYASABKAsyDS5lbmRsZXNzLlRpbGVSBHRpbGU=');

@$core.Deprecated('Use removeTileDescriptor instead')
const RemoveTile$json = {
  '1': 'RemoveTile',
  '2': [
    {'1': 'coords', '3': 1, '4': 1, '5': 11, '6': '.endless.Coordinate', '10': 'coords'},
  ],
};

/// Descriptor for `RemoveTile`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List removeTileDescriptor = $convert.base64Decode(
    'CgpSZW1vdmVUaWxlEisKBmNvb3JkcxgBIAEoCzITLmVuZGxlc3MuQ29vcmRpbmF0ZVIGY29vcm'
    'Rz');

@$core.Deprecated('Use adminRequestDescriptor instead')
const AdminRequest$json = {
  '1': 'AdminRequest',
  '2': [
    {'1': 'client_id', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    {'1': 'add_tile', '3': 2, '4': 1, '5': 11, '6': '.endless.AddTile', '9': 0, '10': 'addTile'},
    {'1': 'remove_tile', '3': 3, '4': 1, '5': 11, '6': '.endless.RemoveTile', '9': 0, '10': 'removeTile'},
  ],
  '8': [
    {'1': 'request'},
  ],
};

/// Descriptor for `AdminRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminRequestDescriptor = $convert.base64Decode(
    'CgxBZG1pblJlcXVlc3QSGwoJY2xpZW50X2lkGAEgASgJUghjbGllbnRJZBItCghhZGRfdGlsZR'
    'gCIAEoCzIQLmVuZGxlc3MuQWRkVGlsZUgAUgdhZGRUaWxlEjYKC3JlbW92ZV90aWxlGAMgASgL'
    'MhMuZW5kbGVzcy5SZW1vdmVUaWxlSABSCnJlbW92ZVRpbGVCCQoHcmVxdWVzdA==');

@$core.Deprecated('Use adminResponseDescriptor instead')
const AdminResponse$json = {
  '1': 'AdminResponse',
  '2': [
    {'1': 'server_id', '3': 1, '4': 1, '5': 9, '10': 'serverId'},
    {'1': 'log', '3': 2, '4': 1, '5': 11, '6': '.endless.Log', '10': 'log'},
    {'1': 'result', '3': 3, '4': 1, '5': 14, '6': '.endless.Result', '10': 'result'},
  ],
};

/// Descriptor for `AdminResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List adminResponseDescriptor = $convert.base64Decode(
    'Cg1BZG1pblJlc3BvbnNlEhsKCXNlcnZlcl9pZBgBIAEoCVIIc2VydmVySWQSHgoDbG9nGAIgAS'
    'gLMgwuZW5kbGVzcy5Mb2dSA2xvZxInCgZyZXN1bHQYAyABKA4yDy5lbmRsZXNzLlJlc3VsdFIG'
    'cmVzdWx0');

const $core.Map<$core.String, $core.dynamic> AdminServiceBase$json = {
  '1': 'Admin',
  '2': [
    {'1': 'Manage', '2': '.endless.AdminRequest', '3': '.endless.AdminResponse', '5': true, '6': true},
  ],
};

@$core.Deprecated('Use adminServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> AdminServiceBase$messageJson = {
  '.endless.AdminRequest': AdminRequest$json,
  '.endless.AddTile': AddTile$json,
  '.endless.Tile': $1.Tile$json,
  '.endless.Coordinate': $1.Coordinate$json,
  '.endless.RemoveTile': RemoveTile$json,
  '.endless.AdminResponse': AdminResponse$json,
  '.endless.Log': $2.Log$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
};

/// Descriptor for `Admin`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List adminServiceDescriptor = $convert.base64Decode(
    'CgVBZG1pbhI7CgZNYW5hZ2USFS5lbmRsZXNzLkFkbWluUmVxdWVzdBoWLmVuZGxlc3MuQWRtaW'
    '5SZXNwb25zZSgBMAE=');
