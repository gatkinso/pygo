from google.protobuf.json_format import MessageToJson
import stencil_pb2
import libagent as agent
import os, threading, datetime, json

def my_func():
    current_trans = transport_pb2.Transport()
    current_trans.string_values["timestamp"] = str(datetime.datetime.utcnow())
    current_trans.string_values["uuid"] = str(uuid.uuid4())
    current_trans.string_values["pid"] = str(os.getpid())
    current_trans.string_values["tid"] = str(threading.get_ident())

    current_request = request._get_current_object()
    current_trans.string_values["base_url"] = current_request.base_url
    current_trans.string_values["endpoint"] = current_request.endpoint
    current_trans.string_values["method"] = current_request.method

    json_trans = MessageToJson(current_trans)

    agent.setmsg(json_trans)

    return None

if __name__ == '__main__':
    my_func()
