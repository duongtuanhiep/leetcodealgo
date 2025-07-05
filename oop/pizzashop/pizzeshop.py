from typing import Callable
from copy import deepcopy


class Order:
    def __init__(self, items: list["PriceFunc"], deals: list["DealFunc"]):
        self.items = items
        self.deals = {}
        for deal in deals:
            self.deals[deal.__name__] = deal

    def add(self, *args, **kwargs):
        pass

    def total_order(self, deals: list["DealFunc"]) -> "Order":
        finalOrder = self
        for deal in deals:
            finalOrder = deal(finalOrder)
        return finalOrder


# order = Order(stuffs)
# deals = [one_pizza_for_free]
# order.total_order(deals)


PriceFunc = Callable[[], int]

DealFunc = Callable[[Order]]


def one_pizza_for_free() -> DealFunc:
    def add_pizza(order: "Order"):
        added = False
        orderLen = len(order.items)
        for i in range(orderLen):
            if order.items[i].type() == FOOD_TYPE_PIZZA and not added:
                newPizze = order.items[i].dup()
                order.items.append(newPizze)
                added = True

    return add_pizza


def free_drink_each_pizza(drink: "Drink") -> DealFunc:
    def add_free_drinks(order: "Order"):
        pizzaCnt = sum(item.type() == FOOD_TYPE_PIZZA for item in order.items)
        for _ in range(pizzaCnt):
            order.items.append(deepcopy(drink))

    return add_free_drinks


FOOD_TYPE_UNDEFINED = "UNDEFINED"
FOOD_TYPE_DRINK = "DRINK"
FOOD_TYPE_PIZZA = "PIZZA"


# Food(ABC)
class Food:
    def __init__(self, name, price, is_free):
        self.name = name
        self.price = price
        self.is_free = is_free

    # @abstractmethod
    def price(self) -> int:
        return self.price

    # @abstractmethod
    def type(self) -> str:
        return FOOD_TYPE_UNDEFINED

    # @abstractmethod
    def as_new(self):
        return Food(self.name, self.price)

    # @abstractmethod
    def is_complementary() -> bool:
        pass


class Drink(Food):
    def __init__(self, name, price):
        self.name = name
        self.price = price

    def price(self):
        return self.price

    def as_new(self) -> "Drink":
        pass


class Pizza(Food):
    def __init__(self, base: "Base", size: "Size", toppings: list["Topping"]):
        self.base = base
        self.size = size
        self.toppings = toppings

    def price(self) -> int:
        return self.base.price + self.size.price + self.toppings.price


class Base:
    def __init__(self, price, name):
        self.price = price
        self.name = name


class Size:
    def __init__(self, price, name):
        self.price = price
        self.name = name


class Topping:
    def __init__(self, price, name):
        self.price = price
        self.name = name
