import random
from flask import Flask


app = Flask(__name__)


@app.route("/")
def index():
    return greeting()


def greeting():
    return random.choice(["hello", "jack", "john", "franke"])
