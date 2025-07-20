from abc import ABC, abstractmethod
from enum import Enum
from datetime import datetime
from collections import defaultdict

"""
Base:
Model the core classes and relationships for a library management system (like a public library or university library system).
The system should support different types of library items (books, magazines, DVDs, etc.) that can be borrowed by library members.
Members should be able to search for items, borrow them, and return them.
The system should track due dates, late fees, and handle renewals.
Include functionality for library staff to manage inventory, check items in/out, and handle member accounts.
Support for reservations when items are currently checked out.

Tips
Think about the relationship between members, library items, and loans/transactions
Consider different types of library materials and their borrowing rules
Design for inventory management and item availability tracking
Handle different user types (members, librarians, administrators)
Consider extensibility for different library policies (loan periods, fine calculations, etc.)
"""


class Item(ABC):
    @abstractmethod
    def get_name(self):
        pass

    @abstractmethod
    def get_id(self):
        pass

    @abstractmethod
    def is_reserved(self):
        pass

    @abstractmethod
    def update_presence(self, status):
        pass

    @abstractmethod
    def update_reserved(self, status):
        pass

    @abstractmethod
    def can_loan(self):
        pass


class Book(Item):
    id: int
    name: str
    is_present: bool
    is_reserved: bool

    def get_id(self):
        return self.id

    def get_name(self):
        return self.name

    def reserved(self):
        return self.is_reserved

    def can_loan(self):
        return self.is_present and not self.is_reserved

    def update_presence(self, status: bool):
        self.is_present = status

    def update_reserved(self, status: bool):
        self.is_reserved = status


class Magazine(Item):
    id: int
    name: str
    is_present: bool
    is_reserved: bool

    def get_id(self):
        return self.id

    def get_name(self):
        return self.name

    def reserved(self):
        return self.is_reserved

    def can_loan(self):
        return self.is_present and not self.is_reserved

    def update_presence(self, status: bool):
        self.is_present = status

    def update_reserved(self, status: bool):
        self.is_reserved = status


class DVDs(Item):
    id: int
    name: str
    is_present: bool
    is_reserved: bool

    def get_id(self):
        return self.id

    def get_name(self):
        return self.name

    def reserved(self):
        return self.is_reserved

    def can_loan(self):
        return self.is_present and not self.is_reserved

    def update_presence(self, status: bool):
        self.is_present = status

    def update_reserved(self, status: bool):
        self.is_reserved = status


class Member:
    id: str
    loans: list["Loan"]


class Loan:
    items: list[Item]
    return_time: datetime
    start_time: datetime

    def item_cnt(self):
        return len(self.items)

    def get_return_time(self):
        return self.return_time


class LoanPayment(ABC):
    @abstractmethod
    def loan_payment(self, loan: "Loan") -> int:
        pass


class LateFeeLoanPayment(LoanPayment):
    def loan_payment(self, loan: "Loan") -> int:
        if datetime.now() > loan.get_return_time():
            return 1.5 * loan.item_cnt()


class ItemsLoan(ABC):
    @abstractmethod
    def make_request(self, items: list["Item"]) -> Loan:
        pass


class ItemsLoaner(ItemsLoan):
    def make_request(self, items) -> Loan:
        return


class ItemReserver(ItemsLoan):
    def make_request(self, items) -> Loan:
        return


class ItemsSearch(ABC):
    @abstractmethod
    def search_item(self, item: Item) -> list[Item]:
        pass


class ItemSearcher(ItemsSearch):
    items: defaultdict[str, list]

    def __init__(self, items):
        self.items

    def search(self, keyword: str) -> list[Item]:
        if keyword in self.items:
            return self.items.get(keyword)


class ItemsCheck(ABC):
    @abstractmethod
    def check_in(self, items: list[Item]):
        pass

    @abstractmethod
    def check_out(self, items: list[Item]):
        pass


class ItemChecker(ItemsCheck):
    current_items: set
    outbound_items: set

    def check_in(self, items, Loan: "Loan"):
        for item in items:
            self.current_items.add(item.get_id())
            self.outbound_items.discard(item.get_id())

    def check_out(self, items, Loan: "Loan"):
        for item in items:
            if item.get_id() in self.current_items:
                self.current_items.discard(item.get_id())
                self.outbound_items.add(item.get_id())


class Library:
    checker: ItemsCheck
    searcher: ItemsSearch
    loaner: ItemsLoan
    loan_payment: Item

    def make_loan(self):
        pass

    def return_load(self):
        pass
