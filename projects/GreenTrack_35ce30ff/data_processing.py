def calculate_emissions(data):
    transport_emissions = data.get('transport', 0)
    energy_emissions = data.get('energy', 0)
    waste_emissions = data.get('waste', 0)
    total_emissions = transport_emissions + energy_emissions + waste_emissions
    return {
        'transport': transport_emissions,
        'energy': energy_emissions,
        'waste': waste_emissions,
        'total': total_emissions
    }