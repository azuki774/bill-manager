import grpc
from proto.api_pb2 import *
from proto import api_pb2_grpc


class grpcClient:
    def open(self):
        self.channel = grpc.insecure_channel("bill-manager-api:9999")
        self.stub = api_pb2_grpc.ElectConsumeServiceStub(self.channel)

    def close(self):
        self.channel.close()

    def ElectConsumePost(self, dateStruct, daytime, nighttime, total):
        req = OnedayElectConsume(
            date=dateStruct, daytime=daytime, nighttime=nighttime, total=total
        )
        print("send + #" + str(dateStruct.day))

        try:
            res = self.stub.ElectConsumePost(req)
        except grpc.RpcError as e:
            print(e.code())
        else:
            print(grpc.StatusCode.OK)
