from product_database import ProductDatabase

class UserInterface:
    def __init__(self):
        self.db = ProductDatabase()

    def display_menu(self):
        print("EcoShopper Menu")
        print("1. Add Product")
        print("2. View Product")
        print("3. Update Product")
        print("4. Delete Product")
        print("5. List All Products")
        print("6. Exit")

    def add_product(self):
        product_id = input("Enter Product ID: ")
        name = input("Enter Product Name: ")
        category = input("Enter Product Category: ")
        sustainability_score = float(input("Enter Sustainability Score: "))
        self.db.add_product(product_id, name, category, sustainability_score)
        print("Product added successfully.")

    def view_product(self):
        product_id = input("Enter Product ID: ")
        product = self.db.get_product(product_id)
        if product:
            print(f"Name: {product['name']}, Category: {product['category']}, "
                  f"Sustainability Score: {product['sustainability_score']}")
        else:
            print("Product not found.")

    def update_product(self):
        product_id = input("Enter Product ID: ")
        name = input("Enter new Product Name (or press Enter to skip): ")
        category = input("Enter new Product Category (or press Enter to skip): ")
        score_input = input("Enter new Sustainability Score (or press Enter to skip): ")
        sustainability_score = float(score_input) if score_input else None
        self.db.update_product(product_id, name, category, sustainability_score)
        print("Product updated successfully.")

    def delete_product(self):
        product_id = input("Enter Product ID: ")
        self.db.delete_product(product_id)
        print("Product deleted successfully.")

    def list_products(self):
        products = self.db.list_products()
        if products:
            for product in products:
                print(f"Name: {product['name']}, Category: {product['category']}, "
                      f"Sustainability Score: {product['sustainability_score']}")
        else:
            print("No products available.")

    def run(self):
        while True:
            self.display_menu()
            choice = input("Select an option: ")
            if choice == '1':
                self.add_product()
            elif choice == '2':
                self.view_product()
            elif choice == '3':
                self.update_product()
            elif choice == '4':
                self.delete_product()
            elif choice == '5':
                self.list_products()
            elif choice == '6':
                sys.exit()
            else:
                print("Invalid option, please try again.")

if __name__ == "__main__":
    ui = UserInterface()