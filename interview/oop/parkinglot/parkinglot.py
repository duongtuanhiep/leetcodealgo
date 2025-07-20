from enum import Enum
from abc import ABC, abstractmethod
from datetime import datetime

# 1. Design a Parking Lot System
# The parking lot can have multiple levels, each with multiple rows and spots of different sizes (compact, large, handicapped).
# Vehicles can be cars, bikes, or trucks, each requiring different spot sizes.
# Support for parking, unparking, and locating vehicles.
# Track available spots and display real-time availability.
# Handle entry/exit gates, ticketing, and payment processing.
# Implement fee calculation based on parking duration and vehicle type.
# Support for reserved spots and EV charging stations


class Size(Enum):
    CAR_SIZE_COMPACT = 1
    CAR_SIZE_LARGE = 2
    CAR_SIZE_HANDICAPPED = 3


class Vehicle(ABC):
    @abstractmethod
    def get_size(self) -> bool:
        pass

    @abstractmethod
    def get_plate(self) -> str:
        pass


class Car(Vehicle):
    pass


class Bike(Vehicle):
    pass


class Truck(Vehicle):
    pass


class Spot:
    def __init__(self, size: "Size", id: str):
        self.__size = size
        self.__id = id

    def size(self) -> "Size":
        return self.__size

    def location(self) -> str:
        return self.__id

    def name(self, name):
        self.__name = name


class Row:
    pass


class Level:
    pass


class ParkingInfo:
    def __init__(self):
        self.plate = ""
        self.checked_in = 0


class ParkingLot:
    def __init__(self):
        self.levels = []
        self.plate_to_spots = {}  # encode K: Plates V: spotsid
        self.available = {}  # K: size, V: list[spots]
        self.gates = []
        self.event_log = []  # K: plate, V:[entry gate, park, unpark, exit gate]

    def park_vehicle(self, vehicle: "Vehicle") -> tuple[str, bool]:
        plate = vehicle.get_plate()
        size = vehicle.get_size()

        # Check if it's in parking lot already
        if plate in self.plate_to_spots:
            return self.plate_to_spots[plate], False

        # Find a spot
        spot_location = ""
        if self.available[size]:
            occupied_spot = self.available[size].pop()

            # Assigned spot to car
            self.plate_to_spots[plate] = occupied_spot
            spot_location = occupied_spot.name()
        else:
            return None, False

        return spot_location, True

    def unpark_vehicle(self, vehicle: "Vehicle") -> bool:
        plate = vehicle.get_plate
        if plate in self.plate_to_spots:
            available_spot = self.plate_to_spots.pop(plate, None)
            self.available[available_spot.size()] = available_spot

        return True

    def locate_vehicle(self, vehicle: "Vehicle") -> str:
        return self.plate_to_spots[vehicle.get_plate()]

    pass


class FeeCalculator:
    def __init__(self):
        self.base_price_minute = {}  # K: vehicle size V: base price per minute

    def calculate(self, vehicle: "Vehicle", start_time, end_time) -> int:
        return (end_time - start_time).min * self.base_price_minute[vehicle.get_size()]
