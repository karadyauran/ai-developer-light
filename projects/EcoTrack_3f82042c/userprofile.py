import json
import os

def load_user_profile():
    if os.path.exists("user_profile.json"):
        with open("user_profile.json", "r") as file:
            return json.load(file)
    else:
        return {"username": "User", "total_footprint": 0, "days_tracked": 0,
                "transport": 0, "electricity": 0, "meat_meals": 0, "vegan_meals": 0}

def save_user_progress(user_data, footprint):
    profile = load_user_profile()
    profile['total_footprint'] += footprint
    profile['days_tracked'] += 1
    profile['transport'] += user_data['transport']
    profile['electricity'] += user_data['electricity']
    profile['meat_meals'] += user_data['meat_meals']
    profile['vegan_meals'] += user_data['vegan_meals']
    with open("user_profile.json", "w") as file:
        json.dump(profile, file)