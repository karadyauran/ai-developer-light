def calculate_footprint(data):
    car_emission_factor = 0.21
    electricity_emission_factor = 0.5
    meat_meal_emission = 2.5
    vegan_meal_emission = 0.5
    transport_emissions = data['transport'] * car_emission_factor
    electricity_emissions = data['electricity'] * electricity_emission_factor
    meal_emissions = (data['meat_meals'] * meat_meal_emission +
                      data['vegan_meals'] * vegan_meal_emission)
    total_footprint = transport_emissions + electricity_emissions + meal_emissions
    return total_footprint