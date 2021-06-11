/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as karlchen_pb from './karlchen_pb';


export class DokoClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoRegister = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.RegisterReply,
    (request: karlchen_pb.RegisterRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.RegisterReply.deserializeBinary
  );

  register(
    request: karlchen_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.RegisterReply>;

  register(
    request: karlchen_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.RegisterReply) => void): grpcWeb.ClientReadableStream<karlchen_pb.RegisterReply>;

  register(
    request: karlchen_pb.RegisterRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.RegisterReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/Register',
        request,
        metadata || {},
        this.methodInfoRegister,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/Register',
    request,
    metadata || {},
    this.methodInfoRegister);
  }

  methodInfoLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.LoginReply,
    (request: karlchen_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.LoginReply.deserializeBinary
  );

  login(
    request: karlchen_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.LoginReply>;

  login(
    request: karlchen_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.LoginReply) => void): grpcWeb.ClientReadableStream<karlchen_pb.LoginReply>;

  login(
    request: karlchen_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.LoginReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/Login',
        request,
        metadata || {},
        this.methodInfoLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/Login',
    request,
    metadata || {},
    this.methodInfoLogin);
  }

  methodInfoCreateTable = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.TableData,
    (request: karlchen_pb.CreateTableRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.TableData.deserializeBinary
  );

  createTable(
    request: karlchen_pb.CreateTableRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.TableData>;

  createTable(
    request: karlchen_pb.CreateTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.TableData) => void): grpcWeb.ClientReadableStream<karlchen_pb.TableData>;

  createTable(
    request: karlchen_pb.CreateTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.TableData) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/CreateTable',
        request,
        metadata || {},
        this.methodInfoCreateTable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/CreateTable',
    request,
    metadata || {},
    this.methodInfoCreateTable);
  }

  methodInfoStartTable = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.MatchState,
    (request: karlchen_pb.StartTableRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.MatchState.deserializeBinary
  );

  startTable(
    request: karlchen_pb.StartTableRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.MatchState>;

  startTable(
    request: karlchen_pb.StartTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.MatchState) => void): grpcWeb.ClientReadableStream<karlchen_pb.MatchState>;

  startTable(
    request: karlchen_pb.StartTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.MatchState) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/StartTable',
        request,
        metadata || {},
        this.methodInfoStartTable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/StartTable',
    request,
    metadata || {},
    this.methodInfoStartTable);
  }

  methodInfoJoinTable = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.TableState,
    (request: karlchen_pb.JoinTableRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.TableState.deserializeBinary
  );

  joinTable(
    request: karlchen_pb.JoinTableRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.TableState>;

  joinTable(
    request: karlchen_pb.JoinTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.TableState) => void): grpcWeb.ClientReadableStream<karlchen_pb.TableState>;

  joinTable(
    request: karlchen_pb.JoinTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.TableState) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/JoinTable',
        request,
        metadata || {},
        this.methodInfoJoinTable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/JoinTable',
    request,
    metadata || {},
    this.methodInfoJoinTable);
  }

  methodInfoPlayCard = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.PlayedCard,
    (request: karlchen_pb.PlayCardRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.PlayedCard.deserializeBinary
  );

  playCard(
    request: karlchen_pb.PlayCardRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.PlayedCard>;

  playCard(
    request: karlchen_pb.PlayCardRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.PlayedCard) => void): grpcWeb.ClientReadableStream<karlchen_pb.PlayedCard>;

  playCard(
    request: karlchen_pb.PlayCardRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.PlayedCard) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/PlayCard',
        request,
        metadata || {},
        this.methodInfoPlayCard,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/PlayCard',
    request,
    metadata || {},
    this.methodInfoPlayCard);
  }

  methodInfoPlaceBid = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.Bid,
    (request: karlchen_pb.PlaceBidRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.Bid.deserializeBinary
  );

  placeBid(
    request: karlchen_pb.PlaceBidRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.Bid>;

  placeBid(
    request: karlchen_pb.PlaceBidRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.Bid) => void): grpcWeb.ClientReadableStream<karlchen_pb.Bid>;

  placeBid(
    request: karlchen_pb.PlaceBidRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.Bid) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/PlaceBid',
        request,
        metadata || {},
        this.methodInfoPlaceBid,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/PlaceBid',
    request,
    metadata || {},
    this.methodInfoPlaceBid);
  }

  methodInfoDeclare = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.Declaration,
    (request: karlchen_pb.DeclareRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.Declaration.deserializeBinary
  );

  declare(
    request: karlchen_pb.DeclareRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.Declaration>;

  declare(
    request: karlchen_pb.DeclareRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.Declaration) => void): grpcWeb.ClientReadableStream<karlchen_pb.Declaration>;

  declare(
    request: karlchen_pb.DeclareRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.Declaration) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/Declare',
        request,
        metadata || {},
        this.methodInfoDeclare,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/Declare',
    request,
    metadata || {},
    this.methodInfoDeclare);
  }

  methodInfoStartNextMatch = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.MatchState,
    (request: karlchen_pb.StartNextMatchRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.MatchState.deserializeBinary
  );

  startNextMatch(
    request: karlchen_pb.StartNextMatchRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.MatchState>;

  startNextMatch(
    request: karlchen_pb.StartNextMatchRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.MatchState) => void): grpcWeb.ClientReadableStream<karlchen_pb.MatchState>;

  startNextMatch(
    request: karlchen_pb.StartNextMatchRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.MatchState) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/StartNextMatch',
        request,
        metadata || {},
        this.methodInfoStartNextMatch,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/StartNextMatch',
    request,
    metadata || {},
    this.methodInfoStartNextMatch);
  }

  methodInfoStartSession = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.Event,
    (request: karlchen_pb.Empty) => {
      return request.serializeBinary();
    },
    karlchen_pb.Event.deserializeBinary
  );

  startSession(
    request: karlchen_pb.Empty,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/api.Doko/StartSession',
      request,
      metadata || {},
      this.methodInfoStartSession);
  }

  methodInfoListTables = new grpcWeb.AbstractClientBase.MethodInfo(
    karlchen_pb.ListTablesResult,
    (request: karlchen_pb.ListTablesRequest) => {
      return request.serializeBinary();
    },
    karlchen_pb.ListTablesResult.deserializeBinary
  );

  listTables(
    request: karlchen_pb.ListTablesRequest,
    metadata: grpcWeb.Metadata | null): Promise<karlchen_pb.ListTablesResult>;

  listTables(
    request: karlchen_pb.ListTablesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: karlchen_pb.ListTablesResult) => void): grpcWeb.ClientReadableStream<karlchen_pb.ListTablesResult>;

  listTables(
    request: karlchen_pb.ListTablesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: karlchen_pb.ListTablesResult) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/ListTables',
        request,
        metadata || {},
        this.methodInfoListTables,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/ListTables',
    request,
    metadata || {},
    this.methodInfoListTables);
  }

}

