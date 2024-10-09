    activities = []
    while True:
        activity = input("Enter activity (or 'done' to finish): ")
        if activity.lower() == 'done':
            break
        duration = float(input("Enter duration in hours: "))
        activities.append({"activity": activity, "duration": duration})