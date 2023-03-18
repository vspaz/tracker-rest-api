import base64
import json
import jsonschema

import flask
from flask import request


app = flask.Flask(__name__)

batch_schema = {
    "type": "object",
    "properties": {
        "batch": {
            'type': 'array',
            'items': {
                'type': 'object',
                'properties': {
                    'integrations': {
                        'type': 'object',
                    },
                    "anonymousId": {
                        "type": ["string", "null"],
                    },
                    "properties": {
                        "type": "object",
                    },
                    "timestamp": {
                        "type": "string",
                    },
                    "context": {
                        "type": "object",
                    },
                    "userId": {
                        "type": "string"
                    },
                    "type": {
                        "type": "string",
                    },
                    "event": {
                        "type": "string",
                    },
                    "messageId": {
                        "type": "string"
                    },
                },
            },
        "sentAt": {
            'type': 'string',
        },
        }
    }
}


def validate_against_schema(payload, schema):
    jsonschema.validate(
        instance=payload,
        schema=schema,
        cls=jsonschema.Draft4Validator,
    )


def get_write_key(encoded_auth_header: str):
    return str(base64.b64decode(encoded_auth_header.split(" ")[1]))


def batch_request():
    print(json.dumps(flask.request.get_json(), indent=4))
    validate_against_schema(
        payload=flask.request.get_json(),
        schema=batch_schema,
    )
    print(request.headers)
    print(get_write_key(request.headers["Authorization"]))
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

@app.route("/v1/t/", methods=['POST'])
def t():
    return batch_request()

@app.route("/v1/identify/", methods=['POST'])
def identify():
    return batch_request()

@app.route("/v1/i/", methods=['POST'])
def i():
    return batch_request()

@app.route("/v1/group/", methods=['POST'])
def group():
    return batch_request()

@app.route("/v1/g/", methods=['POST'])
def g():
    return batch_request()

@app.route("/v1/alias/", methods=['POST'])
def alias():
    return batch_request()

@app.route("/v1/a/", methods=['POST'])
def a():
    return batch_request()

@app.route("/v1/page/", methods=['POST'])
def page():
    return batch_request()

@app.route("/v1/p/", methods=['POST'])
def p():
    return batch_request()

@app.route("/v1/screen/", methods=['POST'])
def screen():
    return batch_request()

@app.route("/v1/s/", methods=['POST'])
def s():
    return batch_request()
