import base64
import json

import flask
from flask import request


app = flask.Flask(__name__)


def batch_request():
    print(json.dumps(flask.request.get_json(), indent=4))
    print(request.headers)
    print(str(base64.b64decode(request.headers["Authorization"].split(" ")[1])))
    return app.response_class(
        response=json.dumps({"status": "200 OK", "message": "OK"}),
        status=200,
        mimetype='application/json'
    )


@app.route("/v1/import/", methods=['POST'])
def _import():
    return batch_request()

@app.route("/v1/batch/", methods=['POST'])
def batch():
    return batch_request()

@app.route("/v1/track/", methods=['POST'])
def track():
    return batch_request()

@app.route("/v1/identify/", methods=['POST'])
def identify():
    return batch_request()

@app.route("/v1/group/", methods=['POST'])
def group():
    return batch_request()

@app.route("/v1/alias/", methods=['POST'])
def alias():
    return batch_request()

@app.route("/v1/page/", methods=['POST'])
def page():
    return batch_request()

@app.route("/v1/screen/", methods=['POST'])
def screen():
    return batch_request()
