from data_input import get_user_activities
from carbon_calculator import calculate_footprint
from data_visualization import generate_report
from db_management import save_data, load_user_data
from notifications import send_recommendations

def main():
    user = authenticate_user()
    if not user:
        print("Authentication failed.")
        return

    user_data = load_user_data(user)
    if not user_data:
        print("Failed to load user data.")
        return

    activities = get_user_activities()
    footprint = calculate_footprint(activities)
    save_data(user, activities, footprint)

    report = generate_report(activities, footprint)
    print(report)

    send_recommendations(user, footprint)

if __name__ == "__main__":