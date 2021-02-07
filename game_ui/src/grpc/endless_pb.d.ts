// package: endless.stream.v1
// file: endless.proto

import * as jspb from "google-protobuf";
import * as input_pb from "./input_pb";
import * as output_pb from "./output_pb";

export class CreateGame extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateGame.AsObject;
  static toObject(includeInstance: boolean, msg: CreateGame): CreateGame.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateGame, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateGame;
  static deserializeBinaryFromReader(message: CreateGame, reader: jspb.BinaryReader): CreateGame;
}

export namespace CreateGame {
  export type AsObject = {
  }
}

export class GameCreated extends jspb.Message {
  getCode(): string;
  setCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameCreated.AsObject;
  static toObject(includeInstance: boolean, msg: GameCreated): GameCreated.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameCreated, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameCreated;
  static deserializeBinaryFromReader(message: GameCreated, reader: jspb.BinaryReader): GameCreated;
}

export namespace GameCreated {
  export type AsObject = {
    code: string,
  }
}

