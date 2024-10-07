from database_handler import DatabaseHandler
from analytics_engine import AnalyticsEngine
from recommendation_module import RecommendationModule

class EcoTrackUI:
    def __init__(self):
        self.db_handler = DatabaseHandler()
        self.analytics_engine = AnalyticsEngine(self.db_handler)
        self.recommendation_module = RecommendationModule(self.analytics_engine)

    def run(self):
        while True:
            print("EcoTrack - Main Menu")
            print("1. Input Data")
            print("2. View Analytics")
            print("3. Get Recommendations")
            print("4. Exit")
            choice = input("Choose an option: ")

            if choice == '1':
                self.input_data()
            elif choice == '2':
                self.view_analytics()
            elif choice == '3':
                self.get_recommendations()
            elif choice == '4':
                sys.exit()
            else:
                print("Invalid choice, please try again.")

    def input_data(self):
        date = input("Enter date (YYYY-MM-DD): ")
        activity = input("Enter activity: ")
        emission = float(input("Enter emission (kg CO2): "))
        self.db_handler.store_data(date, activity, emission)
        print("Data stored successfully.")

    def view_analytics(self):
        total_emissions = self.analytics_engine.calculate_total_emissions()
        print(f"Total Emissions: {total_emissions} kg CO2")

    def get_recommendations(self):
        recommendations = self.recommendation_module.provide_recommendations()
        print("Recommendations:")
        for rec in recommendations:
            print(f"- {rec}")

if __name__ == "__main__":
    app = EcoTrackUI()