from flask import Flask
import flask
import requests


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

app.run(host='0.0.0.0', port=8080)