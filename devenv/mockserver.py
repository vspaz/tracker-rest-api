import base64
import json

import flask
import jsonschema

bp = flask.Blueprint(__name__, "mockserver")

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
    print(flask.request.headers)
    write_key = flask.request.headers.get("Authorization", "")
    if write_key:
        print(get_write_key(flask.request.headers.get("Authorization", "")))
    return app.response_class(
        response=json.dumps({"status": "200 OK", "message": "OK"}),
        status=200,
        mimetype="application/json",
    )

@bp.route("/batch", methods=["POST"])
@bp.route("/import", methods=["POST"])
def batch():
    return batch_request()

@bp.route("/t", methods=["POST"])
@bp.route("/track", methods=["POST"])
def track():
    return batch_request()

@bp.route("/i", methods=["POST"])
@bp.route("/identify", methods=["POST"])
def identify():
    return batch_request()

@bp.route("/i", methods=["POST"])
@bp.route("/group", methods=["POST"])
def group():
    return batch_request()

@bp.route("/a", methods=["POST"])
@bp.route("/alias", methods=["POST"])
def alias():
    return batch_request()

@bp.route("/p", methods=["POST"])
@bp.route("/page", methods=["POST"])
def page():
    return batch_request()

@bp.route("/s", methods=["POST"])
@bp.route("/screen", methods=["POST"])
def screen():
    return batch_request()

if __name__ == "__main__":
    app = flask.Flask(__name__)
    app.url_map.strict_slashes = False
    app.register_blueprint(blueprint=bp, prefix="/api/v1")
    app.run(host="0.0.0.0", port=5000, debug=True)