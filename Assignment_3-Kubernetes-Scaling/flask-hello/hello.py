from flask import Flask
import os

app = Flask(__name__)

@app.route('/')
def index():
    return str(os.getenv('MESSAGE'))

app.run(host='0.0.0.0', port=8080)
