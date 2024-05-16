//
//  Generated code. Do not modify.
//  source: test/test_types.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

/// these types are just for testing, they shouldn't be used in the actual game client
class PingReq extends $pb.GeneratedMessage {
  factory PingReq({
    $core.String? msg,
  }) {
    final $result = create();
    if (msg != null) {
      $result.msg = msg;
    }
    return $result;
  }
  PingReq._() : super();
  factory PingReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PingReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PingReq', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'msg')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PingReq clone() => PingReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PingReq copyWith(void Function(PingReq) updates) => super.copyWith((message) => updates(message as PingReq)) as PingReq;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PingReq create() => PingReq._();
  PingReq createEmptyInstance() => create();
  static $pb.PbList<PingReq> createRepeated() => $pb.PbList<PingReq>();
  @$core.pragma('dart2js:noInline')
  static PingReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PingReq>(create);
  static PingReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get msg => $_getSZ(0);
  @$pb.TagNumber(1)
  set msg($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMsg() => $_has(0);
  @$pb.TagNumber(1)
  void clearMsg() => clearField(1);
}

class PongResp extends $pb.GeneratedMessage {
  factory PongResp({
    $core.String? gsm,
  }) {
    final $result = create();
    if (gsm != null) {
      $result.gsm = gsm;
    }
    return $result;
  }
  PongResp._() : super();
  factory PongResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PongResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PongResp', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'gsm')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PongResp clone() => PongResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PongResp copyWith(void Function(PongResp) updates) => super.copyWith((message) => updates(message as PongResp)) as PongResp;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PongResp create() => PongResp._();
  PongResp createEmptyInstance() => create();
  static $pb.PbList<PongResp> createRepeated() => $pb.PbList<PongResp>();
  @$core.pragma('dart2js:noInline')
  static PongResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PongResp>(create);
  static PongResp? _defaultInstance;

  /// gsm is the Ping.msg msg backwards
  @$pb.TagNumber(1)
  $core.String get gsm => $_getSZ(0);
  @$pb.TagNumber(1)
  set gsm($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGsm() => $_has(0);
  @$pb.TagNumber(1)
  void clearGsm() => clearField(1);
}

class TestStreamRequest extends $pb.GeneratedMessage {
  factory TestStreamRequest({
    $core.int? chunkId,
    $core.String? msg,
  }) {
    final $result = create();
    if (chunkId != null) {
      $result.chunkId = chunkId;
    }
    if (msg != null) {
      $result.msg = msg;
    }
    return $result;
  }
  TestStreamRequest._() : super();
  factory TestStreamRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TestStreamRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'TestStreamRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'chunkId', $pb.PbFieldType.O3)
    ..aOS(2, _omitFieldNames ? '' : 'msg')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TestStreamRequest clone() => TestStreamRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TestStreamRequest copyWith(void Function(TestStreamRequest) updates) => super.copyWith((message) => updates(message as TestStreamRequest)) as TestStreamRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static TestStreamRequest create() => TestStreamRequest._();
  TestStreamRequest createEmptyInstance() => create();
  static $pb.PbList<TestStreamRequest> createRepeated() => $pb.PbList<TestStreamRequest>();
  @$core.pragma('dart2js:noInline')
  static TestStreamRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TestStreamRequest>(create);
  static TestStreamRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get chunkId => $_getIZ(0);
  @$pb.TagNumber(1)
  set chunkId($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasChunkId() => $_has(0);
  @$pb.TagNumber(1)
  void clearChunkId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);
}

class TestStreamResponse extends $pb.GeneratedMessage {
  factory TestStreamResponse({
    $core.int? respId,
    $core.String? gsm,
  }) {
    final $result = create();
    if (respId != null) {
      $result.respId = respId;
    }
    if (gsm != null) {
      $result.gsm = gsm;
    }
    return $result;
  }
  TestStreamResponse._() : super();
  factory TestStreamResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TestStreamResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'TestStreamResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'respId', $pb.PbFieldType.O3)
    ..aOS(2, _omitFieldNames ? '' : 'gsm')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TestStreamResponse clone() => TestStreamResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TestStreamResponse copyWith(void Function(TestStreamResponse) updates) => super.copyWith((message) => updates(message as TestStreamResponse)) as TestStreamResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static TestStreamResponse create() => TestStreamResponse._();
  TestStreamResponse createEmptyInstance() => create();
  static $pb.PbList<TestStreamResponse> createRepeated() => $pb.PbList<TestStreamResponse>();
  @$core.pragma('dart2js:noInline')
  static TestStreamResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TestStreamResponse>(create);
  static TestStreamResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get respId => $_getIZ(0);
  @$pb.TagNumber(1)
  set respId($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRespId() => $_has(0);
  @$pb.TagNumber(1)
  void clearRespId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get gsm => $_getSZ(1);
  @$pb.TagNumber(2)
  set gsm($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasGsm() => $_has(1);
  @$pb.TagNumber(2)
  void clearGsm() => clearField(2);
}

class TestRequest extends $pb.GeneratedMessage {
  factory TestRequest({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  TestRequest._() : super();
  factory TestRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TestRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'TestRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TestRequest clone() => TestRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TestRequest copyWith(void Function(TestRequest) updates) => super.copyWith((message) => updates(message as TestRequest)) as TestRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static TestRequest create() => TestRequest._();
  TestRequest createEmptyInstance() => create();
  static $pb.PbList<TestRequest> createRepeated() => $pb.PbList<TestRequest>();
  @$core.pragma('dart2js:noInline')
  static TestRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TestRequest>(create);
  static TestRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);
}

class TestResponse extends $pb.GeneratedMessage {
  factory TestResponse({
    $core.String? resp,
  }) {
    final $result = create();
    if (resp != null) {
      $result.resp = resp;
    }
    return $result;
  }
  TestResponse._() : super();
  factory TestResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TestResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'TestResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'resp')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TestResponse clone() => TestResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TestResponse copyWith(void Function(TestResponse) updates) => super.copyWith((message) => updates(message as TestResponse)) as TestResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static TestResponse create() => TestResponse._();
  TestResponse createEmptyInstance() => create();
  static $pb.PbList<TestResponse> createRepeated() => $pb.PbList<TestResponse>();
  @$core.pragma('dart2js:noInline')
  static TestResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TestResponse>(create);
  static TestResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get resp => $_getSZ(0);
  @$pb.TagNumber(1)
  set resp($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasResp() => $_has(0);
  @$pb.TagNumber(1)
  void clearResp() => clearField(1);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
