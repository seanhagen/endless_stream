// package: endless.stream.v1
// file: util.proto

import * as jspb from "google-protobuf";

export class Class extends jspb.Message {
  getClass(): ClassType;
  setClass(value: ClassType): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Class.AsObject;
  static toObject(includeInstance: boolean, msg: Class): Class.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Class, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Class;
  static deserializeBinaryFromReader(message: Class, reader: jspb.BinaryReader): Class;
}

export namespace Class {
  export type AsObject = {
    pb_class: ClassType,
  }
}

export enum Type {
  NONE = 0,
  ANIMAL = 10,
  BIRD = 11,
  FISH = 12,
  RAT = 13,
  WOLF = 14,
  INSECT = 15,
  SPIDER = 16,
  DINOSAUR = 20,
  DRAGON = 30,
  ABOMINATION = 40,
  EYE = 50,
  FAE = 60,
  PLANT = 70,
  FUNGUS = 71,
  GOBLIN = 80,
  OGRE = 81,
  TROLL = 82,
  CONSTRUCT = 90,
  GOLEM = 91,
  HYBRID = 92,
  HUMANOID = 100,
  HUMAN = 101,
  SHAPESHIFTER = 102,
  WITCH = 103,
  NAGA = 110,
  SLIME = 120,
  UNDEAD = 130,
  SPIRIT = 131,
  VAMPIRE = 132,
  ELDRITCH = 140,
}

export enum ClassType {
  UNKNOWN = 0,
  STATUS = -2,
  AUDIENCE = -1,
  FIGHTER = 10,
  RANGER = 20,
  CLERIC = 30,
  WIZARD = 40,
}

export enum StatusEffect {
  NORMAL = 0,
  POISONED = 1,
  STUNNED = 2,
  PRONE = 3,
  BLEEDING = 4,
  FRENZIED = 5,
  INVISIBLE = 6,
  INVINCIBLE = 7,
}

export enum Level {
  BLANK = 0,
  FOREST = 1,
  CAVE = 2,
  DUNGEON = 3,
  ICE = 4,
  FIRE = 5,
  VOID = 6,
}

export enum Display {
  SCREENLOADING = 0,
  SCREENCHARSELECT = 1,
  SCREENWAVE = 2,
  SCREENVICTORY = 3,
  SCREENDEAD = 4,
  SCREENGAMEOVER = 5,
  SCREENSTORE = 6,
  SCREENNEWWAVE = 7,
}

