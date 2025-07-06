from abc import ABC, abstractmethod
from typing import Callable
from enum import Enum

"""
Design a Food Delivery System

Base Requirements:
Model the main classes and relationships for a food delivery platform (like Swiggy, DoorDash, or Uber Eats).
The system should support multiple restaurants, each with its own menu.
Customers can browse menus, place orders, and track their delivery status.
Include basic functionality for order placement, assignment to delivery personnel, and status updates (e.g., order placed, being prepared, out for delivery, delivered).
Focus on extensibility and clarity in your class and interface design.
"""


class Restaurant:
    def __init__(self, menu, loc: "Location"):
        self.menu = menu
        self.loc = loc


class Menu:
    def __init__(self, name, description, items):
        self.name = name
        self.description = description
        self.items = items


class FoodItem:
    def __init__(self, name, description, price):
        self.name = name
        self.description = description
        self.price = price


class OrderStatus(Enum):
    ORDER_PLACED = 1
    ORDER_BEING_PREPARED = 2
    ORDER_OUT_FOR_DELIVERY = 3
    ORDER_DELIVERED = 4


class Order:
    def __init__(self, id, restaurant, items):
        self.id = id
        self.restaurant = restaurant
        self.items = items
        self.status = OrderStatus.ORDER_PLACED

    def update(self, status):
        self.status = status

    def assign(self, courier: "Courier"):
        self.courier = courier

    def get_status(self) -> str:
        return f"Order from {self.restaurant} is {self.status} with current loc {self.courier.loc()}"

    def base_price(self) -> int:
        return sum([item.price for item in self.items])

    def update_dynamic_price(self, int):
        self.dynamic_price = int


class OrderTracker:
    def __init__(self):
        self.order = {}

    def add(self, order: "Order"):
        self.order[order.id] = order

    def assign(self, order: "Order", courier: "Courier"):
        self.order[order.id].assign(courier)

    def update(self, order: "Order", status: "OrderStatus"):
        self.order[order.id].update(status)

    def track(self, tracking_order) -> str:
        id = tracking_order.id
        if id in self.order:
            return self.order[id].get_status()
        else:
            return "not found"


class PriceCal:
    def dynamicPrice(originalPrice, courier_availability):
        return 123  # Some math based off courier cnt and price


class Location:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def is_in_range(self, loc: "Location") -> bool:
        return True

    def dist(self, loc: "Location") -> int:
        return 123


class CourierManager:
    def __init__(self):
        pass

    def available_couriers(self) -> int:
        return 10

    def closest_courier(self, loc: Location) -> "Courier":
        # Gets all the courier within X distance and return closest:
        pass


class CourierStatus(Enum):
    COURIER_STATUS_TRANSIT = 1
    COURIER_STATUS_IDLE = 2
    COURIER_STATUS_UNAVAILABLE = 3


class Courier:
    def __init__(self, loc: Location):
        self.status = CourierStatus.COURIER_STATUS_IDLE
        self.loc = loc

    def loc(self):
        return self.loc

    def update_loc(self, loc):
        self.loc = loc

    def update_status(self, status):
        self.status = status


class Customer:
    def __init__(self):
        self.orders = []

    def add_order(self, order):
        self.ordes.append(order)

    def latest_order(self):
        if self.order:
            return self.order[-1]

    def all_order(self):
        return self.order


class OrderSystem:
    order_id = 1

    def __init__(
        self,
        restaurants: list[Restaurant],
        order_tracker: "OrderTracker",
        courier_manager: "CourierManager",
    ):
        self.restaurants = restaurants
        self.order_tracker = order_tracker
        self.courier_manager = courier_manager

    def search(self, user_loc: Location) -> list[Restaurant]:
        return [
            restaurant
            for restaurant in self.restaurants
            if restaurant.loc.is_in_range(user_loc)
        ]

    def place_order(
        self, customer: "Customer", restaurant: "Restaurant", items: list[Restaurant]
    ) -> "Order":
        order = Order(OrderSystem.order_id, restaurant, items)
        OrderSystem.order_id += 1

        self.order_tracker.add(order)
        customer.add_order(order)
        return order

    def assign_courier(self, courier: "Courier", order: "Order"):
        closestCourier = self.courier_manager.closest_courier(order.restaurant.loc)
        self.order_tracker.assign(order, closestCourier)
