from __future__ import print_function

import logging

import grpc
import sys
sys.path.append("./location")

from location import location_pb2
from location import location_pb2_grpc


def get_one_feature(stub):
    feature = stub.GetFeature(location_pb2.Point(lat=35.65861098290943, lng=139.74543289643074))

    if feature.name:
        print(
            "Feature received %r in %s"
            % (feature.name, feature.country)
        )
    else:
        print("Found no location")


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = location_pb2_grpc.RouteGuideStub(channel)
        print("-------------- GetFeature --------------")
        get_one_feature(stub)


if __name__ == "__main__":
    logging.basicConfig()
    run()
