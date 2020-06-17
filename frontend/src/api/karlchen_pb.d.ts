import * as jspb from "google-protobuf"

export class Card extends jspb.Message {
  getSuit(): Suit;
  setSuit(value: Suit): Card;

  getRank(): Rank;
  setRank(value: Rank): Card;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Card.AsObject;
  static toObject(includeInstance: boolean, msg: Card): Card.AsObject;
  static serializeBinaryToWriter(message: Card, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Card;
  static deserializeBinaryFromReader(message: Card, reader: jspb.BinaryReader): Card;
}

export namespace Card {
  export type AsObject = {
    suit: Suit,
    rank: Rank,
  }
}

export class PlayerValue extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): PlayerValue;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerValue.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerValue): PlayerValue.AsObject;
  static serializeBinaryToWriter(message: PlayerValue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerValue;
  static deserializeBinaryFromReader(message: PlayerValue, reader: jspb.BinaryReader): PlayerValue;
}

export namespace PlayerValue {
  export type AsObject = {
    userId: string,
  }
}

export class Declaration extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): Declaration;

  getVorbehalt(): boolean;
  setVorbehalt(value: boolean): Declaration;

  getDefinedgamemode(): Mode | undefined;
  setDefinedgamemode(value?: Mode): Declaration;
  hasDefinedgamemode(): boolean;
  clearDefinedgamemode(): Declaration;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Declaration.AsObject;
  static toObject(includeInstance: boolean, msg: Declaration): Declaration.AsObject;
  static serializeBinaryToWriter(message: Declaration, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Declaration;
  static deserializeBinaryFromReader(message: Declaration, reader: jspb.BinaryReader): Declaration;
}

export namespace Declaration {
  export type AsObject = {
    userId: string,
    vorbehalt: boolean,
    definedgamemode?: Mode.AsObject,
  }
}

export class PlayedCard extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): PlayedCard;

  getCard(): Card | undefined;
  setCard(value?: Card): PlayedCard;
  hasCard(): boolean;
  clearCard(): PlayedCard;

  getTrickWinner(): PlayerValue | undefined;
  setTrickWinner(value?: PlayerValue): PlayedCard;
  hasTrickWinner(): boolean;
  clearTrickWinner(): PlayedCard;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayedCard.AsObject;
  static toObject(includeInstance: boolean, msg: PlayedCard): PlayedCard.AsObject;
  static serializeBinaryToWriter(message: PlayedCard, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayedCard;
  static deserializeBinaryFromReader(message: PlayedCard, reader: jspb.BinaryReader): PlayedCard;
}

export namespace PlayedCard {
  export type AsObject = {
    userId: string,
    card?: Card.AsObject,
    trickWinner?: PlayerValue.AsObject,
  }
}

export class Bid extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): Bid;

  getBid(): BidType;
  setBid(value: BidType): Bid;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Bid.AsObject;
  static toObject(includeInstance: boolean, msg: Bid): Bid.AsObject;
  static serializeBinaryToWriter(message: Bid, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Bid;
  static deserializeBinaryFromReader(message: Bid, reader: jspb.BinaryReader): Bid;
}

export namespace Bid {
  export type AsObject = {
    userId: string,
    bid: BidType,
  }
}

export class EndOfGame extends jspb.Message {
  getWinner(): Party;
  setWinner(value: Party): EndOfGame;

  getValue(): number;
  setValue(value: number): EndOfGame;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EndOfGame.AsObject;
  static toObject(includeInstance: boolean, msg: EndOfGame): EndOfGame.AsObject;
  static serializeBinaryToWriter(message: EndOfGame, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EndOfGame;
  static deserializeBinaryFromReader(message: EndOfGame, reader: jspb.BinaryReader): EndOfGame;
}

export namespace EndOfGame {
  export type AsObject = {
    winner: Party,
    value: number,
  }
}

export class Event extends jspb.Message {
  getWelcome(): UserState | undefined;
  setWelcome(value?: UserState): Event;
  hasWelcome(): boolean;
  clearWelcome(): Event;

  getStart(): MatchState | undefined;
  setStart(value?: MatchState): Event;
  hasStart(): boolean;
  clearStart(): Event;

  getDeclared(): Declaration | undefined;
  setDeclared(value?: Declaration): Event;
  hasDeclared(): boolean;
  clearDeclared(): Event;

