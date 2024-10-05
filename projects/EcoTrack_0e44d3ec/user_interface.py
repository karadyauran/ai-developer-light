from carbon_calculator import calculate_footprint
from data_storage import save_user_data, load_user_data
from tips import get_tips
from analytics import get_analytics

class EcoTrackApp:
    def __init__(self, root):
        self.root = root
        self.root.title("EcoTrack")
        self.activities_frame = tk.Frame(self.root)
        self.activities_frame.pack(pady=20)
        self.distance_label = tk.Label(self.activities_frame, text="Distance Traveled (km):")
        self.distance_label.grid(row=0, column=0)
        self.distance_entry = tk.Entry(self.activities_frame)
        self.distance_entry.grid(row=0, column=1)
        self.electricity_label = tk.Label(self.activities_frame, text="Electricity Used (kWh):")
        self.electricity_label.grid(row=1, column=0)
        self.electricity_entry = tk.Entry(self.activities_frame)
        self.electricity_entry.grid(row=1, column=1)
        self.calculate_button = tk.Button(self.root, text="Calculate Footprint", command=self.calculate)
        self.calculate_button.pack(pady=10)
        self.result_label = tk.Label(self.root, text="")
        self.result_label.pack(pady=10)
        self.tips_button = tk.Button(self.root, text="Get Tips", command=self.show_tips)
        self.tips_button.pack(pady=5)
        self.analytics_button = tk.Button(self.root, text="View Analytics", command=self.show_analytics)
        self.analytics_button.pack(pady=5)
        self.load_user_data()

    def calculate(self):
        distance = float(self.distance_entry.get())
        electricity = float(self.electricity_entry.get())
        footprint = calculate_footprint(distance, electricity)
        self.result_label.config(text=f"Your Carbon Footprint: {footprint} kg CO2")
        save_user_data(distance, electricity)

    def show_tips(self):
        tips = get_tips()
        tips_window = tk.Toplevel(self.root)
        tips_window.title("Eco-Friendly Tips")
        tips_label = tk.Label(tips_window, text=tips)
        tips_label.pack(pady=10, padx=10)

    def show_analytics(self):
        analytics = get_analytics()
        analytics_window = tk.Toplevel(self.root)
        analytics_window.title("Your Analytics")
        analytics_label = tk.Label(analytics_window, text=analytics)
        analytics_label.pack(pady=10, padx=10)

    def load_user_data(self):
        data = load_user_data()
        if data:
            self.distance_entry.insert(0, data['distance'])
            self.electricity_entry.insert(0, data['electricity'])

if __name__ == "__main__":
    root = tk.Tk()
    app = EcoTrackApp(root)