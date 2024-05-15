//
//  Generated code. Do not modify.
//  source: test/test.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'test.pbjson.dart';
import 'test_types.pb.dart' as $5;

export 'test.pb.dart';

abstract class TestServiceBase extends $pb.GeneratedService {
  $async.Future<$5.PongResp> ping($pb.ServerContext ctx, $5.PingReq request);
  $async.Future<$5.TestResponse> clientStream($pb.ServerContext ctx, $5.TestStreamRequest request);
  $async.Future<$5.TestStreamResponse> serverStream($pb.ServerContext ctx, $5.TestRequest request);
  $async.Future<$5.TestStreamResponse> biDiStream($pb.ServerContext ctx, $5.TestStreamRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'Ping': return $5.PingReq();
      case 'ClientStream': return $5.TestStreamRequest();
      case 'ServerStream': return $5.TestRequest();
      case 'BiDiStream': return $5.TestStreamRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'Ping': return this.ping(ctx, request as $5.PingReq);
      case 'ClientStream': return this.clientStream(ctx, request as $5.TestStreamRequest);
      case 'ServerStream': return this.serverStream(ctx, request as $5.TestRequest);
      case 'BiDiStream': return this.biDiStream(ctx, request as $5.TestStreamRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TestServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TestServiceBase$messageJson;
}