  getPlayedCard(): PlayedCard | undefined;
  setPlayedCard(value?: PlayedCard): Event;
  hasPlayedCard(): boolean;
  clearPlayedCard(): Event;

  getPlacedBid(): Bid | undefined;
  setPlacedBid(value?: Bid): Event;
  hasPlacedBid(): boolean;
  clearPlacedBid(): Event;

  getEnded(): EndOfGame | undefined;
  setEnded(value?: EndOfGame): Event;
  hasEnded(): boolean;
  clearEnded(): Event;

  getMember(): MemberEvent | undefined;
  setMember(value?: MemberEvent): Event;
  hasMember(): boolean;
  clearMember(): Event;

  getEventCase(): Event.EventCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    welcome?: UserState.AsObject,
    start?: MatchState.AsObject,
    declared?: Declaration.AsObject,
    playedCard?: PlayedCard.AsObject,
    placedBid?: Bid.AsObject,
    ended?: EndOfGame.AsObject,
    member?: MemberEvent.AsObject,
  }

  export enum EventCase { 
    EVENT_NOT_SET = 0,
    WELCOME = 1,
    START = 2,
    DECLARED = 3,
    PLAYED_CARD = 4,
    PLACED_BID = 5,
    ENDED = 6,
    MEMBER = 7,
  }
}

export class UserState extends jspb.Message {
  getCurrenttable(): TableState | undefined;
  setCurrenttable(value?: TableState): UserState;
  hasCurrenttable(): boolean;
  clearCurrenttable(): UserState;

  getName(): string;
  setName(value: string): UserState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserState.AsObject;
  static toObject(includeInstance: boolean, msg: UserState): UserState.AsObject;
  static serializeBinaryToWriter(message: UserState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserState;
  static deserializeBinaryFromReader(message: UserState, reader: jspb.BinaryReader): UserState;
}

export namespace UserState {
  export type AsObject = {
    currenttable?: TableState.AsObject,
    name: string,
  }
}

export class MemberEvent extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): MemberEvent;

  getName(): string;
  setName(value: string): MemberEvent;

  getType(): MemberEventType;
  setType(value: MemberEventType): MemberEvent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MemberEvent.AsObject;
  static toObject(includeInstance: boolean, msg: MemberEvent): MemberEvent.AsObject;
  static serializeBinaryToWriter(message: MemberEvent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MemberEvent;
  static deserializeBinaryFromReader(message: MemberEvent, reader: jspb.BinaryReader): MemberEvent;
}

export namespace MemberEvent {
  export type AsObject = {
    userId: string,
    name: string,
    type: MemberEventType,
  }
}

export class PlayRequest extends jspb.Message {
  getDeclaration(): GameType;
  setDeclaration(value: GameType): PlayRequest;

  getBid(): BidType;
  setBid(value: BidType): PlayRequest;

  getCard(): Card | undefined;
  setCard(value?: Card): PlayRequest;
  hasCard(): boolean;
  clearCard(): PlayRequest;

  getTable(): string;
  setTable(value: string): PlayRequest;

  getRequestCase(): PlayRequest.RequestCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayRequest): PlayRequest.AsObject;
  static serializeBinaryToWriter(message: PlayRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayRequest;
  static deserializeBinaryFromReader(message: PlayRequest, reader: jspb.BinaryReader): PlayRequest;
}

export namespace PlayRequest {
  export type AsObject = {
    declaration: GameType,
    bid: BidType,
    card?: Card.AsObject,
    table: string,
  }

  export enum RequestCase { 
    REQUEST_NOT_SET = 0,
    DECLARATION = 2,
    BID = 3,
    CARD = 4,
  }
}

export class AuctionState extends jspb.Message {
  getDeclarationsList(): Array<Declaration>;
  setDeclarationsList(value: Array<Declaration>): AuctionState;
  clearDeclarationsList(): AuctionState;
  addDeclarations(value?: Declaration, index?: number): Declaration;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuctionState.AsObject;
  static toObject(includeInstance: boolean, msg: AuctionState): AuctionState.AsObject;
  static serializeBinaryToWriter(message: AuctionState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuctionState;
  static deserializeBinaryFromReader(message: AuctionState, reader: jspb.BinaryReader): AuctionState;
}

export namespace AuctionState {
  export type AsObject = {
    declarationsList: Array<Declaration.AsObject>,
  }
}

export class Mode extends jspb.Message {
  getType(): GameType;
  setType(value: GameType): Mode;

