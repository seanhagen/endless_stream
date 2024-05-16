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

/// Type defines whether the tile is empty, a floor tile, or a wall tile.
class Type extends $pb.ProtobufEnum {
  static const Type Empty = Type._(0, _omitEnumNames ? '' : 'Empty');
  static const Type Floor = Type._(1, _omitEnumNames ? '' : 'Floor');
  static const Type Wall = Type._(2, _omitEnumNames ? '' : 'Wall');

  static const $core.List<Type> values = <Type> [
    Empty,
    Floor,
    Wall,
  ];

  static final $core.Map<$core.int, Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Type? valueOf($core.int value) => _byValue[value];

  const Type._($core.int v, $core.String n) : super(v, n);
}

/// Tileset defines what tileset to use in Godot.
class Tileset extends $pb.ProtobufEnum {
  static const Tileset Dungeon = Tileset._(0, _omitEnumNames ? '' : 'Dungeon');
  static const Tileset Woods = Tileset._(1, _omitEnumNames ? '' : 'Woods');

  static const $core.List<Tileset> values = <Tileset> [
    Dungeon,
    Woods,
  ];

  static final $core.Map<$core.int, Tileset> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Tileset? valueOf($core.int value) => _byValue[value];

  const Tileset._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
