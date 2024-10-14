    "plastic": {"recyclable": True, "instructions": "Place in the plastic recycling bin."},
    "glass": {"recyclable": True, "instructions": "Place in the glass recycling bin."},
    "paper": {"recyclable": True, "instructions": "Place in the paper recycling bin."},
    "metal": {"recyclable": True, "instructions": "Place in the metal recycling bin."},
    "organic": {"recyclable": False, "instructions": "Place in compost bin."},
    "other": {"recyclable": False, "instructions": "Dispose of in general waste."},
}

def get_material_info(item_type):