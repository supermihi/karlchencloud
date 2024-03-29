import * as jspb from 'google-protobuf'



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

export class PartyValue extends jspb.Message {
  getParty(): Party;
  setParty(value: Party): PartyValue;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PartyValue.AsObject;
  static toObject(includeInstance: boolean, msg: PartyValue): PartyValue.AsObject;
  static serializeBinaryToWriter(message: PartyValue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PartyValue;
  static deserializeBinaryFromReader(message: PartyValue, reader: jspb.BinaryReader): PartyValue;
}

export namespace PartyValue {
  export type AsObject = {
    party: Party,
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

  getWinner(): PartyValue | undefined;
  setWinner(value?: PartyValue): PlayedCard;
  hasWinner(): boolean;
  clearWinner(): PlayedCard;

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
    winner?: PartyValue.AsObject,
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

  getMember(): MemberEvent | undefined;
  setMember(value?: MemberEvent): Event;
  hasMember(): boolean;
  clearMember(): Event;

  getNewTable(): TableData | undefined;
  setNewTable(value?: TableData): Event;
  hasNewTable(): boolean;
  clearNewTable(): Event;

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
    member?: MemberEvent.AsObject,
    newTable?: TableData.AsObject,
  }

  export enum EventCase { 
    EVENT_NOT_SET = 0,
    WELCOME = 1,
    START = 2,
    DECLARED = 3,
    PLAYED_CARD = 4,
    PLACED_BID = 5,
    MEMBER = 6,
    NEW_TABLE = 7,
  }
}

export class UserState extends jspb.Message {
  getCurrenttable(): TableState | undefined;
  setCurrenttable(value?: TableState): UserState;
  hasCurrenttable(): boolean;
  clearCurrenttable(): UserState;

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
  }
}

export class MemberEvent extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): MemberEvent;

  getName(): string;
  setName(value: string): MemberEvent;

  getType(): MemberEventType;
  setType(value: MemberEventType): MemberEvent;

  getOnline(): boolean;
  setOnline(value: boolean): MemberEvent;

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
    online: boolean,
  }
}

export class PlayCardRequest extends jspb.Message {
  getTable(): string;
  setTable(value: string): PlayCardRequest;

  getCard(): Card | undefined;
  setCard(value?: Card): PlayCardRequest;
  hasCard(): boolean;
  clearCard(): PlayCardRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayCardRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlayCardRequest): PlayCardRequest.AsObject;
  static serializeBinaryToWriter(message: PlayCardRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayCardRequest;
  static deserializeBinaryFromReader(message: PlayCardRequest, reader: jspb.BinaryReader): PlayCardRequest;
}

export namespace PlayCardRequest {
  export type AsObject = {
    table: string,
    card?: Card.AsObject,
  }
}

export class PlaceBidRequest extends jspb.Message {
  getTable(): string;
  setTable(value: string): PlaceBidRequest;

  getBid(): BidType;
  setBid(value: BidType): PlaceBidRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlaceBidRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlaceBidRequest): PlaceBidRequest.AsObject;
  static serializeBinaryToWriter(message: PlaceBidRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlaceBidRequest;
  static deserializeBinaryFromReader(message: PlaceBidRequest, reader: jspb.BinaryReader): PlaceBidRequest;
}

export namespace PlaceBidRequest {
  export type AsObject = {
    table: string,
    bid: BidType,
  }
}

export class CreateTableRequest extends jspb.Message {
  getPublic(): boolean;
  setPublic(value: boolean): CreateTableRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTableRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTableRequest): CreateTableRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTableRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTableRequest;
  static deserializeBinaryFromReader(message: CreateTableRequest, reader: jspb.BinaryReader): CreateTableRequest;
}

export namespace CreateTableRequest {
  export type AsObject = {
    pb_public: boolean,
  }
}

export class DeclareRequest extends jspb.Message {
  getTable(): string;
  setTable(value: string): DeclareRequest;

  getDeclaration(): GameType;
  setDeclaration(value: GameType): DeclareRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeclareRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeclareRequest): DeclareRequest.AsObject;
  static serializeBinaryToWriter(message: DeclareRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeclareRequest;
  static deserializeBinaryFromReader(message: DeclareRequest, reader: jspb.BinaryReader): DeclareRequest;
}

export namespace DeclareRequest {
  export type AsObject = {
    table: string,
    declaration: GameType,
  }
}

export class ListTablesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTablesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTablesRequest): ListTablesRequest.AsObject;
  static serializeBinaryToWriter(message: ListTablesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTablesRequest;
  static deserializeBinaryFromReader(message: ListTablesRequest, reader: jspb.BinaryReader): ListTablesRequest;
}

export namespace ListTablesRequest {
  export type AsObject = {
  }
}

