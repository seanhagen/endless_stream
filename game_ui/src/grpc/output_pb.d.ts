// package: endless.stream.v1
// file: output.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as util_pb from "./util_pb";

export class Skill extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getLevel(): number;
  setLevel(value: number): void;

  getTarget(): TargetMap[keyof TargetMap];
  setTarget(value: TargetMap[keyof TargetMap]): void;

  getCost(): number;
  setCost(value: number): void;

  getCooldown(): number;
  setCooldown(value: number): void;

  getCooldownLeft(): number;
  setCooldownLeft(value: number): void;

  getActive(): boolean;
  setActive(value: boolean): void;

  getUpgradable(): boolean;
  setUpgradable(value: boolean): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Skill.AsObject;
  static toObject(includeInstance: boolean, msg: Skill): Skill.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Skill, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Skill;
  static deserializeBinaryFromReader(message: Skill, reader: jspb.BinaryReader): Skill;
}

export namespace Skill {
  export type AsObject = {
    id: string,
    name: string,
    level: number,
    target: TargetMap[keyof TargetMap],
    cost: number,
    cooldown: number,
    cooldownLeft: number,
    active: boolean,
    upgradable: boolean,
    description: string,
  }
}

export class Item extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getTarget(): TargetMap[keyof TargetMap];
  setTarget(value: TargetMap[keyof TargetMap]): void;

  getPassive(): boolean;
  setPassive(value: boolean): void;

  getBonusMap(): jspb.Map<string, number>;
  clearBonusMap(): void;
  getCount(): number;
  setCount(value: number): void;

  getActive(): boolean;
  setActive(value: boolean): void;

  getDescription(): string;
  setDescription(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Item.AsObject;
  static toObject(includeInstance: boolean, msg: Item): Item.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Item, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Item;
  static deserializeBinaryFromReader(message: Item, reader: jspb.BinaryReader): Item;
}

export namespace Item {
  export type AsObject = {
    id: string,
    name: string,
    target: TargetMap[keyof TargetMap],
    passive: boolean,
    bonusMap: Array<[string, number]>,
    count: number,
    active: boolean,
    description: string,
  }
}

export class Creature extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getMaxHp(): number;
  setMaxHp(value: number): void;

  getCurHp(): number;
  setCurHp(value: number): void;

  getPosition(): PositionMap[keyof PositionMap];
  setPosition(value: PositionMap[keyof PositionMap]): void;

  clearStatusesList(): void;
  getStatusesList(): Array<util_pb.StatusEffectMap[keyof util_pb.StatusEffectMap]>;
  setStatusesList(value: Array<util_pb.StatusEffectMap[keyof util_pb.StatusEffectMap]>): void;
  addStatuses(value: util_pb.StatusEffectMap[keyof util_pb.StatusEffectMap], index?: number): util_pb.StatusEffectMap[keyof util_pb.StatusEffectMap];

  getStrength(): number;
  setStrength(value: number): void;

  getVitality(): number;
  setVitality(value: number): void;

  getCombatDamageBase(): number;
  setCombatDamageBase(value: number): void;

  getVitalityRegen(): number;
  setVitalityRegen(value: number): void;

  getIntelligence(): number;
  setIntelligence(value: number): void;

  getFocus(): number;
  setFocus(value: number): void;

  getWillpower(): number;
  setWillpower(value: number): void;

  getFocusRegen(): number;
  setFocusRegen(value: number): void;

  getAgility(): number;
  setAgility(value: number): void;

  getEvasion(): number;
  setEvasion(value: number): void;

  getAccuracy(): number;
  setAccuracy(value: number): void;

  getInitiative(): number;
  setInitiative(value: number): void;

  getGold(): number;
  setGold(value: number): void;

  getXp(): number;
  setXp(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Creature.AsObject;
  static toObject(includeInstance: boolean, msg: Creature): Creature.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Creature, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Creature;
  static deserializeBinaryFromReader(message: Creature, reader: jspb.BinaryReader): Creature;
}

