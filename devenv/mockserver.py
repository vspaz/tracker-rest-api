import json
import base64

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


@app.route("/v1/import/", methods=['GET', 'POST'])
def import_deprecated():
    return batch_request()


@app.route("/v1/batch/", methods=['GET', 'POST'])
def batch():
    return batch_request()
