from google.protobuf.json_format import MessageToJson
import stencil_pb2
import libagent as agent
import os, uuid, threading, datetime, json

def my_func():
    current_trans = stencil_pb2.Stencil()
    current_trans.string_values["timestamp"] = str(datetime.datetime.utcnow())
    current_trans.string_values["uuid"] = str(uuid.uuid4())
    current_trans.string_values["pid"] = str(os.getpid())
    current_trans.string_values["tid"] = str(threading.get_ident())

    json_trans = MessageToJson(current_trans)

    agent.setmsg(json_trans)

    return None

if __name__ == '__main__':
    my_func()