  getSoloist(): PlayerValue | undefined;
  setSoloist(value?: PlayerValue): Mode;
  hasSoloist(): boolean;
  clearSoloist(): Mode;

  getSpouse(): PlayerValue | undefined;
  setSpouse(value?: PlayerValue): Mode;
  hasSpouse(): boolean;
  clearSpouse(): Mode;

  getForehand(): string;
  setForehand(value: string): Mode;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Mode.AsObject;
  static toObject(includeInstance: boolean, msg: Mode): Mode.AsObject;
  static serializeBinaryToWriter(message: Mode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Mode;
  static deserializeBinaryFromReader(message: Mode, reader: jspb.BinaryReader): Mode;
}

export namespace Mode {
  export type AsObject = {
    type: GameType,
    soloist?: PlayerValue.AsObject,
    spouse?: PlayerValue.AsObject,
    forehand: string,
  }
}

export class Trick extends jspb.Message {
  getCardsList(): Array<Card>;
  setCardsList(value: Array<Card>): Trick;
  clearCardsList(): Trick;
  addCards(value?: Card, index?: number): Card;

  getUserIdForehand(): string;
  setUserIdForehand(value: string): Trick;

  getUserIdWinner(): PlayerValue | undefined;
  setUserIdWinner(value?: PlayerValue): Trick;
  hasUserIdWinner(): boolean;
  clearUserIdWinner(): Trick;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Trick.AsObject;
  static toObject(includeInstance: boolean, msg: Trick): Trick.AsObject;
  static serializeBinaryToWriter(message: Trick, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Trick;
  static deserializeBinaryFromReader(message: Trick, reader: jspb.BinaryReader): Trick;
}

export namespace Trick {
  export type AsObject = {
    cardsList: Array<Card.AsObject>,
    userIdForehand: string,
    userIdWinner?: PlayerValue.AsObject,
  }
}

export class GameState extends jspb.Message {
  getBidsList(): Array<Bid>;
  setBidsList(value: Array<Bid>): GameState;
  clearBidsList(): GameState;
  addBids(value?: Bid, index?: number): Bid;

  getCompletedTricks(): number;
  setCompletedTricks(value: number): GameState;

  getCurrentTrick(): Trick | undefined;
  setCurrentTrick(value?: Trick): GameState;
  hasCurrentTrick(): boolean;
  clearCurrentTrick(): GameState;

  getPreviousTrick(): Trick | undefined;
  setPreviousTrick(value?: Trick): GameState;
  hasPreviousTrick(): boolean;
  clearPreviousTrick(): GameState;

  getMode(): Mode | undefined;
  setMode(value?: Mode): GameState;
  hasMode(): boolean;
  clearMode(): GameState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameState.AsObject;
  static toObject(includeInstance: boolean, msg: GameState): GameState.AsObject;
  static serializeBinaryToWriter(message: GameState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameState;
  static deserializeBinaryFromReader(message: GameState, reader: jspb.BinaryReader): GameState;
}

export namespace GameState {
  export type AsObject = {
    bidsList: Array<Bid.AsObject>,
    completedTricks: number,
    currentTrick?: Trick.AsObject,
    previousTrick?: Trick.AsObject,
    mode?: Mode.AsObject,
  }
}

export class Cards extends jspb.Message {
  getCardsList(): Array<Card>;
  setCardsList(value: Array<Card>): Cards;
  clearCardsList(): Cards;
  addCards(value?: Card, index?: number): Card;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cards.AsObject;
  static toObject(includeInstance: boolean, msg: Cards): Cards.AsObject;
  static serializeBinaryToWriter(message: Cards, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cards;
  static deserializeBinaryFromReader(message: Cards, reader: jspb.BinaryReader): Cards;
}

export namespace Cards {
  export type AsObject = {
    cardsList: Array<Card.AsObject>,
  }
}

export class MatchState extends jspb.Message {
  getPhase(): MatchPhase;
  setPhase(value: MatchPhase): MatchState;

  getTurn(): PlayerValue | undefined;
  setTurn(value?: PlayerValue): MatchState;
  hasTurn(): boolean;
  clearTurn(): MatchState;

