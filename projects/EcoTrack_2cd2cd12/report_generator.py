import json
import matplotlib.pyplot as plt

def generate_report():
    try:
        with open("user_data.json", "r") as file:
            data = json.load(file)
    except FileNotFoundError:
        return "No data available."

    dates = [entry['date'] for entry in data]
    emissions = [calculate_total_emissions(entry) for entry in data]

    plt.plot(dates, emissions, marker='o')
    plt.title('Carbon Emissions Over Time')
    plt.xlabel('Date')
    plt.ylabel('Emissions (kg CO2)')
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig('emissions_report.png')
    plt.close()
    return "Report generated successfully."

def calculate_total_emissions(entry):
    transport_emissions = float(entry['transport']) * 0.21
    electricity_emissions = float(entry['electricity']) * 0.233
    gas_emissions = float(entry['gas']) * 2.204
    waste_emissions = float(entry['waste']) * 0.1
    return transport_emissions + electricity_emissions + gas_emissions + waste_emissions