export namespace Creature {
  export type AsObject = {
    id: string,
    name: string,
    maxHp: number,
    curHp: number,
    position: PositionMap[keyof PositionMap],
    statusesList: Array<util_pb.StatusEffectMap[keyof util_pb.StatusEffectMap]>,
    strength: number,
    vitality: number,
    combatDamageBase: number,
    vitalityRegen: number,
    intelligence: number,
    focus: number,
    willpower: number,
    focusRegen: number,
    agility: number,
    evasion: number,
    accuracy: number,
    initiative: number,
    gold: number,
    xp: number,
  }
}

export class Player extends jspb.Message {
  hasBase(): boolean;
  clearBase(): void;
  getBase(): Creature | undefined;
  setBase(value?: Creature): void;

  hasCharacter(): boolean;
  clearCharacter(): void;
  getCharacter(): util_pb.Class | undefined;
  setCharacter(value?: util_pb.Class): void;

  getIsAi(): boolean;
  setIsAi(value: boolean): void;

  getLevel(): number;
  setLevel(value: number): void;

  getSkillsMap(): jspb.Map<string, Skill>;
  clearSkillsMap(): void;
  getInventoryMap(): jspb.Map<string, Item>;
  clearInventoryMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Player.AsObject;
  static toObject(includeInstance: boolean, msg: Player): Player.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Player, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Player;
  static deserializeBinaryFromReader(message: Player, reader: jspb.BinaryReader): Player;
}

export namespace Player {
  export type AsObject = {
    base?: Creature.AsObject,
    character?: util_pb.Class.AsObject,
    isAi: boolean,
    level: number,
    skillsMap: Array<[string, Skill.AsObject]>,
    inventoryMap: Array<[string, Item.AsObject]>,
  }
}

export class Monster extends jspb.Message {
  hasBase(): boolean;
  clearBase(): void;
  getBase(): Creature | undefined;
  setBase(value?: Creature): void;

  getType(): util_pb.TypeMap[keyof util_pb.TypeMap];
  setType(value: util_pb.TypeMap[keyof util_pb.TypeMap]): void;

  getIsFlying(): boolean;
  setIsFlying(value: boolean): void;

  getIsBoss(): boolean;
  setIsBoss(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Monster.AsObject;
  static toObject(includeInstance: boolean, msg: Monster): Monster.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Monster, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Monster;
  static deserializeBinaryFromReader(message: Monster, reader: jspb.BinaryReader): Monster;
}

export namespace Monster {
  export type AsObject = {
    base?: Creature.AsObject,
    type: util_pb.TypeMap[keyof util_pb.TypeMap],
    isFlying: boolean,
    isBoss: boolean,
  }
}

export class EventMessage extends jspb.Message {
  getMsgId(): number;
  setMsgId(value: number): void;

  getMsg(): string;
  setMsg(value: string): void;

  getIsError(): boolean;
  setIsError(value: boolean): void;

  getIsAlert(): boolean;
  setIsAlert(value: boolean): void;

  hasPlayerId(): boolean;
  clearPlayerId(): void;
  getPlayerId(): google_protobuf_wrappers_pb.StringValue | undefined;
  setPlayerId(value?: google_protobuf_wrappers_pb.StringValue): void;

  getLogOnly(): boolean;
  setLogOnly(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EventMessage.AsObject;
  static toObject(includeInstance: boolean, msg: EventMessage): EventMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EventMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EventMessage;
  static deserializeBinaryFromReader(message: EventMessage, reader: jspb.BinaryReader): EventMessage;
}

export namespace EventMessage {
  export type AsObject = {
    msgId: number,
    msg: string,
    isError: boolean,
    isAlert: boolean,
    playerId?: google_protobuf_wrappers_pb.StringValue.AsObject,
    logOnly: boolean,
  }
}

export class CharacterSelected extends jspb.Message {
  getSelectedMap(): jspb.Map<string, util_pb.ClassType[keyof util_pb.ClassType]>;
  clearSelectedMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharacterSelected.AsObject;
  static toObject(includeInstance: boolean, msg: CharacterSelected): CharacterSelected.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharacterSelected, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharacterSelected;
  static deserializeBinaryFromReader(message: CharacterSelected, reader: jspb.BinaryReader): CharacterSelected;
}

export namespace CharacterSelected {
  export type AsObject = {
    selectedMap: Array<[string, util_pb.ClassType[keyof util_pb.ClassType]]>,
  }
}

export class Wave extends jspb.Message {
  getNum(): number;
  setNum(value: number): void;

