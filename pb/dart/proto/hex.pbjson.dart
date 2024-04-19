//
//  Generated code. Do not modify.
//  source: proto/hex.proto
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

@$core.Deprecated('Use infoRequestDescriptor instead')
const InfoRequest$json = {
  '1': 'InfoRequest',
};

/// Descriptor for `InfoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List infoRequestDescriptor = $convert.base64Decode(
    'CgtJbmZvUmVxdWVzdA==');

@$core.Deprecated('Use infoResponseDescriptor instead')
const InfoResponse$json = {
  '1': 'InfoResponse',
  '2': [
    {'1': 'version', '3': 1, '4': 1, '5': 9, '10': 'version'},
    {'1': 'build_date', '3': 2, '4': 1, '5': 9, '10': 'buildDate'},
  ],
};

/// Descriptor for `InfoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List infoResponseDescriptor = $convert.base64Decode(
    'CgxJbmZvUmVzcG9uc2USGAoHdmVyc2lvbhgBIAEoCVIHdmVyc2lvbhIdCgpidWlsZF9kYXRlGA'
    'IgASgJUglidWlsZERhdGU=');

@$core.Deprecated('Use getLevelDescriptor instead')
const GetLevel$json = {
  '1': 'GetLevel',
};

/// Descriptor for `GetLevel`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getLevelDescriptor = $convert.base64Decode(
    'CghHZXRMZXZlbA==');

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
    {'1': 'info', '3': 2, '4': 1, '5': 11, '6': '.endless.InfoRequest', '9': 0, '10': 'info'},
    {'1': 'get_level', '3': 3, '4': 1, '5': 11, '6': '.endless.GetLevel', '9': 0, '10': 'getLevel'},
  ],
  '8': [
    {'1': 'request'},
  ],
};

/// Descriptor for `GameRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List gameRequestDescriptor = $convert.base64Decode(
    'CgtHYW1lUmVxdWVzdBIbCgljbGllbnRfaWQYASABKAlSCGNsaWVudElkEioKBGluZm8YAiABKA'
    'syFC5lbmRsZXNzLkluZm9SZXF1ZXN0SABSBGluZm8SMAoJZ2V0X2xldmVsGAMgASgLMhEuZW5k'
    'bGVzcy5HZXRMZXZlbEgAUghnZXRMZXZlbEIJCgdyZXF1ZXN0');

@$core.Deprecated('Use gameResponseDescriptor instead')
const GameResponse$json = {
  '1': 'GameResponse',
  '2': [
    {'1': 'server_id', '3': 1, '4': 1, '5': 9, '10': 'serverId'},
    {'1': 'heartbeat', '3': 2, '4': 1, '5': 11, '6': '.endless.Heartbeat', '9': 0, '10': 'heartbeat'},
    {'1': 'log', '3': 3, '4': 1, '5': 11, '6': '.endless.Log', '9': 0, '10': 'log'},
    {'1': 'info', '3': 4, '4': 1, '5': 11, '6': '.endless.InfoResponse', '9': 0, '10': 'info'},
    {'1': 'level', '3': 5, '4': 1, '5': 11, '6': '.endless.Level', '9': 0, '10': 'level'},
  ],
  '8': [
    {'1': 'message'},
  ],
};

/// Descriptor for `GameResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List gameResponseDescriptor = $convert.base64Decode(
    'CgxHYW1lUmVzcG9uc2USGwoJc2VydmVyX2lkGAEgASgJUghzZXJ2ZXJJZBIyCgloZWFydGJlYX'
    'QYAiABKAsyEi5lbmRsZXNzLkhlYXJ0YmVhdEgAUgloZWFydGJlYXQSIAoDbG9nGAMgASgLMgwu'
    'ZW5kbGVzcy5Mb2dIAFIDbG9nEisKBGluZm8YBCABKAsyFS5lbmRsZXNzLkluZm9SZXNwb25zZU'
    'gAUgRpbmZvEiYKBWxldmVsGAUgASgLMg4uZW5kbGVzcy5MZXZlbEgAUgVsZXZlbEIJCgdtZXNz'
    'YWdl');

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
