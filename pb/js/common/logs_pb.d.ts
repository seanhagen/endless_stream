// @generated by protoc-gen-es v1.9.0
// @generated from file common/logs.proto (package endless, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum endless.LogLevel
 */
export declare enum LogLevel {
  /**
   * @generated from enum value: Info = 0;
   */
  Info = 0,

  /**
   * @generated from enum value: Debug = -1;
   */
  Debug = -1,

  /**
   * @generated from enum value: Warn = 2;
   */
  Warn = 2,

  /**
   * @generated from enum value: Error = 3;
   */
  Error = 3,

  /**
   * @generated from enum value: Fatal = 4;
   */
  Fatal = 4,
}

/**
 * @generated from enum endless.LogSource
 */
export declare enum LogSource {
  /**
   * @generated from enum value: Server = 0;
   */
  Server = 0,

  /**
   * @generated from enum value: Player = 100;
   */
  Player = 100,

  /**
   * @generated from enum value: Game = 200;
   */
  Game = 200,
}

/**
 * @generated from message endless.Log
 */
export declare class Log extends Message<Log> {
  /**
   * @generated from field: string msg = 1;
   */
  msg: string;

  /**
   * @generated from field: google.protobuf.Timestamp at = 2;
   */
  at?: Timestamp;

  /**
   * @generated from field: endless.LogLevel level = 3;
   */
  level: LogLevel;

  /**
   * @generated from field: endless.LogSource source = 4;
   */
  source: LogSource;

  constructor(data?: PartialMessage<Log>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.Log";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Log;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Log;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Log;

  static equals(a: Log | PlainMessage<Log> | undefined, b: Log | PlainMessage<Log> | undefined): boolean;
}
