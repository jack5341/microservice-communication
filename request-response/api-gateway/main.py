from flask import Flask
import flask
import os
import requests

PORT = os.getenv('PORT', 8080)
app = Flask(__name__)

@app.route('/api')
def hello_world():
    username = flask.request.args.get("username")
    r = requests.get(url = "http://localhost:8081/connections?username=" + username)
    data = r.json()

    favoritePosts = {} 
    
    for item in data :
        r = requests.get(url = "http://localhost:8082/posts?username=" + item)
        data = r.json()
        favoritePosts[ item ] = data
    
    return str(favoritePosts)

app.run(host='0.0.0.0', port=PORT)