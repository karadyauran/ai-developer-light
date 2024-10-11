from data_processing import calculate_emissions
from database import save_user_data, get_user_data
from api_integration import fetch_external_data
from report_generator import generate_report

class GreenTrackUI:
    def __init__(self):
        self.user_data = {}

    def main_menu(self):
        while True:
            print("GreenTrack Main Menu")
            print("1. Enter Emissions Data")
            print("2. View Emission Report")
            print("3. Fetch External Data")
            print("4. Exit")
            choice = input("Choose an option: ")

            if choice == '1':
                self.enter_emissions_data()
            elif choice == '2':
                self.view_emission_report()
            elif choice == '3':
                self.fetch_external_data()
            elif choice == '4':
                break
            else:
                print("Invalid choice. Please try again.")

    def enter_emissions_data(self):
        name = input("Enter your name: ")
        transport = float(input("Enter transport emissions in kg: "))
        energy = float(input("Enter energy emissions in kg: "))
        waste = float(input("Enter waste emissions in kg: "))
        self.user_data = {
            'name': name,
            'transport': transport,
            'energy': energy,
            'waste': waste
        }
        save_user_data(self.user_data)
        print("Data saved successfully.")

    def view_emission_report(self):
        name = input("Enter your name: ")
        data = get_user_data(name)
        if data:
            emissions = calculate_emissions(data)
            report = generate_report(emissions)
            print(report)
        else:
            print("No data found for this user.")

    def fetch_external_data(self):
        data = fetch_external_data()
        print("Fetched external data:", data)

if __name__ == "__main__":
    ui = GreenTrackUI()
    ui.main_menu()