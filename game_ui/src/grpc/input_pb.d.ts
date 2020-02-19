// package: endless.stream.v1
// file: input.proto

import * as jspb from "google-protobuf";
import * as util_pb from "./util_pb";

export class Register extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCode(): string;
  setCode(value: string): void;

  getName(): string;
  setName(value: string): void;

  getType(): ClientTypeMap[keyof ClientTypeMap];
  setType(value: ClientTypeMap[keyof ClientTypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Register.AsObject;
  static toObject(includeInstance: boolean, msg: Register): Register.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Register, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Register;
  static deserializeBinaryFromReader(message: Register, reader: jspb.BinaryReader): Register;
}

export namespace Register {
  export type AsObject = {
    id: string,
    code: string,
    name: string,
    type: ClientTypeMap[keyof ClientTypeMap],
  }
}

export class CharSelect extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  hasChoice(): boolean;
  clearChoice(): void;
  getChoice(): util_pb.Class | undefined;
  setChoice(value?: util_pb.Class): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CharSelect.AsObject;
  static toObject(includeInstance: boolean, msg: CharSelect): CharSelect.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CharSelect, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CharSelect;
  static deserializeBinaryFromReader(message: CharSelect, reader: jspb.BinaryReader): CharSelect;
}

export namespace CharSelect {
  export type AsObject = {
    playerId: string,
    choice?: util_pb.Class.AsObject,
  }
}

export class GameStart extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameStart.AsObject;
  static toObject(includeInstance: boolean, msg: GameStart): GameStart.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameStart, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameStart;
  static deserializeBinaryFromReader(message: GameStart, reader: jspb.BinaryReader): GameStart;
}

export namespace GameStart {
  export type AsObject = {
  }
}

export class UseSkill extends jspb.Message {
  getSkillId(): string;
  setSkillId(value: string): void;

  getTargetId(): string;
  setTargetId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UseSkill.AsObject;
  static toObject(includeInstance: boolean, msg: UseSkill): UseSkill.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UseSkill, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UseSkill;
  static deserializeBinaryFromReader(message: UseSkill, reader: jspb.BinaryReader): UseSkill;
}

export namespace UseSkill {
  export type AsObject = {
    skillId: string,
    targetId: string,
  }
}

export class UseItem extends jspb.Message {
  getItemId(): string;
  setItemId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UseItem.AsObject;
  static toObject(includeInstance: boolean, msg: UseItem): UseItem.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UseItem, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UseItem;
  static deserializeBinaryFromReader(message: UseItem, reader: jspb.BinaryReader): UseItem;
}

export namespace UseItem {
  export type AsObject = {
    itemId: string,
  }
}

export class Move extends jspb.Message {
  getDir(): Move.DirMap[keyof Move.DirMap];
  setDir(value: Move.DirMap[keyof Move.DirMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Move.AsObject;
  static toObject(includeInstance: boolean, msg: Move): Move.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Move, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Move;
  static deserializeBinaryFromReader(message: Move, reader: jspb.BinaryReader): Move;
}

export namespace Move {
  export type AsObject = {
    dir: Move.DirMap[keyof Move.DirMap],
  }

  export interface DirMap {
    LEFT: 0;
    RIGHT: 1;
  }

  export const Dir: DirMap;
}

export class ActionComplete extends jspb.Message {
  getCompletedId(): string;
  setCompletedId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActionComplete.AsObject;
  static toObject(includeInstance: boolean, msg: ActionComplete): ActionComplete.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ActionComplete, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActionComplete;
  static deserializeBinaryFromReader(message: ActionComplete, reader: jspb.BinaryReader): ActionComplete;
}

export namespace ActionComplete {
  export type AsObject = {
    completedId: string,
  }
}

export class Purchase extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Purchase.AsObject;
  static toObject(includeInstance: boolean, msg: Purchase): Purchase.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Purchase, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Purchase;
  static deserializeBinaryFromReader(message: Purchase, reader: jspb.BinaryReader): Purchase;
}

export namespace Purchase {
  export type AsObject = {
    id: string,
  }
}

export class EndGame extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EndGame.AsObject;
  static toObject(includeInstance: boolean, msg: EndGame): EndGame.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EndGame, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EndGame;
  static deserializeBinaryFromReader(message: EndGame, reader: jspb.BinaryReader): EndGame;
}

