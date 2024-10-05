import os

data_file = 'user_data.json'

def save_user_data(distance, electricity):
    user_data = {
        'distance': distance,
        'electricity': electricity
    }
    with open(data_file, 'w') as file:
        json.dump(user_data, file)

def load_user_data():
    if os.path.exists(data_file):
        with open(data_file, 'r') as file:
            return json.load(file)