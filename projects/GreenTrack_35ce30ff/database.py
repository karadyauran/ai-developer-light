user_database = {}

def save_user_data(data):
    name = data.get('name')
    if name:
        user_database[name] = data

def get_user_data(name):
    return user_database.get(name)