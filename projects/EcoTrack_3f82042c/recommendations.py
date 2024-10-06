def generate_recommendations(profile):
    tips = []
    if profile['transport'] > 10:
        tips.append("Consider using public transport or carpooling.")
    if profile['electricity'] > 15:
        tips.append("Turn off appliances when not in use to save electricity.")
    if profile['meat_meals'] > 2:
        tips.append("Try to reduce meat consumption in your meals.")
    if profile['vegan_meals'] < 1:
        tips.append("Include more vegan meals in your diet.")
    return tips