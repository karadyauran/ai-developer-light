
def create_report():
    total_emissions = analytics.calculate_average_emissions()
    highest_activity, highest_emission = analytics.highest_emission_activity()
    lowest_activity, lowest_emission = analytics.lowest_emission_activity()
    trend = analytics.emissions_trend()

    report = f"Total Average Emissions: {total_emissions:.2f} kg CO2\n"
    report += f"Highest Emission Activity: {highest_activity} - {highest_emission:.2f} kg CO2\n"
    report += f"Lowest Emission Activity: {lowest_activity} - {lowest_emission:.2f} kg CO2\n"
    report += "Emission Trend:\n"
    for activity, emission in trend:
        report += f"  {activity}: {emission:.2f} kg CO2\n"
    
    return report

def generate_detailed_report():
    factors = analytics.analyze_emission_factors()
    detailed_report = "Emission Factors:\n"
    for activity, factor in factors.items():
        detailed_report += f"  {activity}: {factor:.2f} kg CO2/unit\n"
    return detailed_report

def save_report_to_file(filename):
    report = create_report()
    with open(filename, 'w') as file:
        file.write(report)

def load_report_from_file(filename):
    with open(filename, 'r') as file: