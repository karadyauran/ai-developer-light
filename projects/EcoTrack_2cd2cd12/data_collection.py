import json
import datetime

def collect_data():
    user_data = {}
    user_data['date'] = str(datetime.date.today())
    user_data['transport'] = input("Enter distance traveled by car (km): ")
    user_data['electricity'] = input("Enter electricity usage (kWh): ")
    user_data['gas'] = input("Enter gas usage (cubic meters): ")
    user_data['waste'] = input("Enter waste produced (kg): ")
    store_data(user_data)

def store_data(data):
    try:
        with open("user_data.json", "r") as file:
            existing_data = json.load(file)
    except FileNotFoundError:
        existing_data = []
    existing_data.append(data)
    with open("user_data.json", "w") as file:
        json.dump(existing_data, file, indent=4)