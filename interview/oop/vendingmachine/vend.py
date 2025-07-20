from abc import ABC, abstractmethod
from enum import Enum
from datetime import datetime
from dataclasses import dataclass
from collections import defaultdict

"""
Vending machine questions:
The vending machine has multiple slots/compartments that can hold different products (chips, sodas, candy bars, etc.).
Each product has a price, and customers can insert coins/bills to make purchases.
The machine should accept various denominations of money (quarters, dollars, etc.) and provide change when necessary.
Customers can select a product, and the machine should dispense it if sufficient payment is provided and the item is in stock.
The machine should track inventory levels for each product and prevent sales when items are out of stock.
Handle scenarios where a customer inserts money but cancels the transaction (return money).
Support basic maintenance operations like restocking products and collecting money.
"""


class Denomination(ABC):
    @abstractmethod
    def get_integer_value(self):
        pass


class OneDollar(Denomination):
    def get_integer_value(self):
        return 100


class FiveDollar(Denomination):
    def get_integer_value(self):
        return 500


class Product:
    id: str
    name: str
    price: str

    def __init__(self, id, name, price):
        self.id = id
        self.name = name
        self.price = price

    def get_name(self):
        return self.name

    def get_price(self):
        return self.price


class Chip(Product):
    def __init__(self, id, name, price):
        super().__init__(id, name, price)

    pass


class Drink(Product):
    pass


class ProductStock(ABC):
    @abstractmethod
    def increment_stock(self, product):
        pass

    @abstractmethod
    def has_stock(self, product):
        pass

    @abstractmethod
    def decrement_stock(self, product):
        pass


class ProductManager(ProductStock):
    productToCnt: defaultdict

    def increment_stock(self, product: Product):
        self.productToCnt[product.get_name()] += 1

    def has_stock(self, product: Product):
        return self.productToCnt[product.get_name()]

    def decrement_stock(self, product: Product):
        if self.has_stock(product):
            self.productToCnt[product.get_name()] -= 1


class CashOps(ABC):
    @abstractmethod
    def add_money(self, money: Denomination):
        pass

    @abstractmethod
    def subtract_money(self, int_value: int):
        pass

    @abstractmethod
    def get_balance(self):
        pass


class CashSession(CashOps):
    current: int

    def add_money(self, money: Denomination):
        self.current += money.get_integer_value()

    def subtract_money(self, int_value: int):
        if self.current > int_value:
            self.current -= int_value
            return self.current, True
        else:
            return self.current, False

    def get_balance(self):
        return self.current


class MaintananceOps(ABC):
    @abstractmethod
    def restock(self, products: list[Product]):
        pass

    @abstractmethod
    def withdraw(self, money_int: int):
        pass


class VendingMachineMaintanance(MaintananceOps):
    vending_machine_cash: CashOps
    vending_machine_product: ProductStock

    def restock(self, products: list[Product]):
        for p in products:
            self.vending_machine_product.increment_stock(p)

    def withdraw(self, money_int: int):
        if money_int < self.vending_machine_cash.get_balance():
            self.vending_machine_cash.subtract_money(money_int)


class VendingMachine:
    maintanance: VendingMachineMaintanance
    customer_session: CashOps
    machine_cash_session: CashOps
    inventory: ProductStock

    def customer_add(self, money: Denomination):
        self.customer_session.add_money(money)

    def customer_purchase(self, product: Product):
        product_cost = product.get_price()
        if (
            self.customer_session.get_balance() > product_cost
            and self.inventory.has_stock(product)
        ):
            remains, _ = self.customer_session.subtract_money(product_cost)
            self.machine_cash_session.add_money(product_cost)
            self.inventory.decrement_stock(product)
        # Return product

    def customer_cancel(self):
        balance = self.customer_session.get_balance()
        self.customer_session.subtract_money(balance)
        # Return money
