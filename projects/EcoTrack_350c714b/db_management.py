
def save_data(user, activities, footprint):
    user_data_store[user] = {"activities": activities, "footprint": footprint}

def load_user_data(user):