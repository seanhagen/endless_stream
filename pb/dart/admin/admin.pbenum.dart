//
//  Generated code. Do not modify.
//  source: admin/admin.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

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
