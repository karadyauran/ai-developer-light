
def identify_item(file_path):
    image = Image.open(file_path)
    width, height = image.size
    if width > height:
        return "plastic"
    elif height > width:
        return "glass"
    else: