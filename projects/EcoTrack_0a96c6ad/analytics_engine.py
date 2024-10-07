    def __init__(self, db_handler):
        self.db_handler = db_handler

    def calculate_total_emissions(self):
        data = self.db_handler.fetch_all_data()
        total_emissions = sum(entry[2] for entry in data)
        return total_emissions

    def calculate_daily_average_emissions(self):
        data = self.db_handler.fetch_all_data()
        if not data:
            return 0
        daily_emissions = {}
        for entry in data:
            date = entry[0]
            emission = entry[2]
            if date in daily_emissions:
                daily_emissions[date] += emission
            else:
                daily_emissions[date] = emission
        average_emissions = sum(daily_emissions.values()) / len(daily_emissions)
        return average_emissions

    def calculate_emissions_by_activity(self):
        data = self.db_handler.fetch_all_data()
        activity_emissions = {}
        for entry in data:
            activity = entry[1]
            emission = entry[2]
            if activity in activity_emissions:
                activity_emissions[activity] += emission
            else:
                activity_emissions[activity] = emission
        return activity_emissions

if __name__ == "__main__":
    from database_handler import DatabaseHandler
    db_handler = DatabaseHandler()
    analytics = AnalyticsEngine(db_handler)
    print("Total Emissions:", analytics.calculate_total_emissions())
    print("Daily Average Emissions:", analytics.calculate_daily_average_emissions())
    print("Emissions by Activity:", analytics.calculate_emissions_by_activity())