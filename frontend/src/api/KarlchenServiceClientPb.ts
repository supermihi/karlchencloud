/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!

/* eslint-disable */
// @ts-nocheck

import * as grpcWeb from 'grpc-web';

import {
  Bid,
  Declaration,
  DeclareRequest,
  Empty,
  Event,
  JoinTableRequest,
  MatchState,
  PlaceBidRequest,
  PlayCardRequest,
  PlayedCard,
  RegisterReply,
  TableData,
  TableId,
  TableState,
  UserName,
} from './karlchen_pb';

export class DokoClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string };
  options_: null | { [index: string]: string };

  constructor(
    hostname: string,
    credentials?: null | { [index: string]: string },
    options?: null | { [index: string]: string }
  ) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoRegister = new grpcWeb.AbstractClientBase.MethodInfo(
    RegisterReply,
    (request: UserName) => {
      return request.serializeBinary();
    },
    RegisterReply.deserializeBinary
  );

  register(request: UserName, metadata: grpcWeb.Metadata | null): Promise<RegisterReply>;

  register(
    request: UserName,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: RegisterReply) => void
  ): grpcWeb.ClientReadableStream<RegisterReply>;

  register(
    request: UserName,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: RegisterReply) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/Register',
        request,
        metadata || {},
        this.methodInfoRegister,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/Register',
      request,
      metadata || {},
      this.methodInfoRegister
    );
  }

  methodInfoCheckLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    UserName,
    (request: Empty) => {
      return request.serializeBinary();
    },
    UserName.deserializeBinary
  );

  checkLogin(request: Empty, metadata: grpcWeb.Metadata | null): Promise<UserName>;

  checkLogin(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: UserName) => void
  ): grpcWeb.ClientReadableStream<UserName>;

  checkLogin(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: UserName) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/CheckLogin',
        request,
        metadata || {},
        this.methodInfoCheckLogin,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/CheckLogin',
      request,
      metadata || {},
      this.methodInfoCheckLogin
    );
  }

  methodInfoCreateTable = new grpcWeb.AbstractClientBase.MethodInfo(
    TableData,
    (request: Empty) => {
      return request.serializeBinary();
    },
    TableData.deserializeBinary
  );

  createTable(request: Empty, metadata: grpcWeb.Metadata | null): Promise<TableData>;

  createTable(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: TableData) => void
  ): grpcWeb.ClientReadableStream<TableData>;

  createTable(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: TableData) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/CreateTable',
        request,
        metadata || {},
        this.methodInfoCreateTable,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/CreateTable',
      request,
      metadata || {},
      this.methodInfoCreateTable
    );
  }

  methodInfoStartTable = new grpcWeb.AbstractClientBase.MethodInfo(
    MatchState,
    (request: TableId) => {
      return request.serializeBinary();
    },
    MatchState.deserializeBinary
  );

  startTable(request: TableId, metadata: grpcWeb.Metadata | null): Promise<MatchState>;

  startTable(
    request: TableId,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: MatchState) => void
  ): grpcWeb.ClientReadableStream<MatchState>;

  startTable(
    request: TableId,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: MatchState) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/StartTable',
        request,
        metadata || {},
        this.methodInfoStartTable,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/StartTable',
      request,
      metadata || {},
      this.methodInfoStartTable
    );
  }

  methodInfoJoinTable = new grpcWeb.AbstractClientBase.MethodInfo(
    TableState,
    (request: JoinTableRequest) => {
      return request.serializeBinary();
    },
    TableState.deserializeBinary
  );

  joinTable(request: JoinTableRequest, metadata: grpcWeb.Metadata | null): Promise<TableState>;

  joinTable(
    request: JoinTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: TableState) => void
  ): grpcWeb.ClientReadableStream<TableState>;

  joinTable(
    request: JoinTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: TableState) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/JoinTable',
        request,
        metadata || {},
        this.methodInfoJoinTable,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/JoinTable',
      request,
      metadata || {},
      this.methodInfoJoinTable
    );
  }

  methodInfoPlayCard = new grpcWeb.AbstractClientBase.MethodInfo(
    PlayedCard,
    (request: PlayCardRequest) => {
      return request.serializeBinary();
    },
    PlayedCard.deserializeBinary
  );

  playCard(request: PlayCardRequest, metadata: grpcWeb.Metadata | null): Promise<PlayedCard>;

  playCard(
    request: PlayCardRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: PlayedCard) => void
  ): grpcWeb.ClientReadableStream<PlayedCard>;

  playCard(
    request: PlayCardRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: PlayedCard) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/PlayCard',
        request,
        metadata || {},
        this.methodInfoPlayCard,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/PlayCard',
      request,
      metadata || {},
      this.methodInfoPlayCard
    );
  }

  methodInfoPlaceBid = new grpcWeb.AbstractClientBase.MethodInfo(
    Bid,
    (request: PlaceBidRequest) => {
      return request.serializeBinary();
    },
    Bid.deserializeBinary
  );

  placeBid(request: PlaceBidRequest, metadata: grpcWeb.Metadata | null): Promise<Bid>;

  placeBid(
    request: PlaceBidRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: Bid) => void
  ): grpcWeb.ClientReadableStream<Bid>;

  placeBid(
    request: PlaceBidRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: Bid) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/PlaceBid',
        request,
        metadata || {},
        this.methodInfoPlaceBid,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/PlaceBid',
      request,
      metadata || {},
      this.methodInfoPlaceBid
    );
  }

  methodInfoDeclare = new grpcWeb.AbstractClientBase.MethodInfo(
    Declaration,
    (request: DeclareRequest) => {
      return request.serializeBinary();
    },
    Declaration.deserializeBinary
  );

  declare(request: DeclareRequest, metadata: grpcWeb.Metadata | null): Promise<Declaration>;

  declare(
    request: DeclareRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error, response: Declaration) => void
  ): grpcWeb.ClientReadableStream<Declaration>;

  declare(
    request: DeclareRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error, response: Declaration) => void
  ) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ + '/api.Doko/Declare',
        request,
        metadata || {},
        this.methodInfoDeclare,
        callback
      );
    }
    return this.client_.unaryCall(
      this.hostname_ + '/api.Doko/Declare',
      request,
      metadata || {},
      this.methodInfoDeclare
    );
  }

  methodInfoStartSession = new grpcWeb.AbstractClientBase.MethodInfo(
    Event,
    (request: Empty) => {
      return request.serializeBinary();
    },
    Event.deserializeBinary
  );

  startSession(request: Empty, metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ + '/api.Doko/StartSession',
      request,
      metadata || {},
      this.methodInfoStartSession
    );
  }
}
