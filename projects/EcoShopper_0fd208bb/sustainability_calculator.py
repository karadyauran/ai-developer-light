    def __init__(self, db):
        self.db = db

    def calculate_average_sustainability_score(self):
        products = self.db.list_products()
        if not products:
            return 0
        total_score = sum(product['sustainability_score'] for product in products)
        return total_score / len(products)

    def get_most_sustainable_product(self):
        products = self.db.list_products()
        if not products:
            return None
        return max(products, key=lambda x: x['sustainability_score'])

    def get_least_sustainable_product(self):
        products = self.db.list_products()
        if not products:
            return None
        return min(products, key=lambda x: x['sustainability_score'])

if __name__ == "__main__":
    from product_database import ProductDatabase
    db = ProductDatabase()
    calculator = SustainabilityCalculator(db)
    avg_score = calculator.calculate_average_sustainability_score()
    most_sustainable = calculator.get_most_sustainable_product()
    least_sustainable = calculator.get_least_sustainable_product()
    print(f"Average Sustainability Score: {avg_score}")
    if most_sustainable:
        print(f"Most Sustainable Product: {most_sustainable['name']} "
              f"with score {most_sustainable['sustainability_score']}")
    if least_sustainable:
        print(f"Least Sustainable Product: {least_sustainable['name']} "