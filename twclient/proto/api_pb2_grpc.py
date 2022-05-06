# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from proto import api_pb2 as proto_dot_api__pb2


class ElectConsumeServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ElectConsumePost = channel.unary_unary(
                '/billManagerAPI.ElectConsumeService/ElectConsumePost',
                request_serializer=proto_dot_api__pb2.OnedayElectConsume.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ElectConsumeGet = channel.unary_unary(
                '/billManagerAPI.ElectConsumeService/ElectConsumeGet',
                request_serializer=proto_dot_api__pb2.DateStruct.SerializeToString,
                response_deserializer=proto_dot_api__pb2.OnedayElectConsume.FromString,
                )


class ElectConsumeServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ElectConsumePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ElectConsumeGet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ElectConsumeServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ElectConsumePost': grpc.unary_unary_rpc_method_handler(
                    servicer.ElectConsumePost,
                    request_deserializer=proto_dot_api__pb2.OnedayElectConsume.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ElectConsumeGet': grpc.unary_unary_rpc_method_handler(
                    servicer.ElectConsumeGet,
                    request_deserializer=proto_dot_api__pb2.DateStruct.FromString,
                    response_serializer=proto_dot_api__pb2.OnedayElectConsume.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'billManagerAPI.ElectConsumeService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ElectConsumeService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ElectConsumePost(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/billManagerAPI.ElectConsumeService/ElectConsumePost',
            proto_dot_api__pb2.OnedayElectConsume.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ElectConsumeGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/billManagerAPI.ElectConsumeService/ElectConsumeGet',
            proto_dot_api__pb2.DateStruct.SerializeToString,
            proto_dot_api__pb2.OnedayElectConsume.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)