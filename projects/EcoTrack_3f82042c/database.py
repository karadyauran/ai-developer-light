import json
import os

def load_database():
    if os.path.exists("database.json"):
        with open("database.json", "r") as file:
            return json.load(file)
    else:
        return {}

def save_to_database(user_id, data):
    database = load_database()
    database[user_id] = data
    with open("database.json", "w") as file:
        json.dump(database, file)

def get_user_data(user_id):
    database = load_database()
    return database.get(user_id, None)

def update_user_data(user_id, new_data):
    database = load_database()
    database[user_id] = new_data
    with open("database.json", "w") as file:
        json.dump(database, file)