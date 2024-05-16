//
//  Generated code. Do not modify.
//  source: proxy/proxy.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../common/game.pb.dart' as $4;

class ProxyApi {
  $pb.RpcClient _client;
  ProxyApi(this._client);

  $async.Future<$4.GameResponse> game($pb.ClientContext? ctx, $4.GameRequest request) =>
    _client.invoke<$4.GameResponse>(ctx, 'Proxy', 'Game', request, $4.GameResponse())
  ;
}