  getPlayers(): Players | undefined;
  setPlayers(value?: Players): MatchState;
  hasPlayers(): boolean;
  clearPlayers(): MatchState;

  getSpectator(): Empty | undefined;
  setSpectator(value?: Empty): MatchState;
  hasSpectator(): boolean;
  clearSpectator(): MatchState;

  getOwnCards(): Cards | undefined;
  setOwnCards(value?: Cards): MatchState;
  hasOwnCards(): boolean;
  clearOwnCards(): MatchState;

  getAuctionState(): AuctionState | undefined;
  setAuctionState(value?: AuctionState): MatchState;
  hasAuctionState(): boolean;
  clearAuctionState(): MatchState;

  getGameState(): GameState | undefined;
  setGameState(value?: GameState): MatchState;
  hasGameState(): boolean;
  clearGameState(): MatchState;

  getRoleCase(): MatchState.RoleCase;

  getDetailsCase(): MatchState.DetailsCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MatchState.AsObject;
  static toObject(includeInstance: boolean, msg: MatchState): MatchState.AsObject;
  static serializeBinaryToWriter(message: MatchState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MatchState;
  static deserializeBinaryFromReader(message: MatchState, reader: jspb.BinaryReader): MatchState;
}

export namespace MatchState {
  export type AsObject = {
    phase: MatchPhase,
    turn?: PlayerValue.AsObject,
    players?: Players.AsObject,
    spectator?: Empty.AsObject,
    ownCards?: Cards.AsObject,
    auctionState?: AuctionState.AsObject,
    gameState?: GameState.AsObject,
  }

  export enum RoleCase { 
    ROLE_NOT_SET = 0,
    SPECTATOR = 4,
    OWN_CARDS = 5,
  }

  export enum DetailsCase { 
    DETAILS_NOT_SET = 0,
    AUCTION_STATE = 6,
    GAME_STATE = 7,
  }
}

export class Players extends jspb.Message {
  getUserIdSelf(): string;
  setUserIdSelf(value: string): Players;

  getUserIdLeft(): string;
  setUserIdLeft(value: string): Players;

  getUserIdFace(): string;
  setUserIdFace(value: string): Players;

  getUserIdRight(): string;
  setUserIdRight(value: string): Players;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Players.AsObject;
  static toObject(includeInstance: boolean, msg: Players): Players.AsObject;
  static serializeBinaryToWriter(message: Players, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Players;
  static deserializeBinaryFromReader(message: Players, reader: jspb.BinaryReader): Players;
}

export namespace Players {
  export type AsObject = {
    userIdSelf: string,
    userIdLeft: string,
    userIdFace: string,
    userIdRight: string,
  }
}

export class TableData extends jspb.Message {
  getTableId(): string;
  setTableId(value: string): TableData;

  getOwner(): string;
  setOwner(value: string): TableData;

  getInviteCode(): string;
  setInviteCode(value: string): TableData;

  getMembersList(): Array<TableMember>;
  setMembersList(value: Array<TableMember>): TableData;
  clearMembersList(): TableData;
  addMembers(value?: TableMember, index?: number): TableMember;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TableData.AsObject;
  static toObject(includeInstance: boolean, msg: TableData): TableData.AsObject;
  static serializeBinaryToWriter(message: TableData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TableData;
  static deserializeBinaryFromReader(message: TableData, reader: jspb.BinaryReader): TableData;
}

export namespace TableData {
  export type AsObject = {
    tableId: string,
    owner: string,
    inviteCode: string,
    membersList: Array<TableMember.AsObject>,
  }
}

export class TableState extends jspb.Message {
  getInMatch(): MatchState | undefined;
  setInMatch(value?: MatchState): TableState;
  hasInMatch(): boolean;
  clearInMatch(): TableState;

  getNoMatch(): Empty | undefined;
  setNoMatch(value?: Empty): TableState;
  hasNoMatch(): boolean;
  clearNoMatch(): TableState;

  getData(): TableData | undefined;
  setData(value?: TableData): TableState;
  hasData(): boolean;
  clearData(): TableState;

  getStateCase(): TableState.StateCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TableState.AsObject;
  static toObject(includeInstance: boolean, msg: TableState): TableState.AsObject;
  static serializeBinaryToWriter(message: TableState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TableState;
  static deserializeBinaryFromReader(message: TableState, reader: jspb.BinaryReader): TableState;
}

export namespace TableState {
  export type AsObject = {
    inMatch?: MatchState.AsObject,
    noMatch?: Empty.AsObject,
    data?: TableData.AsObject,
  }

