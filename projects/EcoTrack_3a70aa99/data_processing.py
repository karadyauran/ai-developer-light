from datetime import datetime
import database

activity_emission_factors = {
    "driving": 2.31,
    "electricity": 0.92,
    "meat_consumption": 5.0,
    "public_transport": 0.18
}

def log_activity(activity, amount):
    emission_factor = activity_emission_factors.get(activity.lower(), 0)
    emissions = amount * emission_factor
    activity_data = {
        "activity": activity,
        "amount": amount,
        "emissions": emissions,
        "timestamp": datetime.now().isoformat()
    }
    database.save_activity(activity_data)

def calculate_total_emissions():
    activities = database.load_activities()
    total_emissions = sum(activity['emissions'] for activity in activities)
    return total_emissions

def get_emission_summary():
    activities = database.load_activities()
    summary = {}
    for activity in activities:
        act = activity['activity']
        summary[act] = summary.get(act, 0) + activity['emissions']
    return summary

def load_emission_factors():
    return activity_emission_factors

def save_emission_factors(new_factors):
    global activity_emission_factors
    activity_emission_factors.update(new_factors)
    with open('emission_factors.json', 'w') as f: