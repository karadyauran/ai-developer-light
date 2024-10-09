    recommendations = []
    if footprint > 10:
        recommendations.append("Consider reducing car usage.")
    if footprint > 20:
        recommendations.append("Explore alternative transport modes like cycling.")
    if not recommendations:
        recommendations.append("Great job! Keep maintaining your low footprint.")
    for recommendation in recommendations: