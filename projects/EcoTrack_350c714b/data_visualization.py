    report = "Activity Report:\n"
    for activity in activities:
        report += f"Activity: {activity['activity']}, Duration: {activity['duration']} hours\n"
    report += f"Total Carbon Footprint: {footprint} kg CO2\n"