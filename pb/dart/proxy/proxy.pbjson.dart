//
//  Generated code. Do not modify.
//  source: proxy/proxy.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import '../common/game.pbjson.dart' as $4;
import '../common/logs.pbjson.dart' as $2;
import '../common/tile.pbjson.dart' as $1;
import '../google/protobuf/timestamp.pbjson.dart' as $0;

const $core.Map<$core.String, $core.dynamic> ProxyServiceBase$json = {
  '1': 'Proxy',
  '2': [
    {'1': 'Game', '2': '.endless.GameRequest', '3': '.endless.GameResponse', '5': true, '6': true},
  ],
};

@$core.Deprecated('Use proxyServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> ProxyServiceBase$messageJson = {
  '.endless.GameRequest': $4.GameRequest$json,
  '.endless.Ping': $4.Ping$json,
  '.endless.GameResponse': $4.GameResponse$json,
  '.endless.Heartbeat': $4.Heartbeat$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
  '.endless.Log': $2.Log$json,
  '.endless.Level': $4.Level$json,
  '.endless.Tile': $1.Tile$json,
  '.endless.Coordinate': $1.Coordinate$json,
  '.endless.Pong': $4.Pong$json,
};

/// Descriptor for `Proxy`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List proxyServiceDescriptor = $convert.base64Decode(
    'CgVQcm94eRI3CgRHYW1lEhQuZW5kbGVzcy5HYW1lUmVxdWVzdBoVLmVuZGxlc3MuR2FtZVJlc3'
    'BvbnNlKAEwAQ==');