export class ListTablesResult extends jspb.Message {
  getTablesList(): Array<TableData>;
  setTablesList(value: Array<TableData>): ListTablesResult;
  clearTablesList(): ListTablesResult;
  addTables(value?: TableData, index?: number): TableData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTablesResult.AsObject;
  static toObject(includeInstance: boolean, msg: ListTablesResult): ListTablesResult.AsObject;
  static serializeBinaryToWriter(message: ListTablesResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTablesResult;
  static deserializeBinaryFromReader(message: ListTablesResult, reader: jspb.BinaryReader): ListTablesResult;
}

export namespace ListTablesResult {
  export type AsObject = {
    tablesList: Array<TableData.AsObject>,
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

  getCreated(): Timestamp | undefined;
  setCreated(value?: Timestamp): TableData;
  hasCreated(): boolean;
  clearCreated(): TableData;

  getPublic(): boolean;
  setPublic(value: boolean): TableData;

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
    created?: Timestamp.AsObject,
    pb_public: boolean,
  }
}

export class Timestamp extends jspb.Message {
  getNanos(): number;
  setNanos(value: number): Timestamp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Timestamp.AsObject;
  static toObject(includeInstance: boolean, msg: Timestamp): Timestamp.AsObject;
  static serializeBinaryToWriter(message: Timestamp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Timestamp;
  static deserializeBinaryFromReader(message: Timestamp, reader: jspb.BinaryReader): Timestamp;
}

export namespace Timestamp {
  export type AsObject = {
    nanos: number,
  }
}

export class TableState extends jspb.Message {
  getPhase(): TablePhase;
  setPhase(value: TablePhase): TableState;

  getCurrentMatch(): MatchState | undefined;
  setCurrentMatch(value?: MatchState): TableState;
  hasCurrentMatch(): boolean;
  clearCurrentMatch(): TableState;

  getData(): TableData | undefined;
  setData(value?: TableData): TableState;
  hasData(): boolean;
  clearData(): TableState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TableState.AsObject;
  static toObject(includeInstance: boolean, msg: TableState): TableState.AsObject;
  static serializeBinaryToWriter(message: TableState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TableState;
  static deserializeBinaryFromReader(message: TableState, reader: jspb.BinaryReader): TableState;
}

export namespace TableState {
  export type AsObject = {
    phase: TablePhase,
    currentMatch?: MatchState.AsObject,
    data?: TableData.AsObject,
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

export class RegisterRequest extends jspb.Message {
  getName(): string;
  setName(value: string): RegisterRequest;

  getEmail(): string;
  setEmail(value: string): RegisterRequest;

  getPassword(): string;
  setPassword(value: string): RegisterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    name: string,
    email: string,
    password: string,
  }
}

export class RegisterReply extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): RegisterReply;

  getToken(): string;
  setToken(value: string): RegisterReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterReply.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterReply): RegisterReply.AsObject;
  static serializeBinaryToWriter(message: RegisterReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterReply;
  static deserializeBinaryFromReader(message: RegisterReply, reader: jspb.BinaryReader): RegisterReply;
}

export namespace RegisterReply {
  export type AsObject = {
    userId: string,
    token: string,
  }
}

export class LoginRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    email: string,
    password: string,
  }
}

export class LoginReply extends jspb.Message {
  getToken(): string;
  setToken(value: string): LoginReply;

  getUserId(): string;
  setUserId(value: string): LoginReply;

  getName(): string;
  setName(value: string): LoginReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginReply.AsObject;
  static toObject(includeInstance: boolean, msg: LoginReply): LoginReply.AsObject;
  static serializeBinaryToWriter(message: LoginReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginReply;
  static deserializeBinaryFromReader(message: LoginReply, reader: jspb.BinaryReader): LoginReply;
}

export namespace LoginReply {
  export type AsObject = {
    token: string,
    userId: string,
    name: string,
  }
}

export class JoinTableRequest extends jspb.Message {
  getTableId(): string;
  setTableId(value: string): JoinTableRequest;

  getInviteCode(): string;
  setInviteCode(value: string): JoinTableRequest;

  getTableDescriptionCase(): JoinTableRequest.TableDescriptionCase;

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

  export enum TableDescriptionCase { 
    TABLE_DESCRIPTION_NOT_SET = 0,
    TABLE_ID = 1,
    INVITE_CODE = 2,
  }
}

export class StartTableRequest extends jspb.Message {
  getTableId(): string;
  setTableId(value: string): StartTableRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartTableRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StartTableRequest): StartTableRequest.AsObject;
  static serializeBinaryToWriter(message: StartTableRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartTableRequest;
  static deserializeBinaryFromReader(message: StartTableRequest, reader: jspb.BinaryReader): StartTableRequest;
}

export namespace StartTableRequest {
  export type AsObject = {
    tableId: string,
  }
}

export class StartNextMatchRequest extends jspb.Message {
  getTableId(): string;
  setTableId(value: string): StartNextMatchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartNextMatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StartNextMatchRequest): StartNextMatchRequest.AsObject;
  static serializeBinaryToWriter(message: StartNextMatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartNextMatchRequest;
  static deserializeBinaryFromReader(message: StartNextMatchRequest, reader: jspb.BinaryReader): StartNextMatchRequest;
}

export namespace StartNextMatchRequest {
  export type AsObject = {
    tableId: string,
  }
}

export enum MemberEventType { 
  JOIN_TABLE = 0,
  LEAVE_TABLE = 1,
  GO_OFFLINE = 2,
  GO_ONLINE = 3,
}
export enum TablePhase { 
  NOT_STARTED = 0,
  PLAYING = 1,
  BETWEEN_GAMES = 2,
  TABLE_ENDED = 3,
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