  export enum StateCase { 
    STATE_NOT_SET = 0,
    IN_MATCH = 1,
    NO_MATCH = 2,
  }
}

export class TableMember extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): TableMember;

  getName(): string;
  setName(value: string): TableMember;

  getOnline(): boolean;
  setOnline(value: boolean): TableMember;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TableMember.AsObject;
  static toObject(includeInstance: boolean, msg: TableMember): TableMember.AsObject;
  static serializeBinaryToWriter(message: TableMember, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TableMember;
  static deserializeBinaryFromReader(message: TableMember, reader: jspb.BinaryReader): TableMember;
}

export namespace TableMember {
  export type AsObject = {
    userId: string,
    name: string,
    online: boolean,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class UserName extends jspb.Message {
  getName(): string;
  setName(value: string): UserName;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserName.AsObject;
  static toObject(includeInstance: boolean, msg: UserName): UserName.AsObject;
  static serializeBinaryToWriter(message: UserName, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserName;
  static deserializeBinaryFromReader(message: UserName, reader: jspb.BinaryReader): UserName;
}

export namespace UserName {
  export type AsObject = {
    name: string,
  }
}

export class RegisterReply extends jspb.Message {
  getId(): string;
  setId(value: string): RegisterReply;

  getSecret(): string;
  setSecret(value: string): RegisterReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterReply.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterReply): RegisterReply.AsObject;
  static serializeBinaryToWriter(message: RegisterReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterReply;
  static deserializeBinaryFromReader(message: RegisterReply, reader: jspb.BinaryReader): RegisterReply;
}

export namespace RegisterReply {
  export type AsObject = {
    id: string,
    secret: string,
  }
}

export class TableId extends jspb.Message {
  getValue(): string;
  setValue(value: string): TableId;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TableId.AsObject;
  static toObject(includeInstance: boolean, msg: TableId): TableId.AsObject;
  static serializeBinaryToWriter(message: TableId, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TableId;
  static deserializeBinaryFromReader(message: TableId, reader: jspb.BinaryReader): TableId;
}

export namespace TableId {
  export type AsObject = {
    value: string,
  }
}

export class JoinTableRequest extends jspb.Message {
  getTableId(): string;
  setTableId(value: string): JoinTableRequest;

  getInviteCode(): string;
  setInviteCode(value: string): JoinTableRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JoinTableRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JoinTableRequest): JoinTableRequest.AsObject;
  static serializeBinaryToWriter(message: JoinTableRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JoinTableRequest;
  static deserializeBinaryFromReader(message: JoinTableRequest, reader: jspb.BinaryReader): JoinTableRequest;
}

export namespace JoinTableRequest {
  export type AsObject = {
    tableId: string,
    inviteCode: string,
  }
}

export enum MemberEventType { 
  JOIN_TABLE = 0,
  LEAVE_TABLE = 1,
  GO_OFFLINE = 2,
  GO_ONLINE = 3,
}
export enum MatchPhase { 
  AUCTION = 0,
  GAME = 1,
  FINISHED = 2,
}
export enum Suit { 
  DIAMONDS = 0,
  HEARTS = 1,
  SPADES = 2,
  CLUBS = 3,
}
export enum Rank { 
  NINE = 0,
  JACK = 1,
  QUEEN = 2,
  KING = 3,
  TEN = 4,
  ACE = 5,
}
export enum Party { 
  RE = 0,
  CONTRA = 1,
}
export enum BidType { 
  RE_BID = 0,
  CONTRA_BID = 1,
  RE_NO_NINETY = 2,
  RE_NO_SIXTY = 3,
  RE_NO_THIRTY = 4,
  RE_SCHWARZ = 5,
  CONTRA_NO_NINETY = 6,
  CONTRA_NO_SIXTY = 7,
  CONTRA_NO_THIRTY = 8,
  CONTRA_SCHWARZ = 9,
}
export enum GameType { 
  NORMAL_GAME = 0,
  MARRIAGE = 1,
  DIAMONDS_SOLO = 2,
  HEARTS_SOLO = 3,
  SPADES_SOLO = 4,
  CLUBS_SOLO = 5,
  MEATLESS_SOLO = 6,
}
