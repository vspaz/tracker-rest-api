import base64
import json

import flask
import jsonschema
from flask import request

app = flask.Flask(__name__)

_TYPE_STRING = {
    "type": "string",
}

_TYPE_NUMBER = {
    "type": "number",
}

_CONTEXT = _PROPERTIES = _TYPE_OBJECT = {
    "type": "object",
}

_MESSAGE = {
    "type": "object",
    "properties": {
        "integrations": _TYPE_OBJECT,
        "anonymousId": {
            "type": ["string", "null"],
        },
        "properties": _PROPERTIES,
        "timestamp": _TYPE_STRING,
        "context": _CONTEXT,
        "userId": _TYPE_STRING,
        "type": _TYPE_STRING,
        "event": _TYPE_STRING,
        "messageId": _TYPE_STRING,
    },
}

_BATCH = {
    "type": "array",
    "items": _MESSAGE,
}

batch_schema = {
    "type": "object",
    "properties": {
        "batch": _BATCH,
        "sentAt": _TYPE_STRING,
        "context": _CONTEXT,
        "writeKey": _TYPE_STRING,
        "sequence": _TYPE_NUMBER,
    },
    "required": ["batch"],
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
    write_key = request.headers.get("Authorization", "")
    if write_key:
        print(get_write_key(request.headers.get("Authorization", "")))
    return app.response_class(
        response=json.dumps({"status": "200 OK", "message": "OK"}),
        status=200,
        mimetype="application/json",
    )


@app.route("/v1/import/", methods=["POST"])
def _import():
    return batch_request()

@app.route("/v1/batch/", methods=["POST"])
def batch():
    return batch_request()

@app.route("/v1/track/", methods=["POST"])
def track():
    return batch_request()

@app.route("/v1/t/", methods=["POST"])
def t():
    return batch_request()

@app.route("/v1/identify/", methods=["POST"])
def identify():
    return batch_request()

@app.route("/v1/i/", methods=["POST"])
def i():
    return batch_request()

@app.route("/v1/group/", methods=["POST"])
def group():
    return batch_request()

@app.route("/v1/g/", methods=["POST"])
def g():
    return batch_request()

@app.route("/v1/alias/", methods=["POST"])
def alias():
    return batch_request()

@app.route("/v1/a/", methods=["POST"])
def a():
    return batch_request()

@app.route("/v1/page/", methods=["POST"])
def page():
    return batch_request()

@app.route("/v1/p/", methods=["POST"])
def p():
    return batch_request()

@app.route("/v1/screen/", methods=["POST"])
def screen():
    return batch_request()

@app.route("/v1/s/", methods=["POST"])
def s():
    return batch_request()
