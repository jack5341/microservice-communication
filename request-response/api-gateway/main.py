from flask import Flask
import flask
import os
import requests

PORT = os.getenv('PORT', 8080)
app = Flask(__name__)

@app.route('/api')

def apiGateway():
    # Parse username query
    username = flask.request.args.get("username")
    
    # Fetching connections from connections service with username query 
    r = requests.get(url = "http://localhost:8081/connections?username=" + username)
    data = r.json()

    # Creating In-memory store
    favoritePosts = {} 
    
    # Getting posts by usernames which usernames I take by connections service as []string
    for item in data :
        r = requests.get(url = "http://localhost:8082/posts?username=" + item)
        data = r.json()
        
        # Appending to our In-memory store
        favoritePosts[ item ] = data
    
    # Returning In-memory store
    return str(favoritePosts)

app.run(host='0.0.0.0', port=PORT)
