from datainput import get_user_data
from carboncalculator import calculate_footprint
from recommendations import generate_recommendations
from userprofile import load_user_profile, save_user_progress

def display_menu():
    print("Welcome to EcoTrack")
    print("1. Enter Daily Activities")
    print("2. View Carbon Footprint")
    print("3. Get Recommendations")
    print("4. View Profile")
    print("5. Exit")

def enter_daily_activities():
    user_data = get_user_data()
    footprint = calculate_footprint(user_data)
    save_user_progress(user_data, footprint)
    print(f"Your carbon footprint for today is {footprint} kg CO2")

def view_carbon_footprint():
    profile = load_user_profile()
    print(f"Your cumulative carbon footprint is {profile['total_footprint']} kg CO2")

def get_recommendations():
    profile = load_user_profile()
    recommendations = generate_recommendations(profile)
    print("Here are some tips to reduce your carbon footprint:")
    for tip in recommendations:
        print(f"- {tip}")

def view_profile():
    profile = load_user_profile()
    print(f"User: {profile['username']}")
    print(f"Total Carbon Footprint: {profile['total_footprint']} kg CO2")
    print(f"Days Tracked: {profile['days_tracked']}")

def main():
    while True:
        display_menu()
        choice = input("Choose an option (1-5): ")
        if choice == '1':
            enter_daily_activities()
        elif choice == '2':
            view_carbon_footprint()
        elif choice == '3':
            get_recommendations()
        elif choice == '4':
            view_profile()
        elif choice == '5':
            print("Exiting EcoTrack. Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()