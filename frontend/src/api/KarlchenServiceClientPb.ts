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
  Empty,
  Event,
  JoinTableRequest,
  PlayRequest,
  RegisterReply,
  TableData,
  TableId,
  UserName} from './karlchen_pb';

export class DokoClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
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

  register(
    request: UserName,
    metadata: grpcWeb.Metadata | null): Promise<RegisterReply>;

  register(
    request: UserName,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: RegisterReply) => void): grpcWeb.ClientReadableStream<RegisterReply>;

  register(
    request: UserName,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: RegisterReply) => void) {
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

  methodInfoCheckLogin = new grpcWeb.AbstractClientBase.MethodInfo(
    UserName,
    (request: Empty) => {
      return request.serializeBinary();
    },
    UserName.deserializeBinary
  );

  checkLogin(
    request: Empty,
    metadata: grpcWeb.Metadata | null): Promise<UserName>;

  checkLogin(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: UserName) => void): grpcWeb.ClientReadableStream<UserName>;

  checkLogin(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: UserName) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/CheckLogin',
        request,
        metadata || {},
        this.methodInfoCheckLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/CheckLogin',
    request,
    metadata || {},
    this.methodInfoCheckLogin);
  }

  methodInfoCreateTable = new grpcWeb.AbstractClientBase.MethodInfo(
    TableData,
    (request: Empty) => {
      return request.serializeBinary();
    },
    TableData.deserializeBinary
  );

  createTable(
    request: Empty,
    metadata: grpcWeb.Metadata | null): Promise<TableData>;

  createTable(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: TableData) => void): grpcWeb.ClientReadableStream<TableData>;

  createTable(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: TableData) => void) {
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
    Empty,
    (request: TableId) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  startTable(
    request: TableId,
    metadata: grpcWeb.Metadata | null): Promise<Empty>;

  startTable(
    request: TableId,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void): grpcWeb.ClientReadableStream<Empty>;

  startTable(
    request: TableId,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: Empty) => void) {
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
    Empty,
    (request: JoinTableRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  joinTable(
    request: JoinTableRequest,
    metadata: grpcWeb.Metadata | null): Promise<Empty>;

  joinTable(
    request: JoinTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void): grpcWeb.ClientReadableStream<Empty>;

  joinTable(
    request: JoinTableRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: Empty) => void) {
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

  methodInfoPlay = new grpcWeb.AbstractClientBase.MethodInfo(
    Empty,
    (request: PlayRequest) => {
      return request.serializeBinary();
    },
    Empty.deserializeBinary
  );

  play(
    request: PlayRequest,
    metadata: grpcWeb.Metadata | null): Promise<Empty>;

  play(
    request: PlayRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Empty) => void): grpcWeb.ClientReadableStream<Empty>;

  play(
    request: PlayRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.Doko/Play',
        request,
        metadata || {},
        this.methodInfoPlay,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.Doko/Play',
    request,
    metadata || {},
    this.methodInfoPlay);
  }

  methodInfoStartSession = new grpcWeb.AbstractClientBase.MethodInfo(
    Event,
    (request: Empty) => {
      return request.serializeBinary();
    },
    Event.deserializeBinary
  );

  startSession(
    request: Empty,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/api.Doko/StartSession',
      request,
      metadata || {},
      this.methodInfoStartSession);
  }

}

