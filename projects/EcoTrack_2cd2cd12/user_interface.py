import tkinter as tk
from data_collection import collect_data
from carbon_calculator import calculate_emissions
from tips_generator import generate_tips
from report_generator import generate_report

class EcoTrackApp:
    def __init__(self, root):
        self.root = root
        self.root.title("EcoTrack")
        self.create_widgets()

    def create_widgets(self):
        self.label = tk.Label(self.root, text="EcoTrack: Monitor Your Carbon Footprint")
        self.label.pack(pady=10)

        self.collect_button = tk.Button(self.root, text="Collect Data", command=self.collect_data)
        self.collect_button.pack(pady=5)

        self.calculate_button = tk.Button(self.root, text="Calculate Emissions", command=self.calculate_emissions)
        self.calculate_button.pack(pady=5)

        self.tips_button = tk.Button(self.root, text="Generate Tips", command=self.generate_tips)
        self.tips_button.pack(pady=5)

        self.report_button = tk.Button(self.root, text="Generate Report", command=self.generate_report)
        self.report_button.pack(pady=5)

        self.quit_button = tk.Button(self.root, text="Quit", command=self.root.quit)
        self.quit_button.pack(pady=5)

    def collect_data(self):
        collect_data()
        self.show_message("Data collected successfully.")

    def calculate_emissions(self):
        emissions = calculate_emissions()
        self.show_message(f"Emissions calculated: {emissions} kg CO2")

    def generate_tips(self):
        tips = generate_tips()
        self.show_message(f"Eco Tips: {tips}")

    def generate_report(self):
        report = generate_report()
        self.show_message("Report generated. Check your output folder.")

    def show_message(self, message):
        self.message_label = tk.Label(self.root, text=message)
        self.message_label.pack(pady=5)

if __name__ == "__main__":
    root = tk.Tk()
    app = EcoTrackApp(root)
    root.mainloop()