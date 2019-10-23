# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import config_service_pb2 as config__service__pb2


class ConfigServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SetDesired = channel.unary_unary(
        '/config.ConfigService/SetDesired',
        request_serializer=config__service__pb2.SetDesiredRequest.SerializeToString,
        response_deserializer=config__service__pb2.Response.FromString,
        )
    self.CheckConsistency = channel.unary_unary(
        '/config.ConfigService/CheckConsistency',
        request_serializer=config__service__pb2.CheckConsistencyRequest.SerializeToString,
        response_deserializer=config__service__pb2.Response.FromString,
        )
    self.UpdateReported = channel.unary_unary(
        '/config.ConfigService/UpdateReported',
        request_serializer=config__service__pb2.UpdateReportedRequest.SerializeToString,
        response_deserializer=config__service__pb2.Response.FromString,
        )
    self.GetAllConfig = channel.unary_unary(
        '/config.ConfigService/GetAllConfig',
        request_serializer=config__service__pb2.DeviceEUI.SerializeToString,
        response_deserializer=config__service__pb2.ConfigFields.FromString,
        )
    self.CreateNewConfig = channel.unary_unary(
        '/config.ConfigService/CreateNewConfig',
        request_serializer=config__service__pb2.DeviceEUI.SerializeToString,
        response_deserializer=config__service__pb2.Response.FromString,
        )


class ConfigServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def SetDesired(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CheckConsistency(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateReported(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetAllConfig(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateNewConfig(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ConfigServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SetDesired': grpc.unary_unary_rpc_method_handler(
          servicer.SetDesired,
          request_deserializer=config__service__pb2.SetDesiredRequest.FromString,
          response_serializer=config__service__pb2.Response.SerializeToString,
      ),
      'CheckConsistency': grpc.unary_unary_rpc_method_handler(
          servicer.CheckConsistency,
          request_deserializer=config__service__pb2.CheckConsistencyRequest.FromString,
          response_serializer=config__service__pb2.Response.SerializeToString,
      ),
      'UpdateReported': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateReported,
          request_deserializer=config__service__pb2.UpdateReportedRequest.FromString,
          response_serializer=config__service__pb2.Response.SerializeToString,
      ),
      'GetAllConfig': grpc.unary_unary_rpc_method_handler(
          servicer.GetAllConfig,
          request_deserializer=config__service__pb2.DeviceEUI.FromString,
          response_serializer=config__service__pb2.ConfigFields.SerializeToString,
      ),
      'CreateNewConfig': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNewConfig,
          request_deserializer=config__service__pb2.DeviceEUI.FromString,
          response_serializer=config__service__pb2.Response.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'config.ConfigService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
