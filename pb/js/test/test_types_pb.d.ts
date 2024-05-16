// @generated by protoc-gen-es v1.9.0
// @generated from file test/test_types.proto (package endless, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * these types are just for testing, they shouldn't be used in the actual game client
 *
 * @generated from message endless.PingReq
 */
export declare class PingReq extends Message<PingReq> {
  /**
   * @generated from field: string msg = 1;
   */
  msg: string;

  constructor(data?: PartialMessage<PingReq>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.PingReq";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PingReq;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PingReq;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PingReq;

  static equals(a: PingReq | PlainMessage<PingReq> | undefined, b: PingReq | PlainMessage<PingReq> | undefined): boolean;
}

/**
 * @generated from message endless.PongResp
 */
export declare class PongResp extends Message<PongResp> {
  /**
   * gsm is the Ping.msg msg backwards
   *
   * @generated from field: string gsm = 1;
   */
  gsm: string;

  constructor(data?: PartialMessage<PongResp>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.PongResp";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PongResp;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PongResp;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PongResp;

  static equals(a: PongResp | PlainMessage<PongResp> | undefined, b: PongResp | PlainMessage<PongResp> | undefined): boolean;
}

/**
 * @generated from message endless.TestStreamRequest
 */
export declare class TestStreamRequest extends Message<TestStreamRequest> {
  /**
   * @generated from field: int32 chunk_id = 1;
   */
  chunkId: number;

  /**
   * @generated from field: string msg = 2;
   */
  msg: string;

  constructor(data?: PartialMessage<TestStreamRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.TestStreamRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TestStreamRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TestStreamRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TestStreamRequest;

  static equals(a: TestStreamRequest | PlainMessage<TestStreamRequest> | undefined, b: TestStreamRequest | PlainMessage<TestStreamRequest> | undefined): boolean;
}

/**
 * @generated from message endless.TestStreamResponse
 */
export declare class TestStreamResponse extends Message<TestStreamResponse> {
  /**
   * @generated from field: int32 resp_id = 1;
   */
  respId: number;

  /**
   * @generated from field: string gsm = 2;
   */
  gsm: string;

  constructor(data?: PartialMessage<TestStreamResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.TestStreamResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TestStreamResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TestStreamResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TestStreamResponse;

  static equals(a: TestStreamResponse | PlainMessage<TestStreamResponse> | undefined, b: TestStreamResponse | PlainMessage<TestStreamResponse> | undefined): boolean;
}

/**
 * @generated from message endless.TestRequest
 */
export declare class TestRequest extends Message<TestRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<TestRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.TestRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TestRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TestRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TestRequest;

  static equals(a: TestRequest | PlainMessage<TestRequest> | undefined, b: TestRequest | PlainMessage<TestRequest> | undefined): boolean;
}

/**
 * @generated from message endless.TestResponse
 */
export declare class TestResponse extends Message<TestResponse> {
  /**
   * @generated from field: string resp = 1;
   */
  resp: string;

  constructor(data?: PartialMessage<TestResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "endless.TestResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TestResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TestResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TestResponse;

  static equals(a: TestResponse | PlainMessage<TestResponse> | undefined, b: TestResponse | PlainMessage<TestResponse> | undefined): boolean;
}
