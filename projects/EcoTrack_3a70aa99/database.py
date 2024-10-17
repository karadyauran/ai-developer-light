import os

DATABASE_FILE = 'activities.json'

def save_activity(activity_data):
    activities = load_activities()
    activities.append(activity_data)
    with open(DATABASE_FILE, 'w') as f:
        json.dump(activities, f)

def load_activities():
    if not os.path.exists(DATABASE_FILE):
        return []
    with open(DATABASE_FILE, 'r') as f:
        return json.load(f)

def clear_activities():
    with open(DATABASE_FILE, 'w') as f:
        json.dump([], f)

def delete_activity(timestamp):
    activities = load_activities()
    activities = [activity for activity in activities if activity['timestamp'] != timestamp]
    with open(DATABASE_FILE, 'w') as f:
        json.dump(activities, f)

def get_activity_by_timestamp(timestamp):
    activities = load_activities()
    for activity in activities:
        if activity['timestamp'] == timestamp:
            return activity