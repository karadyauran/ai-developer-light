import json

def load_data():
    try:
        with open("user_data.json", "r") as file:
            return json.load(file)
    except FileNotFoundError:
        return []

def save_data(data):
    with open("user_data.json", "w") as file:
        json.dump(data, file, indent=4)

def add_entry(entry):
    data = load_data()
    data.append(entry)
    save_data(data)

def get_latest_entry():
    data = load_data()
    return data[-1] if data else None

def clear_data():
    save_data([])