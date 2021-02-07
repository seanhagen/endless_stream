/**
 * @fileoverview gRPC-Web generated client stub for endless.stream.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var input_pb = require('./input_pb.js')

var output_pb = require('./output_pb.js')
const proto = {};
proto.endless = {};
proto.endless.stream = {};
proto.endless.stream.v1 = require('./endless_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.endless.stream.v1.GameClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.endless.stream.v1.GamePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.endless.stream.v1.CreateGame,
 *   !proto.endless.stream.v1.GameCreated>}
 */
const methodDescriptor_Game_Create = new grpc.web.MethodDescriptor(
  '/endless.stream.v1.Game/Create',
  grpc.web.MethodType.UNARY,
  proto.endless.stream.v1.CreateGame,
  proto.endless.stream.v1.GameCreated,
  /**
   * @param {!proto.endless.stream.v1.CreateGame} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.endless.stream.v1.GameCreated.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.endless.stream.v1.CreateGame,
 *   !proto.endless.stream.v1.GameCreated>}
 */
const methodInfo_Game_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.endless.stream.v1.GameCreated,
  /**
   * @param {!proto.endless.stream.v1.CreateGame} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.endless.stream.v1.GameCreated.deserializeBinary
);


/**
 * @param {!proto.endless.stream.v1.CreateGame} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.endless.stream.v1.GameCreated)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.endless.stream.v1.GameCreated>|undefined}
 *     The XHR Node Readable Stream
 */
proto.endless.stream.v1.GameClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/endless.stream.v1.Game/Create',
      request,
      metadata || {},
      methodDescriptor_Game_Create,
      callback);
};


/**
 * @param {!proto.endless.stream.v1.CreateGame} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.endless.stream.v1.GameCreated>}
 *     A native promise that resolves to the response
 */
proto.endless.stream.v1.GamePromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/endless.stream.v1.Game/Create',
      request,
      metadata || {},
      methodDescriptor_Game_Create);
};


module.exports = proto.endless.stream.v1;

