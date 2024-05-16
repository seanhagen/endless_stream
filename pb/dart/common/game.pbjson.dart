//
//  Generated code. Do not modify.
//  source: common/game.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pingDescriptor instead')
const Ping$json = {
  '1': 'Ping',
};

/// Descriptor for `Ping`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pingDescriptor = $convert.base64Decode(
    'CgRQaW5n');

@$core.Deprecated('Use pongDescriptor instead')
const Pong$json = {
  '1': 'Pong',
};

/// Descriptor for `Pong`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pongDescriptor = $convert.base64Decode(
    'CgRQb25n');

@$core.Deprecated('Use levelDescriptor instead')
const Level$json = {
  '1': 'Level',
  '2': [
    {'1': 'tiles', '3': 1, '4': 3, '5': 11, '6': '.endless.Tile', '10': 'tiles'},
    {'1': 'tileset', '3': 2, '4': 1, '5': 14, '6': '.endless.Tileset', '10': 'tileset'},
  ],
};

/// Descriptor for `Level`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List levelDescriptor = $convert.base64Decode(
    'CgVMZXZlbBIjCgV0aWxlcxgBIAMoCzINLmVuZGxlc3MuVGlsZVIFdGlsZXMSKgoHdGlsZXNldB'
    'gCIAEoDjIQLmVuZGxlc3MuVGlsZXNldFIHdGlsZXNldA==');

@$core.Deprecated('Use heartbeatDescriptor instead')
const Heartbeat$json = {
  '1': 'Heartbeat',
  '2': [
    {'1': 'beat', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'beat'},
  ],
};

/// Descriptor for `Heartbeat`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List heartbeatDescriptor = $convert.base64Decode(
    'CglIZWFydGJlYXQSLgoEYmVhdBgBIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSBG'
    'JlYXQ=');

@$core.Deprecated('Use gameRequestDescriptor instead')
const GameRequest$json = {
  '1': 'GameRequest',
  '2': [
    {'1': 'client_id', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    {'1': 'ping', '3': 100, '4': 1, '5': 11, '6': '.endless.Ping', '9': 0, '10': 'ping'},
  ],
  '8': [
    {'1': 'request'},
  ],
};

/// Descriptor for `GameRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List gameRequestDescriptor = $convert.base64Decode(
    'CgtHYW1lUmVxdWVzdBIbCgljbGllbnRfaWQYASABKAlSCGNsaWVudElkEiMKBHBpbmcYZCABKA'
    'syDS5lbmRsZXNzLlBpbmdIAFIEcGluZ0IJCgdyZXF1ZXN0');

@$core.Deprecated('Use gameResponseDescriptor instead')
const GameResponse$json = {
  '1': 'GameResponse',
  '2': [
    {'1': 'server_id', '3': 1, '4': 1, '5': 9, '10': 'serverId'},
    {'1': 'heartbeat', '3': 2, '4': 1, '5': 11, '6': '.endless.Heartbeat', '9': 0, '10': 'heartbeat'},
    {'1': 'log', '3': 3, '4': 1, '5': 11, '6': '.endless.Log', '9': 0, '10': 'log'},
    {'1': 'level', '3': 5, '4': 1, '5': 11, '6': '.endless.Level', '9': 0, '10': 'level'},
    {'1': 'pong', '3': 100, '4': 1, '5': 11, '6': '.endless.Pong', '9': 0, '10': 'pong'},
  ],
  '8': [
    {'1': 'message'},
  ],
};

/// Descriptor for `GameResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List gameResponseDescriptor = $convert.base64Decode(
    'CgxHYW1lUmVzcG9uc2USGwoJc2VydmVyX2lkGAEgASgJUghzZXJ2ZXJJZBIyCgloZWFydGJlYX'
    'QYAiABKAsyEi5lbmRsZXNzLkhlYXJ0YmVhdEgAUgloZWFydGJlYXQSIAoDbG9nGAMgASgLMgwu'
    'ZW5kbGVzcy5Mb2dIAFIDbG9nEiYKBWxldmVsGAUgASgLMg4uZW5kbGVzcy5MZXZlbEgAUgVsZX'
    'ZlbBIjCgRwb25nGGQgASgLMg0uZW5kbGVzcy5Qb25nSABSBHBvbmdCCQoHbWVzc2FnZQ==');
