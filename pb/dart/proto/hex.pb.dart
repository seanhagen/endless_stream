//
//  Generated code. Do not modify.
//  source: proto/hex.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/timestamp.pb.dart' as $0;
import 'hex.pbenum.dart';

export 'hex.pbenum.dart';

/// Coordinate is the x,y location in the world for a tile.
class Coordinate extends $pb.GeneratedMessage {
  factory Coordinate({
    $core.int? x,
    $core.int? y,
  }) {
    final $result = create();
    if (x != null) {
      $result.x = x;
    }
    if (y != null) {
      $result.y = y;
    }
    return $result;
  }
  Coordinate._() : super();
  factory Coordinate.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Coordinate.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Coordinate', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'x', $pb.PbFieldType.O3)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'y', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Coordinate clone() => Coordinate()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Coordinate copyWith(void Function(Coordinate) updates) => super.copyWith((message) => updates(message as Coordinate)) as Coordinate;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Coordinate create() => Coordinate._();
  Coordinate createEmptyInstance() => create();
  static $pb.PbList<Coordinate> createRepeated() => $pb.PbList<Coordinate>();
  @$core.pragma('dart2js:noInline')
  static Coordinate getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Coordinate>(create);
  static Coordinate? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get x => $_getIZ(0);
  @$pb.TagNumber(1)
  set x($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasX() => $_has(0);
  @$pb.TagNumber(1)
  void clearX() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get y => $_getIZ(1);
  @$pb.TagNumber(2)
  set y($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasY() => $_has(1);
  @$pb.TagNumber(2)
  void clearY() => clearField(2);
}

/// Tile is a single tile within the world.
class Tile extends $pb.GeneratedMessage {
  factory Tile({
    Type? type,
    Coordinate? coords,
  }) {
    final $result = create();
    if (type != null) {
      $result.type = type;
    }
    if (coords != null) {
      $result.coords = coords;
    }
    return $result;
  }
  Tile._() : super();
  factory Tile.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Tile.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Tile', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..e<Type>(1, _omitFieldNames ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: Type.Empty, valueOf: Type.valueOf, enumValues: Type.values)
    ..aOM<Coordinate>(2, _omitFieldNames ? '' : 'coords', subBuilder: Coordinate.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Tile clone() => Tile()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Tile copyWith(void Function(Tile) updates) => super.copyWith((message) => updates(message as Tile)) as Tile;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Tile create() => Tile._();
  Tile createEmptyInstance() => create();
  static $pb.PbList<Tile> createRepeated() => $pb.PbList<Tile>();
  @$core.pragma('dart2js:noInline')
  static Tile getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Tile>(create);
  static Tile? _defaultInstance;

  /// Type defines whether the tile is empty, a floor, or a wall.
  @$pb.TagNumber(1)
  Type get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(Type v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  /// Coords defines the position of the tile within the world.
  @$pb.TagNumber(2)
  Coordinate get coords => $_getN(1);
  @$pb.TagNumber(2)
  set coords(Coordinate v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasCoords() => $_has(1);
  @$pb.TagNumber(2)
  void clearCoords() => clearField(2);
  @$pb.TagNumber(2)
  Coordinate ensureCoords() => $_ensure(1);
}

/// Level is a single layer of tiles laid out to create a level.
class Level extends $pb.GeneratedMessage {
  factory Level({
    $core.Iterable<Tile>? tiles,
    Tileset? tileset,
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
    ..pc<Tile>(1, _omitFieldNames ? '' : 'tiles', $pb.PbFieldType.PM, subBuilder: Tile.create)
    ..e<Tileset>(2, _omitFieldNames ? '' : 'tileset', $pb.PbFieldType.OE, defaultOrMaker: Tileset.Dungeon, valueOf: Tileset.valueOf, enumValues: Tileset.values)
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
  $core.List<Tile> get tiles => $_getList(0);

  /// Tileset tells Godot what tileset to use.
  @$pb.TagNumber(2)
  Tileset get tileset => $_getN(1);
  @$pb.TagNumber(2)
  set tileset(Tileset v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasTileset() => $_has(1);
  @$pb.TagNumber(2)
  void clearTileset() => clearField(2);
}

class Log extends $pb.GeneratedMessage {
  factory Log({
    $core.String? msg,
    $0.Timestamp? at,
    LogLevel? level,
    LogSource? source,
  }) {
    final $result = create();
    if (msg != null) {
      $result.msg = msg;
    }
    if (at != null) {
      $result.at = at;
    }
    if (level != null) {
      $result.level = level;
    }
    if (source != null) {
      $result.source = source;
    }
    return $result;
  }
  Log._() : super();
  factory Log.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Log.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Log', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'msg')
    ..aOM<$0.Timestamp>(2, _omitFieldNames ? '' : 'at', subBuilder: $0.Timestamp.create)
    ..e<LogLevel>(3, _omitFieldNames ? '' : 'level', $pb.PbFieldType.OE, defaultOrMaker: LogLevel.Info, valueOf: LogLevel.valueOf, enumValues: LogLevel.values)
    ..e<LogSource>(4, _omitFieldNames ? '' : 'source', $pb.PbFieldType.OE, defaultOrMaker: LogSource.Server, valueOf: LogSource.valueOf, enumValues: LogSource.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Log clone() => Log()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Log copyWith(void Function(Log) updates) => super.copyWith((message) => updates(message as Log)) as Log;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Log create() => Log._();
  Log createEmptyInstance() => create();
  static $pb.PbList<Log> createRepeated() => $pb.PbList<Log>();
  @$core.pragma('dart2js:noInline')
  static Log getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Log>(create);
  static Log? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get msg => $_getSZ(0);
  @$pb.TagNumber(1)
  set msg($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMsg() => $_has(0);
  @$pb.TagNumber(1)
  void clearMsg() => clearField(1);

  @$pb.TagNumber(2)
  $0.Timestamp get at => $_getN(1);
  @$pb.TagNumber(2)
  set at($0.Timestamp v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasAt() => $_has(1);
  @$pb.TagNumber(2)
  void clearAt() => clearField(2);
  @$pb.TagNumber(2)
  $0.Timestamp ensureAt() => $_ensure(1);

  @$pb.TagNumber(3)
  LogLevel get level => $_getN(2);
  @$pb.TagNumber(3)
  set level(LogLevel v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasLevel() => $_has(2);
  @$pb.TagNumber(3)
  void clearLevel() => clearField(3);

  @$pb.TagNumber(4)
  LogSource get source => $_getN(3);
  @$pb.TagNumber(4)
  set source(LogSource v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasSource() => $_has(3);
  @$pb.TagNumber(4)
  void clearSource() => clearField(4);
}

class InfoRequest extends $pb.GeneratedMessage {
  factory InfoRequest() => create();
  InfoRequest._() : super();
  factory InfoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory InfoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'InfoRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  InfoRequest clone() => InfoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  InfoRequest copyWith(void Function(InfoRequest) updates) => super.copyWith((message) => updates(message as InfoRequest)) as InfoRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static InfoRequest create() => InfoRequest._();
  InfoRequest createEmptyInstance() => create();
  static $pb.PbList<InfoRequest> createRepeated() => $pb.PbList<InfoRequest>();
  @$core.pragma('dart2js:noInline')
  static InfoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<InfoRequest>(create);
  static InfoRequest? _defaultInstance;
}

class InfoResponse extends $pb.GeneratedMessage {
  factory InfoResponse({
    $core.String? version,
    $core.String? buildDate,
  }) {
    final $result = create();
    if (version != null) {
      $result.version = version;
    }
    if (buildDate != null) {
      $result.buildDate = buildDate;
    }
    return $result;
  }
  InfoResponse._() : super();
  factory InfoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory InfoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'InfoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'version')
    ..aOS(2, _omitFieldNames ? '' : 'buildDate')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  InfoResponse clone() => InfoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  InfoResponse copyWith(void Function(InfoResponse) updates) => super.copyWith((message) => updates(message as InfoResponse)) as InfoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static InfoResponse create() => InfoResponse._();
  InfoResponse createEmptyInstance() => create();
  static $pb.PbList<InfoResponse> createRepeated() => $pb.PbList<InfoResponse>();
  @$core.pragma('dart2js:noInline')
  static InfoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<InfoResponse>(create);
  static InfoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get version => $_getSZ(0);
  @$pb.TagNumber(1)
  set version($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasVersion() => $_has(0);
  @$pb.TagNumber(1)
  void clearVersion() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get buildDate => $_getSZ(1);
  @$pb.TagNumber(2)
  set buildDate($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasBuildDate() => $_has(1);
  @$pb.TagNumber(2)
  void clearBuildDate() => clearField(2);
}

class GetLevel extends $pb.GeneratedMessage {
  factory GetLevel() => create();
  GetLevel._() : super();
  factory GetLevel.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetLevel.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetLevel', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetLevel clone() => GetLevel()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetLevel copyWith(void Function(GetLevel) updates) => super.copyWith((message) => updates(message as GetLevel)) as GetLevel;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetLevel create() => GetLevel._();
  GetLevel createEmptyInstance() => create();
  static $pb.PbList<GetLevel> createRepeated() => $pb.PbList<GetLevel>();
  @$core.pragma('dart2js:noInline')
  static GetLevel getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetLevel>(create);
  static GetLevel? _defaultInstance;
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
  info,
  getLevel,
  notSet
}

class GameRequest extends $pb.GeneratedMessage {
  factory GameRequest({
    $core.String? clientId,
    InfoRequest? info,
    GetLevel? getLevel,
  }) {
    final $result = create();
    if (clientId != null) {
      $result.clientId = clientId;
    }
    if (info != null) {
      $result.info = info;
    }
    if (getLevel != null) {
      $result.getLevel = getLevel;
    }
    return $result;
  }
  GameRequest._() : super();
  factory GameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static const $core.Map<$core.int, GameRequest_Request> _GameRequest_RequestByTag = {
    2 : GameRequest_Request.info,
    3 : GameRequest_Request.getLevel,
    0 : GameRequest_Request.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GameRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..oo(0, [2, 3])
    ..aOS(1, _omitFieldNames ? '' : 'clientId')
    ..aOM<InfoRequest>(2, _omitFieldNames ? '' : 'info', subBuilder: InfoRequest.create)
    ..aOM<GetLevel>(3, _omitFieldNames ? '' : 'getLevel', subBuilder: GetLevel.create)
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

  @$pb.TagNumber(2)
  InfoRequest get info => $_getN(1);
  @$pb.TagNumber(2)
  set info(InfoRequest v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasInfo() => $_has(1);
  @$pb.TagNumber(2)
  void clearInfo() => clearField(2);
  @$pb.TagNumber(2)
  InfoRequest ensureInfo() => $_ensure(1);

  @$pb.TagNumber(3)
  GetLevel get getLevel => $_getN(2);
  @$pb.TagNumber(3)
  set getLevel(GetLevel v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasGetLevel() => $_has(2);
  @$pb.TagNumber(3)
  void clearGetLevel() => clearField(3);
  @$pb.TagNumber(3)
  GetLevel ensureGetLevel() => $_ensure(2);
}

enum GameResponse_Message {
  heartbeat,
  log,
  info,
  level,
  notSet
}

class GameResponse extends $pb.GeneratedMessage {
  factory GameResponse({
    $core.String? serverId,
    Heartbeat? heartbeat,
    Log? log,
    InfoResponse? info,
    Level? level,
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
    if (info != null) {
      $result.info = info;
    }
    if (level != null) {
      $result.level = level;
    }
    return $result;
  }
  GameResponse._() : super();
  factory GameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static const $core.Map<$core.int, GameResponse_Message> _GameResponse_MessageByTag = {
    2 : GameResponse_Message.heartbeat,
    3 : GameResponse_Message.log,
    4 : GameResponse_Message.info,
    5 : GameResponse_Message.level,
    0 : GameResponse_Message.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GameResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'endless'), createEmptyInstance: create)
    ..oo(0, [2, 3, 4, 5])
    ..aOS(1, _omitFieldNames ? '' : 'serverId')
    ..aOM<Heartbeat>(2, _omitFieldNames ? '' : 'heartbeat', subBuilder: Heartbeat.create)
    ..aOM<Log>(3, _omitFieldNames ? '' : 'log', subBuilder: Log.create)
    ..aOM<InfoResponse>(4, _omitFieldNames ? '' : 'info', subBuilder: InfoResponse.create)
    ..aOM<Level>(5, _omitFieldNames ? '' : 'level', subBuilder: Level.create)
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
  Log get log => $_getN(2);
  @$pb.TagNumber(3)
  set log(Log v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasLog() => $_has(2);
  @$pb.TagNumber(3)
  void clearLog() => clearField(3);
  @$pb.TagNumber(3)
  Log ensureLog() => $_ensure(2);

  @$pb.TagNumber(4)
  InfoResponse get info => $_getN(3);
  @$pb.TagNumber(4)
  set info(InfoResponse v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasInfo() => $_has(3);
  @$pb.TagNumber(4)
  void clearInfo() => clearField(4);
  @$pb.TagNumber(4)
  InfoResponse ensureInfo() => $_ensure(3);

  @$pb.TagNumber(5)
  Level get level => $_getN(4);
  @$pb.TagNumber(5)
  set level(Level v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasLevel() => $_has(4);
  @$pb.TagNumber(5)
  void clearLevel() => clearField(5);
  @$pb.TagNumber(5)
  Level ensureLevel() => $_ensure(4);
}

class AddTile extends $pb.GeneratedMessage {
  factory AddTile({
    Tile? tile,
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
    ..aOM<Tile>(1, _omitFieldNames ? '' : 'tile', subBuilder: Tile.create)
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
  Tile get tile => $_getN(0);
  @$pb.TagNumber(1)
  set tile(Tile v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTile() => $_has(0);
  @$pb.TagNumber(1)
  void clearTile() => clearField(1);
  @$pb.TagNumber(1)
  Tile ensureTile() => $_ensure(0);
}

class RemoveTile extends $pb.GeneratedMessage {
  factory RemoveTile({
    Coordinate? coords,
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
    ..aOM<Coordinate>(1, _omitFieldNames ? '' : 'coords', subBuilder: Coordinate.create)
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
  Coordinate get coords => $_getN(0);
  @$pb.TagNumber(1)
  set coords(Coordinate v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCoords() => $_has(0);
  @$pb.TagNumber(1)
  void clearCoords() => clearField(1);
  @$pb.TagNumber(1)
  Coordinate ensureCoords() => $_ensure(0);
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
    Log? log,
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
    ..aOM<Log>(2, _omitFieldNames ? '' : 'log', subBuilder: Log.create)
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
  Log get log => $_getN(1);
  @$pb.TagNumber(2)
  set log(Log v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasLog() => $_has(1);
  @$pb.TagNumber(2)
  void clearLog() => clearField(2);
  @$pb.TagNumber(2)
  Log ensureLog() => $_ensure(1);

  @$pb.TagNumber(3)
  Result get result => $_getN(2);
  @$pb.TagNumber(3)
  set result(Result v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasResult() => $_has(2);
  @$pb.TagNumber(3)
  void clearResult() => clearField(3);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
