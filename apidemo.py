from flask import Flask
from flask_restful import Resource, Api
from kubernetes import client, config
from pprint import pprint
import json
app = Flask(__name__)
api = Api(app)


class List(Resource):
    def get(self):
        output = []
        config.load_kube_config()
        v1 = client.CoreV1Api()
        ret = v1.list_namespaced_service('default')

        for i in ret.items:
            pprint(i)
            if i.metadata.name =="hackerone":
                eachOutput = {}
                eachOutput['name'] = i.metadata.name
                eachOutput['endpoints'] = {}
                for port in i.spec.ports:
                    eachOutput['endpoints'][port.name] = str(i.status.load_balancer.ingress[0].ip) + ":" + str(port.port)
                output.append(eachOutput)

        return output

class Add(Resource):
    def get(self):
        return "Adding a new resource"

class Delete(Resource):
    def get(self):

       return "Deleting"

class Ping(Resource):
    def get(self):
        return "ack"

api.add_resource(List, '/list')
api.add_resource(Add, '/add')
api.add_resource(Delete,'/delete')
api.add_resource(Ping,'/ping')


if __name__ == '__main__':
     app.run(port='5002',host='0.0.0.0')
