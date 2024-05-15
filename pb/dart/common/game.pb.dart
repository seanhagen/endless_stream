//
//  Generated code. Do not modify.
//  source: common/game.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/timestamp.pb.dart' as $0;
import 'logs.pb.dart' as $2;
import 'tile.pb.dart' as $1;
import 'tile.pbenum.dart' as $1;

class Ping extends $pb.GeneratedMessage {
  factory Ping() => create();
  Ping._() : super();
  factory Ping.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Ping.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Ping', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Ping clone() => Ping()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Ping copyWith(void Function(Ping) updates) => super.copyWith((message) => updates(message as Ping)) as Ping;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Ping create() => Ping._();
  Ping createEmptyInstance() => create();
  static $pb.PbList<Ping> createRepeated() => $pb.PbList<Ping>();
  @$core.pragma('dart2js:noInline')
  static Ping getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Ping>(create);
  static Ping? _defaultInstance;
}

class Pong extends $pb.GeneratedMessage {
  factory Pong() => create();
  Pong._() : super();
  factory Pong.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Pong.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Pong', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Pong clone() => Pong()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Pong copyWith(void Function(Pong) updates) => super.copyWith((message) => updates(message as Pong)) as Pong;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Pong create() => Pong._();
  Pong createEmptyInstance() => create();
  static $pb.PbList<Pong> createRepeated() => $pb.PbList<Pong>();
  @$core.pragma('dart2js:noInline')
  static Pong getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Pong>(create);
  static Pong? _defaultInstance;
}

