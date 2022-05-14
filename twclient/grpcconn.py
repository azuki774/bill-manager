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

    def ElectConsumeGet(self, datestruct):
        req = datestruct
        print("GET bill-manager-api:9999")
        try:
            res = self.stub.ElectConsumeGet(req)
        except grpc.RpcError as e:
            print(e.code())
        else:
            print(grpc.StatusCode.OK)
            print("== response ==")
            print(res)
        return res


def get_targetDay():
    # 前日が何月を確認するために実装
    nowadays = datetime.datetime.now() + datetime.timedelta(hours=9)
    yesterday = nowadays - datetime.timedelta(1)
    return DateStruct(year=yesterday.year, month=yesterday.month, day=yesterday.day)
