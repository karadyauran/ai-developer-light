alternative_database = {
    "beef": ["chicken", "tofu", "lentils"],
    "milk": ["almond milk", "soy milk"],
    "rice": ["quinoa", "cauliflower rice"],
    "bread": ["whole grain bread", "gluten-free bread"]
}

def suggest_alternatives(product):
    return alternative_database.get(product['name'], [])