// source: util.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.endless.stream.v1.Class', null, global);
goog.exportSymbol('proto.endless.stream.v1.ClassType', null, global);
goog.exportSymbol('proto.endless.stream.v1.Display', null, global);
goog.exportSymbol('proto.endless.stream.v1.Level', null, global);
goog.exportSymbol('proto.endless.stream.v1.StatusEffect', null, global);
goog.exportSymbol('proto.endless.stream.v1.Type', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.endless.stream.v1.Class = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.endless.stream.v1.Class, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.endless.stream.v1.Class.displayName = 'proto.endless.stream.v1.Class';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.endless.stream.v1.Class.prototype.toObject = function(opt_includeInstance) {
  return proto.endless.stream.v1.Class.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.endless.stream.v1.Class} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.endless.stream.v1.Class.toObject = function(includeInstance, msg) {
  var f, obj = {
    pb_class: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.endless.stream.v1.Class}
 */
proto.endless.stream.v1.Class.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.endless.stream.v1.Class;
  return proto.endless.stream.v1.Class.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.endless.stream.v1.Class} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.endless.stream.v1.Class}
 */
proto.endless.stream.v1.Class.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.endless.stream.v1.ClassType} */ (reader.readEnum());
      msg.setClass(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.endless.stream.v1.Class.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.endless.stream.v1.Class.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.endless.stream.v1.Class} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.endless.stream.v1.Class.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClass();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
};


/**
 * optional ClassType class = 1;
 * @return {!proto.endless.stream.v1.ClassType}
 */
proto.endless.stream.v1.Class.prototype.getClass = function() {
  return /** @type {!proto.endless.stream.v1.ClassType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.endless.stream.v1.ClassType} value
 * @return {!proto.endless.stream.v1.Class} returns this
 */
proto.endless.stream.v1.Class.prototype.setClass = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * @enum {number}
 */
proto.endless.stream.v1.Type = {
  NONE: 0,
  ANIMAL: 10,
  BIRD: 11,
  FISH: 12,
  RAT: 13,
  WOLF: 14,
  INSECT: 15,
  SPIDER: 16,
  DINOSAUR: 20,
  DRAGON: 30,
  ABOMINATION: 40,
  EYE: 50,
  FAE: 60,
  PLANT: 70,
  FUNGUS: 71,
  GOBLIN: 80,
  OGRE: 81,
  TROLL: 82,
  CONSTRUCT: 90,
  GOLEM: 91,
  HYBRID: 92,
  HUMANOID: 100,
  HUMAN: 101,
  SHAPESHIFTER: 102,
  WITCH: 103,
  NAGA: 110,
  SLIME: 120,
  UNDEAD: 130,
  SPIRIT: 131,
  VAMPIRE: 132,
  ELDRITCH: 140,
  HUMANPLAYER: 9999
};

/**
 * @enum {number}
 */
proto.endless.stream.v1.ClassType = {
  UNKNOWN: 0,
  STATUS: -2,
  AUDIENCE: -1,
  FIGHTER: 10,
  RANGER: 20,
  CLERIC: 30,
  WIZARD: 40
};

/**
 * @enum {number}
 */
proto.endless.stream.v1.StatusEffect = {
  NORMAL: 0,
  POISONED: 1,
  STUNNED: 2,
  PRONE: 3,
  BLEEDING: 4,
  FRENZIED: 5,
  INVISIBLE: 6,
  INVINCIBLE: 7
};

/**
 * @enum {number}
 */
proto.endless.stream.v1.Level = {
  BLANK: 0,
  FOREST: 1,
  CAVE: 2,
  DUNGEON: 3,
  ICE: 4,
  FIRE: 5,
  VOID: 6
};

/**
 * @enum {number}
 */
proto.endless.stream.v1.Display = {
  SCREENLOADING: 0,
  SCREENCHARSELECT: 1,
  SCREENWAVE: 2,
  SCREENVICTORY: 3,
  SCREENDEAD: 4,
  SCREENGAMEOVER: 5,
  SCREENSTORE: 6,
  SCREENNEWWAVE: 7
};

goog.object.extend(exports, proto.endless.stream.v1);
