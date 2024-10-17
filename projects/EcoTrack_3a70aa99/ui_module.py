import data_processing
import reporting

class EcoTrackUI:
    def __init__(self, master):
        self.master = master
        master.title("EcoTrack")

        self.label = Label(master, text="Enter Activity Details")
        self.label.pack()

        self.activity_label = Label(master, text="Activity")
        self.activity_label.pack()
        self.activity_entry = Entry(master)
        self.activity_entry.pack()

        self.amount_label = Label(master, text="Amount")
        self.amount_label.pack()
        self.amount_entry = Entry(master)
        self.amount_entry.pack()

        self.submit_button = Button(master, text="Submit", command=self.submit_activity)
        self.submit_button.pack()

        self.report_button = Button(master, text="Generate Report", command=self.generate_report)
        self.report_button.pack()

        self.status_var = StringVar()
        self.status_label = Label(master, textvariable=self.status_var)
        self.status_label.pack()

    def submit_activity(self):
        activity = self.activity_entry.get()
        try:
            amount = float(self.amount_entry.get())
            data_processing.log_activity(activity, amount)
            self.status_var.set(f"Logged: {activity} - {amount}")
            self.activity_entry.delete(0, 'end')
            self.amount_entry.delete(0, 'end')
        except ValueError:
            messagebox.showerror("Invalid Input", "Please enter a numeric value for amount.")

    def generate_report(self):
        report = reporting.create_report()
        messagebox.showinfo("Report", report)

if __name__ == "__main__":
    root = Tk()
    eco_track_ui = EcoTrackUI(root)