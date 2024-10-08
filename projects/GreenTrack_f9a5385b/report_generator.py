import tracker
import calculator

def generate_report():
    data = tracker.get_weekly_data()
    total_emissions = calculator.get_total_emissions(data)
    report = "Weekly Carbon Emissions Report\n"
    report += "-------------------------------\n"
    for activity, amount in data.items():
        emissions = calculator.calculate_emissions(activity, amount)
        report += f"{activity.capitalize()}: {amount} units, {emissions:.2f} kg CO2\n"
    report += f"\nTotal emissions: {total_emissions:.2f} kg CO2\n"
    report += provide_tips(total_emissions)
    return report

def provide_tips(total_emissions):
    tips = "\nEco-Friendly Tips:\n"
    if total_emissions > 100:
        tips += "- Consider reducing car travel.\n"
    if total_emissions > 50:
        tips += "- Use public transportation more often.\n"
    tips += "- Turn off lights when not needed.\n"
    return tips