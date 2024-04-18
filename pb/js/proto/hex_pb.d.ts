// @generated by protoc-gen-es v1.8.0
// @generated from file proto/hex.proto (package endless, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Type defines whether the tile is empty, a floor tile, or a wall tile.
 *
 * @generated from enum endless.Type
 */
export declare enum Type {
  /**
   * Empty is a tile without anything in it.
   *
   * @generated from enum value: Empty = 0;
   */
  Empty = 0,

  /**
   * Floor is a tile that can be walked on.
   *
   * @generated from enum value: Floor = 1;
   */
  Floor = 1,

  /**
   * Wall is a tile that blocks movement and line-of-sight.
   *
   * @generated from enum value: Wall = 2;
   */
  Wall = 2,
}

/**
 * Tileset defines what tileset to use in Godot.
 *
 * @generated from enum endless.Tileset
 */
export declare enum Tileset {
  /**
   * @generated from enum value: Dungeon = 0;
   */
  Dungeon = 0,

  /**
   * @generated from enum value: Woods = 1;
   */
  Woods = 1,
}

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
 * @generated from enum endless.Result
 */
export declare enum Result {
  /**
   * @generated from enum value: Unknown = 0;
   */
  Unknown = 0,

  /**
   * @generated from enum value: Failure = 1;
   */
  Failure = 1,

  /**
   * @generated from enum value: Success = 2;
   */
  Success = 2,
}

/**
 * Coordinate is the x,y location in the world for a tile.
 *
 * @generated from message endless.Coordinate
 */
export declare class Coordinate extends Message<Coordinate> {
  /**
   * @generated from field: int32 x = 1;
   */
  x: number;

  /**
   * @generated from field: int32 y = 2;
   */
  y: number;

  constructor(data?: PartialMessage<Coordinate>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.Coordinate";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Coordinate;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Coordinate;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Coordinate;

  static equals(a: Coordinate | PlainMessage<Coordinate> | undefined, b: Coordinate | PlainMessage<Coordinate> | undefined): boolean;
}

/**
 * Tile is a single tile within the world.
 *
 * @generated from message endless.Tile
 */
export declare class Tile extends Message<Tile> {
  /**
   * Type defines whether the tile is empty, a floor, or a wall.
   *
   * @generated from field: endless.Type type = 1;
   */
  type: Type;

  /**
   * Coords defines the position of the tile within the world.
   *
   * @generated from field: endless.Coordinate coords = 2;
   */
  coords?: Coordinate;

  constructor(data?: PartialMessage<Tile>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.Tile";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Tile;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Tile;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Tile;

  static equals(a: Tile | PlainMessage<Tile> | undefined, b: Tile | PlainMessage<Tile> | undefined): boolean;
}

/**
 * Level is a single layer of tiles laid out to create a level.
 *
 * @generated from message endless.Level
 */
export declare class Level extends Message<Level> {
  /**
   * Tiles is an array of all the tiles laid out on a level.
   *
   * @generated from field: repeated endless.Tile tiles = 1;
   */
  tiles: Tile[];

  /**
   * Tileset tells Godot what tileset to use.
   *
   * @generated from field: endless.Tileset tileset = 2;
   */
  tileset: Tileset;

