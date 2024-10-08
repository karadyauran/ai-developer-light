import json

def load_data():
    try:
        with open('data.json', 'r') as f:
            return json.load(f)
    except FileNotFoundError:
        return {}

def save_data(data):
    with open('data.json', 'w') as f:
        json.dump(data, f)

def log_activity(activity, amount):
    data = load_data()
    if activity in data:
        data[activity] += amount
    else:
        data[activity] = amount
    save_data(data)

def get_weekly_data():
    data = load_data()
    return data