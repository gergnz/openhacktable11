from flask import Flask
from kubernetes import client, config

app = Flask(__name__)

class list:
    def list(self):
        output = ""
        v1 = client.CoreV1Api()
        ret = v1.list_pod_for_all_namespaces(watch=False)
        for i in ret.items:
            output = output + i.status.pod_ip + i.metadata.namespace + i.metadata.name
            break

        return output
