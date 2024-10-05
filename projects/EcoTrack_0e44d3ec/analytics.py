import json

data_file = 'user_data.json'

def get_analytics():
    if not os.path.exists(data_file):
        return "No data available for analytics."
    
    with open(data_file, 'r') as file:
        data = json.load(file)
    
    distance = data.get('distance', 0)
    electricity = data.get('electricity', 0)
    
    analytics = (
        f"Total Distance Recorded: {distance} km\n"
        f"Total Electricity Consumed: {electricity} kWh\n"
    )
    