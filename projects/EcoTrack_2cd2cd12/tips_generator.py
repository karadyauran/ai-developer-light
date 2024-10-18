import random

def generate_tips():
    tips = [
        "Use public transportation to reduce your carbon footprint.",
        "Turn off lights when not in use to save energy.",
        "Recycle and compost to minimize waste.",
        "Use a programmable thermostat to save on heating costs.",
        "Consider installing solar panels.",
        "Use energy-efficient appliances.",
        "Reduce water usage by fixing leaks and taking shorter showers.",
        "Choose reusable products over single-use items.",
        "Plant trees to absorb CO2 and improve air quality.",
        "Support renewable energy projects."
    ]
    return random.choice(tips)