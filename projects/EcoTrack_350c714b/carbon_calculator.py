    footprint_factors = {
        "driving": 2.31,
        "flying": 5.0,
        "cycling": 0.05,
        "walking": 0.0
    }
    total_footprint = 0.0
    for activity in activities:
        factor = footprint_factors.get(activity["activity"], 0.0)
        total_footprint += factor * activity["duration"]