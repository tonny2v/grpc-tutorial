from __future__ import print_function
import grpc
import sys
sys.path.append('../proto')
import add_pb2
import add_pb2_grpc
import signal

def sigint_handler(signum, frame):
    exit()

signal.signal(signal.SIGINT, sigint_handler)

def run(a, b):
    # creds = grpc.ssl_channel_credentials(open('tls.pem', 'rb').read())
    channel = grpc.insecure_channel("192.168.99.105:30463")
    # channel = grpc.secure_channel(str(argv[1]), creds)
    stub = add_pb2_grpc.AddServiceStub(channel)
    request = add_pb2.AddRequest(a=a, b=b)
    response = stub.Compute(request)
    print(response.result)


if __name__ == '__main__':
    import time
    for a in range(10):
        for b in range(10):
            print('calculating ', a, '+', b)
            time.sleep(.5)
            run(a, b)

