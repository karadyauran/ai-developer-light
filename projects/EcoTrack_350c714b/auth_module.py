    users = {"user1": "password1", "user2": "password2"}
    username = input("Enter username: ")
    password = input("Enter password: ")
    if username in users and users[username] == password:
        print("Authentication successful.")
        return username
    else:
        print("Authentication failed.")