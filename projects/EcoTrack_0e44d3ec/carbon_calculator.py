    transport_emission_factor = 0.21
    return distance * transport_emission_factor

def calculate_electricity_footprint(electricity):
    electricity_emission_factor = 0.233
    return electricity * electricity_emission_factor

def calculate_footprint(distance, electricity):
    transport_footprint = calculate_transport_footprint(distance)
    electricity_footprint = calculate_electricity_footprint(electricity)
    total_footprint = transport_footprint + electricity_footprint