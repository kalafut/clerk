import json

CONFIG_FILE = "config.json"


def load(filename=CONFIG_FILE):
    with open(filename) as f:
        cfg = json.load(f)
    return cfg
