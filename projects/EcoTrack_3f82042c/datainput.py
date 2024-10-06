def get_user_data():
    data = {}
    print("Enter your daily activities:")
    data['transport'] = float(input("Distance traveled by car (km): "))
    data['electricity'] = float(input("Electricity used (kWh): "))
    data['meat_meals'] = int(input("Number of meat-based meals: "))
    data['vegan_meals'] = int(input("Number of vegan meals: "))
    return data