  constructor(data?: PartialMessage<Level>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.Level";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Level;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Level;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Level;

  static equals(a: Level | PlainMessage<Level> | undefined, b: Level | PlainMessage<Level> | undefined): boolean;
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

/**
 * @generated from message endless.InfoRequest
 */
export declare class InfoRequest extends Message<InfoRequest> {
  constructor(data?: PartialMessage<InfoRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.InfoRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InfoRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InfoRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InfoRequest;

  static equals(a: InfoRequest | PlainMessage<InfoRequest> | undefined, b: InfoRequest | PlainMessage<InfoRequest> | undefined): boolean;
}

/**
 * @generated from message endless.InfoResponse
 */
export declare class InfoResponse extends Message<InfoResponse> {
  /**
   * @generated from field: string version = 1;
   */
  version: string;

  /**
   * @generated from field: string build_date = 2;
   */
  buildDate: string;

  constructor(data?: PartialMessage<InfoResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.InfoResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InfoResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InfoResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InfoResponse;

  static equals(a: InfoResponse | PlainMessage<InfoResponse> | undefined, b: InfoResponse | PlainMessage<InfoResponse> | undefined): boolean;
}

/**
 * @generated from message endless.GetLevel
 */
export declare class GetLevel extends Message<GetLevel> {
  constructor(data?: PartialMessage<GetLevel>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.GetLevel";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLevel;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLevel;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLevel;

  static equals(a: GetLevel | PlainMessage<GetLevel> | undefined, b: GetLevel | PlainMessage<GetLevel> | undefined): boolean;
}

/**
 * @generated from message endless.GameRequest
 */
export declare class GameRequest extends Message<GameRequest> {
  /**
   * @generated from field: string client_id = 1;
   */
  clientId: string;

  /**
   * @generated from oneof endless.GameRequest.request
   */
  request: {
    /**
     * @generated from field: endless.InfoRequest info = 2;
     */
    value: InfoRequest;
    case: "info";
  } | {
    /**
     * @generated from field: endless.GetLevel get_level = 3;
     */
    value: GetLevel;
    case: "getLevel";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<GameRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.GameRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GameRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GameRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GameRequest;

  static equals(a: GameRequest | PlainMessage<GameRequest> | undefined, b: GameRequest | PlainMessage<GameRequest> | undefined): boolean;
}

/**
 * @generated from message endless.GameResponse
 */
export declare class GameResponse extends Message<GameResponse> {
  /**
   * @generated from field: string server_id = 1;
   */
  serverId: string;

  /**
   * @generated from oneof endless.GameResponse.msesage
   */
  msesage: {
    /**
     * @generated from field: endless.Log log = 2;
     */
    value: Log;
    case: "log";
  } | {
    /**
     * @generated from field: endless.InfoResponse info = 3;
     */
    value: InfoResponse;
    case: "info";
  } | {
    /**
     * @generated from field: endless.Level level = 4;
     */
    value: Level;
    case: "level";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<GameResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.GameResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GameResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GameResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GameResponse;

  static equals(a: GameResponse | PlainMessage<GameResponse> | undefined, b: GameResponse | PlainMessage<GameResponse> | undefined): boolean;
}

/**
 * @generated from message endless.AddTile
 */
export declare class AddTile extends Message<AddTile> {
  /**
   * @generated from field: endless.Tile tile = 1;
   */
  tile?: Tile;

  constructor(data?: PartialMessage<AddTile>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.AddTile";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddTile;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddTile;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddTile;

  static equals(a: AddTile | PlainMessage<AddTile> | undefined, b: AddTile | PlainMessage<AddTile> | undefined): boolean;
}

/**
 * @generated from message endless.RemoveTile
 */
export declare class RemoveTile extends Message<RemoveTile> {
  /**
   * @generated from field: endless.Coordinate coords = 1;
   */
  coords?: Coordinate;

  constructor(data?: PartialMessage<RemoveTile>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.RemoveTile";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveTile;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveTile;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveTile;

  static equals(a: RemoveTile | PlainMessage<RemoveTile> | undefined, b: RemoveTile | PlainMessage<RemoveTile> | undefined): boolean;
}

/**
 * @generated from message endless.AdminRequest
 */
export declare class AdminRequest extends Message<AdminRequest> {
  /**
   * @generated from field: string client_id = 1;
   */
  clientId: string;

  /**
   * @generated from oneof endless.AdminRequest.request
   */
  request: {
    /**
     * @generated from field: endless.AddTile add_tile = 2;
     */
    value: AddTile;
    case: "addTile";
  } | {
    /**
     * @generated from field: endless.RemoveTile remove_tile = 3;
     */
    value: RemoveTile;
    case: "removeTile";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<AdminRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.AdminRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdminRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdminRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdminRequest;

  static equals(a: AdminRequest | PlainMessage<AdminRequest> | undefined, b: AdminRequest | PlainMessage<AdminRequest> | undefined): boolean;
}

/**
 * @generated from message endless.AdminResponse
 */
export declare class AdminResponse extends Message<AdminResponse> {
  /**
   * @generated from field: string server_id = 1;
   */
  serverId: string;

  /**
   * @generated from field: endless.Log log = 2;
   */
  log?: Log;

  /**
   * @generated from field: endless.Result result = 3;
   */
  result: Result;

  constructor(data?: PartialMessage<AdminResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.AdminResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdminResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdminResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdminResponse;

  static equals(a: AdminResponse | PlainMessage<AdminResponse> | undefined, b: AdminResponse | PlainMessage<AdminResponse> | undefined): boolean;
}
