#!/usr/bin/python
from __future__ import print_function

import grpc

from pycf import cf_pb2
from pycf import cf_pb2_grpc



def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = cf_pb2_grpc.GreeterStub(channel)
    response = stub.Add(cf_pb2.CfRequest(num1=1,num2=2))
    print(response.sum)


if __name__ == '__main__':
    run()