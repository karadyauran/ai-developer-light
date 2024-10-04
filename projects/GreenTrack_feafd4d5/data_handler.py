product_database = {
    "apple": {"name": "apple", "carbon_footprint": 0.3},
    "banana": {"name": "banana", "carbon_footprint": 0.2},
    "beef": {"name": "beef", "carbon_footprint": 27.0},
    "chicken": {"name": "chicken", "carbon_footprint": 6.9},
    "milk": {"name": "milk", "carbon_footprint": 1.9},
    "rice": {"name": "rice", "carbon_footprint": 2.7},
    "bread": {"name": "bread", "carbon_footprint": 0.6},
    "carrot": {"name": "carrot", "carbon_footprint": 0.1},
}

def get_product_data(shopping_list):
    product_data = []
    for item in shopping_list:
        product_info = product_database.get(item.lower())
        if product_info:
            product_data.append(product_info)
    return product_data