export namespace EndGame {
  export type AsObject = {
  }
}

export class Continue extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Continue.AsObject;
  static toObject(includeInstance: boolean, msg: Continue): Continue.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Continue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Continue;
  static deserializeBinaryFromReader(message: Continue, reader: jspb.BinaryReader): Continue;
}

export namespace Continue {
  export type AsObject = {
  }
}

export class AudienceYell extends jspb.Message {
  getIsCheering(): boolean;
  setIsCheering(value: boolean): void;

  getIsBooing(): boolean;
  setIsBooing(value: boolean): void;

  getAmount(): number;
  setAmount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AudienceYell.AsObject;
  static toObject(includeInstance: boolean, msg: AudienceYell): AudienceYell.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AudienceYell, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AudienceYell;
  static deserializeBinaryFromReader(message: AudienceYell, reader: jspb.BinaryReader): AudienceYell;
}

export namespace AudienceYell {
  export type AsObject = {
    isCheering: boolean,
    isBooing: boolean,
    amount: number,
  }
}

export class Input extends jspb.Message {
  getPlayerId(): string;
  setPlayerId(value: string): void;

  hasRegister(): boolean;
  clearRegister(): void;
  getRegister(): Register | undefined;
  setRegister(value?: Register): void;

  hasCharSelect(): boolean;
  clearCharSelect(): void;
  getCharSelect(): CharSelect | undefined;
  setCharSelect(value?: CharSelect): void;

  hasGameStart(): boolean;
  clearGameStart(): void;
  getGameStart(): GameStart | undefined;
  setGameStart(value?: GameStart): void;

  hasSkill(): boolean;
  clearSkill(): void;
  getSkill(): UseSkill | undefined;
  setSkill(value?: UseSkill): void;

  hasItem(): boolean;
  clearItem(): void;
  getItem(): UseItem | undefined;
  setItem(value?: UseItem): void;

  hasMove(): boolean;
  clearMove(): void;
  getMove(): Move | undefined;
  setMove(value?: Move): void;

  hasActionComplete(): boolean;
  clearActionComplete(): void;
  getActionComplete(): ActionComplete | undefined;
  setActionComplete(value?: ActionComplete): void;

  hasPurchase(): boolean;
  clearPurchase(): void;
  getPurchase(): Purchase | undefined;
  setPurchase(value?: Purchase): void;

  hasContinue(): boolean;
  clearContinue(): void;
  getContinue(): Continue | undefined;
  setContinue(value?: Continue): void;

  hasEndGame(): boolean;
  clearEndGame(): void;
  getEndGame(): EndGame | undefined;
  setEndGame(value?: EndGame): void;

  hasAudience(): boolean;
  clearAudience(): void;
  getAudience(): AudienceYell | undefined;
  setAudience(value?: AudienceYell): void;

  getInputCase(): Input.InputCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Input.AsObject;
  static toObject(includeInstance: boolean, msg: Input): Input.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Input, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Input;
  static deserializeBinaryFromReader(message: Input, reader: jspb.BinaryReader): Input;
}

export namespace Input {
  export type AsObject = {
    playerId: string,
    register?: Register.AsObject,
    charSelect?: CharSelect.AsObject,
    gameStart?: GameStart.AsObject,
    skill?: UseSkill.AsObject,
    item?: UseItem.AsObject,
    move?: Move.AsObject,
    actionComplete?: ActionComplete.AsObject,
    purchase?: Purchase.AsObject,
    pb_continue?: Continue.AsObject,
    endGame?: EndGame.AsObject,
    audience?: AudienceYell.AsObject,
  }

  export enum InputCase {
    INPUT_NOT_SET = 0,
    REGISTER = 10,
    CHAR_SELECT = 20,
    GAME_START = 30,
    SKILL = 40,
    ITEM = 50,
    MOVE = 60,
    ACTION_COMPLETE = 70,
    PURCHASE = 80,
    CONTINUE = 90,
    END_GAME = 100,
    AUDIENCE = 110,
  }
}

export interface ClientTypeMap {
  CLIENTPLAYER: 0;
  CLIENTAUDIENCE: 1;
  CLIENTDISPLAY: 2;
}

export const ClientType: ClientTypeMap;

