from flask import Flask
import flask
import os
import requests
import json

PORT = os.getenv('PORT', 8081)
app = Flask(__name__)

@app.route('/newpost', methods=["POST"])
def hello_world():
    username = flask.request.args.get("username")
    post = flask.request.args.get("post")

    data= {'username': username,'post': post}
    requests.post(url = "http://localhost:8080/post", data=json.dumps(data))
    return ""

app.run(host='0.0.0.0', port=PORT)