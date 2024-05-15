//
//  Generated code. Do not modify.
//  source: test/test.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'test_types.pb.dart' as $5;

class TestApi {
  $pb.RpcClient _client;
  TestApi(this._client);

  $async.Future<$5.PongResp> ping($pb.ClientContext? ctx, $5.PingReq request) =>
    _client.invoke<$5.PongResp>(ctx, 'Test', 'Ping', request, $5.PongResp())
  ;
  $async.Future<$5.TestResponse> clientStream($pb.ClientContext? ctx, $5.TestStreamRequest request) =>
    _client.invoke<$5.TestResponse>(ctx, 'Test', 'ClientStream', request, $5.TestResponse())
  ;
  $async.Future<$5.TestStreamResponse> serverStream($pb.ClientContext? ctx, $5.TestRequest request) =>
    _client.invoke<$5.TestStreamResponse>(ctx, 'Test', 'ServerStream', request, $5.TestStreamResponse())
  ;
  $async.Future<$5.TestStreamResponse> biDiStream($pb.ClientContext? ctx, $5.TestStreamRequest request) =>
    _client.invoke<$5.TestStreamResponse>(ctx, 'Test', 'BiDiStream', request, $5.TestStreamResponse())
  ;
}
