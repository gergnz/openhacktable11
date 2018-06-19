from flask import Flask
from flask_restful import Resource, Api

app = Flask(__name__)
api = Api(app)


class List(Resource):
    def get(self):
        return "Showing the resources"

class Add(Resource):
    def get(self):
        return "Adding a new resource"

class Delete(Resource):
    def get(self):
        return "Deleting a new resource"

api.add_resource(List, '/list')
api.add_resource(Add, '/add')
api.add_resource(Delete,'/delete')


if __name__ == '__main__':
     app.run(port='5002')