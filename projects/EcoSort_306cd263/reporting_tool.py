    "plastic": 0,
    "glass": 0,
    "paper": 0,
    "metal": 0,
    "organic": 0,
    "other": 0
}

def update_report(item_type):
    if item_type in report_data:
        report_data[item_type] += 1
    else:
        report_data["other"] += 1

def generate_report():
    report_lines = [f"{item_type.capitalize()}: {count}" for item_type, count in report_data.items()]