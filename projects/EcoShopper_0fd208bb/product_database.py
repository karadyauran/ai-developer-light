
class ProductDatabase:
    def __init__(self, db_file='products.json'):
        self.db_file = db_file
        self.products = self.load_products()

    def load_products(self):
        try:
            with open(self.db_file, 'r') as file:
                return json.load(file)
        except FileNotFoundError:
            return {}

    def save_products(self):
        with open(self.db_file, 'w') as file:
            json.dump(self.products, file, indent=4)

    def add_product(self, product_id, name, category, sustainability_score):
        if product_id not in self.products:
            self.products[product_id] = {
                'name': name,
                'category': category,
                'sustainability_score': sustainability_score
            }
            self.save_products()

    def get_product(self, product_id):
        return self.products.get(product_id, None)

    def update_product(self, product_id, name=None, category=None, sustainability_score=None):
        if product_id in self.products:
            if name:
                self.products[product_id]['name'] = name
            if category:
                self.products[product_id]['category'] = category
            if sustainability_score is not None:
                self.products[product_id]['sustainability_score'] = sustainability_score
            self.save_products()

    def delete_product(self, product_id):
        if product_id in self.products:
            del self.products[product_id]
            self.save_products()

    def list_products(self):