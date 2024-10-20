    def __init__(self, db):
        self.db = db
        self.cart = {}

    def add_to_cart(self, product_id, quantity):
        product = self.db.get_product(product_id)
        if product:
            if product_id in self.cart:
                self.cart[product_id]['quantity'] += quantity
            else:
                self.cart[product_id] = {'product': product, 'quantity': quantity}
            print(f"Added {quantity} of {product['name']} to cart.")
        else:
            print("Product not found.")

    def remove_from_cart(self, product_id):
        if product_id in self.cart:
            del self.cart[product_id]
            print("Product removed from cart.")
        else:
            print("Product not in cart.")

    def view_cart(self):
        if not self.cart:
            print("Cart is empty.")
            return
        for item in self.cart.values():
            product = item['product']
            print(f"Name: {product['name']}, Quantity: {item['quantity']}, "
                  f"Sustainability Score: {product['sustainability_score']}")

    def calculate_total_sustainability_score(self):
        if not self.cart:
            return 0
        total_score = sum(item['product']['sustainability_score'] * item['quantity'] for item in self.cart.values())
        total_quantity = sum(item['quantity'] for item in self.cart.values())
        return total_score / total_quantity if total_quantity else 0

if __name__ == "__main__":
    from product_database import ProductDatabase
    db = ProductDatabase()
    cart = ShoppingCart(db)
    cart.add_to_cart('001', 2)
    cart.view_cart()
    print(f"Average Sustainability Score in Cart: {cart.calculate_total_sustainability_score()}")
    cart.remove_from_cart('001')