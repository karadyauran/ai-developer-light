import tracker
import calculator
import report_generator

def display_menu():
    print("GreenTrack")
    print("1. Log Activity")
    print("2. View Weekly Report")
    print("3. Exit")

def get_user_choice():
    choice = input("Enter your choice: ")
    return choice

def log_activity():
    activity = input("Enter activity (e.g., car, bus, walk, electricity): ")
    amount = float(input("Enter amount (e.g., miles, kWh): "))
    tracker.log_activity(activity, amount)
    print("Activity logged.")

def view_weekly_report():
    report = report_generator.generate_report()
    print(report)

def main():
    while True:
        display_menu()
        choice = get_user_choice()
        if choice == '1':
            log_activity()
        elif choice == '2':
            view_weekly_report()
        elif choice == '3':
            print("Exiting GreenTrack.")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == '__main__':
    main()