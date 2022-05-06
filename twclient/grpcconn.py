import grpc
from proto.api_pb2 import *
from proto import api_pb2_grpc
import datetime


class grpcClient:
    def open(self):
        self.channel = grpc.insecure_channel("bill-manager-api:9999")
        self.stub = api_pb2_grpc.ElectConsumeServiceStub(self.channel)

    def close(self):
        self.channel.close()

    def ElectConsumeGet(self):
        req = DateStruct(year=2000, month=10, day=1)
        print("GET bill-manager-api:9999")
        try:
            res = self.stub.ElectConsumeGet(req)
        except grpc.RpcError as e:
            print(e.code())
        else:
            print(grpc.StatusCode.OK)
            print(res)
        return res


def getTargetDay():
    today = datetime.date.today()
    oneday = datetime.timedelta(days=1)
    yesterday = today - oneday
    return DateStruct(year=yesterday.year, month=yesterday.month, day=yesterday.day)
