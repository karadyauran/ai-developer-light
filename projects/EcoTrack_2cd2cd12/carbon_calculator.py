import json

def calculate_emissions():
    try:
        with open("user_data.json", "r") as file:
            data = json.load(file)
    except FileNotFoundError:
        return "No data available."

    latest_entry = data[-1]
    transport_emissions = float(latest_entry['transport']) * 0.21
    electricity_emissions = float(latest_entry['electricity']) * 0.233
    gas_emissions = float(latest_entry['gas']) * 2.204
    waste_emissions = float(latest_entry['waste']) * 0.1

    total_emissions = transport_emissions + electricity_emissions + gas_emissions + waste_emissions
    return round(total_emissions, 2)