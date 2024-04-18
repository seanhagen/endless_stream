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

class LogLevel extends $pb.ProtobufEnum {
  static const LogLevel Info = LogLevel._(0, _omitEnumNames ? '' : 'Info');
  static const LogLevel Debug = LogLevel._(-1, _omitEnumNames ? '' : 'Debug');
  static const LogLevel Warn = LogLevel._(2, _omitEnumNames ? '' : 'Warn');
  static const LogLevel Error = LogLevel._(3, _omitEnumNames ? '' : 'Error');
  static const LogLevel Fatal = LogLevel._(4, _omitEnumNames ? '' : 'Fatal');

  static const $core.List<LogLevel> values = <LogLevel> [
    Info,
    Debug,
    Warn,
    Error,
    Fatal,
  ];

  static final $core.Map<$core.int, LogLevel> _byValue = $pb.ProtobufEnum.initByValue(values);
  static LogLevel? valueOf($core.int value) => _byValue[value];

  const LogLevel._($core.int v, $core.String n) : super(v, n);
}

class LogSource extends $pb.ProtobufEnum {
  static const LogSource Server = LogSource._(0, _omitEnumNames ? '' : 'Server');
  static const LogSource Player = LogSource._(100, _omitEnumNames ? '' : 'Player');
  static const LogSource Game = LogSource._(200, _omitEnumNames ? '' : 'Game');

  static const $core.List<LogSource> values = <LogSource> [
    Server,
    Player,
    Game,
  ];

  static final $core.Map<$core.int, LogSource> _byValue = $pb.ProtobufEnum.initByValue(values);
  static LogSource? valueOf($core.int value) => _byValue[value];

  const LogSource._($core.int v, $core.String n) : super(v, n);
}

class Result extends $pb.ProtobufEnum {
  static const Result Unknown = Result._(0, _omitEnumNames ? '' : 'Unknown');
  static const Result Failure = Result._(1, _omitEnumNames ? '' : 'Failure');
  static const Result Success = Result._(2, _omitEnumNames ? '' : 'Success');

  static const $core.List<Result> values = <Result> [
    Unknown,
    Failure,
    Success,
  ];

  static final $core.Map<$core.int, Result> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Result? valueOf($core.int value) => _byValue[value];

  const Result._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
