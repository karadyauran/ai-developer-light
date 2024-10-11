def generate_report(emissions):
    report = (
        f"Emission Report:\n"
        f"Transport Emissions: {emissions['transport']} kg\n"
        f"Energy Emissions: {emissions['energy']} kg\n"
        f"Waste Emissions: {emissions['waste']} kg\n"
        f"Total Emissions: {emissions['total']} kg\n"
    )
    return report