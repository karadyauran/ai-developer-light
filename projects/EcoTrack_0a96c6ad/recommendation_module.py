    def __init__(self, analytics_engine):
        self.analytics_engine = analytics_engine

    def provide_recommendations(self):
        activity_emissions = self.analytics_engine.calculate_emissions_by_activity()
        recommendations = []
        for activity, emission in activity_emissions.items():
            if emission > 10:
                recommendations.append(f"Consider reducing {activity} to lower emissions.")
            else:
                recommendations.append(f"Continue your efforts in managing {activity}.")
        return recommendations

if __name__ == "__main__":
    from database_handler import DatabaseHandler
    from analytics_engine import AnalyticsEngine
    db_handler = DatabaseHandler()
    analytics = AnalyticsEngine(db_handler)
    recommendations = RecommendationModule(analytics)
    for rec in recommendations.provide_recommendations():
        print(rec)