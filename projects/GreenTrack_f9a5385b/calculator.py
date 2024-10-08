def calculate_emissions(activity, amount):
    emission_factors = {
        'car': 0.411,
        'bus': 0.089,
        'walk': 0.0,
        'electricity': 0.233
    }
    factor = emission_factors.get(activity, 0)
    return amount * factor

def get_total_emissions(data):
    total_emissions = 0
    for activity, amount in data.items():
        total_emissions += calculate_emissions(activity, amount)
    return total_emissions