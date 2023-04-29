import base64
import json

import jsonschema
import sanic


bp = sanic.Blueprint(__name__,  url_prefix="/api/v1")
app = sanic.Sanic(__name__, strict_slashes=False)
app.blueprint(blueprint=bp)

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
    print(sanic.request.Request.json)
    validate_against_schema(
        payload=sanic.request.Request.json,
        schema=batch_schema,
    )
    print(sanic.request.Request.headers)
    write_key = sanic.request.Request.headers.get("Authorization", "")
    if write_key:
        print(get_write_key(sanic.request.Request.headers.get("Authorization", "")))
    return sanic.json(
        body={"status": "200 OK", "message": "OK"},
        status=200,
    )

@bp.route("/batch", name="batch", methods=["POST"])
@bp.route("/import", name="import", methods=["POST"])
async def batch(request):
    return batch_request()

@bp.route("/t", name="t", methods=["POST"])
@bp.route("/track", name="track", methods=["POST"])
async def track(request):
    return batch_request()

@bp.route("/i", name="i", methods=["POST"])
@bp.route("/identify", name="identify", methods=["POST"])
async def identify(request):
    return batch_request()

@bp.route("/g", name="g", methods=["POST"])
@bp.route("/group", name="group", methods=["POST"])
async def group(request):
    return batch_request()

@bp.route("/a", name="a", methods=["POST"])
@bp.route("/alias", name="alias", methods=["POST"])
async def alias(request):
    return batch_request()

@bp.route("/p", name="p", methods=["POST"])
@bp.route("/page", name="page", methods=["POST"])
async def page(request):
    return batch_request()

@bp.route("/s", name="s", methods=["POST"])
@bp.route("/screen", name="screen", methods=["POST"])
async def screen(request):
    return batch_request()



if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)
