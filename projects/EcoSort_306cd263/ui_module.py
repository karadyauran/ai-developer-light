from PIL import Image, ImageTk
import image_recognition
import material_database
import sorting_algorithm
import reporting_tool

class EcoSortApp:
    def __init__(self, root):
        self.root = root
        self.root.title("EcoSort")
        self.label = Label(root, text="Welcome to EcoSort. Click 'Scan Item' to begin.")
        self.label.pack()
        self.scan_button = Button(root, text="Scan Item", command=self.scan_item)
        self.scan_button.pack()
        self.result_label = Label(root, text="")
        self.result_label.pack()
        self.report_button = Button(root, text="View Report", command=self.view_report)
        self.report_button.pack()

    def scan_item(self):
        file_path = filedialog.askopenfilename()
        if file_path:
            item_type = image_recognition.identify_item(file_path)
            material_info = material_database.get_material_info(item_type)
            sorting_instructions = sorting_algorithm.get_sorting_instructions(material_info)
            self.result_label.config(text=sorting_instructions)
            reporting_tool.update_report(item_type)

    def view_report(self):
        report = reporting_tool.generate_report()
        self.result_label.config(text=report)

def main():
    root = Tk()
    app = EcoSortApp(root)
    root.mainloop()

if __name__ == "__main__":