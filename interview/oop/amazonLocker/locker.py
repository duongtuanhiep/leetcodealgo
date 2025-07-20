from abc import ABC, abstractmethod
from enum import Enum
from datetime import datetime
from dataclasses import dataclass

"""
Core Requirements (What's Actually Asked):

Design a system for Amazon Lockers where customers can pick up packages
Support different locker sizes (Small, Medium, Large)
Customers get a pickup code when package is delivered
Packages expire after a certain time period
System should assign appropriate locker based on package size
Handle package delivery and pickup flow

Packeage delivery: 
add_pkg(pkg) -> lockerID + pincode
- Pkg get asked to be put in the system
- System checks availability according to the date
- System accept/deny and return a pin code

- Pkg arrived and get delivered into the locker
- Deliveryman/services ACK the delivery.

Pickup Flow: 
- Customer received msg with pin code and locker
- Customer arrived at the location and put pin code in and open. 
- Customer closed locker and system register it as available again

"""


class ItemSize(Enum):
    SIZE_SMALL = 1
    SIZE_MEDIUM = 2
    SIZE_LARGE = 3


class User:
    id: int
    name: str
    phone: str
    mail: str

    def get_name(self):
        return self.name


class Locker:
    id: int
    size: ItemSize
    is_occupied: bool
    pin: str
    package: "Package" = None  # Optional field with default value

    def get_id(self) -> str:
        return self.id

    def get_size(self) -> ItemSize:
        return self.size

    def get_pin(self) -> str:
        return self.pin

    def update_pin(self, new_pin) -> str:
        self.pin = new_pin

    def get_item_description(self) -> str:
        if self.package:
            return self.package.get_description()
        return "empty"

    def add_package(self, pkg: "Package"):
        self.package = pkg

    def remove_package(self):
        self.package = None


class Package:
    id: int
    description: str
    size: ItemSize
    owner: User

    def get_owner(self):
        return self.owner

    def get_size(self):
        return self.size

    def get_description(self):
        return self.description


class Notify(ABC):
    @abstractmethod
    def notify_user(self, user: User, locker: Locker):
        pass


class SMSNotify(Notify):
    def notify_user(self, user: User, locker: Locker):
        print(f"Sent to {user.get_name()} for locker {locker.get_item_description()}")


class LockersPool(ABC):
    @abstractmethod
    def reserve(self) -> Locker:
        pass

    @abstractmethod
    def free(self, locker: Locker):
        pass

    @abstractmethod
    def refresh(self):
        pass


class LockerReserver(LockersPool):
    freeLocker: dict[str, Locker]
    occupiedLocker: dict[str, Locker]

    def reserve(self):
        if self.freeLocker:
            available_locker = next(iter(self.freeLocker.values()))
            self.freeLocker.pop(available_locker.get_id(), None)
            self.occupiedLocker[available_locker.get_id()] = available_locker
            return available_locker
        return None

    def free(self, locker):
        if locker.get_id() in self.occupiedLocker:
            self.occupiedLocker.pop(locker.get_id(), None)
            self.freeLocker[locker.get_id()] = locker

    def refresh(self):
        for lockerID, locker in self.occupiedLocker:
            # Handle freeing up expired locker here
            pass


class PinGenerator(ABC):
    @abstractmethod
    def generate_pin(self):
        pass


class LockerManager:
    notifier: Notify
    lockers: dict[ItemSize, LockersPool]
    pin: PinGenerator

    def package_arrived(self, pkg: Package) -> tuple[Locker, str]:
        # Refresh locker pool
        for lockerpool in self.lockers.values():
            lockerpool.refresh()

        # Get a matching locker
        size = pkg.get_size()
        locker = self.lockers[size].reserve()
        if not locker:
            print("handle this")

        # Update locker pin
        delivery_pin = self.pin.generate_pin()
        locker.update_pin(delivery_pin)

        return (locker, delivery_pin)

    def package_ready_pickup(self, locker: "Locker", pkg: "Package"):
        locker.add_package(pkg)
        new_pin = self.pin.generate_pin()
        locker.update_pin(new_pin)
        self.notifier.notify_user(pkg.get_owner(), locker)

    def pickup_done(self, locker: "Locker"):
        self.lockers[locker.get_size].free(locker)
        locker.update_pin(self.pin.generate_pin())
