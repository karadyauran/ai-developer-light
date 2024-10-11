import requests

def fetch_external_data():
    try:
        response = requests.get("https://api.example.com/sustainability")
        if response.status_code == 200:
            return response.json()
        else:
            return {"error": "Failed to fetch data"}
    except requests.RequestException as e:
        return {"error": str(e)}