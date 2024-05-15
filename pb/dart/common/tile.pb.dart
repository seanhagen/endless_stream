//
//  Generated code. Do not modify.
//  source: common/tile.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'tile.pbenum.dart';

export 'tile.pbenum.dart';

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


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
