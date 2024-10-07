
class DatabaseHandler:
    def __init__(self):
        self.connection = sqlite3.connect('ecotrack.db')
        self.cursor = self.connection.cursor()
        self.create_table()

    def create_table(self):
        self.cursor.execute('''
            CREATE TABLE IF NOT EXISTS emissions (
                id INTEGER PRIMARY KEY,
                date TEXT NOT NULL,
                activity TEXT NOT NULL,
                emission REAL NOT NULL
            )
        ''')
        self.connection.commit()

    def store_data(self, date, activity, emission):
        self.cursor.execute('''
            INSERT INTO emissions (date, activity, emission)
            VALUES (?, ?, ?)
        ''', (date, activity, emission))
        self.connection.commit()

    def fetch_all_data(self):
        self.cursor.execute('SELECT date, activity, emission FROM emissions')
        return self.cursor.fetchall()

    def fetch_data_by_date(self, date):
        self.cursor.execute('SELECT activity, emission FROM emissions WHERE date = ?', (date,))
        return self.cursor.fetchall()

    def close_connection(self):
        self.connection.close()

if __name__ == "__main__":
    db = DatabaseHandler()
    db.store_data('2023-10-01', 'Driving', 12.5)
    data = db.fetch_all_data()
    print(data)