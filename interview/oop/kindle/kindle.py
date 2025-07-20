from abc import ABC, abstractmethod
from enum import Enum
from datetime import datetime
from dataclasses import dataclass
from collections import defaultdict

"""
Kindle
FR:
- User can have lib of books that they can add or remove from.
- User can "open"/"read" a book.
- Reading App remember user's checkpoint
- Reading app only display a single page at a time

Twist:
- Support Multiple font size 
"""


class Page:
    id: int
    page_num: int
    book_id: int
    data: str


class Book:
    id: int
    name: str
    pages: list[Page]


@dataclass
class UserBookSession:
    book: Book
    current_page: int


class BookLibrary(ABC):
    @abstractmethod
    def add_book(self, book: "Book"):
        pass

    @abstractmethod
    def remove_book(self, book: "Book"):
        pass


class BookLibraryManager(BookLibrary):
    book_library: dict[int, "Book"]

    def __init__(self, books: list[Book]):
        self.book_library = {}
        for book in books:
            self.book_library[book.id] = book

    def add_book(self, book: "Book"):
        if book.id not in self.book_library:
            self.book_library[book.id] = book

    def remove_book(self, book: "Book"):
        if book.id in self.book_library:
            self.book_library.pop(book.id)


class BookNavigation(ABC):
    @abstractmethod
    def load_book(self, book: "Book") -> Page:
        pass

    @abstractmethod
    def prev_page(self) -> Page:
        pass

    @abstractmethod
    def next_page(self) -> Page:
        pass


class DisplaySetting(ABC):
    @abstractmethod
    def display_page(self, page: Page) -> str:
        pass


class PageDisplayer(DisplaySetting):
    font: str
    font_size: str

    def display_page(self, Page) -> str:
        return ""


class BookNavigator(BookNavigation):
    book_sessions: dict[int, UserBookSession]
    active_session: UserBookSession

    def __init__(self):
        self.book_sessions = {}
        self.active_session = None

    def load_book(self, book) -> Page:
        if book.id not in self.book_sessions:
            new_session = UserBookSession(book, 0)
            self.book_sessions[new_session.book.id] = new_session
            self.active_session = new_session
        else:
            self.active_session = self.book_sessions.get(book.id)

        return self.active_session.book.pages[self.active_session.current_page]

    def prev_page(self) -> Page:
        if self.active_session and self.active_session.current_page > 0:
            self.active_session.current_page -= 1
            return self.active_session.book.pages[self.active_session.current_page]
        else:
            raise Exception()  # Todo: Proper err

    def next_page(self) -> Page:
        if (
            self.active_session
            and self.active_session.current_page
            < len(self.active_session.book.pages) - 1
        ):
            self.active_session.current_page += 1
            return self.active_session.book.pages[self.active_session.current_page]
        else:
            raise Exception("")  # TODO: Proper err
