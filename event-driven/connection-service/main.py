from flask import Flask
import flask
import os
import requests
import json

PORT = os.getenv('PORT', 8082)
app = Flask(__name__)

@app.route('/follow', methods=["POST"])
def hello_world():
    follower = flask.request.args.get("follower")
    followed = flask.request.args.get("followed")

    data= {'follower': follower,'followed': followed}
    requests.post(url = "http://localhost:8080/connections", data=json.dumps(data))
    return ""

app.run(host='0.0.0.0', port=PORT)