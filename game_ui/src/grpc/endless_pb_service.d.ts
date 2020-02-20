// package: endless.stream.v1
// file: endless.proto

import * as endless_pb from "./endless_pb";
import * as input_pb from "./input_pb";
import * as output_pb from "./output_pb";
import {grpc} from "@improbable-eng/grpc-web";

type GameCreate = {
  readonly methodName: string;
  readonly service: typeof Game;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof endless_pb.CreateGame;
  readonly responseType: typeof endless_pb.GameCreated;
};

type GameState = {
  readonly methodName: string;
  readonly service: typeof Game;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof input_pb.Input;
  readonly responseType: typeof output_pb.Output;
};

type GameStateOut = {
  readonly methodName: string;
  readonly service: typeof Game;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof input_pb.Input;
  readonly responseType: typeof output_pb.Output;
};

type GameStateIn = {
  readonly methodName: string;
  readonly service: typeof Game;
  readonly requestStream: true;
  readonly responseStream: false;
  readonly requestType: typeof input_pb.Input;
  readonly responseType: typeof output_pb.Output;
};

export class Game {
  static readonly serviceName: string;
  static readonly Create: GameCreate;
  static readonly State: GameState;
  static readonly StateOut: GameStateOut;
  static readonly StateIn: GameStateIn;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class GameClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  create(
    requestMessage: endless_pb.CreateGame,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: endless_pb.GameCreated|null) => void
  ): UnaryResponse;
  create(
    requestMessage: endless_pb.CreateGame,
    callback: (error: ServiceError|null, responseMessage: endless_pb.GameCreated|null) => void
  ): UnaryResponse;
  state(metadata?: grpc.Metadata): BidirectionalStream<input_pb.Input, output_pb.Output>;
  stateOut(requestMessage: input_pb.Input, metadata?: grpc.Metadata): ResponseStream<output_pb.Output>;
  stateIn(metadata?: grpc.Metadata): RequestStream<input_pb.Input>;
}

