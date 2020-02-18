// package: endless.stream.v1
// file: endless.proto

var endless_pb = require("./endless_pb");
var input_pb = require("./input_pb");
var output_pb = require("./output_pb");
var grpc = require("grpc-web-client").grpc;

var Game = (function () {
  function Game() {}
  Game.serviceName = "endless.stream.v1.Game";
  return Game;
}());

Game.Create = {
  methodName: "Create",
  service: Game,
  requestStream: false,
  responseStream: false,
  requestType: endless_pb.CreateGame,
  responseType: endless_pb.GameCreated
};

Game.State = {
  methodName: "State",
  service: Game,
  requestStream: true,
  responseStream: true,
  requestType: input_pb.Input,
  responseType: output_pb.Output
};

exports.Game = Game;

function GameClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

GameClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Game.Create, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

GameClient.prototype.state = function state(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(Game.State, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.end.forEach(function (handler) {
      handler();
    });
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.GameClient = GameClient;

