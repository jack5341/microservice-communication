from flask import Flask
import flask
import os
import requests

PORT = os.getenv('PORT', 8080)
CONNECTION_DOMAIN = os.getenv('CONNECTION_DOMAIN', "localhost:8081")
POST_DOMAIN = os.getenv('POST_DOMAIN', "localhost:8082")
app = Flask(__name__)


@app.route('/api')
def apiGateway():
    # Parse username query
    username = flask.request.args.get("username")

    # Fetching connections from connections service with username query
    r = requests.get(
        url="http://" + CONNECTION_DOMAIN + "/connections?username=" + username)
    data = r.json()

    # Creating In-memory store
    favoritePosts = {}

    # Getting posts by usernames which usernames I take by connections service as []string
    for item in data:
        r = requests.get(url="http://" + POST_DOMAIN +
                         "/posts?username=" + item)
        data = r.json()

        # Appending to our In-memory store
        favoritePosts[item] = data

    # Returning In-memory store
    return str(favoritePosts)

@app.route("/healthz")
def HealthCheck():
    return str("api gateway is alive")

app.run(host='0.0.0.0', port=PORT)