  getHasBoss(): boolean;
  setHasBoss(value: boolean): void;

  getLevel(): util_pb.LevelMap[keyof util_pb.LevelMap];
  setLevel(value: util_pb.LevelMap[keyof util_pb.LevelMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Wave.AsObject;
  static toObject(includeInstance: boolean, msg: Wave): Wave.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Wave, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Wave;
  static deserializeBinaryFromReader(message: Wave, reader: jspb.BinaryReader): Wave;
}

export namespace Wave {
  export type AsObject = {
    num: number,
    hasBoss: boolean,
    level: util_pb.LevelMap[keyof util_pb.LevelMap],
  }
}

export class CurrentState extends jspb.Message {
  clearMonstersList(): void;
  getMonstersList(): Array<Monster>;
  setMonstersList(value: Array<Monster>): void;
  addMonsters(value?: Monster, index?: number): Monster;

  getPlayersMap(): jspb.Map<string, Player>;
  clearPlayersMap(): void;
  hasCurrentPlayer(): boolean;
  clearCurrentPlayer(): void;
  getCurrentPlayer(): google_protobuf_wrappers_pb.StringValue | undefined;
  setCurrentPlayer(value?: google_protobuf_wrappers_pb.StringValue): void;

  getDisplay(): util_pb.DisplayMap[keyof util_pb.DisplayMap];
  setDisplay(value: util_pb.DisplayMap[keyof util_pb.DisplayMap]): void;

  hasCurrentWave(): boolean;
  clearCurrentWave(): void;
  getCurrentWave(): Wave | undefined;
  setCurrentWave(value?: Wave): void;

  getUpcomingWavesMap(): jspb.Map<number, Wave>;
  clearUpcomingWavesMap(): void;
  getAudienceCount(): number;
  setAudienceCount(value: number): void;

  hasSelected(): boolean;
  clearSelected(): void;
  getSelected(): CharacterSelected | undefined;
  setSelected(value?: CharacterSelected): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CurrentState.AsObject;
  static toObject(includeInstance: boolean, msg: CurrentState): CurrentState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CurrentState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CurrentState;
  static deserializeBinaryFromReader(message: CurrentState, reader: jspb.BinaryReader): CurrentState;
}

export namespace CurrentState {
  export type AsObject = {
    monstersList: Array<Monster.AsObject>,
    playersMap: Array<[string, Player.AsObject]>,
    currentPlayer?: google_protobuf_wrappers_pb.StringValue.AsObject,
    display: util_pb.DisplayMap[keyof util_pb.DisplayMap],
    currentWave?: Wave.AsObject,
    upcomingWavesMap: Array<[number, Wave.AsObject]>,
    audienceCount: number,
    selected?: CharacterSelected.AsObject,
  }
}

export class Tick extends jspb.Message {
  hasTime(): boolean;
  clearTime(): void;
  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getProgressMap(): jspb.Map<string, number>;
  clearProgressMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Tick.AsObject;
  static toObject(includeInstance: boolean, msg: Tick): Tick.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Tick, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Tick;
  static deserializeBinaryFromReader(message: Tick, reader: jspb.BinaryReader): Tick;
}

export namespace Tick {
  export type AsObject = {
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    progressMap: Array<[string, number]>,
  }
}

export class JoinedGame extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getAsAudience(): boolean;
  setAsAudience(value: boolean): void;

  getIsVip(): boolean;
  setIsVip(value: boolean): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinedGame.AsObject;
  static toObject(includeInstance: boolean, msg: JoinedGame): JoinedGame.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JoinedGame, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinedGame;
  static deserializeBinaryFromReader(message: JoinedGame, reader: jspb.BinaryReader): JoinedGame;
}

export namespace JoinedGame {
  export type AsObject = {
    id: string,
    asAudience: boolean,
    isVip: boolean,
    name: string,
  }
}

export class Action extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getActorId(): string;
  setActorId(value: string): void;

