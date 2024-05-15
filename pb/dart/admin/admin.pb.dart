//
//  Generated code. Do not modify.
//  source: admin/admin.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../common/logs.pb.dart' as $2;
import '../common/tile.pb.dart' as $1;
import 'admin.pbenum.dart';

export 'admin.pbenum.dart';

class AddTile extends $pb.GeneratedMessage {
  factory AddTile({
    $1.Tile? tile,
  }) {
    final $result = create();
    if (tile != null) {
      $result.tile = tile;
    }
    return $result;
  }
  AddTile._() : super();
  factory AddTile.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddTile.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AddTile', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOM<$1.Tile>(1, _omitFieldNames ? '' : 'tile', subBuilder: $1.Tile.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddTile clone() => AddTile()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddTile copyWith(void Function(AddTile) updates) => super.copyWith((message) => updates(message as AddTile)) as AddTile;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddTile create() => AddTile._();
  AddTile createEmptyInstance() => create();
  static $pb.PbList<AddTile> createRepeated() => $pb.PbList<AddTile>();
  @$core.pragma('dart2js:noInline')
  static AddTile getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddTile>(create);
  static AddTile? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Tile get tile => $_getN(0);
  @$pb.TagNumber(1)
  set tile($1.Tile v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTile() => $_has(0);
  @$pb.TagNumber(1)
  void clearTile() => clearField(1);
  @$pb.TagNumber(1)
  $1.Tile ensureTile() => $_ensure(0);
}

class RemoveTile extends $pb.GeneratedMessage {
  factory RemoveTile({
    $1.Coordinate? coords,
  }) {
    final $result = create();
    if (coords != null) {
      $result.coords = coords;
    }
    return $result;
  }
  RemoveTile._() : super();
  factory RemoveTile.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RemoveTile.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RemoveTile', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOM<$1.Coordinate>(1, _omitFieldNames ? '' : 'coords', subBuilder: $1.Coordinate.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RemoveTile clone() => RemoveTile()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RemoveTile copyWith(void Function(RemoveTile) updates) => super.copyWith((message) => updates(message as RemoveTile)) as RemoveTile;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RemoveTile create() => RemoveTile._();
  RemoveTile createEmptyInstance() => create();
  static $pb.PbList<RemoveTile> createRepeated() => $pb.PbList<RemoveTile>();
  @$core.pragma('dart2js:noInline')
  static RemoveTile getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RemoveTile>(create);
  static RemoveTile? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Coordinate get coords => $_getN(0);
  @$pb.TagNumber(1)
  set coords($1.Coordinate v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCoords() => $_has(0);
  @$pb.TagNumber(1)
  void clearCoords() => clearField(1);
  @$pb.TagNumber(1)
  $1.Coordinate ensureCoords() => $_ensure(0);
}

enum AdminRequest_Request {
  addTile,
  removeTile,
  notSet
}

class AdminRequest extends $pb.GeneratedMessage {
  factory AdminRequest({
    $core.String? clientId,
    AddTile? addTile,
    RemoveTile? removeTile,
  }) {
    final $result = create();
    if (clientId != null) {
      $result.clientId = clientId;
    }
    if (addTile != null) {
      $result.addTile = addTile;
    }
    if (removeTile != null) {
      $result.removeTile = removeTile;
    }
    return $result;
  }
  AdminRequest._() : super();
  factory AdminRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdminRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static const $core.Map<$core.int, AdminRequest_Request> _AdminRequest_RequestByTag = {
    2 : AdminRequest_Request.addTile,
    3 : AdminRequest_Request.removeTile,
    0 : AdminRequest_Request.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AdminRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..oo(0, [2, 3])
    ..aOS(1, _omitFieldNames ? '' : 'clientId')
    ..aOM<AddTile>(2, _omitFieldNames ? '' : 'addTile', subBuilder: AddTile.create)
    ..aOM<RemoveTile>(3, _omitFieldNames ? '' : 'removeTile', subBuilder: RemoveTile.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AdminRequest clone() => AdminRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AdminRequest copyWith(void Function(AdminRequest) updates) => super.copyWith((message) => updates(message as AdminRequest)) as AdminRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AdminRequest create() => AdminRequest._();
  AdminRequest createEmptyInstance() => create();
  static $pb.PbList<AdminRequest> createRepeated() => $pb.PbList<AdminRequest>();
  @$core.pragma('dart2js:noInline')
  static AdminRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdminRequest>(create);
  static AdminRequest? _defaultInstance;

  AdminRequest_Request whichRequest() => _AdminRequest_RequestByTag[$_whichOneof(0)]!;
  void clearRequest() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  AddTile get addTile => $_getN(1);
  @$pb.TagNumber(2)
  set addTile(AddTile v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasAddTile() => $_has(1);
  @$pb.TagNumber(2)
  void clearAddTile() => clearField(2);
  @$pb.TagNumber(2)
  AddTile ensureAddTile() => $_ensure(1);

  @$pb.TagNumber(3)
  RemoveTile get removeTile => $_getN(2);
  @$pb.TagNumber(3)
  set removeTile(RemoveTile v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasRemoveTile() => $_has(2);
  @$pb.TagNumber(3)
  void clearRemoveTile() => clearField(3);
  @$pb.TagNumber(3)
  RemoveTile ensureRemoveTile() => $_ensure(2);
}

class AdminResponse extends $pb.GeneratedMessage {
  factory AdminResponse({
    $core.String? serverId,
    $2.Log? log,
    Result? result,
  }) {
    final $result = create();
    if (serverId != null) {
      $result.serverId = serverId;
    }
    if (log != null) {
      $result.log = log;
    }
    if (result != null) {
      $result.result = result;
    }
    return $result;
  }
  AdminResponse._() : super();
  factory AdminResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdminResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AdminResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'serverId')
    ..aOM<$2.Log>(2, _omitFieldNames ? '' : 'log', subBuilder: $2.Log.create)
    ..e<Result>(3, _omitFieldNames ? '' : 'result', $pb.PbFieldType.OE, defaultOrMaker: Result.Unknown, valueOf: Result.valueOf, enumValues: Result.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AdminResponse clone() => AdminResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AdminResponse copyWith(void Function(AdminResponse) updates) => super.copyWith((message) => updates(message as AdminResponse)) as AdminResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AdminResponse create() => AdminResponse._();
  AdminResponse createEmptyInstance() => create();
  static $pb.PbList<AdminResponse> createRepeated() => $pb.PbList<AdminResponse>();
  @$core.pragma('dart2js:noInline')
  static AdminResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdminResponse>(create);
  static AdminResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get serverId => $_getSZ(0);
  @$pb.TagNumber(1)
  set serverId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerId() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerId() => clearField(1);

  @$pb.TagNumber(2)
  $2.Log get log => $_getN(1);
  @$pb.TagNumber(2)
  set log($2.Log v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasLog() => $_has(1);
  @$pb.TagNumber(2)
  void clearLog() => clearField(2);
  @$pb.TagNumber(2)
  $2.Log ensureLog() => $_ensure(1);

  @$pb.TagNumber(3)
  Result get result => $_getN(2);
  @$pb.TagNumber(3)
  set result(Result v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasResult() => $_has(2);
  @$pb.TagNumber(3)
  void clearResult() => clearField(3);
}

class AdminApi {
  $pb.RpcClient _client;
  AdminApi(this._client);

  $async.Future<AdminResponse> manage($pb.ClientContext? ctx, AdminRequest request) =>
    _client.invoke<AdminResponse>(ctx, 'Admin', 'Manage', request, AdminResponse())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
