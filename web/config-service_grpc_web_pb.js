/**
 * @fileoverview gRPC-Web generated client stub for config
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.config = require('./config-service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.config.ConfigServiceClient =
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
proto.config.ConfigServicePromiseClient =
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
 *   !proto.config.SetDesiredRequest,
 *   !proto.config.Response>}
 */
const methodDescriptor_ConfigService_SetDesired = new grpc.web.MethodDescriptor(
  '/config.ConfigService/SetDesired',
  grpc.web.MethodType.UNARY,
  proto.config.SetDesiredRequest,
  proto.config.Response,
  /**
   * @param {!proto.config.SetDesiredRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.SetDesiredRequest,
 *   !proto.config.Response>}
 */
const methodInfo_ConfigService_SetDesired = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.Response,
  /**
   * @param {!proto.config.SetDesiredRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @param {!proto.config.SetDesiredRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.setDesired =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/SetDesired',
      request,
      metadata || {},
      methodDescriptor_ConfigService_SetDesired,
      callback);
};


/**
 * @param {!proto.config.SetDesiredRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.Response>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.setDesired =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/SetDesired',
      request,
      metadata || {},
      methodDescriptor_ConfigService_SetDesired);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.config.CheckConsistencyRequest,
 *   !proto.config.Response>}
 */
const methodDescriptor_ConfigService_CheckConsistency = new grpc.web.MethodDescriptor(
  '/config.ConfigService/CheckConsistency',
  grpc.web.MethodType.UNARY,
  proto.config.CheckConsistencyRequest,
  proto.config.Response,
  /**
   * @param {!proto.config.CheckConsistencyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.CheckConsistencyRequest,
 *   !proto.config.Response>}
 */
const methodInfo_ConfigService_CheckConsistency = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.Response,
  /**
   * @param {!proto.config.CheckConsistencyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @param {!proto.config.CheckConsistencyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.checkConsistency =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/CheckConsistency',
      request,
      metadata || {},
      methodDescriptor_ConfigService_CheckConsistency,
      callback);
};


/**
 * @param {!proto.config.CheckConsistencyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.Response>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.checkConsistency =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/CheckConsistency',
      request,
      metadata || {},
      methodDescriptor_ConfigService_CheckConsistency);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.config.UpdateReportedRequest,
 *   !proto.config.Response>}
 */
const methodDescriptor_ConfigService_UpdateReported = new grpc.web.MethodDescriptor(
  '/config.ConfigService/UpdateReported',
  grpc.web.MethodType.UNARY,
  proto.config.UpdateReportedRequest,
  proto.config.Response,
  /**
   * @param {!proto.config.UpdateReportedRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.UpdateReportedRequest,
 *   !proto.config.Response>}
 */
const methodInfo_ConfigService_UpdateReported = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.Response,
  /**
   * @param {!proto.config.UpdateReportedRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @param {!proto.config.UpdateReportedRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.updateReported =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/UpdateReported',
      request,
      metadata || {},
      methodDescriptor_ConfigService_UpdateReported,
      callback);
};


/**
 * @param {!proto.config.UpdateReportedRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.Response>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.updateReported =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/UpdateReported',
      request,
      metadata || {},
      methodDescriptor_ConfigService_UpdateReported);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.config.GetConfigByNameRequest,
 *   !proto.config.ConfigField>}
 */
const methodDescriptor_ConfigService_GetConfigByName = new grpc.web.MethodDescriptor(
  '/config.ConfigService/GetConfigByName',
  grpc.web.MethodType.UNARY,
  proto.config.GetConfigByNameRequest,
  proto.config.ConfigField,
  /**
   * @param {!proto.config.GetConfigByNameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.ConfigField.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.GetConfigByNameRequest,
 *   !proto.config.ConfigField>}
 */
const methodInfo_ConfigService_GetConfigByName = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.ConfigField,
  /**
   * @param {!proto.config.GetConfigByNameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.ConfigField.deserializeBinary
);


/**
 * @param {!proto.config.GetConfigByNameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.ConfigField)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.ConfigField>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.getConfigByName =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/GetConfigByName',
      request,
      metadata || {},
      methodDescriptor_ConfigService_GetConfigByName,
      callback);
};


/**
 * @param {!proto.config.GetConfigByNameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.ConfigField>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.getConfigByName =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/GetConfigByName',
      request,
      metadata || {},
      methodDescriptor_ConfigService_GetConfigByName);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.config.Identifier,
 *   !proto.config.ConfigFields>}
 */
const methodDescriptor_ConfigService_GetAllConfig = new grpc.web.MethodDescriptor(
  '/config.ConfigService/GetAllConfig',
  grpc.web.MethodType.UNARY,
  proto.config.Identifier,
  proto.config.ConfigFields,
  /**
   * @param {!proto.config.Identifier} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.ConfigFields.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.Identifier,
 *   !proto.config.ConfigFields>}
 */
const methodInfo_ConfigService_GetAllConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.ConfigFields,
  /**
   * @param {!proto.config.Identifier} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.ConfigFields.deserializeBinary
);


/**
 * @param {!proto.config.Identifier} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.ConfigFields)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.ConfigFields>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.getAllConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/GetAllConfig',
      request,
      metadata || {},
      methodDescriptor_ConfigService_GetAllConfig,
      callback);
};


/**
 * @param {!proto.config.Identifier} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.ConfigFields>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.getAllConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/GetAllConfig',
      request,
      metadata || {},
      methodDescriptor_ConfigService_GetAllConfig);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.config.Identifier,
 *   !proto.config.Response>}
 */
const methodDescriptor_ConfigService_CreatePendingConfig = new grpc.web.MethodDescriptor(
  '/config.ConfigService/CreatePendingConfig',
  grpc.web.MethodType.UNARY,
  proto.config.Identifier,
  proto.config.Response,
  /**
   * @param {!proto.config.Identifier} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.Identifier,
 *   !proto.config.Response>}
 */
const methodInfo_ConfigService_CreatePendingConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.Response,
  /**
   * @param {!proto.config.Identifier} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @param {!proto.config.Identifier} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.createPendingConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/CreatePendingConfig',
      request,
      metadata || {},
      methodDescriptor_ConfigService_CreatePendingConfig,
      callback);
};


/**
 * @param {!proto.config.Identifier} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.Response>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.createPendingConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/CreatePendingConfig',
      request,
      metadata || {},
      methodDescriptor_ConfigService_CreatePendingConfig);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.config.NewConfigRequest,
 *   !proto.config.Response>}
 */
const methodDescriptor_ConfigService_CreateNewConfig = new grpc.web.MethodDescriptor(
  '/config.ConfigService/CreateNewConfig',
  grpc.web.MethodType.UNARY,
  proto.config.NewConfigRequest,
  proto.config.Response,
  /**
   * @param {!proto.config.NewConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.config.NewConfigRequest,
 *   !proto.config.Response>}
 */
const methodInfo_ConfigService_CreateNewConfig = new grpc.web.AbstractClientBase.MethodInfo(
  proto.config.Response,
  /**
   * @param {!proto.config.NewConfigRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.config.Response.deserializeBinary
);


/**
 * @param {!proto.config.NewConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.config.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.config.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.config.ConfigServiceClient.prototype.createNewConfig =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/config.ConfigService/CreateNewConfig',
      request,
      metadata || {},
      methodDescriptor_ConfigService_CreateNewConfig,
      callback);
};


/**
 * @param {!proto.config.NewConfigRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.config.Response>}
 *     A native promise that resolves to the response
 */
proto.config.ConfigServicePromiseClient.prototype.createNewConfig =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/config.ConfigService/CreateNewConfig',
      request,
      metadata || {},
      methodDescriptor_ConfigService_CreateNewConfig);
};


module.exports = proto.config;
