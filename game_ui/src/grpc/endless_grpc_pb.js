// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var endless_pb = require('./endless_pb.js');
var input_pb = require('./input_pb.js');
var output_pb = require('./output_pb.js');

function serialize_endless_stream_v1_CreateGame(arg) {
  if (!(arg instanceof endless_pb.CreateGame)) {
    throw new Error('Expected argument of type endless.stream.v1.CreateGame');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_endless_stream_v1_CreateGame(buffer_arg) {
  return endless_pb.CreateGame.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_endless_stream_v1_GameCreated(arg) {
  if (!(arg instanceof endless_pb.GameCreated)) {
    throw new Error('Expected argument of type endless.stream.v1.GameCreated');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_endless_stream_v1_GameCreated(buffer_arg) {
  return endless_pb.GameCreated.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_endless_stream_v1_Input(arg) {
  if (!(arg instanceof input_pb.Input)) {
    throw new Error('Expected argument of type endless.stream.v1.Input');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_endless_stream_v1_Input(buffer_arg) {
  return input_pb.Input.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_endless_stream_v1_Output(arg) {
  if (!(arg instanceof output_pb.Output)) {
    throw new Error('Expected argument of type endless.stream.v1.Output');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_endless_stream_v1_Output(buffer_arg) {
  return output_pb.Output.deserializeBinary(new Uint8Array(buffer_arg));
}


var GameService = exports.GameService = {
  create: {
    path: '/endless.stream.v1.Game/Create',
    requestStream: false,
    responseStream: false,
    requestType: endless_pb.CreateGame,
    responseType: endless_pb.GameCreated,
    requestSerialize: serialize_endless_stream_v1_CreateGame,
    requestDeserialize: deserialize_endless_stream_v1_CreateGame,
    responseSerialize: serialize_endless_stream_v1_GameCreated,
    responseDeserialize: deserialize_endless_stream_v1_GameCreated,
  },
  state: {
    path: '/endless.stream.v1.Game/State',
    requestStream: true,
    responseStream: true,
    requestType: input_pb.Input,
    responseType: output_pb.Output,
    requestSerialize: serialize_endless_stream_v1_Input,
    requestDeserialize: deserialize_endless_stream_v1_Input,
    responseSerialize: serialize_endless_stream_v1_Output,
    responseDeserialize: deserialize_endless_stream_v1_Output,
  },
};

exports.GameClient = grpc.makeGenericClientConstructor(GameService);