/// Level is a single layer of tiles laid out to create a level.
class Level extends $pb.GeneratedMessage {
  factory Level({
    $core.Iterable<$1.Tile>? tiles,
    $1.Tileset? tileset,
  }) {
    final $result = create();
    if (tiles != null) {
      $result.tiles.addAll(tiles);
    }
    if (tileset != null) {
      $result.tileset = tileset;
    }
    return $result;
  }
  Level._() : super();
  factory Level.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Level.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Level', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..pc<$1.Tile>(1, _omitFieldNames ? '' : 'tiles', $pb.PbFieldType.PM, subBuilder: $1.Tile.create)
    ..e<$1.Tileset>(2, _omitFieldNames ? '' : 'tileset', $pb.PbFieldType.OE, defaultOrMaker: $1.Tileset.Dungeon, valueOf: $1.Tileset.valueOf, enumValues: $1.Tileset.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Level clone() => Level()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Level copyWith(void Function(Level) updates) => super.copyWith((message) => updates(message as Level)) as Level;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Level create() => Level._();
  Level createEmptyInstance() => create();
  static $pb.PbList<Level> createRepeated() => $pb.PbList<Level>();
  @$core.pragma('dart2js:noInline')
  static Level getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Level>(create);
  static Level? _defaultInstance;

  /// Tiles is an array of all the tiles laid out on a level.
  @$pb.TagNumber(1)
  $core.List<$1.Tile> get tiles => $_getList(0);

  /// Tileset tells Godot what tileset to use.
  @$pb.TagNumber(2)
  $1.Tileset get tileset => $_getN(1);
  @$pb.TagNumber(2)
  set tileset($1.Tileset v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTileset() => $_has(1);
  @$pb.TagNumber(2)
  void clearTileset() => clearField(2);
}

class Heartbeat extends $pb.GeneratedMessage {
  factory Heartbeat({
    $0.Timestamp? beat,
  }) {
    final $result = create();
    if (beat != null) {
      $result.beat = beat;
    }
    return $result;
  }
  Heartbeat._() : super();
  factory Heartbeat.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Heartbeat.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Heartbeat', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOM<$0.Timestamp>(1, _omitFieldNames ? '' : 'beat', subBuilder: $0.Timestamp.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Heartbeat clone() => Heartbeat()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Heartbeat copyWith(void Function(Heartbeat) updates) => super.copyWith((message) => updates(message as Heartbeat)) as Heartbeat;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Heartbeat create() => Heartbeat._();
  Heartbeat createEmptyInstance() => create();
  static $pb.PbList<Heartbeat> createRepeated() => $pb.PbList<Heartbeat>();
  @$core.pragma('dart2js:noInline')
  static Heartbeat getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Heartbeat>(create);
  static Heartbeat? _defaultInstance;

  @$pb.TagNumber(1)
  $0.Timestamp get beat => $_getN(0);
  @$pb.TagNumber(1)
  set beat($0.Timestamp v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBeat() => $_has(0);
  @$pb.TagNumber(1)
  void clearBeat() => clearField(1);
  @$pb.TagNumber(1)
  $0.Timestamp ensureBeat() => $_ensure(0);
}

enum GameRequest_Request {
  ping,
  notSet
}

class GameRequest extends $pb.GeneratedMessage {
  factory GameRequest({
    $core.String? clientId,
    Ping? ping,
  }) {
    final $result = create();
    if (clientId != null) {
      $result.clientId = clientId;
    }
    if (ping != null) {
      $result.ping = ping;
    }
    return $result;
  }
  GameRequest._() : super();
  factory GameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static const $core.Map<$core.int, GameRequest_Request> _GameRequest_RequestByTag = {
    100 : GameRequest_Request.ping,
    0 : GameRequest_Request.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GameRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..oo(0, [100])
    ..aOS(1, _omitFieldNames ? '' : 'clientId')
    ..aOM<Ping>(100, _omitFieldNames ? '' : 'ping', subBuilder: Ping.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GameRequest clone() => GameRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GameRequest copyWith(void Function(GameRequest) updates) => super.copyWith((message) => updates(message as GameRequest)) as GameRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GameRequest create() => GameRequest._();
  GameRequest createEmptyInstance() => create();
  static $pb.PbList<GameRequest> createRepeated() => $pb.PbList<GameRequest>();
  @$core.pragma('dart2js:noInline')
  static GameRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GameRequest>(create);
  static GameRequest? _defaultInstance;

  GameRequest_Request whichRequest() => _GameRequest_RequestByTag[$_whichOneof(0)]!;
  void clearRequest() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(100)
  Ping get ping => $_getN(1);
  @$pb.TagNumber(100)
  set ping(Ping v) { setField(100, v); }
  @$pb.TagNumber(100)
  $core.bool hasPing() => $_has(1);
  @$pb.TagNumber(100)
  void clearPing() => clearField(100);
  @$pb.TagNumber(100)
  Ping ensurePing() => $_ensure(1);
}

enum GameResponse_Message {
  heartbeat,
  log,
  level,
  pong,
  notSet
}

class GameResponse extends $pb.GeneratedMessage {
  factory GameResponse({
    $core.String? serverId,
    Heartbeat? heartbeat,
    $2.Log? log,
    Level? level,
    Pong? pong,
  }) {
    final $result = create();
    if (serverId != null) {
      $result.serverId = serverId;
    }
    if (heartbeat != null) {
      $result.heartbeat = heartbeat;
    }
    if (log != null) {
      $result.log = log;
    }
    if (level != null) {
      $result.level = level;
    }
    if (pong != null) {
      $result.pong = pong;
    }
    return $result;
  }
  GameResponse._() : super();
  factory GameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static const $core.Map<$core.int, GameResponse_Message> _GameResponse_MessageByTag = {
    2 : GameResponse_Message.heartbeat,
    3 : GameResponse_Message.log,
    5 : GameResponse_Message.level,
    100 : GameResponse_Message.pong,
    0 : GameResponse_Message.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GameResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..oo(0, [2, 3, 5, 100])
    ..aOS(1, _omitFieldNames ? '' : 'serverId')
    ..aOM<Heartbeat>(2, _omitFieldNames ? '' : 'heartbeat', subBuilder: Heartbeat.create)
    ..aOM<$2.Log>(3, _omitFieldNames ? '' : 'log', subBuilder: $2.Log.create)
    ..aOM<Level>(5, _omitFieldNames ? '' : 'level', subBuilder: Level.create)
    ..aOM<Pong>(100, _omitFieldNames ? '' : 'pong', subBuilder: Pong.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GameResponse clone() => GameResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GameResponse copyWith(void Function(GameResponse) updates) => super.copyWith((message) => updates(message as GameResponse)) as GameResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GameResponse create() => GameResponse._();
  GameResponse createEmptyInstance() => create();
  static $pb.PbList<GameResponse> createRepeated() => $pb.PbList<GameResponse>();
  @$core.pragma('dart2js:noInline')
  static GameResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GameResponse>(create);
  static GameResponse? _defaultInstance;

  GameResponse_Message whichMessage() => _GameResponse_MessageByTag[$_whichOneof(0)]!;
  void clearMessage() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get serverId => $_getSZ(0);
  @$pb.TagNumber(1)
  set serverId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerId() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerId() => clearField(1);

  @$pb.TagNumber(2)
  Heartbeat get heartbeat => $_getN(1);
  @$pb.TagNumber(2)
  set heartbeat(Heartbeat v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasHeartbeat() => $_has(1);
  @$pb.TagNumber(2)
  void clearHeartbeat() => clearField(2);
  @$pb.TagNumber(2)
  Heartbeat ensureHeartbeat() => $_ensure(1);

  @$pb.TagNumber(3)
  $2.Log get log => $_getN(2);
  @$pb.TagNumber(3)
  set log($2.Log v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasLog() => $_has(2);
  @$pb.TagNumber(3)
  void clearLog() => clearField(3);
  @$pb.TagNumber(3)
  $2.Log ensureLog() => $_ensure(2);

  @$pb.TagNumber(5)
  Level get level => $_getN(3);
  @$pb.TagNumber(5)
  set level(Level v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasLevel() => $_has(3);
  @$pb.TagNumber(5)
  void clearLevel() => clearField(5);
  @$pb.TagNumber(5)
  Level ensureLevel() => $_ensure(3);

  @$pb.TagNumber(100)
  Pong get pong => $_getN(4);
  @$pb.TagNumber(100)
  set pong(Pong v) { setField(100, v); }
  @$pb.TagNumber(100)
  $core.bool hasPong() => $_has(4);
  @$pb.TagNumber(100)
  void clearPong() => clearField(100);
  @$pb.TagNumber(100)
  Pong ensurePong() => $_ensure(4);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