  clearTargetIdsList(): void;
  getTargetIdsList(): Array<string>;
  setTargetIdsList(value: Array<string>): void;
  addTargetIds(value: string, index?: number): string;

  getMsg(): string;
  setMsg(value: string): void;

  hasSkill(): boolean;
  clearSkill(): void;
  getSkill(): Skill | undefined;
  setSkill(value?: Skill): void;

  getValue(): number;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Action.AsObject;
  static toObject(includeInstance: boolean, msg: Action): Action.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Action, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Action;
  static deserializeBinaryFromReader(message: Action, reader: jspb.BinaryReader): Action;
}

export namespace Action {
  export type AsObject = {
    id: string,
    actorId: string,
    targetIdsList: Array<string>,
    msg: string,
    skill?: Skill.AsObject,
    value: number,
  }
}

export class StoreInventory extends jspb.Message {
  getInventoryMap(): jspb.Map<string, Item>;
  clearInventoryMap(): void;
  getPricesMap(): jspb.Map<string, number>;
  clearPricesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StoreInventory.AsObject;
  static toObject(includeInstance: boolean, msg: StoreInventory): StoreInventory.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StoreInventory, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StoreInventory;
  static deserializeBinaryFromReader(message: StoreInventory, reader: jspb.BinaryReader): StoreInventory;
}

export namespace StoreInventory {
  export type AsObject = {
    inventoryMap: Array<[string, Item.AsObject]>,
    pricesMap: Array<[string, number]>,
  }
}

export class Output extends jspb.Message {
  hasState(): boolean;
  clearState(): void;
  getState(): CurrentState | undefined;
  setState(value?: CurrentState): void;

  hasJoined(): boolean;
  clearJoined(): void;
  getJoined(): JoinedGame | undefined;
  setJoined(value?: JoinedGame): void;

  hasTick(): boolean;
  clearTick(): void;
  getTick(): Tick | undefined;
  setTick(value?: Tick): void;

  hasMsg(): boolean;
  clearMsg(): void;
  getMsg(): EventMessage | undefined;
  setMsg(value?: EventMessage): void;

  hasAction(): boolean;
  clearAction(): void;
  getAction(): Action | undefined;
  setAction(value?: Action): void;

  hasSelected(): boolean;
  clearSelected(): void;
  getSelected(): CharacterSelected | undefined;
  setSelected(value?: CharacterSelected): void;

  getDataCase(): Output.DataCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Output.AsObject;
  static toObject(includeInstance: boolean, msg: Output): Output.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Output, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Output;
  static deserializeBinaryFromReader(message: Output, reader: jspb.BinaryReader): Output;
}

export namespace Output {
  export type AsObject = {
    state?: CurrentState.AsObject,
    joined?: JoinedGame.AsObject,
    tick?: Tick.AsObject,
    msg?: EventMessage.AsObject,
    action?: Action.AsObject,
    selected?: CharacterSelected.AsObject,
  }

  export enum DataCase {
    DATA_NOT_SET = 0,
    STATE = 1,
    JOINED = 2,
    TICK = 3,
    MSG = 4,
    ACTION = 5,
    SELECTED = 6,
  }
}

export interface TargetMap {
  SELF: 0;
  MELEE: 1;
  RANGED: 2;
  AOE: 3;
}

export const Target: TargetMap;

export interface PositionMap {
  LEFT: 0;
  MIDDLE: 1;
  RIGHT: 2;
}

export const Position: PositionMap;

export interface ProgressBarTypeMap {
  COUNTDOWNGAMESTARTING: 0;
  COUNTDOWNTURNEND: 1;
}

export const ProgressBarType: ProgressBarTypeMap;

