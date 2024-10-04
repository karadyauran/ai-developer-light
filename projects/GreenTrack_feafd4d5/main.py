from data_handler import get_product_data
from carbon_calculator import calculate_carbon_footprint
from recommendations import suggest_alternatives

def display_menu():
    print("Welcome to GreenTrack")
    print("1. Enter Shopping List")
    print("2. Exit")
    choice = input("Choose an option: ")
    return choice

def enter_shopping_list():
    shopping_list = []
    print("Enter your shopping items (type 'done' to finish):")
    while True:
        item = input("Item: ")
        if item.lower() == 'done':
            break
        shopping_list.append(item)
    return shopping_list

def analyze_shopping_list(shopping_list):
    product_data = get_product_data(shopping_list)
    total_footprint = 0
    for product in product_data:
        footprint = calculate_carbon_footprint(product)
        total_footprint += footprint
        print(f"{product['name']}: {footprint} kg CO2")
    print(f"Total Carbon Footprint: {total_footprint} kg CO2")
    return product_data

def show_recommendations(product_data):
    print("Eco-friendly Alternatives:")
    for product in product_data:
        alternatives = suggest_alternatives(product)
        if alternatives:
            print(f"Alternatives for {product['name']}: {', '.join(alternatives)}")

def main():
    while True:
        choice = display_menu()
        if choice == '1':
            shopping_list = enter_shopping_list()
            product_data = analyze_shopping_list(shopping_list)
            show_recommendations(product_data)
        elif choice == '2':
            print("Thank you for using GreenTrack!")
            break

if __name__ == "__main__":
    main()