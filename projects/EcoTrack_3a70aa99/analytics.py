
def calculate_average_emissions():
    activities = data_processing.get_emission_summary()
    total_emissions = sum(activities.values())
    average_emission = total_emissions / len(activities) if activities else 0
    return average_emission

def highest_emission_activity():
    activities = data_processing.get_emission_summary()
    if not activities:
        return None, 0
    highest_activity = max(activities, key=activities.get)
    return highest_activity, activities[highest_activity]

def lowest_emission_activity():
    activities = data_processing.get_emission_summary()
    if not activities:
        return None, 0
    lowest_activity = min(activities, key=activities.get)
    return lowest_activity, activities[lowest_activity]

def emissions_trend():
    activities = data_processing.get_emission_summary()
    sorted_activities = sorted(activities.items(), key=lambda x: x[1], reverse=True)
    return sorted_activities

def analyze_emission_factors():