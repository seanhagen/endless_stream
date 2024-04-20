//
//  Generated code. Do not modify.
//  source: proto/hex_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'hex.pb.dart' as $1;

class HexApi {
  $pb.RpcClient _client;
  HexApi(this._client);

  $async.Future<$1.InfoResponse> info($pb.ClientContext? ctx, $1.InfoRequest request) =>
    _client.invoke<$1.InfoResponse>(ctx, 'Hex', 'Info', request, $1.InfoResponse())
  ;
  $async.Future<$1.GameResponse> game($pb.ClientContext? ctx, $1.GameRequest request) =>
    _client.invoke<$1.GameResponse>(ctx, 'Hex', 'Game', request, $1.GameResponse())
  ;
}

class AdminApi {
  $pb.RpcClient _client;
  AdminApi(this._client);

  $async.Future<$1.AdminResponse> manage($pb.ClientContext? ctx, $1.AdminRequest request) =>
    _client.invoke<$1.AdminResponse>(ctx, 'Admin', 'Manage', request, $1.AdminResponse())
  ;
}

class TestApi {
  $pb.RpcClient _client;
  TestApi(this._client);

  $async.Future<$1.PongResp> ping($pb.ClientContext? ctx, $1.PingReq request) =>
    _client.invoke<$1.PongResp>(ctx, 'Test', 'Ping', request, $1.PongResp())
  ;
  $async.Future<$1.TestResponse> clientStream($pb.ClientContext? ctx, $1.TestStreamRequest request) =>
    _client.invoke<$1.TestResponse>(ctx, 'Test', 'ClientStream', request, $1.TestResponse())
  ;
  $async.Future<$1.TestStreamResponse> serverStream($pb.ClientContext? ctx, $1.TestRequest request) =>
    _client.invoke<$1.TestStreamResponse>(ctx, 'Test', 'ServerStream', request, $1.TestStreamResponse())
  ;
  $async.Future<$1.TestStreamResponse> biDiStream($pb.ClientContext? ctx, $1.TestStreamRequest request) =>
    _client.invoke<$1.TestStreamResponse>(ctx, 'Test', 'BiDiStream', request, $1.TestStreamResponse())
  ;
}
