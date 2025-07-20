from typing import Protocol, Callable
from enum import Enum

"""
Second iteration of pizza:
- Base, size and toppings, calculate price of pizza.

Twist:
- Now we have full order (pizza/drinks/etc...).
- Now we have deals (1 for 1, free drinks with each pizza, free most expensive toppings)

"""

# class PizzaComponent(Protocol):
#     def get_component_price(self) -> int: ...


class PizzaComponent:
    def __init__(self, name, price, is_complementary):
        self.name = name
        self.price = price
        self.is_complementary = is_complementary


class Base(PizzaComponent):
    def __init__(self, name, price, is_complementary):
        super().__init__(name, price, is_complementary)


class Size(PizzaComponent):
    def __init__(self, name, price, is_complementary):
        super().__init__(name, price, is_complementary)


class Topping(PizzaComponent):
    def __init__(self, name, price, is_complementary):
        super().__init__(name, price, is_complementary)


class FoodType(Enum):
    FOOD_PIZZA = 1
    FOOD_DRINK = 2


class Food(Protocol):
    def get_price(self) -> int: ...
    def is_complementary(self) -> bool: ...
    def get_food_type(self) -> FoodType: ...


class Drink:
    def __init__(self, name, price, complementary):
        self.name = name
        self.price = price
        self.complementary = complementary

    def get_price(self):
        return self.price

    def is_complementary(self):
        return self.complementary

    def get_food_type(self):
        return FoodType.FOOD_DRINK


class Pizza:
    def __init__(self, size, base, toppings):
        self.size = size
        self.base = base
        self.toppings = toppings

    def get_price(self):
        return (
            self.size.price
            + self.base.price
            + sum([topping.price for topping in self.toppings])
        )

    def get_food_type(self):
        return FoodType.FOOD_PIZZA

    def is_complementary(self):
        return self.complementary


class Deal(Protocol):
    def apply_deal(self, order: "Order"): ...


class FreePizza:
    def __init__(self, free_pizza):
        self.free_pizza = free_pizza

    def apply_deal(self, order: "Order"):
        order.items.append()


class Order:
    def __init__(self, items: list[Food], applied_deals: list[Deal]):
        self.items = items
        self.deals = applied_deals

    def original_price(self):
        return sum(
            [item.get_price() for item in self.items if not item.is_complementary()]
        )
