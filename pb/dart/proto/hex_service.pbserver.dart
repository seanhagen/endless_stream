//
//  Generated code. Do not modify.
//  source: proto/hex_service.proto
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

import 'hex.pb.dart' as $1;
import 'hex_service.pbjson.dart';

export 'hex_service.pb.dart';

abstract class HexServiceBase extends $pb.GeneratedService {
  $async.Future<$1.InfoResponse> info($pb.ServerContext ctx, $1.InfoRequest request);
  $async.Future<$1.GameResponse> game($pb.ServerContext ctx, $1.GameRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'Info': return $1.InfoRequest();
      case 'Game': return $1.GameRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'Info': return this.info(ctx, request as $1.InfoRequest);
      case 'Game': return this.game(ctx, request as $1.GameRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => HexServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => HexServiceBase$messageJson;
}

abstract class AdminServiceBase extends $pb.GeneratedService {
  $async.Future<$1.AdminResponse> manage($pb.ServerContext ctx, $1.AdminRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'Manage': return $1.AdminRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'Manage': return this.manage(ctx, request as $1.AdminRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => AdminServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => AdminServiceBase$messageJson;
}
